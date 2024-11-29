package mock

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestMockNFTDetector(t *testing.T) {
	detector := NewMockNFTDetector()

	t.Run("Known contract", func(t *testing.T) {
		detector.Reset()
		address := common.HexToAddress("0xBAYC")
		detector.AddKnownContract(address)

		// Create transaction to the known contract address
		tx := types.NewTransaction(
			0,
			address,
			big.NewInt(0),
			21000,
			big.NewInt(1),
			nil,
		)

		// Add debug assertions
		toAddr := tx.To()
		if toAddr == nil {
			t.Fatal("Transaction To address is nil")
		}

		// Verify the addresses match
		assert.Equal(t, address, *toAddr, "Transaction address should match known contract")

		// Check if transaction is detected as NFT
		result := detector.IsNFTTransaction(tx)
		assert.True(t, result, "Transaction to known NFT contract address should return true")
	})
}
