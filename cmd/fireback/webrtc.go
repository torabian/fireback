package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pion/rtp"
	webrtc "github.com/pion/webrtc/v3"
)

// e.POST("/webrtc/offer", WebRTCOfferHandler)

func WebRTCOfferHandler(c *gin.Context) {
	var offer webrtc.SessionDescription
	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Basic PeerConnection config
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	})

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio)
	if err != nil {
		log.Fatal(err)
	}

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		fmt.Println("Track:", track.Kind())
		if track.Kind() == webrtc.RTPCodecTypeAudio {
			go func() {
				// Save to a file (raw Opus RTP payloads)
				file, err := os.Create("audio.opus")
				if err != nil {
					panic(err)
				}
				defer file.Close()

				rtpBuf := make([]byte, 1500)
				for {
					n, _, err := track.Read(rtpBuf)
					if err != nil {
						break
					}
					packet := &rtp.Packet{}
					if err := packet.Unmarshal(rtpBuf[:n]); err != nil {
						continue
					}

					// Save just the audio payload (not the full RTP packet)
					file.Write(packet.Payload)
				}
			}()
		}
	})

	// Handle ICE candidates if needed
	peerConnection.OnICECandidate(func(cand *webrtc.ICECandidate) {
		fmt.Println("On OnICECandidate", cand)
		// You can implement trickle ICE here if needed
	})

	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 🔥 Wait for ICE candidates to be gathered
	<-gatherComplete

	c.JSON(http.StatusOK, *peerConnection.LocalDescription())

}
