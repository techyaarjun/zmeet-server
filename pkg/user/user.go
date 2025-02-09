package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/pion/webrtc/v4"
	"sync"
	"zmeet/pkg/pion"
)

type ZMeetUser struct {
	mu             sync.RWMutex
	id             uuid.UUID
	name           string
	peerConnection *pion.PC
	dataChannel    *webrtc.DataChannel
	connected      bool
	ctx            context.Context
	cancel         context.CancelFunc
}

func NewZMeetUser(id uuid.UUID, name string, pc *pion.PC, ctx context.Context, ctxCancel context.CancelFunc) *ZMeetUser {
	return &ZMeetUser{
		id:             id,
		peerConnection: pc,
		dataChannel:    nil,
		connected:      false,
		name:           name,
		ctx:            ctx,
		cancel:         ctxCancel,
	}
}

func (z *ZMeetUser) ID() uuid.UUID {
	z.mu.RLock()
	defer z.mu.RUnlock()
	return z.id
}

func (z *ZMeetUser) Name() string {
	z.mu.RLock()
	defer z.mu.RUnlock()
	return z.name
}

func (z *ZMeetUser) Connected() bool {
	z.mu.RLock()
	defer z.mu.RUnlock()
	return z.connected
}

func (z *ZMeetUser) SetConnected(state bool) {
	z.mu.RLock()
	defer z.mu.RUnlock()
	z.connected = state
}

func (z *ZMeetUser) CTX() context.Context {
	z.mu.RLock()
	defer z.mu.RUnlock()
	return z.ctx
}

func (z *ZMeetUser) Cancel() context.CancelFunc {
	z.mu.RLock()
	defer z.mu.RUnlock()
	return z.cancel
}

func (z *ZMeetUser) PeerConnection() *pion.PC {
	z.mu.RLock()
	defer z.mu.RUnlock()
	return z.peerConnection
}

func (z *ZMeetUser) DataChannel() *webrtc.DataChannel {
	z.mu.RLock()
	defer z.mu.RUnlock()
	return z.dataChannel
}
