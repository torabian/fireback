package fireback

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
)

func init() {
	// Override the implementation with our actual code.
	MousePositionCaptureActionImp = MousePositionCaptureAction
}

func MousePositionCaptureAction(
	req *MousePositionCaptureActionReqDto,
	q QueryDSL,
) (*MousePositionCaptureActionResDto,
	*IError,
) {

	fmt.Println(req)
	session, err := HandleWebRTC(req.Offer, func(peerConnection *webrtc.PeerConnection) {
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
						fmt.Println("Read", n)
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

		onRam := func(dc *webrtc.DataChannel) {
			fmt.Println("Data channel opened:", dc.Label())
			dc.OnMessage(func(msg webrtc.DataChannelMessage) {
				var cursorPosition struct {
					UsedJSHeapSize int64 `json:"usedJSHeapSize"`
				}
				err := json.Unmarshal(msg.Data, &cursorPosition)
				if err != nil {
					log.Println("Error decoding cursor position:", err)
					return
				}
				fmt.Printf("Ram usage: %v \n", string(msg.Data))
			})
		}

		onMouse := func(dc *webrtc.DataChannel) {
			fmt.Println("Data channel opened:", dc.Label())
			dc.OnMessage(func(msg webrtc.DataChannelMessage) {
				var cursorPosition map[string]float64
				err := json.Unmarshal(msg.Data, &cursorPosition)
				if err != nil {
					log.Println("Error decoding cursor position:", err)
					return
				}
				fmt.Printf("Cursor position: x = %f, y = %f\n", cursorPosition["x"], cursorPosition["y"])
			})
		}

		// Handle data channel for cursor position
		peerConnection.OnDataChannel(func(dc *webrtc.DataChannel) {
			if dc.Label() == "ram" {
				onRam(dc)
			}
			if dc.Label() == "mouse" {
				onMouse(dc)
			}

		})
	})

	return &MousePositionCaptureActionResDto{
		SessionDescription: session,
	}, CastToIError(err)
}
