package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jetlum/WalletAlertService/config"
	"github.com/Jetlum/WalletAlertService/database"
	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	nfts "github.com/Jetlum/WalletAlertService/nft"
	"github.com/Jetlum/WalletAlertService/repository"
	"github.com/Jetlum/WalletAlertService/services"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		cancel()
	}()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	// Close the database connection gracefully
	if err := database.CloseDB(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize database
	if err := database.InitDB(cfg.DatabaseURL); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize services
	eventRepo := repository.NewEventRepository(database.DB)
	userPrefRepo := repository.NewUserPreferenceRepository(database.DB)
	emailNotification := services.NewEmailNotification(cfg.SendGridAPIKey)
	nftDetector := nfts.NewNFTDetector()

	// Connect to Ethereum node with retry mechanism
	client, err := connectWithRetry(ctx, cfg.InfuraProjectID)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}
	defer client.Close()

	return processBlocks(ctx, client, nftDetector, emailNotification, eventRepo, userPrefRepo)
}

func connectWithRetry(ctx context.Context, infuraID string) (*ethclient.Client, error) {
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled while connecting: %w", ctx.Err())
		default:
			client, err := ethclient.Dial(fmt.Sprintf("wss://mainnet.infura.io/ws/v3/%s", infuraID))
			if err == nil {
				return client, nil
			}
			log.Printf("Failed to connect to Ethereum node, retrying in 5s...")

			// Use context for sleep
			timer := time.NewTimer(5 * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				return nil, fmt.Errorf("context cancelled during retry: %w", ctx.Err())
			case <-timer.C:
				continue
			}
		}
	}
	return nil, fmt.Errorf("failed to connect after 3 attempts")
}

func processBlocks(
	ctx context.Context,
	client *ethclient.Client,
	nftDetector nfts.INFTDetector,
	emailNotification services.EmailNotifier,
	eventRepo repository.EventRepositoryInterface,
	userPrefRepo repository.UserPreferenceRepositoryInterface,
) error {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new blocks: %w", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %w", err)
		case header := <-headers:
			if err := processBlock(client, header, nftDetector, emailNotification, eventRepo, userPrefRepo); err != nil {
				log.Printf("Error processing block: %v", err)
			}
		}
	}
}

func processBlock(
	client *ethclient.Client,
	header *types.Header,
	nftDetector nfts.INFTDetector,
	emailNotification services.EmailNotifier,
	eventRepo repository.EventRepositoryInterface,
	userPrefRepo repository.UserPreferenceRepositoryInterface,
) error {
	block, err := client.BlockByHash(context.Background(), header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get block: %w", err)
	}

	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue // Skip contract creation transactions
		}

		event := createEvent(tx, client)

		// Determine event type
		if nftDetector.IsNFTTransaction(tx) {
			event.EventType = "NFT_TRANSFER"
		} else if isLargeTransfer(tx) {
			event.EventType = "LARGE_TRANSFER"
		} else {
			continue // Skip if not relevant
		}

		// Save event
		if err := eventRepo.Create(event); err != nil {
			log.Printf("Error saving event: %v", err)
			continue
		}

		// Notify users
		if err := notifyUsers(event, userPrefRepo, emailNotification); err != nil {
			log.Printf("Error notifying users: %v", err)
		}

		// Get matching preferences and notify users
		preferences, err := userPrefRepo.GetMatchingPreferences(event)
		if err != nil {
			log.Printf("Error getting matching preferences: %v", err)
			continue
		}

		for _, pref := range preferences {
			if pref.EmailNotification {
				if err := emailNotification.Send(event, &pref); err != nil {
					log.Printf("Error sending notification: %v", err)
				}
			}
		}
	}

	return nil
}

func createEvent(tx *types.Transaction, client mock.EthClient) *models.Event {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	signer := types.NewEIP155Signer(chainID)
	fromAddress, err := types.Sender(signer, tx)
	if err != nil {
		log.Fatalf("Failed to get sender address: %v", err)
	}

	return &models.Event{
		TxHash:      tx.Hash().Hex(),
		FromAddress: fromAddress.Hex(),
		ToAddress:   tx.To().Hex(),
		Value:       tx.Value().String(),
	}
}

func isLargeTransfer(tx *types.Transaction) bool {
	threshold := big.NewInt(1000000000000000000) // 1 Ether in Wei
	return tx.Value().Cmp(threshold) >= 0
}

func notifyUsers(
	event *models.Event,
	userPrefRepo repository.UserPreferenceRepositoryInterface,
	emailNotification services.EmailNotifier,
) error {
	// Get matching preferences and notify users
	preferences, err := userPrefRepo.GetMatchingPreferences(event)
	if err != nil {
		return fmt.Errorf("error getting matching preferences: %w", err)
	}

	for _, pref := range preferences {
		if pref.EmailNotification {
			if err := emailNotification.Send(event, &pref); err != nil {
				log.Printf("Error sending notification: %v", err)
			}
		}
	}
	return nil
}
