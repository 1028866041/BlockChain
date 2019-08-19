package src

import "fmt"

type BlockChain struct{
    Blocks []*Block
}

func (bc *BlockChain) ApendBlock(newBlock *Block) {
    if len(bc.Blocks) == 0 {
        bc.Blocks = append(bc.Blocks, newBlock)
        return
    }
    if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
        bc.Blocks = append(bc.Blocks, newBlock)
    }
}

func (bc *BlockChain)SendData(data string){
    preBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := GenerateNewBlock(*preBlock, data)
    bc.ApendBlock(&newBlock)
}

func isValid(newBlock Block, oldBlock Block) bool{
    if newBlock.index -1 != oldBlock.index{
        return false
    }
    if newBlock.preHash != oldBlock.hash{
        return false
    }
    return true
}

func NewBlockCHain() *BlockChain {
    block := GenerateGenesisBlock()
    blockChain := BlockChain{}
    blockChain.ApendBlock(&block)
    return &blockChain
}

func (bc *BlockChain) Print(){
    for _,block := range bc.Blocks {
        fmt.Printf("index: %d\n", block.index)
        fmt.Printf("pre: %s\n", block.preHash)
        fmt.Printf("hash: %s\n", block.hash)
        fmt.Printf("blockData: %s\n", block.blockData)
        fmt.Println()
    }
}