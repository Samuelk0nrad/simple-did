package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"net"
	"net/rpc"

	publicregister "github.com/Samuelk0nrad/simple-did/instances/issuer/publicRegister"
	"github.com/Samuelk0nrad/simple-did/lib/did"
)

var (
	PrivateKey rsa.PrivateKey
	Did        did.DID
)

type Verifier struct{}

type SignedData struct {
	Data      []byte
	Signature []byte
	PublicKey rsa.PublicKey
	Issuer    did.DID
}

func (v *Verifier) GetAccess(sdata *SignedData, reply *bool) error {
	var err error
	hash := sha256.Sum256(sdata.Data)
	validate := rsa.VerifyPKCS1v15(&sdata.PublicKey, crypto.SHA256, hash[:], sdata.Signature)
	if validate != nil {
		err = validate
	}

	*reply = true
	return err
}

func main() {
	// generate Private for public DID
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("and error accured while creating the private Key: %s", err)
		return
	}
	PrivateKey = *privateKey

	// Register Public DID and Document
	publicKey := privateKey.PublicKey
	document, err := publicregister.RegisterPublicDID(&publicKey, "issuer")
	if err != nil {
		fmt.Printf("and error accured while creating the public DID: %s", err)
		return
	}
	Did = document.Did

	verifier := new(Verifier)
	err = rpc.Register(verifier)
	if err != nil {
		fmt.Printf("error registering: %v", err)
	}

	listener, err := net.Listen("tcp", ":5800")
	if err != nil {
		fmt.Printf("error registering: %v", err)
	}
	defer listener.Close()

	rpc.Accept(listener)
}
