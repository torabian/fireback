import { useEffect, useRef } from "react";
declare const bodymovin: any;

export function AnimatedLogo(props: any) {
  const container = useRef(null);

  const renderAnimation = () => {
    const t = typeof bodymovin === "undefined";
    if (t) {
      return;
    }

    const anim = bodymovin.loadAnimation({
      container: container.current,
      renderer: "canvas",
      loop: props.loop || false,
      path: props.path,
    });

    anim.play();
  };

  useEffect(() => {
    const t = typeof bodymovin === "undefined";

    if (t === false) {
      renderAnimation();
      return () => {
        container.current = null;
      };
    }

    function onLottieLoad() {
      renderAnimation();
    }

    document.body.addEventListener("lottieLoadEvent", onLottieLoad, false);
    return () => {
      document.body.removeEventListener("lottieLoadEvent", onLottieLoad);
    };
  }, []);

  return <div ref={container} {...props}></div>;
}
