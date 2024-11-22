package nftservices

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type NFTDetector struct {
	nftContracts map[common.Address]bool
}

func NewNFTDetector() *NFTDetector {
	return &NFTDetector{
		nftContracts: map[common.Address]bool{
			common.HexToAddress("0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D"): true, // BAYC
			common.HexToAddress("0x23581767a106ae21c074b2276D25e5C3e136a68b"): true, // Moonbirds
		},
	}
}

func (d *NFTDetector) IsNFTTransaction(tx *types.Transaction) bool {
	if tx.To() == nil {
		return false
	}
	_, exists := d.nftContracts[*tx.To()]
	return exists
}
