package main

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Setup test context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userPref := models.UserPreference{
		UserID:            "test@example.com",
		WalletAddress:     "0xto",
		MinEtherValue:     "500000000000000000", // 0.5 ETH
		EmailNotification: true,
	}

	mockUserPrefRepo := &mock.MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(e *models.Event) ([]models.UserPreference, error) {
			return []models.UserPreference{userPref}, nil
		},
	}

	// Create mock Ethereum client
	mockClient := mock.NewMockEthClient()

	// Create and use fromAddress in assertions
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	mockEventRepo := &mock.MockEventRepository{
		CreateFunc: func(e *models.Event) error {
			assert.Equal(t, fromAddress.Hex(), e.FromAddress, "From address should match")
			return nil
		},
	}
	// Create mock transaction (unsigned)
	mockTx := types.NewTransaction(
		0,
		common.HexToAddress("0xto"),
		big.NewInt(2000000000000000000), // 2 ETH (above threshold)
		21000,
		big.NewInt(1),
		nil,
	)

	// Sign the transaction
	chainID := big.NewInt(1) // Use the same chain ID in the mock client
	signer := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(mockTx, signer, privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	// Update NetworkIDFunc to return the correct chain ID
	mockClient.NetworkIDFunc = func(ctx context.Context) (*big.Int, error) {
		return chainID, nil
	}

	// Fix block.WithBody call
	mockClient.BlockByHashFunc = func(ctx context.Context, hash common.Hash) (*types.Block, error) {
		header := &types.Header{
			Number:     big.NewInt(1),
			ParentHash: common.HexToHash("0x123"),
			Time:       uint64(time.Now().Unix()),
		}

		block := types.NewBlockWithHeader(header)
		return block.WithBody(types.Body{
			Transactions: []*types.Transaction{signedTx},
			Uncles:       []*types.Header{},
		}), nil
	}

	// Use mock NFT detector
	mockNFTDetector := &mock.MockNFTDetector{
		IsNFTTransactionFunc: func(tx *types.Transaction) bool {
			return false
		},
	}

	emailSent := false
	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(event *models.Event, userPref *models.UserPreference) error {
			emailSent = true
			// Add assertions if needed
			return nil
		},
	}

	// Create mock header
	mockHeader := &types.Header{
		Number: big.NewInt(1),
	}

	done := make(chan bool)
	go func() {
		processBlock(
			mockClient,
			mockHeader,
			mockNFTDetector,
			mockEmailNotification,
			mockEventRepo,
			mockUserPrefRepo,
		)
		done <- true
	}()

	select {
	case <-ctx.Done():
		t.Fatal("Test timed out")
	case <-done:
		assert.True(t, emailSent, "Email notification should have been sent")
	}
}
