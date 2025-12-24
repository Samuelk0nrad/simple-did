package main

import (
	"errors"
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

// GetDIDDocument retuns the DIDDocument for the assosiated did
func (v *VAR) GetDIDDocument(didString string, reply *did.DIDDocument) error {
	var result did.DIDDocument
	var error error

	if !did.ValidadeDID(didString) {
		message := fmt.Sprintf(`string "%s" is not a valide DID, please check DID syntax at: [https://www.w3.org/TR/did-1.0/#did-syntax]`, didString)
		error = errors.New(message)
	}

	var found bool
	for _, doc := range storage {
		docDidString, err := doc.Did.GetDID()
		if err == nil && docDidString == didString {
			result = doc
			found = true
			break
		}
	}

	if !found {
		message := fmt.Sprintf("no did Document found with the did: %s", didString)
		error = errors.New(message)
	}

	*reply = result
	return error
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
