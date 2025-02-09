package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pion/webrtc/v4"
	"net/http"
	"zmeet/pkg/pion"
	"zmeet/pkg/store"
	"zmeet/pkg/user"
)

type SDPRequest struct {
	Type string `json:"type"`
	SDP  string `json:"sdp"`
}

func Offer(id string, s *store.Store, c *gin.Context) {

	pc := pion.NewPeerConnection(s.CustomLogger())

	userID, _ := uuid.Parse(id)
	u := user.NewZMeetUser(userID, "Hello", pc)

	init := webrtc.RTPTransceiverInit{
		Direction: webrtc.RTPTransceiverDirectionSendrecv,
	}

	_, _ = pc.PeerConnection().AddTransceiverFromKind(webrtc.RTPCodecTypeAudio, init)

	offer, err := pc.CreateOffer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	pc.SetLocalDescription(offer)

	s.AddZMeetUser(u)
	c.JSON(http.StatusCreated, offer)
}

func Answer(id string, s *store.Store, c *gin.Context) {

	userID, _ := uuid.Parse(id)
	u := s.GetZMeetUser(userID)
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	var sdp webrtc.SessionDescription
	if err := c.ShouldBindJSON(&sdp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid SDP answer"})
		return
	}

	pc := u.PeerConnection()
	pc.SetRemoteDescription(sdp)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func HandleHandshake(s *store.Store, c *gin.Context) {

	var sdp webrtc.SessionDescription
	if err := c.ShouldBindJSON(&sdp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid SDP offer"})
		return
	}

	pc := pion.NewPeerConnection(s.CustomLogger())
	pc.SetRemoteDescription(sdp)
	answer, err := pc.CreateAnswer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	pc.SetLocalDescription(answer)
	c.JSON(http.StatusOK, answer)
}
