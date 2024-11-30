package mock

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MockNFTDetector struct {
	IsNFTTransactionFunc func(tx *types.Transaction) bool
	knownContracts       sync.Map
	callCount            int
}

func NewMockNFTDetector() *MockNFTDetector {
	detector := &MockNFTDetector{}
	detector.knownContracts = sync.Map{}

	// Initialize with test contracts
	testContracts := map[string]bool{
		"0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D": true,
		"0x23581767a106ae21c074b2276D25e5C3e136a68b": true,
	}

	for addr := range testContracts {
		detector.knownContracts.Store(common.HexToAddress(addr), true)
	}

	return detector
}

func (m *MockNFTDetector) IsNFTTransaction(tx *types.Transaction) bool {
	m.callCount++

	// Custom function takes precedence
	if m.IsNFTTransactionFunc != nil {
		return m.IsNFTTransactionFunc(tx)
	}

	if tx.To() == nil {
		return false
	}

	toAddr := tx.To()
	// Check known contracts
	val, exists := m.knownContracts.Load(*toAddr)
	if !exists {
		return false
	}
	return val.(bool)
}

// Helper methods for testing
func (m *MockNFTDetector) GetCallCount() int {
	return m.callCount
}

func (m *MockNFTDetector) AddKnownContract(address common.Address) {
	m.knownContracts.Store(address, true)
}

func (m *MockNFTDetector) RemoveKnownContract(address common.Address) {
	m.knownContracts.Delete(address)
}

func (m *MockNFTDetector) Reset() {
	m.callCount = 0
	m.knownContracts = sync.Map{}
}
