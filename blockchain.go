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
	"log"
//	"time"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain{} struct
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// BlockchainIterator{} struct
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// SetTargetBits()
// Update for v0.0.8
//func (bc *Blockchain) SetTargetBits() int {
	// Sets time in seconds per Block
//	var targetTime = 60
	// This number will be modified over time, initally targetBits
//      var newTargetBits = 16
        // Sets length of blocks for PoW Difficulty Adjustment
//    var targetBlocks = 64
        // -1 since we are immediately getting lastBlock of Blockchain
//targetBlocks--
////        bci := bc.Iterator()
//        lastBlock := bci.Next() // sets lastBlock
//        lastDiff := lastBlock.Difficulty // Returns Last Difficulty
//        timeMeow := time.Now().Unix() // Yes a Super Troopers Reference
        // finds the last timestamp of the targetBlock
        // say targetBlock was 10, we only need 9 (see above)
        // if you hit Genesis, code ends
//        var i int
  //      var timeThen int64
    //    for i = 0; i < targetBlocks; i++ {
      //          block := bci.Next()
        //        timeThen = block.Timestamp
          //      if len(block.PrevBlockHash) == 0 {
            //            break
              //  }
//        }
  //      // a is either equal to or less than orginal targetBlocks
    //    targetBlocks++ // now at orginal value
      //  i++ // sets count to proper number of blocks
        //if i < targetBlocks {
		// set as old const targetBits
//                newTargetBits = 16
  //      } else {
    //            // sets time difference
      //          tTime := timeMeow - timeThen
	//	totalTime := int(tTime)
          //      // calculates seconds per block
            //    spb := totalTime/targetBlocks
              //  if spb < targetTime {
                //        newTargetBits = lastDiff + 1
//                }
  //              if spb > targetTime {
    //                    newTargetBits = lastDiff - 1
      //          }
        //}
//	return newTargetBits
//}

// bc.AddBlock(trans *Transaction)
// Opens blockchain.db, pulls lashHash
// Calls NewBlock(trans, lastHash)
// Calls Serialize() and adds to blocksBucket
func (bc *Blockchain) AddBlock(trans *Transaction) {
	var lastHash []byte
	// For Height
	var lastBlock *Block

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		// For Height
		lastBlock = DeserializeBlock(b.Get(lastHash))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	// For Height
	lastHeight := lastBlock.Height
	// For newTargetBits
//	newTarget := bc.SetTargetBits()
	newBlock := NewBlock(trans, lastHash, lastHeight)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

// Iterator()
// Returns BlockchainIterator{} struct
// lastHash or newBlockHash is bc.tip
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

