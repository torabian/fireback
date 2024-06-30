import classNames from "classnames";
import { useCallback, useRef, useState } from "react";
import PhoneInput from "react-phone-input-2";

import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import { useLocale } from "@/fireback/hooks/useLocale";

export interface FormTextProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: any) => void;

  secureTextEntry?: boolean;
  Icon?: any;
  dir?: string;
  errorMessage?: string;
  autoFocus?: boolean;
  validMessage?: string;
  value?: any | null;
  type?: "text" | "password" | "number" | "phonenumber" | "email";
  focused?: boolean;
  getInputRef?: (ref: any) => void;
  pattern?: string;
  children?: any;
}

// & React.InputHTMLAttributes<HTMLInputElement>;

export const FormText = (props: FormTextProps) => {
  const { region } = useLocale();
  const {
    placeholder,
    label,
    getInputRef,
    secureTextEntry,
    Icon,
    onChange,
    value,
    children,
    errorMessage,
    type,
    focused: f = false,
    autoFocus,
    ...restProps
  } = props;

  const [focused, setFocused] = useState(false);
  const ref = useRef<HTMLInputElement | null>();
  const onClick = useCallback(() => {
    ref.current?.focus();
  }, [ref.current]);

  let innerValue = value === undefined ? "" : value;

  if (type === "number") {
    innerValue = +value;
  }

  const onChangeHandler = (e: any) => {
    if (!onChange) {
      return;
    }
    if (type === "number") {
      onChange(+e.target.value);
    } else {
      onChange(e.target.value);
    }
  };

  return (
    <BaseFormElement focused={focused} onClick={onClick} {...props}>
      {props.type === "phonenumber" ? (
        <PhoneInput
          country={region}
          value={innerValue}
          containerClass="form-phone-input"
          onChange={(e) => onChange && onChange(e)}
        />
      ) : (
        <input
          {...restProps}
          ref={(el) => (ref.current = el)}
          value={innerValue}
          autoFocus={autoFocus}
          className={classNames(
            "form-control",
            props.errorMessage && "is-invalid",
            props.validMessage && "is-valid"
          )}
          type={type || "text"}
          onChange={onChangeHandler}
          onBlur={() => setFocused(false)}
          onFocus={() => setFocused(true)}
        />
      )}
      {children}
    </BaseFormElement>
  );
};
