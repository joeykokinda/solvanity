package main

import (
	"crypto/ed25519"
	crypto_rand "crypto/rand"
	"fmt"
	"math/rand"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/mr-tron/base58"
)

func main() {
	iterations := 1000000

	// 1. solana.NewWallet()
	start := time.Now()
	for i := 0; i < iterations; i++ {
		solana.NewWallet()
	}
	fmt.Println("1. solana.NewWallet():      ", time.Since(start))

	// 2. raw ed25519 with crypto/rand
	start = time.Now()
	for i := 0; i < iterations; i++ {
		ed25519.GenerateKey(crypto_rand.Reader)
	}
	fmt.Println("2. raw ed25519 crypto/rand: ", time.Since(start))

	// 3. raw ed25519 with math/rand
	rng := rand.New(rand.NewSource(42))
	start = time.Now()
	for i := 0; i < iterations; i++ {
		ed25519.GenerateKey(rng)
	}
	fmt.Println("3. raw ed25519 math/rand:   ", time.Since(start))

	// 4. raw ed25519 math/rand + base58 encode
	rng2 := rand.New(rand.NewSource(42))
	start = time.Now()
	for i := 0; i < iterations; i++ {
		pub, _, _ := ed25519.GenerateKey(rng2)
		base58.Encode(pub)
	}
	fmt.Println("4. ed25519 math/rand+b58:   ", time.Since(start))

	// 5. raw ed25519 math/rand + base58 + prefix check
	rng3 := rand.New(rand.NewSource(42))
	prefix := "abc"
	prefixLen := len(prefix)
	start = time.Now()
	for i := 0; i < iterations; i++ {
		pub, _, _ := ed25519.GenerateKey(rng3)
		full := base58.Encode(pub)
		_ = len(full) >= prefixLen && full[:prefixLen] == prefix
	}
	fmt.Println("5. ed25519 math/rand+check: ", time.Since(start))
}
