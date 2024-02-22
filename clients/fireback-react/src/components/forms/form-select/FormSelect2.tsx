import AsyncSelect from "react-select/async";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import classNames from "classnames";
import { useT } from "@/hooks/useT";

export interface FormSelect2Props<T> extends BaseFormElementProps {
  label?: string;
  placeholder?: string;
  value?: T;
  errorMessage?: string;
  onChange?: (value: T) => void;
  iconName?: string;
  itemsInitializer?: () => void;
  iconColor?: string;
  options?: T[];
  testID?: string;
  fnLabelFormat?: (t: any) => string;
  keyExtractor?: (t: any) => string;
  fnLoadOptions?: (keword: string) => Promise<T[]>;
}

export function FormSelect2<T>(props: FormSelect2Props<T>) {
  const t = useT();
  return (
    <BaseFormElement {...props}>
      <AsyncSelect
        value={props.value}
        onChange={(newValue) => {
          props.onChange && props.onChange(newValue as any);
        }}
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
        onMenuOpen={props.itemsInitializer}
        defaultOptions={props.options}
        placeholder={t.searchplaceholder}
        noOptionsMessage={() => t.noOptions}
        getOptionValue={props.keyExtractor}
        formatOptionLabel={props.fnLabelFormat}
        loadOptions={props.fnLoadOptions}
      />
    </BaseFormElement>
  );
}
