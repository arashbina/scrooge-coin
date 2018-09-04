package main

import (
	"fmt"

	"github.com/arashbina/scrooge-coin/internal/transaction"
)

func main() {

	tx := transaction.Transaction{}
	tx.Hash = []byte("asdf")
	fmt.Println("Hello Scrooge transaction: ", tx.Hash)
}
