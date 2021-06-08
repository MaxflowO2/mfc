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
//	"time"
)

// CLI responsible for processing command line arguments
type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock - with multiple 'bs-transactions'")
//	fmt.Println("  autogen - creates a block every 15 seconds with nulltrans")
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
	var toBlock []*Transaction
	a := bsTransaction()
	b := bsTransaction()
	toBlock = append(toBlock, a, b)
        cli.bc.AddBlock(toBlock)
        fmt.Println("Success!")
}

//func (cli *CLI) autoGen(t time.Time) {
//	cli.bc.AddBlock(bsTransaction())
//}

func (cli *CLI) bsTrans() {
	bsTransaction()
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("Block Height: %v\n", block.Height)
		fmt.Printf("Previous Hash:\n%x\n", block.PrevBlockHash)
		fmt.Printf("Transactions in Block:\n%v\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Difficulty: %v\n", block.Difficulty)
		fmt.Printf("Hashed By: %s\n", block.HashBy)
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

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
//	autoGen := flag.NewFlagSet("autogen", flag.ExitOnError)
	bsTrans := flag.NewFlagSet("bstrans", flag.ExitOnError)
//	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
//	case "autogen":
//		err := autoGen.Parse(os.Args[2:])
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

//	if autoGen.Parsed() {
//	        Repeat(15*time.Second, cli.autoGen)
//	}

        if bsTrans.Parsed() {
        	cli.bsTrans()
        }

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}

