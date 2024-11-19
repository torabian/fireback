import { useEffect, useState } from "react";

// Define types for the threshold and callback functions
type Threshold = { name: string; value: number };
type Callback = (thresholdName: string) => void;

function useResponsiveThresholds(
  selector: string,
  thresholds: Threshold[],
  onEnter?: Callback,
  onLeave?: Callback
): string | null {
  const [currentThreshold, setCurrentThreshold] = useState<string | null>(null);

  useEffect(() => {
    const element = document.querySelector(selector);
    if (!element) return;

    let lastThreshold: string | null = null;

    const resizeObserver = new ResizeObserver((entries) => {
      for (const entry of entries) {
        const width = entry.contentRect.width;
        let newThreshold: string | null = null;

        // Determine the current threshold based on width
        for (const { name, value } of thresholds) {
          if (width < value) {
            newThreshold = name;
            break;
          }
        }

        // If the new threshold is different from the last, trigger enter/leave
        if (newThreshold !== lastThreshold) {
          if (newThreshold && onEnter) onEnter(newThreshold); // Entering a new threshold
          if (lastThreshold && onLeave) onLeave(lastThreshold); // Leaving the last threshold
          setCurrentThreshold(newThreshold);
          lastThreshold = newThreshold;
        }
      }
    });

    resizeObserver.observe(element);

    // Clean up when component unmounts
    return () => {
      resizeObserver.unobserve(element);
      resizeObserver.disconnect();
    };
  }, [selector, thresholds, onEnter, onLeave]);

  return currentThreshold;
}

export default useResponsiveThresholds;
