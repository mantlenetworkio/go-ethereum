package mantle

import (
	"context"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type MtInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction) error
	BlockChain() *core.BlockChain
	MtNode() interface{}
}
