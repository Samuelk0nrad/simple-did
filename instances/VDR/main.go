package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/Samuelk0nrad/simple-did/lib/did"
)

const EntityIdentifier string = "vdr"

type VAR struct{}

var (
	registry []did.DID
	storage  []did.DIDDocument
)

type CreateDIDArgs struct {
	Name      string
	PublicKey string
	Address   string
}

// CreatePublicDID adds a new DID to the registry
func (v *VAR) CreatePublicDID(args *CreateDIDArgs, reply *did.DID) error {
	var result did.DID

	name := args.Name
	if did.CompareDIDs(&registry, "did", EntityIdentifier, name) {
		name += "12"
	}

	result = did.DID{
		Scheme:     "did",
		Method:     EntityIdentifier,
		Identifier: name,
	}

	registry = append(registry, result)

	*reply = result
	return nil
}

type CreateDIDDocumentArgs struct {
	Did       string
	PublicKey string
}

// CreateDIDDocument adds a new DID Document to the storage with the public key of the did
func (v *VAR) CreateDIDDocument(args *CreateDIDDocumentArgs, reply *did.DIDDocument) error {
	didString, err := did.ParseStringToDID(args.Did)
	if err != nil {
		return err
	}

	result := did.DIDDocument{
		Did:                didString,
		PublicKeyMultibase: args.PublicKey,
	}

	storage = append(storage, result)

	*reply = result
	return nil
}

func main() {
	api := new(VAR)

	err := rpc.Register(api)
	if err != nil {
		fmt.Printf("an error accured: %s", err)
	}

	listener, err := net.Listen("tcp", ":5600")
	if err != nil {
		fmt.Printf("an error accured: %s", err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 6500...")

	rpc.Accept(listener)
}
