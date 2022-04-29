package broadcast

import (
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p-core/peer"
)

// broadcasterHook defines the interface for broadcaster hooks. This is not
// exported as it should only be used for testing.
type broadcasterHook interface {
	// Run before a raw message is broadcasted and can be used to modify the
	// message.
	BeforeBroadcastRawMessage(b *Broadcaster, target peer.ID, pb *proto.Message)
}

// NoOpBroadcasterHook is a broadcasterHook that does nothing.
type noOpBroadcasterHook struct{}

var _ broadcasterHook = (*noOpBroadcasterHook)(nil)

func (h *noOpBroadcasterHook) BeforeBroadcastRawMessage(b *Broadcaster, target peer.ID, pb *proto.Message) {
}
