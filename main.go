package main

import (
	"fmt"
	//"context"
	"strings"

	"github.com/gagliardetto/solana-go"
  //"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	
	prefix := "ab"
	//found := 0
	for {
		account := solana.NewWallet()
		pubkey := account.PublicKey().String()


	  fmt.Println("Account pub key: ", pubkey)

		if strings.HasPrefix(pubkey, prefix) {
			fmt.Println("Match found")
			fmt.Println("Account pub key: ", pubkey)
			fmt.Println("Account private key: ", account.PrivateKey)
			break
		}
	}

}
