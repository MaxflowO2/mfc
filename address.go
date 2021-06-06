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
//	"bytes"
//	"encoding/hex"
	"golang.org/x/crypto/sha3"
//	"fmt"
//	"encoding/json"
//	"io/ioutil"
)

// Struct to be saved to database
//type MFCAddress struct {
//	MFCxAddy string
//	Address []byte
//	PublicKey ed25519.PublicKey
//}

// LoadAddress()
// Opens MFCKeys.JSON and returns []byte Address
func LoadAddress() []byte {
	keys := LoadKeys()
	pre := MakeAddress(keys)
	addy := pre[:]
	return addy
}

// MakeAddress(MFCKeys)
// Takes MFCKeys {} and returns []byte Address
func MakeAddress(mfc MFCKeys) []byte {
	pre := sha3.Sum256(mfc.PublicKey)
        addy := pre[:]
	return addy
}

// For saving struct to JSON
//func main() {
//	mfcx := "MFCx"
//      file, _ := ioutil.ReadFile("MFCKeys.json")
//      keys := MFCKeys{}
//      _ = json.Unmarshal([]byte(file), &keys)
//	addy := MakeAddress(keys)
//	addyString := hex.EncodeToString(addy)
//	mfcxAddy := mfcx + addyString
//	newMFCAddress := MFCAddress{}
//	newMFCAddress.MFCxAddy = mfcxAddy
//	newMFCAddress.Address = addy
//	newMFCAddress.PublicKey = keys.PublicKey
//	fmt.Println(newMFCAddress.MFCxAddy)
//        fmt.Println(newMFCAddress.Address)
//       fmt.Println(newMFCAddress.PublicKey)
//}
