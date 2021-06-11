// (db.go) - Sets up the BoltBD "mfc.db" all bolt commands for ./mfc
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
	"github.com/boltdb/bolt"
)

const dbFile = "mfc.db"
const blocksBucket = "blocks"
const addressBucket = "address"
const addressHexBucket = "addresshex"
const transactionBucket = "transactions"
const transInBlock = "transinblock"

// setupBD()
// Initializes mfc.db
// Sets up Block, Block{} struct stored as JSON
// Sets up Block.Transaction, Transaction{} within the block (to find Transactions within a block)
// Sets up Transaction, Transaction{} "Mempool edition"
// Sets up Address (string) and Address ([]Byte), MFCAddress{}
// Returns db and error
func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not open BoltDB, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		blocks, err := tx.CreateBucketIfNotExists([]byte(blocksBucket))
		if err != nil {
			return fmt.Errorf("Could not make Block Bucket: %v", err)
		}
                _, err = blocks.CreateBucketIfNotExists([]byte(transInBlock))
                if err != nil {
                        return fmt.Errorf("Could not make Transactions in Block Bucket: %v", err)
                }
                _, err = tx.CreateBucketIfNotExists([]byte(transactionBucket))
                if err != nil {
                        return fmt.Errorf("Could not make Transaction Bucket: %v", err)
                }
                _, err = tx.CreateBucketIfNotExists([]byte(addressBucket))
                if err != nil {
                        return fmt.Errorf("Could not make Address by String Bucket: %v", err)
                }
                _, err = tx.CreateBucketIfNotExists([]byte(addressHexBucket))
                if err != nil {
                        return fmt.Errorf("Could not make Address by Hex Bucket: %v", err)
                }
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Could not set up buckets, %v", err)
	}
	fmt.Println("mfc.db is now setup")
	return db, nil
}

// startDB()
// Basically opens and closes mfc.db
func startDB() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
