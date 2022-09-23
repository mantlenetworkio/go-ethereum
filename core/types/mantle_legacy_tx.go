package types

import (
	"bytes"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type MantleLegacyTxData struct {
	LegacyTx
	HashOverride      common.Hash // Hash cannot be locally computed from other fields
	EffectiveGasPrice uint64
	L1BlockNumber     uint64
	Sender            *common.Address `rlp:"optional,nil"` // only used in unsigned Txs
}

func NewMantleLegacyTx(origTx *Transaction, hashOverride common.Hash, effectiveGas uint64, l1Block uint64, senderOverride *common.Address) (*Transaction, error) {
	if origTx.Type() != LegacyTxType {
		return nil, errors.New("attempt to mantle-wrap non-legacy transaction")
	}
	legacyPtr := origTx.GetInner().(*LegacyTx)
	inner := MantleLegacyTxData{
		LegacyTx:          *legacyPtr,
		HashOverride:      hashOverride,
		EffectiveGasPrice: effectiveGas,
		L1BlockNumber:     l1Block,
		Sender:            senderOverride,
	}
	return NewTx(&inner), nil
}

func (tx *MantleLegacyTxData) copy() TxData {
	legacyCopy := tx.LegacyTx.copy().(*LegacyTx)
	var sender *common.Address
	if tx.Sender != nil {
		sender = new(common.Address)
		*sender = *tx.Sender
	}
	return &MantleLegacyTxData{
		LegacyTx:          *legacyCopy,
		HashOverride:      tx.HashOverride,
		EffectiveGasPrice: tx.EffectiveGasPrice,
		L1BlockNumber:     tx.L1BlockNumber,
		Sender:            sender,
	}
}

func (tx *MantleLegacyTxData) txType() byte { return MantleLegacyTxType }

func (tx *MantleLegacyTxData) EncodeOnlyLegacyInto(w *bytes.Buffer) {
	rlp.Encode(w, tx.LegacyTx)
}
