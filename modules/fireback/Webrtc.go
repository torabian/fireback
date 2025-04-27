package fireback

import (
	"fmt"

	"github.com/pion/webrtc/v3"
)

func HandleWebRTC(offer *webrtc.SessionDescription, peerSetup func(peerConnection *webrtc.PeerConnection)) (*webrtc.SessionDescription, error) {
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		BundlePolicy: webrtc.BundlePolicyMaxBundle,
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	})
	if err != nil {
		return nil, err
	}

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State changed: %s\n", connectionState.String())
	})

	_, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio)
	if err != nil {
		return nil, err
	}

	if peerSetup != nil {
		peerSetup(peerConnection)
	}

	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		// Handle trickle ICE if needed
	})

	if err := peerConnection.SetRemoteDescription(*offer); err != nil {
		return nil, err
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		return nil, err
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	if err := peerConnection.SetLocalDescription(answer); err != nil {
		return nil, err
	}

	<-gatherComplete

	return peerConnection.LocalDescription(), nil
}
