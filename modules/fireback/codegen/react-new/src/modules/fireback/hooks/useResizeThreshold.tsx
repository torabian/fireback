import { useEffect, useRef } from "react";

export function useResizeThreshold(threshold, callback, debounceDelay = 100) {
  const lastWidth = useRef(window.innerWidth);

  useEffect(() => {
    let timeout;

    const checkThreshold = () => {
      const currentWidth = window.innerWidth;
      const isBelow = currentWidth < threshold;

      if (
        (lastWidth.current >= threshold && isBelow) || // Crossed from above to below
        (lastWidth.current < threshold && !isBelow) // Crossed from below to above
      ) {
        callback(isBelow);
      }

      lastWidth.current = currentWidth;
    };

    const debouncedResize = () => {
      clearTimeout(timeout);
      timeout = setTimeout(checkThreshold, debounceDelay);
    };

    window.addEventListener("resize", debouncedResize);

    // Initial check
    checkThreshold();

    // Cleanup
    return () => {
      window.removeEventListener("resize", debouncedResize);
      clearTimeout(timeout);
    };
  }, [threshold, callback, debounceDelay]);
}
