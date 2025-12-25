package publicregister

import (
	"crypto"
	"net/rpc"

	"github.com/Samuelk0nrad/simple-did/lib/did"
)

type CreateDIDDocumentArgs struct {
	Did       string
	PublicKey crypto.PublicKey
}

func RegisterPublicDID(publicKey *crypto.PublicKey, name string) (did.DIDDocument, error) {
	client, err := rpc.Dial("tcp", "localhost:5435")
	var id did.DID
	var document did.DIDDocument
	if err != nil {
		return document, err
	}
	defer client.Close()

	err = client.Call("VDR.CreatePublicDID", name, &id)
	if err != nil {
		return document, err
	}

	didString, err := id.GetDID()
	if err != nil {
		return document, err
	}
	didDocumentARgs := CreateDIDDocumentArgs{
		Did:       didString,
		PublicKey: publicKey,
	}

	err = client.Call("VDR.CreateDIDDocument", didDocumentARgs, &document)
	if err != nil {
		return document, err
	}

	return document, nil
}
