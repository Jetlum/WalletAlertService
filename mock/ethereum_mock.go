package mock

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EthClient defines the interface for Ethereum client operations
type EthClient interface {
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockNumber(ctx context.Context) (uint64, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	NetworkID(ctx context.Context) (*big.Int, error)
}

type MockEthClient struct {
	BlockByHashFunc    func(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockNumberFunc    func(ctx context.Context) (uint64, error)
	HeaderByNumberFunc func(ctx context.Context, number *big.Int) (*types.Header, error)
	NetworkIDFunc      func(ctx context.Context) (*big.Int, error)
}

func NewMockEthClient() *MockEthClient {
	return &MockEthClient{
		BlockByHashFunc: func(ctx context.Context, hash common.Hash) (*types.Block, error) {
			header := &types.Header{
				Number:     big.NewInt(1),
				ParentHash: common.HexToHash("0x123"),
				Time:       uint64(time.Now().Unix()),
			}

			block := types.NewBlockWithHeader(header)
			return block.WithBody(types.Body{
				Transactions: make([]*types.Transaction, 0),
				Uncles:       make([]*types.Header, 0),
			}), nil
		},
		NetworkIDFunc: func(ctx context.Context) (*big.Int, error) {
			return big.NewInt(1), nil
		},
	}
}

func (m *MockEthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return m.BlockByHashFunc(ctx, hash)
}

func (m *MockEthClient) BlockNumber(ctx context.Context) (uint64, error) {
	if m.BlockNumberFunc != nil {
		return m.BlockNumberFunc(ctx)
	}
	return 0, nil
}

func (m *MockEthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	if m.HeaderByNumberFunc != nil {
		return m.HeaderByNumberFunc(ctx, number)
	}
	return nil, nil
}

func (m *MockEthClient) NetworkID(ctx context.Context) (*big.Int, error) {
	if m.NetworkIDFunc != nil {
		return m.NetworkIDFunc(ctx)
	}
	return big.NewInt(1), nil
}
