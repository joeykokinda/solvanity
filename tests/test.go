package main

import (
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/mr-tron/base58"
)

func main() {
	prefix := "ab"

	for i := 0; i < 5; i++ {
		account := solana.NewWallet()
		pubkeyBytes := account.PublicKey()

		full := base58.Encode(pubkeyBytes[:])

		fmt.Println("full:        ", full)
		fmt.Println("first N chars:", full[:len(prefix)])
		fmt.Println("prefix:      ", prefix)
		fmt.Println("match:", full[:len(prefix)] == prefix)
		fmt.Println("---")
	}
}
