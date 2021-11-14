package store

import (
	"github.com/reachout-chat/axolotl-go/groups/state/record"
	"github.com/reachout-chat/axolotl-go/protocol"
)

type SenderKey interface {
	StoreSenderKey(senderKeyName *protocol.SenderKeyName, keyRecord *record.SenderKey)
	LoadSenderKey(senderKeyName *protocol.SenderKeyName) *record.SenderKey
}
