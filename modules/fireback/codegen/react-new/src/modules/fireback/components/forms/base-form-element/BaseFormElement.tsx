import classNames from "classnames";
import React from "react";

export interface CommonFormElementProps<T> {
  value?: T;
  onChange?: (value: T) => void;
  disabled?: boolean;
  label?: string;
}

export interface BaseFormElementProps extends CommonFormElementProps<any> {
  Icon?: any;
  errorMessage?: string;
  validMessage?: string;
  className?: string;
  value?: any | null;
  hasAnimation?: boolean;
  focused?: boolean;
  hint?: string;
  onClick?: () => void;
  getInputRef?: (ref: any) => void;
  children?: React.ReactNode;
  displayValue?: string | null;
}

function errorMessageAsString(possibleErrorMessage: string | Object) {
  if (typeof possibleErrorMessage === "string") {
    return possibleErrorMessage;
  }

  if (Array.isArray(possibleErrorMessage)) {
    return possibleErrorMessage.join(", ");
  }

  return JSON.stringify(possibleErrorMessage);
}

export const BaseFormElement = ({
  label,
  getInputRef,
  displayValue,
  Icon,
  children,
  errorMessage,
  validMessage,
  value,
  hint,
  onClick,
  onChange,
  className,
  focused = false,
  hasAnimation,
}: BaseFormElementProps) => {
  return (
    <div
      style={{ position: "relative" }}
      className={classNames("mb-3", className)}
    >
      {label && <label className="form-label">{label}</label>}
      {children}

      <div className="form-text">{hint}</div>
      <div className="invalid-feedback">{errorMessage}</div>
      <div className="valid-feedback">{validMessage}</div>
    </div>
  );
};
