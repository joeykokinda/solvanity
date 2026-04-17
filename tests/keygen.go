package main

import (
	"fmt"
	"time"
	"github.com/gagliardetto/solana-go"
	"github.com/mr-tron/base58"
)

func main() {
	account := solana.NewWallet()
	pubkeyBytes := account.PublicKey()

	// benchmark just base58 encode
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		base58.Encode(pubkeyBytes[:])
	}
	fmt.Println("1M base58 encodes:", time.Since(start))

	// benchmark just keygen
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		solana.NewWallet()
	}
	fmt.Println("1M keygens:", time.Since(start))
}
