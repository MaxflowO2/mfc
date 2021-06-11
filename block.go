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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	//	"github.com/boltdb/bolt"
)

// Block{} struct
// Used throughout code
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
	Difficulty    int
	HashBy        []byte
	Signed        []byte
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

// NewBlock(trans *Transaction, prevBlockHash []Byte)
// Adds Timestamp, *Transaction, BlockHash, Hash, Nonce to *Block{}
// Preforms PoW
// Returns *Block{}
func NewBlock(trans []*Transaction, prevBlockHash []byte, prevHeight int) *Block {
	block := &Block{time.Now().Unix(), trans, prevBlockHash, []byte{}, 0, prevHeight, 0, LoadAddressHex(), []byte{}}
	pow := NewProofOfWork(block)
	nonce, hash, diff := pow.Run()

	prevHeight++
	block.Height = prevHeight
	block.Hash = hash[:]
	block.Nonce = nonce
	block.Difficulty = diff
	block.Signed = Sign(block.Hash)

	header := "alpha/block/"
	dotblock := ".block"
	filename := header + hex.EncodeToString(block.Hash) + dotblock
	file, err := json.MarshalIndent(block, "", " ")
	if err != nil {
		fmt.Errorf("No Marshal\n")
	}
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		fmt.Errorf("No file\n")
	}

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

func AlphaGenesisBlock() *Block {
	var alphaTrans []*Transaction
	theOne := AlphaGenesis()
	alphaTrans = append(alphaTrans, theOne)
	alpha := &Block{1623289682, alphaTrans, []byte{}, []byte{}, 0, 1, 0, []byte{}, []byte{}}
	// been hashed, values below are correct
	alpha.Hash = []byte{0, 0, 91, 237, 75, 239, 186, 156, 203, 254, 5, 66, 134, 202, 179, 200, 24, 123, 177, 62, 127, 223, 166, 39, 79, 139, 178, 237, 146, 253, 100, 214}
	alpha.Nonce = 55995
	alpha.Difficulty = 16

	header := "alpha/block/"
	dotblock := ".block"
	filename := header + hex.EncodeToString(alpha.Hash) + dotblock
	file, err := json.MarshalIndent(alpha, "", " ")
	if err != nil {
		fmt.Errorf("No Marshal\n")
	}
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		fmt.Errorf("No file\n")
	}

	return alpha
}

//	data, err := json.Marshal(alpha)
//	if err != nil {
//		fmt.Errorf("Couldn't Marshal AlphaNet Genesis Block, %v\n", err)
//	}

//	transdata, err := json.Marshal(theOne)
//	if err != nil{
//		fmt.Errorf("Couldn't Marshal the One Transaction in AlphaNet Genesis Block, %v\n", err)
//	}

//	db, err := setupDB()
//	if err != nil {
//		fmt.Errorf("Couldn't open mfc.db, %v\n", err)
//	}
//	defer db.Close()

//	err = db.Update(func (tx *bolt.Tx) error {
//		err := tx.Bucket([]byte(blocksBucket)).Put(alpha.Hash, alpha.Serialize())
//		if err != nil {
//			return fmt.Errorf("Alpha Genesis did not insert into mfc.db blocksBucket, code: %v\n", err)
//			}

//		err = tx.Bucket([]byte(blocksBucket)).Put([]byte("1"), alpha.Hash)
//		if err != nil {
//			return fmt.Errorf("Pointer Key not inserted at byte '1', code: %v\n", err)
//			}

//		err = tx.Bucket([]byte(blocksBucket)).Bucket([]byte(transInBlock)).Put(theOne.Hash, theOne.Serialize())
//		if err != nil {
//			return fmt.Errorf("Alpha Genesis Transaction did not insert into mfc.db blocksBucket.transInBlock, code: %v\n", err)
//			} else {
//			tx.Bucket([]byte(transactionBucket)).Delete(theOne.Hash)
//			}
//		return nil

//	})

//	return alpha
//}

// DeserializeBlock(d []byte)
// Deserialize a block
// Returns *Block
func DeserializeBlock(d []byte) *Block {
	var block Block

	err := json.Unmarshal(d, &block)

	if err != nil {
		fmt.Errorf("%v, couldn't Unmarshal\n", &block)
	}

	return &block
}
