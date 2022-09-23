package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var MtosAddress = common.HexToAddress("0xa4b05")
var MtSysAddress = common.HexToAddress("0x64")
var MtGasInfoAddress = common.HexToAddress("0x6c")
var MtRetryableTxAddress = common.HexToAddress("0x6e")
var NodeInterfaceAddress = common.HexToAddress("0xc8")
var NodeInterfaceDebugAddress = common.HexToAddress("0xc9")

type mantleSigner struct{ Signer }

func NewMantleSigner(signer Signer) Signer {
	return mantleSigner{Signer: signer}
}

func (s mantleSigner) Sender(tx *Transaction) (common.Address, error) {
	switch inner := tx.inner.(type) {
	case *MantleUnsignedTx:
		return inner.From, nil
	case *MantleContractTx:
		return inner.From, nil
	case *MantleDepositTx:
		return inner.From, nil
	case *MantleInternalTx:
		return MtosAddress, nil
	case *MantleRetryTx:
		return inner.From, nil
	case *MantleSubmitRetryableTx:
		return inner.From, nil
	case *MantleLegacyTxData:
		legacyData := tx.inner.(*MantleLegacyTxData)
		if legacyData.Sender != nil {
			return *legacyData.Sender, nil
		}
		fakeTx := NewTx(&legacyData.LegacyTx)
		return s.Signer.Sender(fakeTx)
	default:
		return s.Signer.Sender(tx)
	}
}

func (s mantleSigner) Equal(s2 Signer) bool {
	x, ok := s2.(mantleSigner)
	return ok && x.Signer.Equal(s.Signer)
}

func (s mantleSigner) SignatureValues(tx *Transaction, sig []byte) (R, S, V *big.Int, err error) {
	switch tx.inner.(type) {
	case *MantleUnsignedTx:
		return bigZero, bigZero, bigZero, nil
	case *MantleContractTx:
		return bigZero, bigZero, bigZero, nil
	case *MantleDepositTx:
		return bigZero, bigZero, bigZero, nil
	case *MantleInternalTx:
		return bigZero, bigZero, bigZero, nil
	case *MantleRetryTx:
		return bigZero, bigZero, bigZero, nil
	case *MantleSubmitRetryableTx:
		return bigZero, bigZero, bigZero, nil
	case *MantleLegacyTxData:
		legacyData := tx.inner.(*MantleLegacyTxData)
		fakeTx := NewTx(&legacyData.LegacyTx)
		return s.Signer.SignatureValues(fakeTx, sig)
	default:
		return s.Signer.SignatureValues(tx, sig)
	}
}

// Hash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (s mantleSigner) Hash(tx *Transaction) common.Hash {
	if legacyData, isMtLegacy := tx.inner.(*MantleLegacyTxData); isMtLegacy {
		fakeTx := NewTx(&legacyData.LegacyTx)
		return s.Signer.Hash(fakeTx)
	}
	return s.Signer.Hash(tx)
}
