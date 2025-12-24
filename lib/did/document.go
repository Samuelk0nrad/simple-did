package did

// DIDDocument is a simplifid DID Document to store the public
// key of a did
type DIDDocument struct {
	did                DID
	publicKeyMultibase string
}
