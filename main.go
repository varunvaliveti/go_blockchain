package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

/* Only using significant information of a block on the blockchain

TimeStamp: Represents the time of the creation of the 'block'

Data: The information inside of the block

PrevBlockHash: Stores the hash value of the block before

Hash: The hash value of the block

*/

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

type Blockchain struct {
	blocks []*Block
}

func main() {
	blockchain := NewBlockChain()

	blockchain.AddBlocks("Send 50 DOGE to Andy")
	blockchain.AddBlocks("Send 10 more DOGE to Andy")

	for _, block := range blockchain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

/*
	 Here we are calculating the hashes
		On the concatenated combination, we calculate a SHA256 hash
*/
func (block *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timeStamp}, []byte{})
	hash := sha256.Sum256(headers)

	block.Hash = hash[:]
}

/*
Here we are simply creating a new block and returning it, and calling the set hash function
Read More: https://en.bitcoin.it/wiki/Block_hashing_algorithm
*/
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block

}

/*
Function to add blocks to the chain
*/
func (bc *Blockchain) AddBlocks(data string) {
	previousBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, previousBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

/*
Because every blockchain requires an 'initial' block, this is just our first block
*/
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block: ", []byte{})
}

/*
function to create a new block chain
*/
func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
