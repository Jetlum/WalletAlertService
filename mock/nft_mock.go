package mock

import (
	nfts "github.com/Jetlum/WalletAlertService/nft"
	"github.com/ethereum/go-ethereum/core/types"
)

type MockNFTDetector struct {
	IsNFTTransactionFunc func(tx *types.Transaction) bool
}

func NewMockNFTDetector() nfts.INFTDetector {
	return &MockNFTDetector{
		IsNFTTransactionFunc: func(tx *types.Transaction) bool {
			return false
		},
	}
}

func (m *MockNFTDetector) IsNFTTransaction(tx *types.Transaction) bool {
	if m.IsNFTTransactionFunc != nil {
		return m.IsNFTTransactionFunc(tx)
	}
	return false
}
