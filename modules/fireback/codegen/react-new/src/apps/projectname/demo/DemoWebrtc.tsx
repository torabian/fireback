import { useWebrtcMouseposition } from "@/modules/fireback/sdk/modules/fireback/useWebrtcMouseposition";
import { debounce } from "lodash";
import { useEffect } from "react";

export function DemoWebrtc() {
  const { dataChannel, state } = useWebrtcMouseposition({});

  const sendRamUsage = () => {
    const re = setInterval(() => {
      const data = (performance as any).memory;
      dataChannel.current?.ram.send(
        JSON.stringify({ usedJSHeapSize: data.usedJSHeapSize })
      );
    }, 1000);

    return () => {
      clearInterval(re);
    };
  };

  useEffect(() => {
    return sendRamUsage();
  }, []);

  return (
    <div>
      <br />
      <h2>Webrtc ({state})</h2>
      <p>Let's demo the web rtc here</p>
    </div>
  );
}
