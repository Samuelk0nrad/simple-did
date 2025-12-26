package did

import (
	"crypto/rsa"
)

// DIDDocument is a simplifid DID Document to store the public
// key of a did
type DIDDocument struct {
	Did                DID
	PublicKeyMultibase rsa.PublicKey
}
