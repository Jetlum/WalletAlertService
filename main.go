package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Jetlum/WalletAlertService/config"
	"github.com/Jetlum/WalletAlertService/database"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/Jetlum/WalletAlertService/nfts"
	"github.com/Jetlum/WalletAlertService/repository"
	"github.com/Jetlum/WalletAlertService/services"
)

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	err = database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize services and repositories
	nftDetector := nfts.NewNFTDetector()
	emailNotification := services.NewEmailNotification(cfg.SendGridAPIKey)
	eventRepo := repository.NewEventRepository()
	userPrefRepo := repository.NewUserPreferenceRepository()

	// Connect to Ethereum node
	client, err := ethclient.Dial(fmt.Sprintf("wss://mainnet.infura.io/ws/v3/%s", cfg.InfuraProjectID))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fmt.Println("Connected to Ethereum network")

	// Subscribe to new head (block) events
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
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
	client *ethclient.Client,
	header *types.Header,
	nftDetector *nfts.NFTDetector,
	emailNotification *services.EmailNotification,
	eventRepo *repository.EventRepository,
	userPrefRepo *repository.UserPreferenceRepository,
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

		event := createEvent(tx)

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

func createEvent(tx *types.Transaction) *models.Event {
	return &models.Event{
		TxHash:      tx.Hash().String(),
		FromAddress: tx.From().String(),
		ToAddress:   tx.To().String(),
		Value:       tx.Value().String(),
	}
}

func isLargeTransfer(tx *types.Transaction) bool {
	threshold := big.NewInt(1000000000000000000) // 1 Ether in Wei
	return tx.Value().Cmp(threshold) >= 0
}

func notifyUsers(
	event *models.Event,
	userPrefRepo *repository.UserPreferenceRepository,
	emailNotification *services.EmailNotification,
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
