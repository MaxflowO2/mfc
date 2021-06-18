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
//	"strconv"
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
	fmt.Println("  powtest - uses 10 transactions, with forced block gen at 15 seconds for PoW testing")
	fmt.Println("  powvisatest - runs visatest every 35 seconds")
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
	timeOne := time.Now().Unix()
	// Visa's 1,734 per second * 15 seconds
	for i := 0; i < 26010; i++ {
		fill := bsTransaction()
		sendit = append(sendit, fill)
		j := i+1
		fmt.Printf(" of 26010 Transactions of visa test")
		fmt.Printf("\r%v", j)
	}
	fmt.Printf("\n")
	timeTwo := time.Now().Unix()
	cli.bc.AddBlock(sendit)
	timeThree := time.Now().Unix()
//	visaTime := 15
	timeProcess := timeTwo - timeOne
	timeTotal := timeThree - timeOne
	numb := 26010
//	avgProcess := timeProcess / 26010
//	avgTime := timeTotal / 26010
	fmt.Printf("Time to process %v transactions is %v\n", numb, timeProcess)
	fmt.Printf("Time to in block %v transactions is %v\n", numb, timeTotal)
	fmt.Println("Success!")
}

func (cli *CLI) powTest(t time.Time) {
	var sendit []*Transaction
	for i := 0; i < 10; i++ {
		fill := bsTransaction()
		sendit = append(sendit, fill)
//		fmt.Printf("Transaction %v of 10 done\n", i)
	}
	cli.bc.AddBlock(sendit)
}

//func (cli *CLI) addToDB() {
//        addy := LoadAddress()
//        AddAddress(addy)
//}

func (cli *CLI) powVisaTest(t time.Time) {
	cli.addBlock()
}

func (cli *CLI) printChain() {
	// throwing errors
	bci := cli.bc.Iterator()
	//	var addy *MFCAddress
	for {
		block := bci.Next()
		fmt.Printf("Block Height: %v\n", block.Height)
//		fmt.Printf("Previous Hash:\n%x\n", block.PrevBlockHash)
//		fmt.Printf("Transactions in Block:\n%v\n", block.Transactions)
//		fmt.Printf("Hash: %x\n", block.Hash)
//		pow := NewProofOfWork(block)
//		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Difficulty: %v\n", block.Difficulty)
//		fmt.Printf("Hashed By: %s\n", block.HashBy)
//		fmt.Printf("Signature: %x\n", block.Signed)
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
	powTest := flag.NewFlagSet("powtest", flag.ExitOnError)
	powVisaTest := flag.NewFlagSet("powvisatest", flag.ExitOnError)
	//	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "visatest":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "powtest":
		err := powTest.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

		//	case "addtodb":
		//		err := addToDB.Parse(os.Args[2:])
		//		if err != nil {
		//			log.Panic(err)
		//		}

	case "powvisatest":
		err := powVisaTest.Parse(os.Args[2:])
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

	if powTest.Parsed() {
		Repeat(1*time.Millisecond, cli.powTest)
	}

	//	if addToDB.Parsed() {
	//		cli.addToDB()
	//	}

	if powVisaTest.Parsed() {
		Repeat(35*time.Second, cli.powVisaTest)
	}

	if printChainCmd.Parsed() {
		// throwing errors
		cli.printChain()
	}

}
