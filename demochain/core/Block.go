package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

/*
	1.新建工程demochain
	2.创建Block文件
	3.创建Block结构体与相关函数
*/

type Block struct {
	//区块头
	Index        int64  // 区块的位置
	Timestamp    int64  // 区块链的时间 使用秒数 避免失去
	PrevBlocHash string // 上一个区块的哈希值
	Hash         string // 当前区块的哈希值
	//区块体
	Data string // 区块数据 很多交易
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + string(b.PrevBlocHash) + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Data = data
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlocHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func GenerateisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
