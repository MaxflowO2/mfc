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

import(
	"fmt"
        "bytes"
        "time"
        "math"
        "math/big"
	"crypto/ed25519"

        "golang.org/x/crypto/sha3"
)

type Transaction struct {
        Timestamp       int64
	Sender		[]byte
	Receiver	[]byte
	Amount		uint64
	Signature	[]byte
	Hash		[]byte
	Nonce		int
}

type powTransaction struct{
	transaction *Transaction
	target *big.Int
}

var maxTransNonce = math.MaxInt64
var targetTrans = 8

func NewPOWTrans (t *Transaction) *powTransaction {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetTrans))

	powT := &powTransaction{t, target}

	return powT
}

func (powT *powTransaction) prepareTransData(tnonce int) []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(int64(powT.transaction.Timestamp)),
			powT.transaction.Sender,
			powT.transaction.Receiver,
			IntToHex(int64(powT.transaction.Amount)),
			powT.transaction.Signature,
			IntToHex(int64(targetTrans)),
			IntToHex(int64(tnonce)),
		},
		[]byte{},
	)

	return data
}

func (powT *powTransaction) RunTrans() (int, []byte) {
        var hashInt big.Int
        var hash [32]byte
        tnonce := 0

        fmt.Println("Mining the Transaction")
        for tnonce < maxTransNonce {
                data := powT.prepareTransData(tnonce)

                hash =sha3.Sum256(data)
                fmt.Printf("\r%x", hash)
                hashInt.SetBytes(hash[:])

                if hashInt.Cmp(powT.target) == -1 {
                        break
                } else {
                        tnonce++
                }
        }
        fmt.Print("\n")

        return tnonce, hash[:]
}

func (powT *powTransaction) ValidateTrans() bool {
        var hashInt big.Int

        data := powT.prepareTransData(powT.transaction.Nonce)
        hash := sha3.Sum256(data)
        hashInt.SetBytes(hash[:])

        isValid := hashInt.Cmp(powT.target) == -1

        return isValid
}
// null Transaction for "filling blocks"
func nullTransaction() *Transaction {
	a := KeyGen()
	b := KeyGen()
	sender := RandomAddress(a)
	receiver := RandomAddress(b)
	var amount uint64 = 0
	message := bytes.Join(
                [][]byte{
                        sender,
                        receiver,
                        IntToHex(int64(amount)),
                },
                []byte{},
        )
	signed := ed25519.Sign(a.PrivateKey, message)
	nullTransaction := &Transaction{time.Now().Unix(), sender, receiver, amount, signed, []byte{}, 0}

        fmt.Printf("Timestamp: %x\n", nullTransaction.Timestamp)
        fmt.Printf("Sender Address: %x\n", nullTransaction.Sender)
        fmt.Printf("Receiver Address: %x\n", nullTransaction.Receiver)
        fmt.Println("Amount: " , nullTransaction.Amount)
        fmt.Printf("Signtaure:\n%x\n", nullTransaction.Signature)

	powT := NewPOWTrans(nullTransaction)
	nonce, hash := powT.RunTrans()

	nullTransaction.Hash = hash[:]
	nullTransaction.Nonce = nonce

	fmt.Println("Nonce: ", nullTransaction.Nonce)
	// Println success!
	return nullTransaction
}

