import { usePostMouseposition } from "@/modules/fireback/sdk/modules/fireback/usePostMouseposition";
import { debounce } from "lodash";
import { useEffect } from "react";

export function DemoWebrtc() {
  const { dataChannel, state } = usePostMouseposition({});

  const handleMouseMove = debounce(
    (e: MouseEvent) => {
      try {
        const value = { x: e.clientX, y: e.clientY };
        console.log(value);
        dataChannel.current?.mouse.send(JSON.stringify(value));
      } catch (err) {}
    },
    100,
    { maxWait: 300 }
  );

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
    document.onmousemove = (e) => {
      handleMouseMove(e);
    };

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
