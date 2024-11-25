package models

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (up *UserPreference) Validate() error {
	if up.UserID == "" {
		return fmt.Errorf("user ID is required")
	}

	if up.WalletAddress == "" {
		return fmt.Errorf("wallet address is required")
	}

	if !common.IsHexAddress(up.WalletAddress) {
		return fmt.Errorf("invalid wallet address format")
	}

	if up.MinEtherValue != "" {
		value, ok := new(big.Int).SetString(up.MinEtherValue, 10)
		if !ok {
			return fmt.Errorf("invalid minimum ether value")
		}
		if value.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("minimum ether value cannot be negative")
		}
	}

	return nil
}
