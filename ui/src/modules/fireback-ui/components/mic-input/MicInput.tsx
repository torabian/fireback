import React, { useState, useEffect, useRef } from "react";

/**
 * Example of the microphone input for react, which can be hooked into
 * reactive (websocket) functions to stream live audio
 * @returns
 */

export const MicInput = ({
  onData,
}: {
  onData: (data: ArrayBuffer) => void;
}) => {
  const [isStreaming, setIsStreaming] = useState(false);
  const audioContextRef = useRef<AudioContext | null>(null);
  const streamRef = useRef<MediaStream | null>(null);
  const processorRef = useRef<ScriptProcessorNode | null>(null);

  const startStreaming = async () => {
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
    streamRef.current = stream;

    const audioContext = new (window.AudioContext ||
      (window as any).webkitAudioContext)();
    audioContextRef.current = audioContext;

    const source = audioContext.createMediaStreamSource(stream);
    const processor = audioContext.createScriptProcessor(4096, 1, 1);
    processorRef.current = processor;

    processor.onaudioprocess = (e) => {
      const input = e.inputBuffer.getChannelData(0);
      onData(input.buffer.slice(0)); // slice to ensure it's a copy
    };

    source.connect(processor);
    processor.connect(audioContext.destination);

    setIsStreaming(true);
  };

  const stopStreaming = () => {
    processorRef.current?.disconnect();
    audioContextRef.current?.close();
    streamRef.current?.getTracks().forEach((t) => t.stop());

    processorRef.current = null;
    audioContextRef.current = null;
    streamRef.current = null;

    setIsStreaming(false);
  };

  return (
    <div>
      <h1>Microphone Audio Streaming</h1>
      <button
        className="btn btn-secondary"
        type="button"
        onClick={isStreaming ? stopStreaming : startStreaming}
      >
        {isStreaming ? "Stop Streaming" : "Start Streaming"}
      </button>
    </div>
  );
};
