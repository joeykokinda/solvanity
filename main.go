package main

import (
	"fmt"
	"crypto/ed25519"
	"sync/atomic"
	"time"
	"runtime"
	"math/rand"
	"github.com/mr-tron/base58"
  //"github.com/gagliardetto/solana-go/rpc"
)

type Result struct {
    pubkey  string
    privkey ed25519.PrivateKey
}

func generateWallet(prefix string, total *int64, resultChan chan Result, stopChan chan struct{}) {
	//checkBytes := len(prefix)*5/4 + 4
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefixLen := len(prefix)
	for {
		select {
		case <-stopChan:
			return
		default:
		}

		pub, priv, _ := ed25519.GenerateKey(rng)
		atomic.AddInt64(total, 1)

		full := base58.Encode(pub)
		if len(full) >= prefixLen && full[:prefixLen] == prefix {
			resultChan <- Result{pubkey: full, privkey: priv}
			return
		}
	}
}

func main() {
	resultChan := make(chan Result)
	stopChan := make(chan struct{})
	var total int64
	start := time.Now()
  ticker := time.NewTicker(time.Second)

	fmt.Printf("starting %d workers\n", runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go generateWallet("joey", &total, resultChan, stopChan)
	}

	go func() {
    for range ticker.C {
      t := atomic.LoadInt64(&total)
      elapsed := time.Since(start).Seconds()
      fmt.Printf("%d wallets | %.2f per second\n", t, float64(t)/elapsed)
    }
  }()

	result := <-resultChan
	close(stopChan)
	fmt.Println("Match found")
	fmt.Println("pub:", result.pubkey)
	fmt.Println("priv:", base58.Encode(result.privkey))
}
