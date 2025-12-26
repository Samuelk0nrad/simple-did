package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/rpc"

	publicregister "github.com/Samuelk0nrad/simple-did/instances/issuer/publicRegister"
	"github.com/Samuelk0nrad/simple-did/lib/did"
)

var (
	PrivateKey rsa.PrivateKey
	Did        did.DID
)

type Data struct {
	Age  int
	Name string
}

type SignedData struct {
	Data      []byte
	Signature []byte
	PublicKey rsa.PublicKey
	Issuer    did.DID
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

	// -- get Signature from issuer --
	// create data to sign
	data := Data{Age: 18, Name: "Simon"}
	dataStirng, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error stringifing data: %v", err)
	}

	// create client to connect to issuer a ...any
	client, err := rpc.Dial("tcp", "localhost:5700")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer client.Close()

	// get signed data
	var signed SignedData

	err = client.Call("Con.SignData", dataStirng, &signed)
	if err != nil {
		fmt.Printf("error signing data: %v ", err)
	}

	jsonSign, err := json.Marshal(signed)
	if err != nil {
		fmt.Printf("error stringifing signature: %v", err)
	}

	fmt.Printf("\njson Signature: %v,\n data: %v", string(jsonSign), string(signed.Data))
}
