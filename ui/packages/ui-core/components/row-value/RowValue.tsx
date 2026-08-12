import React from "react";
import { type KeyValue } from "../../types/KeyValue";

export const RowValue = (props: KeyValue) => {
  return (
    <div>
      <span>{props.label}</span>
      <span>{props.value}</span>
    </div>
  );
};
