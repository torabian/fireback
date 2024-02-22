import React from "react";
import { KeyValue } from "@/definitions/definitions";

export const RowValue = (props: KeyValue) => {
  return (
    <div>
      <span>{props.label}</span>
      <span>{props.value}</span>
    </div>
  );
};
