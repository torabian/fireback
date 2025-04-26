import { usePostMouseposition } from "@/modules/fireback/sdk/modules/fireback/usePostMouseposition";
import { debounce } from "lodash";
import { useEffect, useRef } from "react";

export function DemoWebrtc() {
  const { submit } = usePostMouseposition({});
  let pc = useRef<RTCPeerConnection>();
  let dataChannel = useRef<RTCDataChannel>();

  const handleMouseMove = debounce(
    (e: MouseEvent) => {
      try {
        const value = { x: e.clientX, y: e.clientY };
        console.log(value);
        dataChannel.current?.send(JSON.stringify(value));
      } catch (err) {}
    },
    100,
    { maxWait: 300 }
  );

  const createWebrtcOffer = async () => {
    // let's initiate the webrtc
    pc.current = new RTCPeerConnection({
      iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
    });

    dataChannel.current = pc.current.createDataChannel("cursorData");
    dataChannel.current.onopen = () => console.log("Data channel opened!");
    dataChannel.current.onclose = () => console.log("Data channel closed!");

    const offer = await pc.current.createOffer();
    await pc.current.setLocalDescription(offer);

    await new Promise((resolve) => {
      if (pc.current.iceGatheringState === "complete") {
        resolve(true);
      } else {
        pc.current.onicegatheringstatechange = () => {
          if (pc.current.iceGatheringState === "complete") resolve(true);
        };
      }
    });

    const answer = await submit({ offer: offer } as any);
    console.log(2, answer);
    await pc.current.setRemoteDescription(answer.data.sessionDescription);
    console.log("done");
  };

  useEffect(() => {
    document.onmousemove = (e) => {
      handleMouseMove(e);
    };

    createWebrtcOffer();
  }, []);

  return (
    <div>
      <br />
      <h2>Webrtc</h2>
      <p>Let's demo the web rtc here</p>
    </div>
  );
}
