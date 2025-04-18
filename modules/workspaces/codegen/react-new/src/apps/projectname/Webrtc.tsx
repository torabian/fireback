// client/src/App.js
import React, { useEffect } from "react";

export const Webrtc = () => {
  const start = async () => {
    const pc = new RTCPeerConnection({
      iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
    });
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
    pc.oniceconnectionstatechange = (e) => console.log(pc.iceConnectionState);

    pc.ontrack = (event) => console.log("on track called!");

    stream.getTracks().forEach((track) => {
      console.log(22, track);
      return pc.addTrack(track, stream);
    });

    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);

    console.log(68, pc.localDescription);

    const res = await fetch("http://localhost:4500/webrtc/offer", {
      method: "POST",
      body: JSON.stringify(pc.localDescription),
      headers: { "Content-Type": "application/json" },
    });

    const answer = await res.json();
    await pc.setRemoteDescription(answer);
  };

  return (
    <div>
      Sending Mic Audio...
      <button onClick={() => start()}>Start</button>
    </div>
  );
};
