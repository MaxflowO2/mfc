// (cli.go) - Contains the Command Line Interface of ./mfc
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
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
//	"encoding/hex"
)

// CLI responsible for processing command line arguments
type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  visatest - 15 seconds worth of Visa Transactions")
//	fmt.Println("  addtodb - sends your Address to BoltDB")
	fmt.Println("  autogen - creates a massive block every 60 seconds with nulltrans")
	fmt.Println("  bstrans, creates a 'bs-transaction'")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock() {
	var sendit []*Transaction
	// Visa's 1,734 per second * 15 seconds
	for i := 0; i < 26010; i++ {
		fill := bsTransaction()
		sendit = append(sendit, fill)
	}
	cli.bc.AddBlock(sendit)
	fmt.Println("Success!")
}

func (cli *CLI) autoGen(t time.Time) {
	var sendit []*Transaction
	// Visa's 1,734 per second * 15 seconds
	for i := 0; i < 100; i++ {
		fill := bsTransaction()
		sendit = append(sendit, fill)
	}
	cli.bc.AddBlock(sendit)
}

//func (cli *CLI) addToDB() {
//        addy := LoadAddress()
//        AddAddress(addy)
//}

func (cli *CLI) bsTrans() {
	bsTransaction()
}

func (cli *CLI) printChain() {
	// throwing errors
	bci := cli.bc.Iterator()
	//	var addy *MFCAddress
	for {
		block := bci.Next()
		fmt.Printf("Block Height: %v\n", block.Height)
		fmt.Printf("Previous Hash:\n%x\n", block.PrevBlockHash)
//		fmt.Printf("Transactions in Block:\n%v\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Difficulty: %v\n", block.Difficulty)
		//hangs
		//		addy = RetreiveMFCAddressHex(block.HashBy)
		fmt.Printf("Hashed By: %x\n", block.HashBy)
		fmt.Printf("Signature: %x\n", block.Signed)
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("visatest", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	//	addToDB := flag.NewFlagSet("addtodb", flag.ExitOnError)
	autoGen := flag.NewFlagSet("autogen", flag.ExitOnError)
	bsTrans := flag.NewFlagSet("bstrans", flag.ExitOnError)
	//	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "visatest":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "autogen":
		err := autoGen.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

		//	case "addtodb":
		//		err := addToDB.Parse(os.Args[2:])
		//		if err != nil {
		//			log.Panic(err)
		//		}

	case "bstrans":
		err := bsTrans.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		cli.addBlock()
	}

	if autoGen.Parsed() {
		Repeat(15*time.Second, cli.autoGen)
	}

	//	if addToDB.Parsed() {
	//		cli.addToDB()
	//	}

	if bsTrans.Parsed() {
		cli.bsTrans()
	}

	if printChainCmd.Parsed() {
		// throwing errors
		cli.printChain()
	}

}
