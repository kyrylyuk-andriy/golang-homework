package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

var Blockchain []Block

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(prevBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}

func IsBlockValid(newBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}
	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func GetBlocks(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(Blockchain)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func MineBlock(w http.ResponseWriter, r *http.Request) {
	prevBlock := Blockchain[len(Blockchain)-1]
	data := "this is a new data into a blockchain"
	newBlock := GenerateBlock(prevBlock, data)

	if IsBlockValid(newBlock, prevBlock) {
		Blockchain = append(Blockchain, newBlock)
		response, _ := json.Marshal(newBlock)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	} else {
		http.Error(w, "Invalid Block", http.StatusForbidden)
	}
}
