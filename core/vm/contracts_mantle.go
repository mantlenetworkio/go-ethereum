package vm

import "github.com/ethereum/go-ethereum/common"

var (
	PrecompiledContractsMantle = make(map[common.Address]PrecompiledContract)
	PrecompiledAddressesMantle []common.Address
)
