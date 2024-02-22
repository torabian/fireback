import { useEffect, useState } from "react";
import { FormSelect2, FormSelect2Props } from "./FormSelect2";

interface FormEntitySelect2Props<T> extends FormSelect2Props<T> {
  options?: any;
  fnLoadOptions?: (keword: string) => Promise<T[]>;
}

export function FormEntitySelect2<T>(
  props: FormEntitySelect2Props<T> & {
    labelFn?: (item: any) => void;
  }
) {
  const [options, setOptions] = useState<any>([]);
  // useEffect(() => {
  //   if (props.fnLoadOptions)
  //     props.fnLoadOptions("").then((items) => {
  //       setOptions(items);
  //     });
  // }, []);
  const itemsInitializer = () => {
    if (props.fnLoadOptions)
      props.fnLoadOptions("").then((items) => {
        setOptions(items);
      });
  };

  return (
    <div>
      <FormSelect2
        {...props}
        fnLabelFormat={props.labelFn as any}
        placeholder="asda"
        keyExtractor={(t) => t["uniqueId"]}
        fnLoadOptions={props.fnLoadOptions}
        itemsInitializer={itemsInitializer}
        label={props.label}
        options={options}
      />
    </div>
  );
}
