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

		tx := types.NewTransaction(
			0,
			address, // Remove & operator, use Address directly
			big.NewInt(0),
			21000,
			big.NewInt(1),
			nil,
		)
		result := detector.IsNFTTransaction(tx)
		assert.True(t, result)
	})
}
