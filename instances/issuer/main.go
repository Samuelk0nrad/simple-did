package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	publicregister "github.com/Samuelk0nrad/simple-did/instances/issuer/publicRegister"
	"github.com/Samuelk0nrad/simple-did/lib/did"
)

var (
	PrivateKey rsa.PrivateKey
	Did        did.DID
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("and error accured while creating the private Key: %s", err)
		return
	}

	PrivateKey = *privateKey

	publicKey := PrivateKey.Public()

	document, err := publicregister.RegisterPublicDID(&publicKey, "issuer")
	if err != nil {
		fmt.Printf("and error accured while creating the public DID: %s", err)
		return
	}
	Did = document.Did
}
