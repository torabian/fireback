import { useCallback, useRef, useState } from "react";

import { useLocale } from "../../../hooks/useLocale";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";

export interface FormTimeProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: string) => void;
  secureTextEntry?: boolean;
  Icon?: any;
  dir?: string;
  errorMessage?: string;
  autoFocus?: boolean;
  validMessage?: string;
  value?: any | null;
  type?: "jalali" | "european";
  focused?: boolean;
  inputProps?: any;
  getInputRef?: (ref: any) => void;
  pattern?: string;
}

export const FormTime = (props: FormTimeProps) => {
  const { region } = useLocale();
  const {
    placeholder,
    label,
    getInputRef,
    secureTextEntry,
    Icon,
    onChange,
    value,
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

  return (
    <BaseFormElement focused={focused} onClick={onClick} {...props}>
      <input
        type="time"
        className="form-control"
        value={props.value}
        onChange={(e) => props.onChange && props.onChange(e.target.value)}
        {...props.inputProps}
      />
    </BaseFormElement>
  );
};
