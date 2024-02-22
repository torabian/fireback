import classNames from "classnames";
import { BaseFormElementProps } from "../base-form-element/BaseFormElement";

export interface FormButtonProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: string) => void;
  secureTextEntry?: boolean;
  Icon?: any;
  errorMessage?: string;
  isSubmitting?: boolean;
  value?: any | null;
  focused?: boolean;
  type?: "primary" | "secondary";
  getInputRef?: (ref: any) => void;
}

export const FormButton = (props: FormButtonProps) => {
  const {
    placeholder,
    label,
    getInputRef,
    secureTextEntry,
    Icon,
    isSubmitting,
    errorMessage,
    onChange,
    value,
    disabled,
    type,
    focused: f = false,
    className,
    ...restProps
  } = props;

  return (
    <button
      onClick={props.onClick}
      type="submit"
      disabled={disabled}
      className={classNames("btn mb-3", `btn-${type || "primary"}`, className)}
    >
      {props.label}
    </button>
  );
};
