// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type MantleChainParams struct {
	EnableMtOS                bool
	AllowDebugPrecompiles     bool
	DataAvailabilityCommittee bool
	InitialMtOSVersion        uint64
	InitialChainOwner         common.Address
	GenesisBlockNum           uint64
}

func (c *ChainConfig) IsMantle() bool {
	return c.MantleChainParams.EnableMtOS
}

func (c *ChainConfig) IsMantleMantle(num *big.Int) bool {
	return c.IsMantle() && isForked(new(big.Int).SetUint64(c.MantleChainParams.GenesisBlockNum), num)
}

func (c *ChainConfig) DebugMode() bool {
	return c.MantleChainParams.AllowDebugPrecompiles
}

func (c *ChainConfig) checkMantleCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if c.IsMantle() != newcfg.IsMantle() {
		// This difference applies to the entire chain, so report that the genesis block is where the difference appears.
		return newCompatError("isMantle", common.Big0, common.Big0)
	}
	if !c.IsMantle() {
		return nil
	}
	cMt := &c.MantleChainParams
	newMt := &newcfg.MantleChainParams
	if cMt.GenesisBlockNum != newMt.GenesisBlockNum {
		return newCompatError("genesisblocknum", new(big.Int).SetUint64(cMt.GenesisBlockNum), new(big.Int).SetUint64(newMt.GenesisBlockNum))
	}
	return nil
}

func MantleOneParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                true,
		AllowDebugPrecompiles:     false,
		DataAvailabilityCommittee: false,
		InitialMtOSVersion:        6,
		InitialChainOwner:         common.HexToAddress("0xd345e41ae2cb00311956aa7109fc801ae8c81a52"),
	}
}

func MantleNovaParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                true,
		AllowDebugPrecompiles:     false,
		DataAvailabilityCommittee: true,
		InitialMtOSVersion:        1,
		InitialChainOwner:         common.HexToAddress("0x9C040726F2A657226Ed95712245DeE84b650A1b5"),
	}
}

func MantleRollupGoerliTestnetParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                true,
		AllowDebugPrecompiles:     false,
		DataAvailabilityCommittee: false,
		InitialMtOSVersion:        2,
		InitialChainOwner:         common.HexToAddress("0x186B56023d42B2B4E7616589a5C62EEf5FCa21DD"),
	}
}

func MantleRinkebyTestParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                true,
		AllowDebugPrecompiles:     false,
		DataAvailabilityCommittee: false,
		InitialMtOSVersion:        3,
		InitialChainOwner:         common.HexToAddress("0x06C7DBC804D7BcD881D7b86b667893736b8e0Be2"),
	}
}

func MantleDevTestParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                true,
		AllowDebugPrecompiles:     true,
		DataAvailabilityCommittee: false,
		InitialMtOSVersion:        7,
		InitialChainOwner:         common.Address{},
	}
}

func MantleDevTestDASParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                true,
		AllowDebugPrecompiles:     true,
		DataAvailabilityCommittee: true,
		InitialMtOSVersion:        7,
		InitialChainOwner:         common.Address{},
	}
}

func MantleAnytrustGoerliTestnetParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                true,
		AllowDebugPrecompiles:     false,
		DataAvailabilityCommittee: true,
		InitialMtOSVersion:        2,
		InitialChainOwner:         common.HexToAddress("0x186B56023d42B2B4E7616589a5C62EEf5FCa21DD"),
	}
}

func DisableMantleParams() MantleChainParams {
	return MantleChainParams{
		EnableMtOS:                false,
		AllowDebugPrecompiles:     false,
		DataAvailabilityCommittee: false,
		InitialMtOSVersion:        0,
		InitialChainOwner:         common.Address{},
	}
}

func MantleOneChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:             big.NewInt(42161),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		MantleChainParams:   MantleOneParams(),
		Clique: &CliqueConfig{
			Period: 0,
			Epoch:  0,
		},
	}
}

func MantleNovaChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:             big.NewInt(42170),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		MantleChainParams:   MantleNovaParams(),
		Clique: &CliqueConfig{
			Period: 0,
			Epoch:  0,
		},
	}
}

func MantleRollupGoerliTestnetChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:             big.NewInt(421613),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		MantleChainParams:   MantleRollupGoerliTestnetParams(),
		Clique: &CliqueConfig{
			Period: 0,
			Epoch:  0,
		},
	}
}

func MantleDevTestChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:             big.NewInt(412346),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		MantleChainParams:   MantleDevTestParams(),
		Clique: &CliqueConfig{
			Period: 0,
			Epoch:  0,
		},
	}
}

func MantleDevTestDASChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:             big.NewInt(412347),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		MantleChainParams:   MantleDevTestDASParams(),
		Clique: &CliqueConfig{
			Period: 0,
			Epoch:  0,
		},
	}
}

func MantleAnytrustGoerliTestnetChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:             big.NewInt(421703),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		MantleChainParams:   MantleAnytrustGoerliTestnetParams(),
		Clique: &CliqueConfig{
			Period: 0,
			Epoch:  0,
		},
	}
}

func MantleRinkebyTestnetChainConfig() *ChainConfig {
	return &ChainConfig{
		ChainID:             big.NewInt(421611),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		MantleChainParams:   MantleRinkebyTestParams(),
		Clique: &CliqueConfig{
			Period: 0,
			Epoch:  0,
		},
	}
}

var MantleSupportedChainConfigs = []*ChainConfig{
	MantleOneChainConfig(),
	MantleNovaChainConfig(),
	MantleRollupGoerliTestnetChainConfig(),
	MantleDevTestChainConfig(),
	MantleDevTestDASChainConfig(),
	MantleAnytrustGoerliTestnetChainConfig(),
	MantleRinkebyTestnetChainConfig(),
}
