package main

import (
	"fmt"

	"summitto.com/txsigner/pkg/transactions"
)

func main() {
	fmt.Println("Generating ed25519 Private/Public Key Pair:")

	publicKeyEncoded, privateKeyEncoded, _, _, err := transactions.GenerateKeyPair()

	if err != nil {
		fmt.Println("There was an error while generating the key pair")
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("Public Key:", publicKeyEncoded)
		fmt.Println("Private Key:", privateKeyEncoded)
	}
}
