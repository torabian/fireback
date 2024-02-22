import { IndeterminateCheck } from "@/definitions/definitions";
import { HTMLProps, useEffect, useRef } from "react";

export function Checkbox({
  value,
  onChange,
  ...props
}: {
  onChange: (value: IndeterminateCheck) => void;
  value: IndeterminateCheck;
} & HTMLProps<HTMLInputElement>) {
  const cRef = useRef<any>();

  useEffect(() => {
    cRef.current.indeterminate = value === "indeterminate";
  }, [cRef, value]);

  return (
    <input
      {...props}
      type="checkbox"
      ref={cRef}
      onChange={(e) => {
        onChange("checked");
      }}
      checked={value === "checked"}
      className="form-check-input"
    />
  );
}
