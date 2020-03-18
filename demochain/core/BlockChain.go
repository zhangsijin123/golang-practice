package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}

func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

// 添加区块
func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}

}

// 区块校验
func isValid(newBlock, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlocHash != oldBlock.Hash {
		return false
	}
	if newBlock.Hash != calculateHash(newBlock) {
		return false
	}
	return true
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index:%d\n", block.Index)
		fmt.Printf("Prev.Hash:%s\n", block.PrevBlocHash)
		fmt.Printf("Curr.Hash:%s\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Timestamp:%d\n", block.Timestamp)
	}
}
