// (balance.go) - Contains all the Address.State commands in ./mfc
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
//        "io/ioutil"
//        "encoding/hex"
//        "encoding/json"
//      "github.com/boltdb/bolt"
)

type Balance struct {
	Type	string
	Name	string
	Amount	float64
}

func GetBal(u string, t string) float64 {
	var value *MFCAddress
	// retreive MFCAddress 'u' from Bolt.DB
	value = AddyFromBoltDB(u)
	max := len(value.State)
	var number float64
	for i := 0; i < max ; i++ {
		if value.State[i].Type == t {
			number = value.State[i].Amount
		} else {
			number = 0
		}
	}
	return number
}

func AdjustBal(u string, t string, v float64) {
	var value *MFCAddress
	// retreive MFCAddress of 'u' from Bolt.DB
	value = AddyFromBoltDB(u)
        // to read all of state and change it
	max := len(value.State)
	for i := 0 ; i <= max ; i++ {
		// add or subtract if you have it
		if value.State[i].Type == t {
			original := value.State[i].Amount
			value.State[i].Amount = original + v // use negatives for withdrawl
		// add value if none exists, positive v only
		} else if i == max && v > 0 && value.State[i].Type != t {
			add := &Balance{t, GetName(t), v}
			value.State = append(value.State, add)
		// this is the error, either variant should have occured
		} else {
			fmt.Printf("Coding error\n")
		}
	}
	// now state is changed
	// update database aka write to Bolt.DB
	value.ToBoltDB()
}

func GetName(t string) string {
	data := AddyFromBoltDB(t)
	return data.Name
}

