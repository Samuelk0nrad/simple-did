package did

import (
	"errors"
	"fmt"
)

// DID is a simple text string consisting of three parts:
//  1. **Scheme:** the `did` URI scheme identifier, mostly / allwaies "did"
//  2. **Method:** the identifier of the DID method, the did Method is the
//     profider where the did is been created, if the did is created on a
//     holder / subject for peer to peer use the method is "peer"
//     else it is the identifier of the creator
//  3. **Identifier:** the DID method-specific identitier, this identifier
//     is unique ofer all the method DIDs for `peer` DIDs it follow a
//     strict scheme descriped under:
//     (Peer DID method spec)[https://identity.foundation/peer-did-method-spec/]
type DID struct {
	Scheme     string
	Method     string
	Identifier string
}

// GetDID returns the string representation of the DID following the
// standard schema for DIDs descriped here:
// (Decentralized Identifiers v1.0)[https://www.w3.org/TR/did-1.0/]
//
// # The defauld Scheme is `did` if no Scheme is set. If the Method or
// the Identifier is not set, a error will get returned
//
// Stucture:
//
//	`Scheme`:`Method`:`Identifier`
//
// example:
//
//	did:example:123456789abcdefghi
func (d *DID) GetDID() (string, error) {
	var result string

	var error error

	if d.Scheme != "" {
		result += fmt.Sprintf("%s:", d.Scheme)
	} else {
		result += "did:"
	}

	if d.Method != "" {
		result += fmt.Sprintf("%s:", d.Method)
	} else {
		error = errors.New("no Method set for DID")
	}

	if d.Identifier != "" {
		result += fmt.Sprintf("%s:", d.Identifier)
	} else {
		error = errors.New("no Identifier set for DID")
	}

	return result, error
}

// CompareDID compares a DID with a method and / or identifer and / or scheme.
// Method, identifer and scheme are optional if they are are empty they will
// not get compared
func (d *DID) CompareDID(scheme string, method string, identifier string) bool {
	if method != d.Method && method != "" {
		return false
	}

	if identifier != d.Identifier && identifier != "" {
		return false
	}

	if scheme != d.Scheme && scheme != "" {
		return false
	}

	return true
}

// CompareDIDs checks if there is a DID with a method and / or identifer
// and / or scheme. Method, identifer and scheme are optional if they are
// are empty they will not get compared
func CompareDIDs(ds *[]DID, scheme string, method string, identifier string) bool {
	var found bool

	// iterate ofer the slice and compare the DID
	for _, u := range *ds {
		if u.CompareDID(scheme, method, identifier) {
			found = true
			break
		}
	}

	return found
}
