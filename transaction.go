// (transaction.go) - Contains all the Transaction commans in ./mfc
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
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"io/ioutil"
	"math"
	"math/big"
	"time"
//	"github.com/boltdb/bolt"
	"github.com/MaxflowO2/mfc/K12"
)

// Transaction{} struct
// Basic state change of blockchain/db
// Used throughout as *Transaction
// Remove in v0.0.10
type Transaction struct {
	Timestamp int64
	Sender    string
	Receiver  string
	Amount    uint64
	Message   string
	Signature []byte
	Hash      []byte
	Nonce     int
}

// Mempool struct
// Used to create transactions within ./mfc
// Will be transmitted on ./mfc p2p
// v0.0.10
type Mempool struct {
	Timestamp	int64
	Sender		string
	Receiver	string
	Amount		uint64
	Type		string
	Name		string
	//ATPCost	uint
	Signature	[]byte
}

// MinedMempool struct
// Is a mined Mempool struct using PoW functions below
// Will be inserted into a block as a string
// v0.0.10
type MinedMempool struct {
	mempool		Mempool
	Timestamp	int64
	IsValid		bool
	Rewards		string
	//ATPUsed	uint
	Nonce		int
	Hash		[]byte
	Signature	[]byte
}

// powMempool struct
// Used as a struct for mining
// v0.0.10
type powMempool struct {
	toMine		*MinedMempool
	target		*big.Int
}

// powTransaction {} struct
// Used for powT functions below
// Remove for v0.0.10
type powTransaction struct {
	transaction *Transaction
	target      *big.Int
}

// Set Variables for powT functions below
var maxTransNonce = math.MaxInt64
var targetTrans = 8

// newPowMempool
// Forced hash requiring 00 at start
// Returns *powMempool with new target and Mempool Struct
// v0.0.10
func newPowMempool(m *MinedMempool) *powMempool {
	target := big.NewInt(1)
	target.Lsh(target, uint(248)) // forced two 0's

	pow := &powMempool{m, target}

	return pow
}

// NewPOWTrans(transaction *Transaction)
// sets Transaction.target of powTransaction{} from targetTrans
// Returns *powTransaction{}
// Removed in v0.0.10
func NewPOWTrans(t *Transaction) *powTransaction {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetTrans))

	powT := &powTransaction{t, target}

	return powT
}

// pow.prepMempool(nonce)
// Joins all elements of MinedMempool, to make a []byte
// Returns []byte
// v0.0.10
func (pow *powMempool) prepMempool(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(pow.toMine.mempool.Timestamp),
			[]byte(pow.toMine.mempool.Sender),
			[]byte(pow.toMine.mempool.Receiver),
			IntToHex(int64(pow.toMine.mempool.Amount)),
			[]byte(pow.toMine.mempool.Type),
			[]byte(pow.toMine.mempool.Name),
			//IntToHex(int64(pow.toMine.mempool.ATPCost)),
			pow.toMine.mempool.Signature,
			IntToHex(pow.toMine.Timestamp),
			[]byte(strconv.FormatBool(pow.toMine.IsValid)),
			[]byte(pow.toMine.Rewards),
			//IntToHex(int64(pow.toMine.ATPUsed)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// powT.prepareTransData(tnonce)
// Joins all elements of *Transaction into bytes for hash
// Returns []byte
// deleting
func (powT *powTransaction) prepareTransData(tnonce int) []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(int64(powT.transaction.Timestamp)),
			// update Sender/Reciever to string v0.0.8
			[]byte(powT.transaction.Sender),
			[]byte(powT.transaction.Receiver),
			IntToHex(int64(powT.transaction.Amount)),
			[]byte(powT.transaction.Message),
			powT.transaction.Signature,
			IntToHex(int64(targetTrans)),
			IntToHex(int64(tnonce)),
		},
		[]byte{},
	)

	return data
}

// powT.RunTrans()
// K12.Sum256 Hash of Transaction data
// Returns Nonce, Hash
func (powT *powTransaction) RunTrans() (int, []byte) {
	var hashInt big.Int
	var hash []byte
	tnonce := 0

//	fmt.Println("Mining the Transaction")
	for tnonce < maxTransNonce {
		data := powT.prepareTransData(tnonce)

		hash = K12.Sum256(data)
//		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash)

		if hashInt.Cmp(powT.target) == -1 {
			break
		} else {
			tnonce++
		}
	}
//	fmt.Print("\n")

	return tnonce, hash
}

// powT.ValidateTrans()
// part of powT *Transaction
// Returns bool
func (powT *powTransaction) ValidateTrans() bool {
	var hashInt big.Int

	data := powT.prepareTransData(powT.transaction.Nonce)
	hash := K12.Sum256(data)
	hashInt.SetBytes(hash)

	isValid := hashInt.Cmp(powT.target) == -1

	return isValid
}

// SliceTransaction(t *Transaction)
// Add to slice of []*Transaction
// Don't know if i'll continue to use
// Returs []*Transaction
func SliceTransaction(t *Transaction, st []*Transaction) []*Transaction {
	st = append(st, t)
	return st
}

// null Transaction for "filling blocks"
// testing only
// will need updates on v0.0.8 (string for sender/receiver)
func bsTransaction() *Transaction {
	a := KeyGen()
	b := KeyGen()
	sender := "MFCx" + hex.EncodeToString(HashKeys(a))
	receiver := "MFCx" + hex.EncodeToString(HashKeys(b))
	var amount uint64 = 0
	message := bytes.Join(
		[][]byte{
			[]byte(sender),
			[]byte(receiver),
			IntToHex(int64(amount)),
		},
		[]byte{},
	)
	signed := ed25519.Sign(a.PrivateKey, message)
	var signature []byte
	verify := ed25519.Verify(a.PublicKey, message, signed)
	if verify == true {
		signature = signed
	}
	bs := &Transaction{time.Now().Unix(), sender, receiver, amount, "", signature, []byte{}, 0}

//	fmt.Println("START OF TRANSACTION")
//	fmt.Printf("Timestamp: %x\n", bs.Timestamp)
//	fmt.Printf("Sender Address: %x\n", bs.Sender)
//	fmt.Printf("Receiver Address: %x\n", bs.Receiver)
//	fmt.Println("Amount: ", bs.Amount)
//	fmt.Printf("Signtaure:\n%x\n", bs.Signature)

	powT := NewPOWTrans(bs)
	nonce, hash := powT.RunTrans()

	bs.Hash = hash
	bs.Nonce = nonce

//	fmt.Println("Nonce: ", bs.Nonce)
//	fmt.Println("END OF TRANSACTION")

	return bs
}

// AlphaNet Genesis use
func AlphaGenesis() *Transaction {
	var alpha *Transaction
	alpha = &Transaction{1623289682, "", "", 0, "AlphaNet of MaxFlowChain, created for testing purposes on 6/9/2021, www.nytimes.com/2021/06/09/technology/bitcoin-untraceable-pipeline-ransomware.html issues 101", []byte{}, []byte{}, 0}

	// New K12.Sum256 values
	alpha.Hash = []byte{0, 72, 127, 147, 61, 30, 56, 31, 247, 87, 166, 17, 121, 200, 25, 195, 11, 41, 49, 181, 182, 216, 30, 13, 203, 207, 161, 213, 112, 34, 212, 108}
	alpha.Nonce = 5

	// Old sha3.Sum256 values
	//alpha.Hash = []byte{0, 193, 197, 91, 204, 202, 150, 0, 152, 178, 150, 35, 108, 152, 68, 106, 19, 114, 152, 94, 9, 131, 80, 44, 246, 98, 103, 106, 207, 218, 75, 96}
	//alpha.Nonce = 314

	header := "alpha/trans/"
	dotblock := ".trans"
	filename := header + hex.EncodeToString(alpha.Hash) + dotblock
	file, _ := json.MarshalIndent(alpha, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)

	//        data, err := json.Marshal(alpha)
	//        if err != nil {
	//                fmt.Errorf("Couldn't Marshal AlphaNet Genesis Transaction, %v", err)
	//        	}
	//
	//        db, err := setupDB()
	//        if err != nil {
	//                fmt.Errorf("Couldn't open mfc.db, %v", err)
	//        	}
	//        defer db.Close()

	//        err = db.Update(func (tx *bolt.Tx) error {
	//                err := tx.Bucket([]byte(transactionBucket)).Put(alpha.Hash, data)
	//                if err != nil {
	//                        return fmt.Errorf("Alpha Genesis did not insert into transactionBucket, code: %v", err)
	//                	}
	//                return nil
	//        })
	//
	return alpha
}

// t.SerializeTrans()
// *Transaction to JSON for Bolt.DB
// Returns []byte
func (t *Transaction) Serialize() []byte {
	value, err := json.Marshal(t)

	if err != nil {
		fmt.Errorf("%v, couldn't Marshal\n", t)
	}

	return value
}

// DeserializeTrans(d []byte)
// JSON to *Transacation for Bolt.DB
// Returns *Transaction
func DeserializeTrans(d []byte) *Transaction {
	var trans Transaction

	err := json.Unmarshal(d, &trans)

	if err != nil {
		fmt.Errorf("%v, couldn't Unmarshal\n", &trans)
	}

	return &trans
}
