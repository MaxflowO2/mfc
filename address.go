// (address.go) - Contains all the Address commands in ./mfc
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
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"bytes"
	"encoding/gob"
	"log"
	"encoding/json"
	"io/ioutil"
)

const addressBucket = "address"

// MFCAddress {}
// Struct will be used throughout code
// Will save to DB under Address Basket
type MFCAddress struct {
	MFCxAddy string
	PublicKey ed25519.PublicKey
}

// LoadAddress()
// Opens MFCKeys.JSON and returns []byte Address
// v0.0.8 update to String
func LoadAddress() []byte {
	keys := LoadKeys()
	pre := MakeAddress(keys)
	addy := pre[:]
	return addy
}

// MakeAddress(MFCKeys)
// Takes MFCKeys {} and returns []byte Address
// v0.0.8 update to string
func MakeAddress(mfc MFCKeys) []byte {
	pre := sha3.Sum256(mfc.PublicKey)
        addy := pre[:]
	return addy
}

// SaveAddress()
// Opens MFCKeys.JSON and makes MFCAddress{}
func SaveAddress() {
	mfcx := "MFCx"
	keys := LoadKeys()
	addypre := MakeAddress(keys)
	addy := addypre[12:]
	addyString := hex.EncodeToString(addy)
	mfcxaddy := mfcx + addyString
	newaddy := MFCAddress{}
	newaddy.MFCxAddy = mfcxaddy
	newaddy.PublicKey = keys.PublicKey
        file, _ := json.MarshalIndent(newaddy, "", " ")
        _ = ioutil.WriteFile("MFCAddress.json", file, 0644)
}

// LoadKeys()
// Opens MFCKeys.JSON and returns MFCKeys{}
func LoadAddy() string {
        file, _ := ioutil.ReadFile("MFCAddress.json")
        addy := MFCAddress{}
        _ = json.Unmarshal([]byte(file), &addy)
        return addy.MFCxAddy
}

func LoadMFCAddy() MFCAddress {
        file, _ := ioutil.ReadFile("MFCAddress.json")
        addy := MFCAddress{}
        _ = json.Unmarshal([]byte(file), &addy)
        return addy
}

// Database functions below

// a.Serialized()
// Serialized address for Bolt.DB
// Returns []byte
func (a *MFCAddress) Serialize() []byte {
        var result bytes.Buffer
        encoder := gob.NewEncoder(&result)

        err := encoder.Encode(a)
        if err != nil {
                log.Panic(err)
        }

        return result.Bytes()
}

// DeserializeAddress(d []byte)
// Deserialize an MFCAddress
// Returns *MFCAddress
func DeserializeAddress(d []byte) *MFCAddress {
        var addy MFCAddress

        decoder := gob.NewDecoder(bytes.NewReader(d))
        err := decoder.Decode(&addy)
        if err != nil {
                log.Panic(err)
        }

        return &addy
}

// AddAddress()
// Adds MFCAddress to Bolt.DB
//func AddAddress() {
//	db, err := blot.Open(dbFile, 0600, nil)
//	if err != nil {
//		log.Panic(err)
//	}
//	err = db.Update(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(addressBucket))
