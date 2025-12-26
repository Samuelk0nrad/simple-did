package publicregister

import (
	"crypto/rsa"
	"fmt"
	"net/rpc"

	"github.com/Samuelk0nrad/simple-did/lib/did"
)

type CreateDIDDocumentArgs struct {
	Did       string
	PublicKey rsa.PublicKey
}

func RegisterPublicDID(publicKey *rsa.PublicKey, name string) (did.DIDDocument, error) {
	client, err := rpc.Dial("tcp", "localhost:5600")
	var id did.DID
	var document did.DIDDocument
	if err != nil {
		return document, err
	}
	defer client.Close()

	err = client.Call("VDR.CreatePublicDID", name, &id)
	if err != nil {
		return document, err
	} else {
		fmt.Printf("successfully registered public DID %v", id)
	}

	didString, err := id.GetDID()
	if err != nil {
		return document, err
	}
	didDocumentARgs := CreateDIDDocumentArgs{
		Did:       didString,
		PublicKey: *publicKey,
	}

	err = client.Call("VDR.CreateDIDDocument", didDocumentARgs, &document)
	if err != nil {
		return document, err
	}

	return document, nil
}
