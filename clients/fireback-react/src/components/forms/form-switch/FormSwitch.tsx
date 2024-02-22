import classNames from "classnames";
import { useCallback, useRef, useState } from "react";
import "react-phone-input-2/lib/style.css";

import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";

export interface FormCheckboxProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: boolean) => void;
  secureTextEntry?: boolean;
  Icon?: any;
  errorMessage?: string;
  autoFocus?: boolean;
  validMessage?: string;
  value?: boolean | null;
  type?: "text" | "password" | "number" | "phonenumber";
  focused?: boolean;
  getInputRef?: (ref: any) => void;
}

export const FormCheckbox = (props: FormCheckboxProps) => {
  const {
    placeholder,
    label,
    getInputRef,
    secureTextEntry,
    Icon,
    onChange,
    value,
    disabled,
    focused: f = false,
    errorMessage,
    autoFocus,
    ...restProps
  } = props;

  const [focused, setFocused] = useState(false);
  const ref = useRef<HTMLInputElement | null>();
  const onClick = useCallback(() => {
    ref.current?.focus();
  }, [ref.current]);

  return (
    <BaseFormElement focused={focused} onClick={onClick} {...props} label="">
      <label className="form-label mr-2">
        <input
          {...restProps}
          ref={(el) => (ref.current = el)}
          checked={!!value}
          type={"checkbox"}
          onChange={(e) => onChange && onChange(!value)}
          onBlur={() => setFocused(false)}
          onFocus={() => setFocused(true)}
          className="form-checkbox"
        />
        {label}
      </label>
    </BaseFormElement>
  );
};
