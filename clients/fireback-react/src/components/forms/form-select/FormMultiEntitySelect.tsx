import { useEffect, useState } from "react";
import { FormMultiSelect, FormMultiSelectProps } from "./FormMultiSelect";

interface FormMultiEntitySelectProps<T> extends FormMultiSelectProps<T> {
  options?: any;
  fnLoadOptions?: (keword: string) => Promise<T[]>;
}

export function FormMultiEntitySelect<T>(
  props: FormMultiEntitySelectProps<T> & {
    labelFn?: (item: any) => void;
  }
) {
  const [options, setOptions] = useState<any>([]);
  useEffect(() => {
    if (props.fnLoadOptions)
      props.fnLoadOptions("").then((items) => {
        setOptions(items);
      });
  }, []);
  return (
    <div>
      <FormMultiSelect
        {...props}
        fnLabelFormat={props.labelFn as any}
        keyExtractor={(t) => t["uniqueId"]}
        fnLoadOptions={props.fnLoadOptions}
        label={props.label}
        options={options}
      />
    </div>
  );
}
