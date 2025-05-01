import { useState } from "react";

export function useArrayValue<T = any>() {
  const [items, setItems] = useState<T[]>([]);

  const append = () => {
    setItems((items: any) => {
      return [...items, {}];
    });
  };

  const remove = (index: number) => {
    setItems((data) => {
      return data.filter((l, index2) => {
        if (index === index2) {
          return false;
        }
        return true;
      });
    });
  };

  const update = (item: T, index: number) => {
    setItems((data) => {
      return data.map((l, index2) => {
        if (index === index2) {
          return item;
        }
        return l;
      });
    });
  };

  return { items, append, remove, update };
}
