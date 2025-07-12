import classNames from "classnames";
import { useCallback, useEffect, useRef, useState } from "react";
import CurrencyInput from "react-currency-input-field";

import { useLocale } from "../../../hooks/useLocale";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";

export interface FormCurrencyProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: { amount: number; currency: string } | null) => void;
  secureCurrencyEntry?: boolean;
  Icon?: any;
  dir?: string;
  errorMessage?: string;
  autoFocus?: boolean;
  validMessage?: string;
  value?: { amount: number; currency: string } | null;
  focused?: boolean;
  getInputRef?: (ref: any) => void;
  pattern?: string;
}

// & React.InputHTMLAttributes<HTMLInputElement>;

export const FormCurrency = (props: FormCurrencyProps) => {
  const { placeholder, onChange, value, ...restProps } = props;

  const [inputValue, setInputValue] = useState(() =>
    value?.amount != null ? value.amount.toString() : ""
  );

  useEffect(() => {
    const externalStr = value?.amount != null ? value.amount.toString() : "";
    if (externalStr !== inputValue) {
      setInputValue(externalStr);
    }
  }, [value?.amount]);

  const commitValue = (str: string) => {
    const parsed = parseFloat(str);
    if (!isNaN(parsed)) {
      onChange?.({ ...value, amount: parsed });
    }
  };

  const handleChange = (val: string | undefined) => {
    const str = val || "";
    setInputValue(str);
    if (str.trim() !== "") {
      commitValue(str);
    }
  };

  return (
    <BaseFormElement {...restProps}>
      <div
        className="flex gap-2 items-center"
        style={{ flexDirection: "row", display: "flex" }}
      >
        <CurrencyInput
          placeholder={placeholder}
          value={inputValue}
          decimalsLimit={2}
          className={classNames(
            "form-control",
            props.errorMessage && "is-invalid",
            props.validMessage && "is-valid"
          )}
          onValueChange={handleChange}
        />
        <select
          value={value?.currency}
          onChange={(e) => onChange?.({ ...value, currency: e.target.value })}
          className="form-select w-24"
          style={{ width: "110px" }}
        >
          <option value="USD">USD</option>
          <option value="PLN">PLN</option>
          <option value="EUR">EUR</option>
        </select>
      </div>
    </BaseFormElement>
  );
};
