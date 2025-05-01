import classNames from "classnames";
import { BaseFormElementProps } from "../base-form-element/BaseFormElement";
import { UseMutationResult } from "react-query";

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
  children?: React.ReactNode;
  mutation?: UseMutationResult<any, any, Partial<any>, any>;
}

export const FormButton = (
  props: React.ButtonHTMLAttributes<HTMLButtonElement> & FormButtonProps
) => {
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
    mutation,
    ...restProps
  } = props;

  const isLoading = mutation?.isLoading;

  return (
    <button
      onClick={props.onClick}
      type="submit"
      disabled={disabled || isLoading}
      className={classNames("btn mb-3", `btn-${type || "primary"}`, className)}
      {...props}
    >
      {props.children || props.label}
    </button>
  );
};
