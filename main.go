package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"interface.social/config"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"interface.social/database"
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
	// Connect to Ethereum node
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID")
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
			// Get the block
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			// Process transactions
			for _, tx := range block.Transactions() {
				// Filter based on user preferences (e.g., large transfers)
				if isLargeTransfer(tx) {
					// Index the event
					indexEvent(tx)
					// Send notification
					sendNotification(tx)
				}
			}
		}
	}

	nftDetector := services.NewNFTDetector()
	emailNotification := services.NewEmailNotification(cfg.SendGridAPIKey)
	eventRepo := repository.NewEventRepository()
	userPrefRepo := repository.NewUserPreferenceRepository()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Printf("Error getting block: %v", err)
				continue
			}

			for _, tx := range block.Transactions() {
				if tx.To() == nil {
					continue
				}

				event := &models.Event{
					TxHash:      tx.Hash().String(),
					FromAddress: tx.From().String(),
					ToAddress:   tx.To().String(),
					Value:       tx.Value().String(),
				}

				if nftDetector.IsNFTTransaction(tx) {
					event.EventType = "NFT_TRANSFER"
				} else if isLargeTransfer(tx) {
					event.EventType = "LARGE_TRANSFER"
				} else {
					continue
				}

				// Save event to database
				if err := eventRepo.Create(event); err != nil {
					log.Printf("Error saving event: %v", err)
					continue
				}

				// Notify users
				preferences, _ := userPrefRepo.GetMatchingPreferences(event)
				for _, pref := range preferences {
					if pref.EmailNotification {
						if err := emailNotification.Send(event, &pref); err != nil {
							log.Printf("Error sending notification: %v", err)
						}
					}
				}
			}
		}
	}
}

func isLargeTransfer(tx *types.Transaction) bool {
	// Implement logic to check if the transaction is a large transfer
	// For example, check if the value exceeds a certain threshold
	threshold := big.NewInt(1000000000000000000) // 1 Ether in Wei
	return tx.Value().Cmp(threshold) >= 0
}

func indexEvent(tx *types.Transaction) {
	// Implement indexing logic (e.g., store in a database)
}

func sendNotification(tx *types.Transaction) {
	// Implement notification logic (e.g., send email or SMS)
}
