import { useCallback, useRef, useState } from "react";
import { DateRangePicker } from "react-date-range";
import { enUS } from "react-date-range/dist/locale";

import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";

import "react-date-range/dist/styles.css"; // main style file
import "react-date-range/dist/theme/default.css"; // theme css file

export interface FormDateTimeRangeProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: {
    startDate?: string | Date | null;
    endDate?: string | Date | null;
  }) => void;
  secureTextEntry?: boolean;
  Icon?: any;
  dir?: string;
  errorMessage?: string;
  autoFocus?: boolean;
  validMessage?: string;
  value?: {
    startDate?: string | Date | null;
    endDate?: string | Date | null;
  };
  type?: "jalali" | "european";
  focused?: boolean;
  inputProps?: any;
  getInputRef?: (ref: any) => void;
  pattern?: string;
}

export const FormDateTimeRange = (props: FormDateTimeRangeProps) => {
  const {
    placeholder,
    label,
    getInputRef,
    secureTextEntry,
    Icon,
    onChange,
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
      <div>
        <DateRangePicker
          locale={enUS}
          date={props.value}
          months={2}
          showSelectionPreview={true}
          direction="horizontal"
          moveRangeOnFirstSelection={false}
          ranges={[{ ...(props.value || {}), key: "selection" }]}
          onChange={(value) => {
            props.onChange?.(value.selection);
          }}
        />
      </div>
    </BaseFormElement>
  );
};
