package main

import (
	"fmt"

	"github.com/portto/solana-go-sdk/types"
)

func main() {
	// create a new wallet using types.NewAccount()
	wallet := types.NewAccount()

	// display the wallet public and private keys
	fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())
	fmt.Println("Private Key:", wallet.PrivateKey)
}
