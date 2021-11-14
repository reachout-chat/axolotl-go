package tests

import (
	"fmt"
	"testing"

	"github.com/reachout-chat/axolotl-go/fingerprint"
)

// TestFingerprint will test printing key fingerprints.
func TestFingerprint(t *testing.T) {

	// Create a serializer object that will be used to encode/decode data.
	serializer := newSerializer()

	// Create our users who will talk to each other.
	alice := newUser("Alice", 1, serializer)
	bob := newUser("Bob", 2, serializer)

	fp := fingerprint.NewDisplay(
		alice.identityKeyPair.PublicKey().Serialize(),
		bob.identityKeyPair.PublicKey().Serialize(),
	)

	fmt.Println(fp.DisplayText())

}
