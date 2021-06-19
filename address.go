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
	"fmt"
	"io/ioutil"
	"encoding/hex"
	"encoding/json"
	"crypto/ed25519"
	"github.com/MaxflowO2/mfc/K12"
//	"log"
	"github.com/boltdb/bolt"
)

// MFCAddress struct {}
// Struct will be used throughout code
// Will save to DB under Address Basket
type MFCAddress struct {
	MFCxAddy	string
	PublicKey	ed25519.PublicKey
	// New in v0.0.10+ Name is for contracts, State is your balance
	Name 		string
	State 		[]*Balance
}

// HashKeys(MFCKeys)
// Takes MFCKeys {}
// Returns []byte Hash
func HashKeys(mfc MFCKeys) []byte {
	addy := K12.Sum256(mfc.PublicKey)
	return addy
}

// MakeAddress()
// Makes MFCAddress{}, and calls .ToFile() for saving
func MakeAddress() *MFCAddress {
	var value MFCAddress
	mfcx := "MFCx"
	keys := LoadKeys()
	addy := HashKeys(keys)
	addyString := hex.EncodeToString(addy)
	value.MFCxAddy = mfcx + addyString
	value.PublicKey = keys.PublicKey
	value.ToFile()
	return &value
}

// a.ToFile()
// Saves MFCAddress to MFCAddress.JSON
func (a *MFCAddress) ToFile() {
	file, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		fmt.Errorf("Could not Marshal Indent %v!\n", a)
	}
	err = ioutil.WriteFile("MFCAddress.json", file, 0644)
	if err != nil {
		fmt.Errorf("Could not save MFCAddress.json\n")
	}
}

// AddyFromFile()
// Opens MFCAddress.JSON and returns MFCAddress{}
func AddyFromFile() MFCAddress {
	var addy MFCAddress
	file, err := ioutil.ReadFile("MFCAddress.json")
	if err != nil {
		fmt.Errorf("Could not load MFCAddress.json\n")
	}
	err = json.Unmarshal([]byte(file), &addy)
	if err != nil {
		fmt.Errorf("Could not Unmarshal MFCAddress.json\n")
	}
	return addy
}

// LoadAddress()
// Opens MFCAddress.JSON and returns String Address
func LoadAddress() string {
	file, _ := ioutil.ReadFile("MFCAddress.json")
	addy := MFCAddress{}
	_ = json.Unmarshal([]byte(file), &addy)
	return addy.MFCxAddy
}

// Bolt.DB functions below

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

var addressBucket = "Address"

// AddAddress(mfc MFCAddress)
// Adds MFCAddress to Bolt.DB
// Method: Key/Value as []byte(MFCxAddy)/Serialize() in Bucket "Address"
func (mfc *MFCAddress) ToBoltDB() {
	db, err := bolt.Open(dbFile, 0644, nil)
		if err != nil {
			fmt.Errorf("Could not load database file, %v\n", err)
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
		fmt.Errorf("Bucket error: %v\n", err)
	}
}

// RetrieveMFCAddress(s string)
// Opens DB, finds address string s in StringBucket
// Returns MFCAddress of user
func AddyFromBoltDB(s string) *MFCAddress {
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
