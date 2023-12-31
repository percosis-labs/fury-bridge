package signer

import (
	"github.com/fury-labs/fury-bridge/relayer/broadcast"
	"github.com/fury-labs/fury-bridge/relayer/broadcast/pending_store"
	"github.com/fury-labs/fury-bridge/relayer/broadcast/types"
)

type BroadcastHandler struct {
	signer *Signer
}

func NewBroadcastHandler(signer *Signer) *BroadcastHandler {
	return &BroadcastHandler{
		signer: signer,
	}
}

var _ broadcast.BroadcastHandler = (*BroadcastHandler)(nil)

func (h *BroadcastHandler) RawMessage(msg pending_store.MessageWithPeerMetadata) {
}

func (h *BroadcastHandler) ValidatedMessage(msg types.BroadcastMessage) {
	h.signer.handleBroadcastMessage(msg)
}

func (h *BroadcastHandler) MismatchMessage(msg pending_store.MessageWithPeerMetadata) {}
