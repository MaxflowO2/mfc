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
	"crypto/ecdsa"
 	"crypto/elliptic"
	"golang.org/x/crypto/sha3"
 	"crypto/rand"
 	"fmt"
 	"hash"
 	"io"
 	"math/big"
 	"os"
)

func keys() {
	// see http://golang.org/pkg/crypto/elliptic/#P256
 	pubkeyCurve := elliptic.P256() 
	// Generates new private Key
 	privatekey := new(ecdsa.PrivateKey)
 	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	var pubkey ecdsa.PublicKey
 	pubkey = privatekey.PublicKey

 	fmt.Println("Private Key :")
 	fmt.Printf("%T\n", privatekey)
	fmt.Println(privatekey)
 	fmt.Println("Public Key :")
 	fmt.Printf("%T \n", pubkey)
	fmt.Println(pubkey)
 	// Sign ecdsa style

 	var h hash.Hash
 	h = sha3.New256()
 	r := big.NewInt(10)
 	s := big.NewInt(0)

 	io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
 	signhash := h.Sum(nil)

 	r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
 	if serr != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	signature := r.Bytes()
 	signature = append(signature, s.Bytes()...)

 	fmt.Printf("Signature : %x\n", signature)

 	// Verify
 	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
 	fmt.Println(verifystatus) // should be true
 }
