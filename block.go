// (block.go) - Contains the Block struct and Block commands in ./mfc
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
	"fmt"
	"math"
	"math/big"
	"time"
	"bytes"
	"io/ioutil"
	"encoding/hex"
	"encoding/json"
	"github.com/MaxflowO2/mfc/K12"
//	"github.com/boltdb/bolt"
)

// Block struct {}
// Used throughout code
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
	Difficulty    int
	HashBy        string
	Signed        []byte
}

// NewBlock(trans *Transaction, prevBlockHash []Byte)
// Adds Timestamp, *Transaction, BlockHash, Hash, Nonce to *Block{}
// Preforms PoW
// Returns *Block{}
func NewBlock(trans []*Transaction, lastHash []byte, lastHeight int, newTarget int) *Block {
	block := &Block{time.Now().Unix(), trans, lastHash, []byte{}, 0, lastHeight, newTarget, LoadAddress(), []byte{}}
	pow := NewProofOfWork(block)
	nonce, hash, diff := pow.Run()

	lastHeight++
	block.Height = lastHeight
	block.Hash = hash
	block.Nonce = nonce
	block.Difficulty = diff
	block.Signed = Sign(block.Hash)

	block.ToFile()

	return block
}

// NewGenesisBlock()
// Obviously the start
// **Future Update** Genesis struct{} and pass gen Genesis{} to this
// Why? For --alphanet --betanet --mainnet (ect ect) from .JSON file
// Returns *Block{}
func NewGenesisBlock() *Block {
	var genesis []*Transaction
	return NewBlock(genesis, []byte{}, 0, 16)
}

func AlphaGenesisBlock() *Block {
	var alphaTrans []*Transaction
	theOne := AlphaGenesis()
	alphaTrans = append(alphaTrans, theOne)
	alpha := &Block{1623289682, alphaTrans, []byte{}, []byte{}, 0, 1, 0, "", []byte{}}
	// New K12.Sum256 values
	alpha.Hash = []byte{0, 0, 131, 168, 228, 219, 228, 184, 223, 179, 126, 55, 55, 36, 55, 171, 23, 131, 204, 236, 181, 229, 18, 188, 113, 30, 105, 184, 71, 38, 246, 130}
	alpha.Nonce = 62078
	alpha.Difficulty = 16

	// Old sha3.Sum256 values
	//alpha.Hash = []byte{0, 0, 91, 237, 75, 239, 186, 156, 203, 254, 5, 66, 134, 202, 179, 200, 24, 123, 177, 62, 127, 223, 166, 39, 79, 139, 178, 237, 146, 253, 100, 214}
	//alpha.Nonce = 55995
	//alpha.Difficulty = 16

	alpha.ToFile()

	return alpha
}

// PoW functions below
//const targetBits = 16

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
	newTargetBits := b.Difficulty
	target.Lsh(target, uint(256-newTargetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// This is a bullshit version of MerkleRoot.
// sliceHash([]*Trans)
// Gets the bytes of all hashes, in []*Transaction
// Returns Sha3.Sum256 of []*Transaction
func (pow *ProofOfWork) sliceHash() []byte {
	max := len(pow.block.Transactions)
	var temp *Transaction
	var toHash []byte
	for i := 0; i < max; i++ {
		temp = pow.block.Transactions[i]
		toHash = bytes.Join(
			[][]byte{
				toHash,
				temp.Hash,
			},
			[]byte{},
		)
		// add a "toblock bool" - add to struct
		// add a "database addition" - Main database
		// add a "database withdrawl" - Mempool
	}
	resultpre := K12.Sum256(toHash)
	result := resultpre[:]
	return result
}

// pow.prepareData(nonce int)
// Joins block data into []byte
// Returns []byte
func (pow *ProofOfWork) prepareData(nonce int, mroot []byte) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			mroot,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(pow.block.Difficulty)),
			IntToHex(int64(nonce)),
			[]byte(pow.block.HashBy),
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
	var hash []byte
	nonce := 0
	mroot := pow.sliceHash()
//	fmt.Printf("Mining the block containing:\n \"%v\"\n", pow.block.Transactions)
	for nonce < maxNonce {

		data := pow.prepareData(nonce, mroot)

		hash = K12.Sum256(data)
//		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash)

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("Block #%v mined\n\n", pow.block.Height)
	return nonce, hash[:], pow.block.Difficulty
}

// pow.Validate()
// Validates Hash
// Returns bool
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce, pow.sliceHash())
	hash := K12.Sum256(data)
	hashInt.SetBytes(hash)

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

// Bolt.DB functions

// b.ToFile()
// Saves block to alpha/block as (Hash).block
func (b *Block) ToFile() {
        header := "alpha/block/"
        dotblock := ".block"
        filename := header + hex.EncodeToString(b.Hash) + dotblock
        file, err := json.MarshalIndent(b, "", " ")
        if err != nil {
                fmt.Errorf("%s did not Marshal\n", filename)
        }
        err = ioutil.WriteFile(filename, file, 0644)
        if err != nil {
                fmt.Errorf("%s did not save\n", filename)
        }
}

// b.Serialize()
// Serialized block for Bolt.DB
// Returns []byte
func (b *Block) Serialize() []byte {
        value, err := json.Marshal(b)
        if err != nil {
                fmt.Errorf("%v did not Marshal\n", b)
        }
        return value
}

// DeserializeBlock(d []byte)
// Deserialize a block
// Returns *Block
func DeserializeBlock(d []byte) *Block {
	var block Block
	err := json.Unmarshal(d, &block)
	if err != nil {
		fmt.Errorf("Could not Unmarshal\n")
	}
	return &block
}
