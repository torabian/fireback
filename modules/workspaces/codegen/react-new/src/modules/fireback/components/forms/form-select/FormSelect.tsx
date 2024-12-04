import { KeyValue } from "../../../definitions/definitions";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import classNames from "classnames";
import { useT } from "../../../hooks/useT";

export interface FormSelectProps extends BaseFormElementProps {
  label?: string;
  placeholder?: string;
  value?: any;
  errorMessage?: string;
  onChange: (value: string | number) => void;
  iconName?: string;
  iconColor?: string;
  options?: KeyValue[];
  testID?: string;
  name?: string;
  type?: string;
  valueType?: "string" | "number";
}

function VerboseSelect(props: FormSelectProps) {
  return (
    <BaseFormElement {...props}>
      <div className="form-select-verbos">
        {props.options?.map((option) => {
          return (
            <label key={option.value}>
              <input
                name={props.name}
                type="radio"
                onClick={(t) => {
                  let v: any = option.value;

                  if (props.valueType === "number") {
                    v = +v;
                  }

                  props.onChange(v);
                }}
                value={option.value}
                checked={option.value === props.value}
              />
              {option.label}
            </label>
          );
        })}
      </div>
    </BaseFormElement>
  );
}

export function FormSelect(props: FormSelectProps) {
  const t = useT();
  if (props.type === "verbose") {
    return <VerboseSelect {...props} />;
  }

  return (
    <BaseFormElement {...props}>
      <select
        onChange={(t) => {
          let v: any = t.target.value;

          if (props.valueType === "number") {
            v = +v;
          }

          props.onChange(v);
        }}
        className={classNames(
          "form-select",
          props.errorMessage && "is-invalid",
          props.validMessage && "is-valid"
        )}
        disabled={props.disabled}
        aria-label="Default select example"
        value={props.value}
      >
        <option key={undefined} value={""}>
          {t.selectPlaceholder}
        </option>
        {props.options?.map((t) => (
          <option key={t.value} value={t.value}>
            {t.label}
          </option>
        ))}
      </select>
    </BaseFormElement>
  );
}
