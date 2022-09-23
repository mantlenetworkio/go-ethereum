// Copyright 2021 The go-ethereum Authors
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

package logger

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

func (*AccessListTracer) CaptureMantleTransfer(env *vm.EVM, from, to *common.Address, value *big.Int, before bool, purpose string) {
}
func (*JSONLogger) CaptureMantleTransfer(env *vm.EVM, from, to *common.Address, value *big.Int, before bool, purpose string) {
}
func (*StructLogger) CaptureMantleTransfer(env *vm.EVM, from, to *common.Address, value *big.Int, before bool, purpose string) {
}
func (*mdLogger) CaptureMantleTransfer(env *vm.EVM, from, to *common.Address, amount *big.Int, before bool, purpose string) {
}

func (*AccessListTracer) CaptureMantleStorageGet(key common.Hash, depth int, before bool) {}
func (*JSONLogger) CaptureMantleStorageGet(key common.Hash, depth int, before bool)       {}
func (*StructLogger) CaptureMantleStorageGet(key common.Hash, depth int, before bool)     {}
func (*mdLogger) CaptureMantleStorageGet(key common.Hash, depth int, before bool)         {}

func (*AccessListTracer) CaptureMantleStorageSet(key, value common.Hash, depth int, before bool) {}
func (*JSONLogger) CaptureMantleStorageSet(key, value common.Hash, depth int, before bool)       {}
func (*StructLogger) CaptureMantleStorageSet(key, value common.Hash, depth int, before bool)     {}
func (*mdLogger) CaptureMantleStorageSet(key, value common.Hash, depth int, before bool)         {}
