package pion

import (
	"fmt"
	"github.com/pion/webrtc/v4"
	"log"
	"sync"
	"zmeet/pkg/logger"
)

type PC struct {
	mu              sync.RWMutex
	peerConnection  *webrtc.PeerConnection
	customLogger    *logger.Logger
	iceConnected    bool
	connectionState bool
}

func NewPeerConnection(l *logger.Logger) *PC {

	mediaEngine := webrtc.MediaEngine{}
	_ = mediaEngine.RegisterDefaultCodecs()

	api := webrtc.NewAPI(webrtc.WithMediaEngine(&mediaEngine))

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	pc, err := api.NewPeerConnection(config)
	if err != nil {
		l.Debug("Failed to create peer connection")
		return nil
	}

	newPc := &PC{
		peerConnection:  pc,
		customLogger:    l,
		iceConnected:    false,
		connectionState: false,
	}

	go newPc.ICEConnectionStateChangeNotify()
	go newPc.ConnectionStateChangeNotify()

	return newPc
}

func (p *PC) SetICEConnected(state bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.iceConnected = state
}

func (p *PC) SetConnectionState(state bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.connectionState = state
}

func (p *PC) PeerConnection() *webrtc.PeerConnection {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.peerConnection
}

func (p *PC) CustomLogger() *logger.Logger {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.customLogger
}

func (p *PC) ICEConnected() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.iceConnected
}

func (p *PC) ConnectionState() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.connectionState
}

func (p *PC) CreateOffer() (webrtc.SessionDescription, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	sdp, err := p.peerConnection.CreateOffer(nil)
	if err != nil {
		p.customLogger.Debug("Failed to create offer")
		return webrtc.SessionDescription{}, err
	}

	return sdp, nil
}

func (p *PC) CreateAnswer() (webrtc.SessionDescription, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	sdp, err := p.peerConnection.CreateAnswer(nil)
	if err != nil {
		p.customLogger.Debug("Failed to create answer")
		return webrtc.SessionDescription{}, err
	}

	return sdp, nil
}

func (p *PC) SetLocalDescription(desc webrtc.SessionDescription) {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.peerConnection.SetLocalDescription(desc)
	if err != nil {
		p.customLogger.Debug("Failed to set local description")
	}
}

func (p *PC) SetRemoteDescription(desc webrtc.SessionDescription) {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.peerConnection.SetRemoteDescription(desc)
	if err != nil {
		log.Println(err)
		p.customLogger.Debug("Failed to set remote description")
	}
}

func (p *PC) AddICECandidate(ice webrtc.ICECandidateInit) {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.peerConnection.AddICECandidate(ice)
	if err != nil {
		p.customLogger.Debug("Failed to add ICECandidate")
	}
}

func (p *PC) ICEConnectionStateChangeNotify() {
	p.PeerConnection().OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		mess := fmt.Sprintf("ICE Connection State has changed to: %s", state.String())
		p.CustomLogger().Info(mess)
		if state == webrtc.ICEConnectionStateConnected {
			p.SetICEConnected(true)
		}
	})
}

func (p *PC) SignalingStateChangeNotify() {
	p.PeerConnection().OnSignalingStateChange(func(state webrtc.SignalingState) {
		mess := fmt.Sprintf("Signaling State has changed to: %s", state.String())
		p.CustomLogger().Info(mess)
	})
}

func (p *PC) ConnectionStateChangeNotify() {
	p.PeerConnection().OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		mess := fmt.Sprintf("Connection State has changed to: %s", state.String())
		p.CustomLogger().Info(mess)
		if state == webrtc.PeerConnectionStateConnected {
			p.SetConnectionState(true)
		}
	})
}

func (p *PC) OnICECandidate() {
	p.PeerConnection().OnICECandidate(func(candidate *webrtc.ICECandidate) {
		mess := fmt.Sprintf("Candiate : %s", candidate)
		p.CustomLogger().Info(mess)
	})
}
