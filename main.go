package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain - the chain
type BlockChain struct {
	blocks []*Block
}

// AddBlock - add a block
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Block - the block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash - derive the hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock - create the block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Genesis -
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain -
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash:  %x\n", block.PrevHash)
		fmt.Printf("Data in block:  %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("")
	}
}
