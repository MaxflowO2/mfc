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
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/boltdb/bolt"
)

// MFCAddress {}
// Struct will be used throughout code
// Will save to DB under Address Basket
type MFCAddress struct {
	MFCxAddy 	string
	MFCxHex		[]byte
	PublicKey 	ed25519.PublicKey
}

// HashKeys(MFCKeys)
// Takes MFCKeys {} and returns []byte Hash
func HashKeys(mfc MFCKeys) []byte {
	pre := sha3.Sum256(mfc.PublicKey)
        addy := pre[:]
	return addy
}

// MakeMFCAddyString()
// Returns MFCx String for address
func MakeMFCAddyString() string {
	mfcx := "MFCx"
	keys := LoadKeys()
	addypre := HashKeys(keys)
	addy := addypre[12:]
	addyString := hex.EncodeToString(addy)
	mfcxaddy := mfcx + addyString
	return mfcxaddy
}

// MakeMFCAddyByte()
// Returns MFCx []Byte for address
func MakeMFCAddyHex() []byte {
	mfcx := "MFCx"
	keys := LoadKeys()
	addypre := HashKeys(keys)
	addy := addypre[12:]
	mfcxhex := append([]byte(mfcx), addy[:]...)
	return mfcxhex
}

// SaveAddress()
// Opens MFCKeys.JSON and makes MFCAddress{}
func SaveAddress() {
	keys := LoadKeys()
	newaddy := MFCAddress{}
	newaddy.MFCxAddy = MakeMFCAddyString()
	newaddy.MFCxHex = MakeMFCAddyHex()
	newaddy.PublicKey = keys.PublicKey
        file, _ := json.MarshalIndent(newaddy, "", " ")
        _ = ioutil.WriteFile("MFCAddress.json", file, 0644)
}

// LoadAddress()
// Opens MFCAddress.JSON and returns MFCAddress{}
func LoadAddress() MFCAddress {
        file, _ := ioutil.ReadFile("MFCAddress.json")
        addy := MFCAddress{}
        _ = json.Unmarshal([]byte(file), &addy)
        return addy
}

// LoadAddressString()
// Opens MFCAddress.JSON and returns String Address
func LoadAddressString() string {
        file, _ := ioutil.ReadFile("MFCAddress.json")
        addy := MFCAddress{}
        _ = json.Unmarshal([]byte(file), &addy)
        return addy.MFCxAddy
}

// LoadAddressHex()
// Opens MFCAddress.JSON and returns Hex Address
func LoadAddressHex() []byte {
	file, _ := ioutil.ReadFile("MFCAddress.json")
	addy := MFCAddress{}
	_ = json.Unmarshal([]byte(file), &addy)
	return addy.MFCxHex
}

// Database functions below

// a.Serialize()
// *MFCAddress to JSON for Bolt.DB
// Returns []byte
func (a *MFCAddress) Serialize() []byte {
        value, err := json.Marshal(a)

	if err != nil {
		fmt.Errorf("Cannot JSON Marshal %v\n", a)
	}

        return value
}

// DeserializeAddy(d []byte)
// JSON to *MFCAddress for Bolt.DB
// Returns MFCAddress
func DeserializeAddy(d []byte) *MFCAddress {
	var addy MFCAddress

	err := json.Unmarshal(d, &addy)

	if err != nil {
		fmt.Errorf("Couldn't Unmarshal %v\n", &addy)
	}

        return &addy
}

// AddAddress(mfc MFCAddress)
// Adds MFCAddress to Bolt.DB
// Method 1: MFCxAddy (string) as Key/Value as Serialize() in StringBucket
// Method 2: MFCxHex ([]byte) as Key/Value as Serialize() in HexBucket
// Logic, will print verbose in string, will hash with []byte
func AddAddress(mfc MFCAddress) {
	db, err := bolt.Open(dbFile, 0644, nil)
		if err != nil {
			log.Fatal(err)
		}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
 		bucket, err := tx.CreateBucketIfNotExists([]byte(addressBucket))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(mfc.MFCxAddy), mfc.Serialize())
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

        err = db.Update(func(tx *bolt.Tx) error {
                bucket, err := tx.CreateBucketIfNotExists([]byte(addressHexBucket))
                if err != nil {
                        return err
                }

                err = bucket.Put(mfc.MFCxHex, mfc.Serialize())
                if err != nil {
                        return err
                }
                return nil
        })

        if err != nil {
                log.Fatal(err)
        }

}

// RetrieveMFCAddress(s string)
// Opens DB, finds address string s in StringBucket
// Returns MFCAddress of user
func RetreiveMFCAddress(s string) *MFCAddress {
	db, err := bolt.Open(dbFile, 0644, nil)
		if err != nil {
			fmt.Errorf("Could not load %s, %v\n", dbFile, err)
		}
	defer db.Close()

	var mfcaddy *MFCAddress
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(addressBucket))
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!\n", addressBucket)
		}

		mfcaddy = DeserializeAddy(bucket.Get([]byte(s)))

		return nil
	})

	if err != nil {
		fmt.Errorf("Did not View(%s), code: %v\n", addressBucket, err)
	}

	return mfcaddy
}

// RetrieveMFCAddressHex(b []byte)
// Opens DB, finds address []byte b in HexBucket
// Returns MFCAddress of user
func RetreiveMFCAddressHex(b []byte) *MFCAddress {
	db, err := bolt.Open(dbFile, 0644, nil)

	if err != nil {
		fmt.Errorf("Could not load %s, %v\n", dbFile, err)
	}

	defer db.Close()

	var mfcaddy *MFCAddress

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(addressHexBucket))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found!\n", addressHexBucket)
		}

        	mfcaddy = DeserializeAddy(bucket.Get(b))

		return nil
	})

	if err != nil {
		fmt.Errorf("Could not View(%s), %v\n", addressHexBucket, err)
	}

	return mfcaddy
}
