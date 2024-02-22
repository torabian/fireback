import classNames from "classnames";
import { useCallback, useRef, useState } from "react";
import CurrencyInput from "react-currency-input-field";
import "react-phone-input-2/lib/style.css";

import { useLocale } from "@/hooks/useLocale";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";

export interface FormCurrencyProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: string) => void;
  secureCurrencyEntry?: boolean;
  Icon?: any;
  dir?: string;
  errorMessage?: string;
  autoFocus?: boolean;
  validMessage?: string;
  value?: any | null;
  type?: "Currency" | "password" | "number" | "phonenumber" | "email";
  focused?: boolean;
  getInputRef?: (ref: any) => void;
  pattern?: string;
}

// & React.InputHTMLAttributes<HTMLInputElement>;

export const FormCurrency = (props: FormCurrencyProps) => {
  const { region } = useLocale();
  const {
    placeholder,
    label,
    getInputRef,
    secureCurrencyEntry,
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
      <CurrencyInput
        placeholder={placeholder}
        value={value}
        decimalsLimit={2}
        className={classNames(
          "form-control",
          props.errorMessage && "is-invalid",
          props.validMessage && "is-valid"
        )}
        onValueChange={(value, name) => {
          onChange && onChange("" + value);
          console.log(value, name);
        }}
      />
    </BaseFormElement>
  );
};
