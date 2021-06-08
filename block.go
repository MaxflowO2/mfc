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
	"encoding/gob"
	"log"
	"time"
//	"fmt"
)

// Block{} struct
// Used throughout code
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height	      int
	Difficulty    int
	HashBy      string
	Signed	      []byte
}

// b.Serialize()
// Serialized block for Bolt.DB
// Returns []byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// NewBlock(trans *Transaction, prevBlockHash []Byte)
// Adds Timestamp, *Transaction, BlockHash, Hash, Nonce to *Block{}
// Preforms PoW
// Returns *Block{}
func NewBlock(trans []*Transaction, prevBlockHash []byte, prevHeight int) *Block {
	block := &Block{time.Now().Unix(), trans, prevBlockHash, []byte{}, 0, prevHeight, 0, "fix", []byte{}}
	pow := NewProofOfWork(block)
	nonce, hash, diff := pow.Run()

	prevHeight++
	block.Height =prevHeight
	block.Hash = hash[:]
	block.Nonce = nonce
	block.Difficulty = diff
	block.HashBy = LoadAddy()
	block.Signed = Sign(block.Hash)

	return block
}

// NewGenesisBlock()
// Obviously the start
// **Future Update** Genesis struct{} and pass gen Genesis{} to this
// Why? For --alphanet --betanet --mainnet (ect ect) from .JSON file
// Returns *Block{}
func NewGenesisBlock() *Block {
	var genesis []*Transaction
	return NewBlock(genesis, []byte{}, 0)
}

// DeserializeBlock(d []byte)
// Deserialize a block
// Returns *Block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

