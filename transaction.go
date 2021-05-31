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
