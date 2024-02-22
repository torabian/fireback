import { useEffect, useState } from "react";

export function useDomLoaded() {
  const [domLoaded, setDomLoaded] = useState(false);

  useEffect(() => {
    setDomLoaded(true);
  }, []);

  return domLoaded;
}
