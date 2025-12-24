package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/Samuelk0nrad/simple-did/lib/did"
)

const EntityIdentifier string = "vdr"

type VAR struct{}

var registry []did.DID

type CreateDIDArgs struct {
	Name      string
	PublicKey string
	Address   string
}

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
