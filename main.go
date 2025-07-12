// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/Jetlum/WalletAlertService/config"
	"github.com/Jetlum/WalletAlertService/database"
	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	nfts "github.com/Jetlum/WalletAlertService/nft"
	"github.com/Jetlum/WalletAlertService/repository"
	"github.com/Jetlum/WalletAlertService/services"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClientWrapper struct {
	*ethclient.Client
}

func (w *EthClientWrapper) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return w.Client.BlockByHash(ctx, hash)
}

func (w *EthClientWrapper) BlockNumber(ctx context.Context) (uint64, error) {
	return w.Client.BlockNumber(ctx)
}

func (w *EthClientWrapper) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return w.Client.HeaderByNumber(ctx, number)
}

func (w *EthClientWrapper) NetworkID(ctx context.Context) (*big.Int, error) {
	return w.Client.NetworkID(ctx)
}

func init() {
	// Skip completely in test mode
	if os.Getenv("GO_ENV") == "test" {
		database.SetupMockDB() // Set mock mode
		return
	}

	// Only run DB initialization in non-test mode
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	if _, err := database.InitDB(cfg.DatabaseURL); err != nil {
		// Only log error in non-test mode
		if !database.IsMockMode {
			log.Fatal("Failed to initialize database:", err)
		}
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	var eventRepo repository.EventRepositoryInterface
	var userPrefRepo repository.UserPreferenceRepositoryInterface

	if os.Getenv("GO_ENV") == "test" {
		eventRepo = mock.NewMockEventRepository()
		userPrefRepo = mock.NewMockUserPreferenceRepository()
	} else {
		eventRepo = repository.NewEventRepository(database.DB)
		userPrefRepo = repository.NewUserPreferenceRepository(database.DB)
	}

	emailNotification := services.NewEmailNotification(cfg.SendGridAPIKey)
	nftDetector := nfts.NewNFTDetector()

	// Initialize price monitoring services
	priceMonitor := services.NewPriceMonitor(cfg.CoinGeckoAPIKey)
	priceAlertRepo := repository.NewPriceAlertRepository(database.DB)
	priceAlertService := services.NewPriceAlertService(
		priceMonitor,
		priceAlertRepo,
		emailNotification,
	)

	// Start monitoring
	priceMonitor.StartMonitoring(1 * time.Minute)
	priceAlertService.StartMonitoring()

	// Connect to Ethereum node
	ethClient, err := ethclient.Dial(fmt.Sprintf("wss://mainnet.infura.io/ws/v3/%s", cfg.InfuraProjectID))
	if err != nil {
		log.Fatal(err)
	}
	defer ethClient.Close()

	client := &EthClientWrapper{Client: ethClient}

	fmt.Println("Connected to Ethereum network")

	// Subscribe to new head (block) events
	headers := make(chan *types.Header)
	// Update subscription
	sub, err := ethClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			processBlock(client, header, nftDetector, emailNotification, eventRepo, userPrefRepo)
		}
	}
}

func processBlock(
	client mock.EthClient,
	header *types.Header,
	nftDetector nfts.INFTDetector,
	emailNotification services.EmailNotifier,
	eventRepo repository.EventRepositoryInterface,
	userPrefRepo repository.UserPreferenceRepositoryInterface,
) {
	block, err := client.BlockByHash(context.Background(), header.Hash())
	if err != nil {
		log.Printf("Error getting block: %v", err)
		return
	}

	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue
		}

		event := createEvent(tx, client)

		if nftDetector.IsNFTTransaction(tx) {
			event.EventType = "NFT_TRANSFER"
		} else if isLargeTransfer(tx) {
			event.EventType = "LARGE_TRANSFER"
		} else {
			continue
		}

		if err := eventRepo.Create(event); err != nil {
			log.Printf("Error saving event: %v", err)
			continue
		}

		notifyUsers(event, userPrefRepo, emailNotification)
	}
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
) {
	preferences, err := userPrefRepo.GetMatchingPreferences(event)
	if err != nil {
		log.Printf("Error getting matching preferences: %v", err)
		return
	}

	for _, pref := range preferences {
		if pref.EmailNotification {
			if err := emailNotification.Send(event, &pref); err != nil {
				log.Printf("Error sending notification: %v", err)
			}
		}
	}
}
