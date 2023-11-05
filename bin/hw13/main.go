package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	blockchain "github.com/kyrylyuk-andriy/golang-homework/hw13"
)

func main() {
	r := mux.NewRouter()

	genesisBlock := blockchain.Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
	}

	genesisBlock.Hash = blockchain.CalculateHash(genesisBlock)
	blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)

	r.HandleFunc("/blocks", blockchain.GetBlocks).Methods("GET")
	r.HandleFunc("/mine", blockchain.MineBlock).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Listen on port 8082...")
	http.ListenAndServe(":8082", nil)
}
