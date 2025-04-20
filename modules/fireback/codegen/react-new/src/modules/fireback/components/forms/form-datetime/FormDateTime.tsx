import { useCallback, useRef, useState } from "react";

import { useLocale } from "../../../hooks/useLocale";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import "react-datetime/css/react-datetime.css";
import Datetime from "react-datetime";
import moment from "moment";

export interface FormDateTimeProps extends BaseFormElementProps {
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

export const FormDateTime = (props: FormDateTimeProps) => {
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
      <Datetime
        value={moment(props.value)}
        onChange={(e) => props.onChange && props.onChange(e as string)}
        {...props.inputProps}
      />
    </BaseFormElement>
  );
};
