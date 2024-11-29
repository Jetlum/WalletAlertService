package mock

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

// mock/nft_mock_test.go
func TestMockNFTDetector(t *testing.T) {
	t.Run("Known contract", func(t *testing.T) {
		detector := NewMockNFTDetector() // Create fresh instance

		contractAddr := common.HexToAddress("0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")

		// Verify initial state
		val, exists := detector.knownContracts.Load(contractAddr)
		assert.True(t, exists, "Contract should be in initial known contracts")
		assert.True(t, val.(bool), "Contract should be marked as valid")

		tx := types.NewTransaction(
			0,
			contractAddr,
			big.NewInt(0),
			21000,
			big.NewInt(1),
			nil,
		)

		result := detector.IsNFTTransaction(tx)
		assert.True(t, result, "Transaction to known NFT contract should be detected")
	})
}
