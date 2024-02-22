import { KeyValue } from "@/definitions/definitions";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import AsyncSelect from "react-select/async";

export interface FormMultiSelectProps<T> extends BaseFormElementProps {
  label: string;
  placeholder?: string;
  value?: T[];
  errorMessage?: string;
  onChange: (value: T[]) => void;
  iconName?: string;
  iconColor?: string;
  options?: T[];
  testID?: string;
  fnLabelFormat?: (t: any) => string;
  keyExtractor?: (t: any) => string;
  fnLoadOptions?: (keword: string) => Promise<T[]>;
}

export function FormMultiSelect<T>(props: FormMultiSelectProps<T>) {
  return (
    <BaseFormElement {...props}>
      <AsyncSelect
        isMulti
        value={props.value}
        onChange={(newValue) => {
          props.onChange((newValue || []) as any[]);
        }}
        defaultOptions={props.options}
        getOptionValue={props.keyExtractor}
        formatOptionLabel={props.fnLabelFormat}
        loadOptions={props.fnLoadOptions}
      />
    </BaseFormElement>
  );
}
