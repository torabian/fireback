/**
 * Often it happens that we want to stream and Audio over socket.
 * Fireback reactive actions generate the necessary socket sdk for such operation,
 * and you can bind the custom handler after creating a sourceBuffer object
 * 
 * Assuming you have created a fireback 'reactive' action called 'audio',
 * you can use the following code to playback the stream
 * 
 * 
  const { operate } = useReactiveAudio({});
  
  useEffect(() => {
    playAudioStreamedOverSocket().then(({ sourceBuffer }) => {
      operate([], (evt) => {
        if (!sourceBuffer.updating) {
          sourceBuffer.appendBuffer(new Uint8Array(evt.data));
        }
      });
    });
  }, []);

 * 
 * @returns
 */
export async function playAudioStreamedOverSocket(): Promise<{
  sourceBuffer: SourceBuffer;
  mediaSource: MediaSource;
  audio: HTMLAudioElement;
}> {
  return new Promise((resolve, reject) => {
    const audio = new Audio();
    audio.autoplay = true;
    audio.controls = false;
    audio.style.display = "none";
    document.body.appendChild(audio);

    const mediaSource = new MediaSource();
    audio.src = URL.createObjectURL(mediaSource);

    mediaSource.addEventListener("sourceopen", () => {
      try {
        const mime = "audio/mpeg";
        const sourceBuffer = mediaSource.addSourceBuffer(mime);
        resolve({ audio, sourceBuffer, mediaSource });
      } catch (e) {
        reject(e);
      }
    });

    mediaSource.addEventListener("error", (e) => {
      reject(e);
    });
  });
}
