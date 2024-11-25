package mock

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type MockNFTDetector struct {
	IsNFTTransactionFunc func(tx *types.Transaction) bool
}

func NewMockNFTDetector() *MockNFTDetector {
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
