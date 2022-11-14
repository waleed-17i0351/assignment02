package main

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}

type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
	var hash string
	var1 := fmt.Sprintf("%v", inputBlock.Data.Transactions)
	hash = fmt.Sprintf("%v", sha256.Sum256([]byte(var1)))
	return hash
}
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	var temp *Block = new(Block)

	if chainHead == nil {
		temp.Data = dataToInsert
		temp.PrevHash = ""
		temp.PrevPointer = nil
		var chash string
		chash = CalculateHash(temp)
		temp.CurrentHash = chash
		fmt.Printf("%s", "Added genesis block")
		fmt.Printf("%s", "\n") 
	} else {
		temp.Data = dataToInsert
		temp.PrevHash = chainHead.CurrentHash
		temp.PrevPointer = chainHead
		var chash string
		chash = CalculateHash(temp)
		temp.CurrentHash = chash
	}

	return temp
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	var temp *Block = chainHead

	for {
		if temp.PrevPointer == nil {
			trans := temp.Data.Transactions
			for a := 0; a < len(trans); a++ {
				if trans[a] == oldTrans {
					trans[a] = newTrans
				}
			}
			break
		} else if temp.PrevPointer != nil {
			trans := temp.Data.Transactions
			for a := 0; a < len(trans); a++ {
				if trans[a] == oldTrans {
					trans[a] = newTrans
				}
			}
		}
		temp = temp.PrevPointer
	}
}

func ListBlocks(chainHead *Block) {
	var temp *Block = chainHead
	for {
		if temp != nil {
			fmt.Printf("%s", temp.Data.Transactions)

			fmt.Printf("%s", "\n")
		}
		if temp == nil {
			fmt.Printf("%s", "Listing Completed")
			fmt.Printf("%s", "\n")
			break
		}

		temp = temp.PrevPointer
	}

}

func VerifyChain(chainHead *Block) {
	var phash string
	var chash string
	var temp *Block = chainHead

	for {
		if temp == nil {
			break
		} else if temp.PrevPointer != nil {
			phash = CalculateHash(temp.PrevPointer)
			chash = CalculateHash(temp)
			if temp.PrevHash == phash && temp.CurrentHash == chash {
				fmt.Printf("%s", "Integrity: Verified")
				fmt.Printf("%s", "\n")
			} else if temp.PrevHash != phash || temp.CurrentHash != chash {
				fmt.Printf("%s", "Integrity: Compromised")
				fmt.Printf("%s", "\n")
			}
		}
		temp = temp.PrevPointer
	}
}
func main() {

	var chainHead *Block
	genesis := BlockData{Transactions: []string{"Block", "Chain"}}
	chainHead = InsertBlock(genesis, chainHead)
	secondBlock := BlockData{Transactions: []string{"Lemon", "Mango", "Banana"}}
	chainHead = InsertBlock(secondBlock, chainHead)

	ListBlocks(chainHead)
	ListBlocks(chainHead)
	VerifyChain(chainHead)
}
