// (pow.go) - Contains the pow functionality of Block in ./mfc
// Copyright (C) 2021 MaxflowO2, the only author of Max Flow Chain
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"

	"golang.org/x/crypto/sha3"
)

const targetBits = 16

// Variable set throughout pow.go
var (
	maxNonce = math.MaxInt64
)

// ProofOfWork {} struct
// Used for pow functions below
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork(b *Block)
// Uses targetBits to set difficulty
// Returns *ProofOfWork{}
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// pow.prepareData(nonce int)
// Joins block data into []byte
// Returns []byte
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
//			pow.block.Transactions.Hash,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// pow.Run()
// Preforms Sha3.Sum256 hash of block data
// Returns Nonce, Hash
func (pow *ProofOfWork) Run() (int, []byte, int) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	
	fmt.Printf("Mining the block containing:\n \"%v\"\n", pow.block.Transactions)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash =sha3.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n")
	return nonce, hash[:], targetBits
}

// pow.Validate()
// Validates Hash
// Returns bool
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha3.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

