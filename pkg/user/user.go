package user

import (
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
}

func NewZMeetUser(id uuid.UUID, name string, pc *pion.PC) *ZMeetUser {
	return &ZMeetUser{
		id:             id,
		peerConnection: pc,
		dataChannel:    nil,
		connected:      false,
		name:           name,
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
