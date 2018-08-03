package main

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

type block struct {
	blockNumber        int
	data               string
	previousBlockHash  uint64
	hashOfCurrentBlock uint64
}

func newBlock(c []block, data string) block {
	nextBlockNumber := len(c)
	if nextBlockNumber == 0 {
		//create start of chain
	} else {
		fmt.Println("Creating New Block number", nextBlockNumber)
		nextBlock := block{
			blockNumber:       nextBlockNumber,
			data:              data,
			previousBlockHash: c[0].hashOfCurrentBlock,
		}
		return nextBlock
	}
	return block{}
}

func appendBlockToChain(chain []block, data ...block) []block {
	m := len(chain)
	n := m + len(data)
	if n > cap(chain) {
		newchain := make([]block, (n+1)*2)
		copy(newchain, chain)
		chain = newchain
	}
	chain = chain[0:n]
	copy(chain[m:n], data)
	return chain
}

func (b *block) hashBlock() {
	fmt.Println("hashing block number ", b.blockNumber)
	hash, err := hashstructure.Hash(b, nil)
	if err != nil {
		panic(err)
	}
	b.hashOfCurrentBlock = hash
	fmt.Println("block hash: ", b.hashOfCurrentBlock)
}

func main() {
	fmt.Println("Welcome to the blockchain app")

	block0 := block{
		blockNumber:       0,
		data:              "data for firsdt block",
		previousBlockHash: 0,
	}

	block0.hashBlock()
	myChain := []block{}
	myChain = appendBlockToChain(myChain, block0)
	block1 := newBlock(myChain, "more data to add")
	block1.hashBlock()
	myChain = appendBlockToChain(myChain, block1)

	fmt.Println("MyChain: ", myChain)

	// hashing function not working
	// define a better struct for chain
	// build initial block creation into new block function
}
