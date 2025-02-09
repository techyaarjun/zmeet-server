package pion

import (
	"fmt"
	"github.com/pion/webrtc/v4"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"zmeet/pkg/logger"
)

type PeerConnectionService interface {
	CreateOffer()
	CreateAnswer()
	SetLocalDescription()
	SetRemoteDescription()
	ICEConnectionStateChangeNotify()
	SignalingStateChangeNotify()
}

type PC struct {
	mu                    sync.RWMutex
	peerConnection        *webrtc.PeerConnection
	peerConnectionService PeerConnectionService
	customLogger          *logger.Logger
}

func parseUint16(value string) uint16 {
	port, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(port)
}

func NewPeerConnection(l *logger.Logger) *PC {

	mediaEngine := webrtc.MediaEngine{}
	_ = mediaEngine.RegisterDefaultCodecs()

	ips := strings.Split(os.Getenv("IP_LIST"), ",")
	settingEngine := webrtc.SettingEngine{}

	settingEngine.SetNAT1To1IPs(ips, webrtc.ICECandidateTypeHost)
	_ = settingEngine.SetEphemeralUDPPortRange(parseUint16(os.Getenv("PORT_MIN")), parseUint16(os.Getenv("PORT_MAX")))
	settingEngine.SetNetworkTypes([]webrtc.NetworkType{webrtc.NetworkTypeUDP4})

	api := webrtc.NewAPI(webrtc.WithMediaEngine(&mediaEngine), webrtc.WithSettingEngine(settingEngine))

	config := webrtc.Configuration{}
	pc, err := api.NewPeerConnection(config)
	if err != nil {
		l.Debug("Failed to create peer connection")
		return nil
	}

	newPc := &PC{
		peerConnection: pc,
		customLogger:   l,
	}

	newPc.ICEConnectionStateChangeNotify()
	//newPc.SignalingStateChangeNotify()
	newPc.ConnectionStateChangeNotify()

	return newPc
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

func (p *PC) ICEConnectionStateChangeNotify() {
	p.PeerConnection().OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		mess := fmt.Sprintf("ICE Connection State has changed to: %s", state.String())
		p.CustomLogger().Info(mess)
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
	})
}
