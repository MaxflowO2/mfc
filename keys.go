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
	"fmt"
)

type MFCKeys struct{
	PublicKey ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

func KeyGen() {
	pub, priv, _ := ed25519.GenerateKey(nil)

	// now in struct
	a := MFCKeys{
		PublicKey: pub,
		PrivateKey: priv,
	}

//	fmt.Println("Public Key:")
//	fmt.Printf("%x\n", yours.PublicKey)
//	fmt.Printf("%T\n", yours.PublicKey)
//	fmt.Println("Private Key:")
//	fmt.Printf("%x\n", yours.PrivateKey)
//      fmt.Printf("%T\n", yours.PrivateKey)

	// Add RETURN
//	return a
}
