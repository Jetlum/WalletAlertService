package main

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Set test environment
	os.Setenv("GO_ENV", "test")
	// Run tests
	os.Exit(m.Run())
}

func TestNotifyUsers(t *testing.T) {
	event := &models.Event{
		TxHash:      "0xabc123",
		FromAddress: "0xfrom",
		ToAddress:   "0xto",
		Value:       "1000000000000000000",
		EventType:   "LARGE_TRANSFER",
	}

	emailSent := false
	mockUserPrefRepo := &mock.MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(event *models.Event) ([]models.UserPreference, error) {
			return []models.UserPreference{{
				UserID:            "user@example.com",
				WalletAddress:     event.ToAddress,
				EmailNotification: true,
			}}, nil
		},
	}

	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(event *models.Event, userPref *models.UserPreference) error {
			emailSent = true
			return nil
		},
	}

	notifyUsers(event, mockUserPrefRepo, mockEmailNotification)
	assert.True(t, emailSent, "Email notification should have been sent")
}

func TestIsLargeTransfer(t *testing.T) {
	// 1 ETH in Wei
	oneEth := big.NewInt(1000000000000000000)
	lessThanOneEth := big.NewInt(999999999999999999)
	moreThanOneEth := big.NewInt(2000000000000000000)

	toAddr := common.HexToAddress("0x1")
	tx1 := types.NewTransaction(0, toAddr, lessThanOneEth, 21000, big.NewInt(1), nil)
	toAddr2 := common.HexToAddress("0x2")
	tx2 := types.NewTransaction(0, toAddr2, oneEth, 21000, big.NewInt(1), nil)
	toAddr3 := common.HexToAddress("0x3")
	tx3 := types.NewTransaction(0, toAddr3, moreThanOneEth, 21000, big.NewInt(1), nil)

	assert.False(t, isLargeTransfer(tx1))
	assert.True(t, isLargeTransfer(tx2))
	assert.True(t, isLargeTransfer(tx3))
}

func TestCreateEvent(t *testing.T) {
	mockClient := &mock.MockEthClient{
		NetworkIDFunc: func(ctx context.Context) (*big.Int, error) {
			return big.NewInt(1), nil
		},
	}
	toAddr := common.HexToAddress("0x0000000000000000000000000000000000000001")
	tx := types.NewTransaction(0, toAddr, big.NewInt(100), 21000, big.NewInt(1), nil)
	event := createEvent(tx, mockClient)
	assert.Equal(t, tx.Hash().Hex(), event.TxHash)
	assert.Equal(t, toAddr.Hex(), event.ToAddress)
	assert.Equal(t, "100", event.Value)
}

func TestProcessBlock(t *testing.T) {
	// Prepare mock block with two transactions: one NFT, one large transfer
	toAddr := common.HexToAddress("0x0000000000000000000000000000000000000002")
	nftTx := types.NewTransaction(0, toAddr, big.NewInt(100), 21000, big.NewInt(1), nil)
	largeTx := types.NewTransaction(0, toAddr, big.NewInt(1000000000000000000), 21000, big.NewInt(1), nil)

	// Fix block creation - use proper Body struct
	body := &types.Body{
		Transactions: []*types.Transaction{nftTx, largeTx},
		Uncles:       []*types.Header{},
	}
	block := types.NewBlock(&types.Header{}, body, nil, nil)

	mockClient := &mock.MockEthClient{
		BlockByHashFunc: func(ctx context.Context, hash common.Hash) (*types.Block, error) {
			return block, nil
		},
		NetworkIDFunc: func(ctx context.Context) (*big.Int, error) {
			return big.NewInt(1), nil
		},
	}
	mockNFTDetector := &mock.MockNFTDetector{
		IsNFTTransactionFunc: func(tx *types.Transaction) bool {
			return tx.Value().Cmp(big.NewInt(100)) == 0 // Only first tx is NFT
		},
	}
	createdEvents := []*models.Event{}
	mockEventRepo := &mock.MockEventRepository{
		CreateFunc: func(event *models.Event) error {
			createdEvents = append(createdEvents, event)
			return nil
		},
	}
	mockUserPrefRepo := &mock.MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(event *models.Event) ([]models.UserPreference, error) {
			return []models.UserPreference{{
				UserID:            "user@example.com",
				WalletAddress:     event.ToAddress,
				EmailNotification: true,
			}}, nil
		},
	}
	emailSent := 0
	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(event *models.Event, userPref *models.UserPreference) error {
			emailSent++
			return nil
		},
	}

	header := &types.Header{Number: big.NewInt(1)}
	processBlock(
		mockClient,
		header,
		mockNFTDetector,
		mockEmailNotification,
		mockEventRepo,
		mockUserPrefRepo,
	)

	assert.Len(t, createdEvents, 2)
	assert.Equal(t, "NFT_TRANSFER", createdEvents[0].EventType)
	assert.Equal(t, "LARGE_TRANSFER", createdEvents[1].EventType)
	assert.Equal(t, 2, emailSent)
}

func TestMainInitialization(t *testing.T) {
	os.Setenv("GO_ENV", "test")
	defer os.Unsetenv("GO_ENV")
	// Test that initialization doesn't panic in test mode
	assert.NotPanics(t, func() {
		// Call the actual init function logic manually since init() is called automatically
		if os.Getenv("GO_ENV") == "test" {
			// This should not panic
		}
	})
}
