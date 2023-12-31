package pending_store

import (
	"context"

	"github.com/fury-labs/fury-bridge/relayer/broadcast/types"
	"github.com/libp2p/go-libp2p-core/peer"
)

// MessageWithPeerMetadata is a message with metadata about the peer that sent it.
type MessageWithPeerMetadata struct {
	BroadcastMessage types.BroadcastMessage

	// Not transmitted over wire, added when received.
	PeerID peer.ID

	Context context.Context
}
