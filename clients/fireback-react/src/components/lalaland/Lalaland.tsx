import { useState } from "react";
import BasicLayout from "./Grid";

export function Lalaland() {
  const [value, setValue] = useState<any>();

  return (
    <>
      <div>
        {/* <p>I put all experimental components here. Love you, Ali :)</p> */}
        {/* <NumberPicker canvasId="item1" onChange={(val) => setValue(val)} /> */}
        <BasicLayout />
        {/* <Chart value={value} /> */}
        <span>{value}</span>
      </div>
    </>
  );
}
