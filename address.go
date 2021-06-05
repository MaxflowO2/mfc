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
//      "encoding/hex"
        "golang.org/x/crypto/sha3"
//        "fmt"
        "encoding/json"
        "io/ioutil"
)

// makes an address type
type Address []byte

// New Address genesis
func LoadAddress() Address {
        file, _ := ioutil.ReadFile("MFCKeys.json")
        keys := MFCKeys{}
        _ = json.Unmarshal([]byte(file), &keys)
        pkhash := sha3.Sum256(keys.PublicKey)
	slice := pkhash[12:]
	mfcx := []byte("MFCx")
	address :=  bytes.Join([][]byte{mfcx, slice},[]byte{})
	return address
}

// For Make Address w/o loading keys
func MakeAddress(mfc MFCKeys) Address {
        pkhash := sha3.Sum256(mfc.PublicKey)
        slice := pkhash[12:]
        mfcx := []byte("MFCx")
        address :=  bytes.Join([][]byte{mfcx, slice},[]byte{})
        return address
}

