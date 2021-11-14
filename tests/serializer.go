package tests

import (
	"github.com/reachout-chat/axolotl-go/serialize"
)

// newSerializer will return a JSON serializer for testing.
func newSerializer() *serialize.Serializer {
	serializer := serialize.NewProtoBufSerializer()

	return serializer
}
