import Select from "react-select";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import classNames from "classnames";
import { useT } from "@/hooks/useT";

export interface FormSelect3Props<T> extends BaseFormElementProps {
  label?: string;
  placeholder?: string;
  value?: T;
  errorMessage?: string;
  onChange: (value: T) => void;
  iconName?: string;
  itemsInitializer?: () => void;
  iconColor?: string;
  options?: T[];
  testID?: string;
  children?: any;
  fnLabelFormat?: (t: any) => string;
  onInputChange?: (t: string) => void;
  keyExtractor?: (t: any) => string;
  fnLoadOptions?: (keword: string) => Promise<T[]>;
  multiple?: boolean;
  convertToNative?: boolean;
}

export function FormSelect3<T>(props: FormSelect3Props<T>) {
  const t = useT();

  return (
    <BaseFormElement {...props}>
      {props.children}
      {props.convertToNative ? (
        <select
          onChange={(e) => {
            const item = props.options?.find(
              (t: any) => t.uniqueId === e.target.value
            ) as any;

            props.onChange(item);
          }}
          className={classNames(
            "form-select",
            props.errorMessage && "is-invalid",
            props.validMessage && "is-valid"
          )}
          disabled={props.disabled}
          aria-label="Default select example"
          value={
            typeof props?.value === "string"
              ? props.value
              : (props.value as any)?.uniqueId
          }
        >
          <option key={undefined} value={""}>
            {t.selectPlaceholder}
          </option>
          {props.options?.filter(Boolean).map((t: any) => (
            <option key={t.uniqueId} value={t.uniqueId}>
              {props.fnLabelFormat ? props.fnLabelFormat(t) : t.uniqueId}
            </option>
          ))}
        </select>
      ) : (
        <Select
          value={props.value}
          onChange={(newValue) => {
            props.onChange(newValue as any);
          }}
          isMulti={props.multiple}
          classNames={{
            container(propsx: any) {
              return classNames(
                props.errorMessage &&
                  " form-control form-control-no-padding is-invalid",
                props.validMessage && "is-valid"
              );
            },
            control(props2: any) {
              return classNames("form-control form-control-no-padding");
            },
            menu(props) {
              return "react-select-menu-area";
            },
          }}
          isSearchable
          onMenuOpen={props.itemsInitializer}
          options={props.options}
          placeholder={t.searchplaceholder}
          noOptionsMessage={() => t.noOptions}
          filterOption={null}
          getOptionValue={props.keyExtractor}
          formatOptionLabel={props.fnLabelFormat}
          onInputChange={props.onInputChange}
        />
      )}
    </BaseFormElement>
  );
}
