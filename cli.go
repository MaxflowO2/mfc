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
)

// CLI responsible for processing command line arguments
type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock, with nulltrans")
	fmt.Println("  autogen, creates a block every 15 seconds with nulltrans")
	fmt.Println("  nulltrans, creates a basic 'null' transaction")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock() {
        cli.bc.AddBlock(nullTransaction())
        fmt.Println("Success!")
}

func (cli *CLI) autoGen(t time.Time) {
	cli.bc.AddBlock(nullTransaction())
}

func (cli *CLI) nullTrans() {
	nullTransaction()
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Transactions: %x\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
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
	autoGen := flag.NewFlagSet("autogen", flag.ExitOnError)
	nullTrans := flag.NewFlagSet("nulltrans", flag.ExitOnError)
//	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "autogen":
		err := autoGen.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

        case "nulltrans":
                err := nullTrans.Parse(os.Args[2:])
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

        if nullTrans.Parsed() {
        	cli.nullTrans()
        }

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

