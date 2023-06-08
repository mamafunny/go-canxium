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

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type MiningTx struct {
	ChainID   *big.Int
	Nonce     uint64   // sender nonce
	GasTipCap *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap *big.Int // a.k.a. maxFeePerGas
	Gas       uint64
	To        common.Address // mining reward receiver
	Value     *big.Int       // value should equal difficulty * consensus reward per difficulty hash
	Data      []byte

	// mining fields
	Algorithm  uint8
	Difficulty *big.Int
	MixDigest  common.Hash
	Seed       uint64 // mining nonce

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *MiningTx) copy() TxData {
	cpy := &MiningTx{
		Nonce: tx.Nonce,
		To:    tx.To,
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		Value:     new(big.Int),
		ChainID:   new(big.Int),
		GasTipCap: new(big.Int),
		GasFeeCap: new(big.Int),
		// mining fields
		Algorithm:  tx.Algorithm,
		Difficulty: new(big.Int),
		Seed:       tx.Seed,
		MixDigest:  tx.MixDigest,
		// signature
		V: new(big.Int),
		R: new(big.Int),
		S: new(big.Int),
	}

	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasTipCap != nil {
		cpy.GasTipCap.Set(tx.GasTipCap)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.Difficulty != nil {
		cpy.Difficulty.Set(tx.Difficulty)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *MiningTx) txType() byte           { return MiningTxType }
func (tx *MiningTx) chainID() *big.Int      { return tx.ChainID }
func (tx *MiningTx) accessList() AccessList { return nil }
func (tx *MiningTx) data() []byte           { return tx.Data }
func (tx *MiningTx) gas() uint64            { return tx.Gas }
func (tx *MiningTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *MiningTx) gasTipCap() *big.Int    { return tx.GasTipCap }
func (tx *MiningTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *MiningTx) value() *big.Int        { return tx.Value }
func (tx *MiningTx) nonce() uint64          { return tx.Nonce }
func (tx *MiningTx) to() *common.Address    { return &tx.To }

// mining fields
func (tx *MiningTx) algorithm() byte        { return tx.Algorithm }
func (tx *MiningTx) difficulty() *big.Int   { return tx.Difficulty }
func (tx *MiningTx) seed() uint64           { return tx.Seed }
func (tx *MiningTx) mixDigest() common.Hash { return tx.MixDigest }
func (tx *MiningTx) setPow(seed uint64, mixDigest common.Hash) {
	tx.Seed = seed
	tx.MixDigest = mixDigest
}

func (tx *MiningTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	if baseFee == nil {
		return dst.Set(tx.GasFeeCap)
	}
	tip := dst.Sub(tx.GasFeeCap, baseFee)
	if tip.Cmp(tx.GasTipCap) > 0 {
		tip.Set(tx.GasTipCap)
	}
	return tip.Add(tip, baseFee)
}

func (tx *MiningTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *MiningTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}
