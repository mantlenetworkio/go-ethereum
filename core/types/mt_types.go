package types

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/common"
)

type fallbackError struct {
}

var fallbackErrorMsg = "missing trie node 0000000000000000000000000000000000000000000000000000000000000000 (path ) <nil>"
var fallbackErrorCode = -32000

func SetFallbackError(msg string, code int) {
	fallbackErrorMsg = msg
	fallbackErrorCode = code
	log.Debug("setting fallback error", "msg", msg, "code", code)
}

func (f fallbackError) ErrorCode() int { return fallbackErrorCode }
func (f fallbackError) Error() string  { return fallbackErrorMsg }

var ErrUseFallback = fallbackError{}

type FallbackClient interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

var bigZero = big.NewInt(0)

func (tx *LegacyTx) isFake() bool     { return false }
func (tx *AccessListTx) isFake() bool { return false }
func (tx *DynamicFeeTx) isFake() bool { return false }

type MantleUnsignedTx struct {
	ChainId *big.Int
	From    common.Address

	Nonce     uint64          // nonce of sender account
	GasFeeCap *big.Int        // wei per gas
	Gas       uint64          // gas limit
	To        *common.Address `rlp:"nil"` // nil means contract creation
	Value     *big.Int        // wei amount
	Data      []byte          // contract invocation input data
}

func (tx *MantleUnsignedTx) txType() byte { return MantleUnsignedTxType }

func (tx *MantleUnsignedTx) copy() TxData {
	cpy := &MantleUnsignedTx{
		ChainId:   new(big.Int),
		Nonce:     tx.Nonce,
		GasFeeCap: new(big.Int),
		Gas:       tx.Gas,
		From:      tx.From,
		To:        nil,
		Value:     new(big.Int),
		Data:      common.CopyBytes(tx.Data),
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.To != nil {
		tmp := *tx.To
		cpy.To = &tmp
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	return cpy
}

func (tx *MantleUnsignedTx) chainID() *big.Int      { return tx.ChainId }
func (tx *MantleUnsignedTx) accessList() AccessList { return nil }
func (tx *MantleUnsignedTx) data() []byte           { return tx.Data }
func (tx *MantleUnsignedTx) gas() uint64            { return tx.Gas }
func (tx *MantleUnsignedTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *MantleUnsignedTx) gasTipCap() *big.Int    { return bigZero }
func (tx *MantleUnsignedTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *MantleUnsignedTx) value() *big.Int        { return tx.Value }
func (tx *MantleUnsignedTx) nonce() uint64          { return tx.Nonce }
func (tx *MantleUnsignedTx) to() *common.Address    { return tx.To }
func (tx *MantleUnsignedTx) isFake() bool           { return false }

func (tx *MantleUnsignedTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}

func (tx *MantleUnsignedTx) setSignatureValues(chainID, v, r, s *big.Int) {

}

type MantleContractTx struct {
	ChainId   *big.Int
	RequestId common.Hash
	From      common.Address

	GasFeeCap *big.Int        // wei per gas
	Gas       uint64          // gas limit
	To        *common.Address `rlp:"nil"` // nil means contract creation
	Value     *big.Int        // wei amount
	Data      []byte          // contract invocation input data
}

func (tx *MantleContractTx) txType() byte { return MantleContractTxType }

func (tx *MantleContractTx) copy() TxData {
	cpy := &MantleContractTx{
		ChainId:   new(big.Int),
		RequestId: tx.RequestId,
		GasFeeCap: new(big.Int),
		Gas:       tx.Gas,
		From:      tx.From,
		To:        nil,
		Value:     new(big.Int),
		Data:      common.CopyBytes(tx.Data),
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.To != nil {
		tmp := *tx.To
		cpy.To = &tmp
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	return cpy
}

func (tx *MantleContractTx) chainID() *big.Int      { return tx.ChainId }
func (tx *MantleContractTx) accessList() AccessList { return nil }
func (tx *MantleContractTx) data() []byte           { return tx.Data }
func (tx *MantleContractTx) gas() uint64            { return tx.Gas }
func (tx *MantleContractTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *MantleContractTx) gasTipCap() *big.Int    { return bigZero }
func (tx *MantleContractTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *MantleContractTx) value() *big.Int        { return tx.Value }
func (tx *MantleContractTx) nonce() uint64          { return 0 }
func (tx *MantleContractTx) to() *common.Address    { return tx.To }
func (tx *MantleContractTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}
func (tx *MantleContractTx) setSignatureValues(chainID, v, r, s *big.Int) {}
func (tx *MantleContractTx) isFake() bool                                 { return true }

type MantleRetryTx struct {
	ChainId *big.Int
	Nonce   uint64
	From    common.Address

	GasFeeCap           *big.Int        // wei per gas
	Gas                 uint64          // gas limit
	To                  *common.Address `rlp:"nil"` // nil means contract creation
	Value               *big.Int        // wei amount
	Data                []byte          // contract invocation input data
	TicketId            common.Hash
	RefundTo            common.Address
	MaxRefund           *big.Int // the maximum refund sent to RefundTo (the rest goes to From)
	SubmissionFeeRefund *big.Int // the submission fee to refund if successful (capped by MaxRefund)
}

func (tx *MantleRetryTx) txType() byte { return MantleRetryTxType }

func (tx *MantleRetryTx) copy() TxData {
	cpy := &MantleRetryTx{
		ChainId:             new(big.Int),
		Nonce:               tx.Nonce,
		GasFeeCap:           new(big.Int),
		Gas:                 tx.Gas,
		From:                tx.From,
		To:                  nil,
		Value:               new(big.Int),
		Data:                common.CopyBytes(tx.Data),
		TicketId:            tx.TicketId,
		RefundTo:            tx.RefundTo,
		MaxRefund:           new(big.Int),
		SubmissionFeeRefund: new(big.Int),
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.To != nil {
		tmp := *tx.To
		cpy.To = &tmp
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.MaxRefund != nil {
		cpy.MaxRefund.Set(tx.MaxRefund)
	}
	if tx.SubmissionFeeRefund != nil {
		cpy.SubmissionFeeRefund.Set(tx.SubmissionFeeRefund)
	}
	return cpy
}

func (tx *MantleRetryTx) chainID() *big.Int      { return tx.ChainId }
func (tx *MantleRetryTx) accessList() AccessList { return nil }
func (tx *MantleRetryTx) data() []byte           { return tx.Data }
func (tx *MantleRetryTx) gas() uint64            { return tx.Gas }
func (tx *MantleRetryTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *MantleRetryTx) gasTipCap() *big.Int    { return bigZero }
func (tx *MantleRetryTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *MantleRetryTx) value() *big.Int        { return tx.Value }
func (tx *MantleRetryTx) nonce() uint64          { return tx.Nonce }
func (tx *MantleRetryTx) to() *common.Address    { return tx.To }
func (tx *MantleRetryTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}
func (tx *MantleRetryTx) setSignatureValues(chainID, v, r, s *big.Int) {}
func (tx *MantleRetryTx) isFake() bool                                 { return true }

type MantleSubmitRetryableTx struct {
	ChainId   *big.Int
	RequestId common.Hash
	From      common.Address
	L1BaseFee *big.Int

	DepositValue     *big.Int
	GasFeeCap        *big.Int        // wei per gas
	Gas              uint64          // gas limit
	RetryTo          *common.Address `rlp:"nil"` // nil means contract creation
	RetryValue       *big.Int        // wei amount
	Beneficiary      common.Address
	MaxSubmissionFee *big.Int
	FeeRefundAddr    common.Address
	RetryData        []byte // contract invocation input data
}

func (tx *MantleSubmitRetryableTx) txType() byte { return MantleSubmitRetryableTxType }

func (tx *MantleSubmitRetryableTx) copy() TxData {
	cpy := &MantleSubmitRetryableTx{
		ChainId:          new(big.Int),
		RequestId:        tx.RequestId,
		DepositValue:     new(big.Int),
		L1BaseFee:        new(big.Int),
		GasFeeCap:        new(big.Int),
		Gas:              tx.Gas,
		From:             tx.From,
		RetryTo:          tx.RetryTo,
		RetryValue:       new(big.Int),
		Beneficiary:      tx.Beneficiary,
		MaxSubmissionFee: new(big.Int),
		FeeRefundAddr:    tx.FeeRefundAddr,
		RetryData:        common.CopyBytes(tx.RetryData),
	}
	if tx.ChainId != nil {
		cpy.ChainId.Set(tx.ChainId)
	}
	if tx.DepositValue != nil {
		cpy.DepositValue.Set(tx.DepositValue)
	}
	if tx.L1BaseFee != nil {
		cpy.L1BaseFee.Set(tx.L1BaseFee)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.RetryTo != nil {
		tmp := *tx.RetryTo
		cpy.RetryTo = &tmp
	}
	if tx.RetryValue != nil {
		cpy.RetryValue.Set(tx.RetryValue)
	}
	if tx.MaxSubmissionFee != nil {
		cpy.MaxSubmissionFee.Set(tx.MaxSubmissionFee)
	}
	return cpy
}

func (tx *MantleSubmitRetryableTx) chainID() *big.Int      { return tx.ChainId }
func (tx *MantleSubmitRetryableTx) accessList() AccessList { return nil }
func (tx *MantleSubmitRetryableTx) gas() uint64            { return tx.Gas }
func (tx *MantleSubmitRetryableTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *MantleSubmitRetryableTx) gasTipCap() *big.Int    { return big.NewInt(0) }
func (tx *MantleSubmitRetryableTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *MantleSubmitRetryableTx) value() *big.Int        { return common.Big0 }
func (tx *MantleSubmitRetryableTx) nonce() uint64          { return 0 }
func (tx *MantleSubmitRetryableTx) to() *common.Address    { return &MtRetryableTxAddress }
func (tx *MantleSubmitRetryableTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}
func (tx *MantleSubmitRetryableTx) setSignatureValues(chainID, v, r, s *big.Int) {}
func (tx *MantleSubmitRetryableTx) isFake() bool                                 { return true }

func (tx *MantleSubmitRetryableTx) data() []byte {
	var retryTo common.Address
	if tx.RetryTo != nil {
		retryTo = *tx.RetryTo
	}
	data := make([]byte, 0)
	data = append(data, tx.RequestId.Bytes()...)
	data = append(data, math.U256Bytes(tx.L1BaseFee)...)
	data = append(data, math.U256Bytes(tx.DepositValue)...)
	data = append(data, math.U256Bytes(tx.RetryValue)...)
	data = append(data, math.U256Bytes(tx.GasFeeCap)...)
	data = append(data, math.U256Bytes(new(big.Int).SetUint64(tx.Gas))...)
	data = append(data, math.U256Bytes(tx.MaxSubmissionFee)...)
	data = append(data, make([]byte, 12)...)
	data = append(data, tx.FeeRefundAddr.Bytes()...)
	data = append(data, make([]byte, 12)...)
	data = append(data, tx.Beneficiary.Bytes()...)
	data = append(data, make([]byte, 12)...)
	data = append(data, retryTo.Bytes()...)
	offset := len(data) + 32
	data = append(data, math.U256Bytes(big.NewInt(int64(offset)))...)
	data = append(data, math.U256Bytes(big.NewInt(int64(len(tx.RetryData))))...)
	data = append(data, tx.RetryData...)
	extra := len(tx.RetryData) % 32
	if extra > 0 {
		data = append(data, make([]byte, 32-extra)...)
	}
	data = append(hexutil.MustDecode("0xc9f95d32"), data...)
	return data
}

type MantleDepositTx struct {
	ChainId     *big.Int
	L1RequestId common.Hash
	From        common.Address
	To          common.Address
	Value       *big.Int
}

func (d *MantleDepositTx) txType() byte {
	return MantleDepositTxType
}

func (d *MantleDepositTx) copy() TxData {
	tx := &MantleDepositTx{
		ChainId:     new(big.Int),
		L1RequestId: d.L1RequestId,
		From:        d.From,
		To:          d.To,
		Value:       new(big.Int),
	}
	if d.ChainId != nil {
		tx.ChainId.Set(d.ChainId)
	}
	if d.Value != nil {
		tx.Value.Set(d.Value)
	}
	return tx
}

func (d *MantleDepositTx) chainID() *big.Int      { return d.ChainId }
func (d *MantleDepositTx) accessList() AccessList { return nil }
func (d *MantleDepositTx) data() []byte           { return nil }
func (d *MantleDepositTx) gas() uint64            { return 0 }
func (d *MantleDepositTx) gasPrice() *big.Int     { return bigZero }
func (d *MantleDepositTx) gasTipCap() *big.Int    { return bigZero }
func (d *MantleDepositTx) gasFeeCap() *big.Int    { return bigZero }
func (d *MantleDepositTx) value() *big.Int        { return d.Value }
func (d *MantleDepositTx) nonce() uint64          { return 0 }
func (d *MantleDepositTx) to() *common.Address    { return &d.To }
func (d *MantleDepositTx) isFake() bool           { return true }

func (d *MantleDepositTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}

func (d *MantleDepositTx) setSignatureValues(chainID, v, r, s *big.Int) {

}

type MantleInternalTx struct {
	ChainId *big.Int
	Data    []byte
}

func (t *MantleInternalTx) txType() byte {
	return MantleInternalTxType
}

func (t *MantleInternalTx) copy() TxData {
	return &MantleInternalTx{
		new(big.Int).Set(t.ChainId),
		common.CopyBytes(t.Data),
	}
}

func (t *MantleInternalTx) chainID() *big.Int      { return t.ChainId }
func (t *MantleInternalTx) accessList() AccessList { return nil }
func (t *MantleInternalTx) data() []byte           { return t.Data }
func (t *MantleInternalTx) gas() uint64            { return 0 }
func (t *MantleInternalTx) gasPrice() *big.Int     { return bigZero }
func (t *MantleInternalTx) gasTipCap() *big.Int    { return bigZero }
func (t *MantleInternalTx) gasFeeCap() *big.Int    { return bigZero }
func (t *MantleInternalTx) value() *big.Int        { return common.Big0 }
func (t *MantleInternalTx) nonce() uint64          { return 0 }
func (t *MantleInternalTx) to() *common.Address    { return &MtosAddress }
func (t *MantleInternalTx) isFake() bool           { return true }

func (d *MantleInternalTx) rawSignatureValues() (v, r, s *big.Int) {
	return bigZero, bigZero, bigZero
}

func (d *MantleInternalTx) setSignatureValues(chainID, v, r, s *big.Int) {

}

type HeaderInfo struct {
	SendRoot          common.Hash
	SendCount         uint64
	L1BlockNumber     uint64
	MtOSFormatVersion uint64
}

func (info HeaderInfo) extra() []byte {
	return info.SendRoot[:]
}

func (info HeaderInfo) mixDigest() [32]byte {
	mixDigest := common.Hash{}
	binary.BigEndian.PutUint64(mixDigest[:8], info.SendCount)
	binary.BigEndian.PutUint64(mixDigest[8:16], info.L1BlockNumber)
	binary.BigEndian.PutUint64(mixDigest[16:24], info.MtOSFormatVersion)
	return mixDigest
}

func (info HeaderInfo) UpdateHeaderWithInfo(header *Header) {
	header.MixDigest = info.mixDigest()
	header.Extra = info.extra()
}

func DeserializeHeaderExtraInformation(header *Header) (HeaderInfo, error) {
	if header.BaseFee == nil || header.BaseFee.Sign() == 0 || len(header.Extra) == 0 {
		// imported blocks have no base fee
		// The genesis block doesn't have an MtOS encoded extra field
		return HeaderInfo{}, nil
	}
	if len(header.Extra) != 32 {
		return HeaderInfo{}, fmt.Errorf("unexpected header extra field length %v", len(header.Extra))
	}
	extra := HeaderInfo{}
	copy(extra.SendRoot[:], header.Extra)
	extra.SendCount = binary.BigEndian.Uint64(header.MixDigest[:8])
	extra.L1BlockNumber = binary.BigEndian.Uint64(header.MixDigest[8:16])
	extra.MtOSFormatVersion = binary.BigEndian.Uint64(header.MixDigest[16:24])
	return extra, nil
}
