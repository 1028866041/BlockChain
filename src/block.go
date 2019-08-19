package src

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
)

type Block struct{
    index int64
    timeStamp int64
    preHash string
    hash string
    blockData string
}

func calculateHash(b Block) string{
    blockData := string(b.index) + string(b.timeStamp) + b.preHash + b.hash + b.blockData
    hashBytes := sha256.Sum256([]byte(blockData))
    return hex.EncodeToString(hashBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block{
    newBlock := Block{}
    newBlock.index = preBlock.index + 1
    newBlock.timeStamp = time.Now().Unix()
    newBlock.preHash = preBlock.hash
    newBlock.blockData = data
    newBlock.hash = calculateHash(newBlock)
    return newBlock
}

func GenerateGenesisBlock() Block{
    preBlock := Block{}
    preBlock.index = -1
    preBlock.hash = "NULL"
    return GenerateNewBlock(preBlock, "Genesis")
}


