import { useState, useEffect } from "react";

function useLocalStorageState<T>(key: string, defaultValue: T) {
  const [state, setState] = useState<T>(() => {
    try {
      const storedValue = localStorage.getItem(key);
      if (storedValue === null) return defaultValue; // No stored value, use default
      return JSON.parse(storedValue) as T;
    } catch (error) {
      console.error(`Error parsing localStorage key "${key}":`, error);
      return defaultValue; // Fallback to default value
    }
  });

  useEffect(() => {
    try {
      localStorage.setItem(key, JSON.stringify(state));
    } catch (error) {
      console.error(`Error saving to localStorage key "${key}":`, error);
    }
  }, [key, state]);

  return [state, setState] as const;
}

export default useLocalStorageState;
