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
//	"golang.org/x/crypto/sha3"
//	"fmt"
	"encoding/json"
	"io/ioutil"
)

// MFCKeys{} - Our Key struct that will be used throughout code
type MFCKeys struct{
	PublicKey ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

// KeyGen()
// Takes nil and passes to ed25519.GenerateKey
// Returns MFCKeys{} of ed25519 circle
func KeyGen() MFCKeys {
	// pass nil to get keys
	pub, priv, _ := ed25519.GenerateKey(nil)

	// keys added to MFCKeys{}
	keys := MFCKeys{
		PublicKey: pub,
		PrivateKey: priv,
	}

	// Add RETURN of keys as MFCKeys{}
	return keys
}

// KeySave(mfc MFCKeys)
// Takes MFCKeys{} and saves to MFCKeys.JSON
func KeySave(mfc MFCKeys) {
	file, _ := json.MarshalIndent(mfc, "", " ")
	_ = ioutil.WriteFile("MFCKeys.json", file, 0644)
}

// LoadKeys()
// Opens MFCKeys.JSON and returns MFCKeys{}
func LoadKeys() MFCKeys {
	file, _ := ioutil.ReadFile("MFCKeys.json")
	keys := MFCKeys{}
	_ = json.Unmarshal([]byte(file), &keys)
	return keys
}
