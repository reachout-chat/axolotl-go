package tests

import (
	"testing"

	"github.com/reachout-chat/axolotl-go/logger"
	"github.com/reachout-chat/axolotl-go/serialize"
	"github.com/reachout-chat/axolotl-go/util/keyhelper"
)

// TestPreKeys checks generating prekeys.
func TestPreKeys(t *testing.T) {

	// Create a serializer object that will be used to encode/decode data.
	serializer := serialize.NewSerializer()
	serializer.SignalMessage = &serialize.ProtoBufSignalMessageSerializer{}
	serializer.PreKeySignalMessage = &serialize.ProtoBufPreKeySignalMessageSerializer{}
	serializer.SignedPreKeyRecord = &serialize.JSONSignedPreKeyRecordSerializer{}

	logger.Info("Testing prekey generation...")
	identityKeyPair, err := keyhelper.GenerateIdentityKeyPair()
	if err != nil {
		t.Error(err)
	}

	logger.Info("Generating prekeys")
	preKeys, _ := keyhelper.GeneratePreKeys(1, 100, serializer.PreKeyRecord)
	logger.Info("Generated PreKeys: ", preKeys)

	logger.Info("Generating Signed PreKey")
	signedPreKey, _ := keyhelper.GenerateSignedPreKey(identityKeyPair, 1, serializer.SignedPreKeyRecord)
	logger.Info("Signed PreKey: ", signedPreKey)
}
