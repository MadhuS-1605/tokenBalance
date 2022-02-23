package main

import (
	"fmt"
	"log"
	"net/http"

	Handlers "tokenBal/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	// Create a client instance to connect to our providr
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a9126c60ad5346e79ff1061816be1dd4")

	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// We will define a single endpoint
	r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))
}