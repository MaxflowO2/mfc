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
	"crypto/ed25519"
//	"encoding/hex"
	"golang.org/x/crypto/sha3"
//	"fmt"
	"encoding/json"
	"io/ioutil"
)

type MFCKeys struct{
	PublicKey ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

type Address []byte

func KeyGen() MFCKeys {
	pub, priv, _ := ed25519.GenerateKey(nil)

	// now in struct
	a := MFCKeys{
		PublicKey: pub,
		PrivateKey: priv,
	}

	// Add RETURN of MFCKeys
	return a
}

// Saves Keys to MFCKeys.json
func KeySave(mfc MFCKeys) {
	file, _ := json.MarshalIndent(mfc, "", " ")
	_ = ioutil.WriteFile("MFCKeys.json", file, 0644)
}

// Loads Keys from MFCKeys.json
func LoadKey() MFCKeys {
	file, _ := ioutil.ReadFile("MFCKeys.json")
	keys := MFCKeys{}
	_ = json.Unmarshal([]byte(file), &keys)
	return keys
}

// Generate an address
func LoadAddress() Address {
	// loads Keys from MFCKeys.json
	keys := LoadKey()
	// sha3/Sum256 hash of Public Key
	pkhash := sha3.Sum256(keys.PublicKey)
	// slice of last 20
	slice := pkhash[12:]
	return slice
}

// For Make Address w/o loading keys
func MakeAddress(mfc MFCKeys) Address {
	pkhash := sha3.Sum256(mfc.PublicKey)
	slice := pkhash[12:]
	return slice
}

//func main() {
//        loaded := LoadKey()
//        fmt.Println("Keys loaded from MFCKeys.json")
//        fmt.Println("Public Key:")
//        fmt.Printf("%x\n", loaded.PublicKey)
//        fmt.Printf("%T\n", loaded.PublicKey)
//        fmt.Println("Private Key:")
//        fmt.Printf("%x\n", loaded.PrivateKey)
//        fmt.Printf("%T\n", loaded.PrivateKey)
//	addy := LoadAddress()
//	fmt.Println("Address:")
//	fmt.Printf("%x\n" , addy)
//	a := KeyGen()
//	b := KeyGen()
//	aaddy := RandomAddress(a)
//	baddy := RandomAddress(b)
//	fmt.Println("a, priv/pub/addy")
//	fmt.Printf("%x\n", a.PublicKey)
//	fmt.Printf("%x\n", a.PrivateKey)
//	fmt.Printf("%x\n", aaddy)
//        fmt.Println("b, priv/pub/addy")
//        fmt.Printf("%x\n", b.PublicKey)
//        fmt.Printf("%x\n", b.PrivateKey)
//        fmt.Printf("%x\n", baddy)
//}//
