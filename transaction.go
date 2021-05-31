package main

import(
	"fmt"
	"time"
)

type Transaction struct {
        Time            int64
	Sender		[]byte
	Reciever	[]byte
	Amount		uint64
	Signature	[]byte
}

// null Transaction for "filling blocks"
func nullTransaction() *Transaction {
	nullTransaction := &Transaction{time.Now().Unix(), []byte("null"), []byte("null"), 0, []byte("null")}

	fmt.Println(nullTransaction)
	// Println success!
	return nullTransaction
}
