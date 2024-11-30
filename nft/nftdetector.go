package nfts

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// INFTDetector defines the interface for NFT detection
type INFTDetector interface {
	IsNFTTransaction(tx *types.Transaction) bool
}

// NFTDetector implements the INFTDetector interface
type NFTDetector struct {
	nftContracts sync.Map
}

func NewNFTDetector() INFTDetector {
	detector := &NFTDetector{}

	// Initialize known contracts
	knownContracts := map[string]bool{
		"0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D": true, // BAYC
		"0x23581767a106ae21c074b2276D25e5C3e136a68b": true, // Moonbirds
	}

	for addr := range knownContracts {
		detector.nftContracts.Store(common.HexToAddress(addr), true)
	}

	return detector
}

func (d *NFTDetector) IsNFTTransaction(tx *types.Transaction) bool {
	if tx.To() == nil {
		return false
	}

	if isContract, ok := d.nftContracts.Load(*tx.To()); ok {
		return isContract.(bool)
	}

	return false
}
