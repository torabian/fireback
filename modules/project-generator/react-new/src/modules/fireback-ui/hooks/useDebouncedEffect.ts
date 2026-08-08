import { useRef } from "react";

export const useDebouncedEffect = () => {
  const debounceRef = useRef<NodeJS.Timeout>();

  const withDebounce = (fn: () => void, delay: number) => {
    if (debounceRef.current) {
      clearTimeout(debounceRef.current);
    }

    debounceRef.current = setTimeout(fn, delay);
  };

  return { withDebounce };
};
