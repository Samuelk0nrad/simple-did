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

type Con struct {
	OwnDID          string
	OwnPubKey       string
	ConnectorDID    string
	ConnectorPubKey string
}

type SignedData struct {
	Data      []byte
	Signature []byte
	PublicKey rsa.PublicKey
	Issuer    did.DID
}

// SignData signs the data with the private Key of the issuer and returs the data, the Signature, the Public Key, and the Issuer Did
func (c *Con) SignData(data []byte, reply *SignedData) error {
	var err error
	var result SignedData
	fmt.Printf("data: %v", string(data))

	// sign hashed data (bacause only small data can be hashed)
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, &PrivateKey, crypto.SHA256, hash[:])

	if err == nil {
		result = SignedData{
			Data:      data,
			Signature: signature,
			PublicKey: PrivateKey.PublicKey,
			Issuer:    Did,
		}
	}

	fmt.Println("signed new data")

	*reply = result
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
	publicKey := PrivateKey.PublicKey
	document, err := publicregister.RegisterPublicDID(&publicKey, "issuer")
	if err != nil {
		fmt.Printf("and error accured while creating the public DID: %s", err)
		return
	}
	Did = document.Did

	// Allow connection to Holder

	connection := new(Con)

	err = rpc.Register(connection)
	if err != nil {
		fmt.Printf("error Registering Con: %v", err)
	}

	listener, err := net.Listen("tcp", ":5700")
	if err != nil {
		fmt.Printf("error listening: %v", err)
	}
	defer listener.Close()

	fmt.Println("listening on port 5700 ... ")
	rpc.Accept(listener)
}
