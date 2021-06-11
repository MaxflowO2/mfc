// (main.go) - Must have file for ./mfc
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
//	"time"
//	"fmt"
	"os"
)

// fileExists(path string)
// File check function
// Returns err
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {
	// checks for MFCKeys and makes it if not there
	var mfckeys string = "MFCKeys.json"
	keysExist := fileExists(mfckeys)

	if keysExist {
//		fmt.Println("MFCKeys.json found...")
//		keys := LoadKeys()
//		fmt.Printf("MFC Public Keys is:\n%x\n", keys.PublicKey)
//		fmt.Printf("MFC Private Key is:\n%x\n", keys.PrivateKey)
//		fmt.Printf("DO NOT HAND OUT YOUR PRIVATE KEY!\n\n")
	} else {

//		fmt.Println("MFCKeys.json was not found, generating Key Set...")
		newKeys := KeyGen()
//		fmt.Printf("MFC Public Keys is:\n%x\n", newKeys.PublicKey)
//		fmt.Printf("MFC Private Key is:\n%x\n", newKeys.PrivateKey)
//		fmt.Println("DO NOT HAND OUT YOUR PRIVATE KEY!")
//		fmt.Printf("Saving keys to MFCKeys.json!\n\n")
		KeySave(newKeys)
	}

	// checks for MFCAddress and makes it if not there
	var mfcaddy string = "MFCAddress.json"
	addyExist := fileExists(mfcaddy)

	if addyExist {
//		fmt.Println("MFCAddress.json found...")
//		addy := LoadAddress()
//		fmt.Printf("MFC Address is:\n%s\n\n", addy)
	} else {
//		fmt.Println("MFCAddress.json was not found, generating Address.")
		SaveAddress()
//		addy := LoadAddress()
//		fmt.Printf("MFC Address is:\n%s\n\n", addy)
	}

	// Sends MFCAddress to mfc.db, crutial step - passed
//	addy := LoadAddress()
//	AddAddress(addy)
//	fmt.Println("added to database")
//	testOne := RetreiveMFCAddress(addy.MFCxAddy)
//	fmt.Printf("from one: %s\n", testOne.MFCxAddy)
//	testTwo := RetreiveMFCAddressHex(addy.MFCxHex)
//	fmt.Printf("from two: %s\n", testTwo.MFCxAddy)
//	fmt.Printf("Asked from this: %s\n", addy.MFCxAddy)

//	startDB()
//	AlphaGenesisBlock()
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
