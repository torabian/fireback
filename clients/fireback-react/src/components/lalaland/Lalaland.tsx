import { useState } from "react";

export function Lalaland() {
  const [value, setValue] = useState<any>();

  return (
    <>
      <div>
        <span>{value}</span>
      </div>
    </>
  );
}
