package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type block struct {
	blockNumber        int
	data               string
	nonce              int
	previousBlockHash  []byte
	hashOfCurrentBlock []byte
}

func newBlock(chain []block, data string) block {
	nextBlockNumber := len(chain)
	if nextBlockNumber == 0 {
		fmt.Println(`Creating Origin Block`)
		originBlock := block{
			blockNumber:        0,
			data:               data,
			nonce:              1,
			previousBlockHash:  []byte{0},
			hashOfCurrentBlock: []byte{0},
		}
		originBlock.mineBlock()
		return originBlock
	}
	fmt.Println(`Creating New Block number`, nextBlockNumber)
	nextBlock := block{
		blockNumber:       nextBlockNumber,
		data:              data,
		previousBlockHash: chain[nextBlockNumber-1].hashOfCurrentBlock,
	}
	nextBlock.mineBlock()
	return nextBlock

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

func (block *block) mineBlock() {
	nonceSolved := false
	for !nonceSolved {
		newHash := sha256.New()
		numberAndData := append([]byte(strconv.Itoa(block.blockNumber)), []byte(block.data)...)
		numberDataAndNonce := append(numberAndData, []byte(strconv.Itoa(block.nonce))...)
		numberDataNonceAndPrevBlock := append(numberDataAndNonce, block.previousBlockHash...)
		newHash.Write(numberDataNonceAndPrevBlock)
		block.hashOfCurrentBlock = newHash.Sum(nil)

		if block.hashOfCurrentBlock[0] == 0 && block.hashOfCurrentBlock[1] == 0 && block.hashOfCurrentBlock[2] == 0 && block.hashOfCurrentBlock[3] == 0 {
			// if block.hashOfCurrentBlock[0] == 0 && block.hashOfCurrentBlock[1] == 0 && block.hashOfCurrentBlock[2] == 0 {
			// if block.hashOfCurrentBlock[0] == 0 && block.hashOfCurrentBlock[1] == 0 {
			fmt.Println(`block: `, block.blockNumber, ` nonce: `, block.nonce)
			nonceSolved = true
		} else {
			block.nonce++
		}
	}

	fmt.Println(`block NUMBER : `, block.blockNumber, `HASH: `, block.hashOfCurrentBlock)
}

func main() {
	fmt.Println(`Welcome to the blockchain app`)
	start := time.Now()
	myChain := []block{}
	block0 := newBlock(myChain, `This Is the Data that is in Block Zero`)
	myChain = appendBlockToChain(myChain, block0)
	block1 := newBlock(myChain, `This is the data that will be in Block 1`)
	myChain = appendBlockToChain(myChain, block1)
	block2 := newBlock(myChain, `Here is my data for block 2`)
	myChain = appendBlockToChain(myChain, block2)
	block3 := newBlock(myChain, `Finally I have a third block and this data is good`)
	myChain = appendBlockToChain(myChain, block3)
	block4 := newBlock(myChain, `The last block of data`)
	myChain = appendBlockToChain(myChain, block4)

	fmt.Println(`MyChain: length `, len(myChain))
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(`Mining took `, elapsed/5, ` seconds per block`)
}
