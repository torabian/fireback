import { useState } from "react";
import { useQueryClient } from "react-query";
import { FormSelect2Props } from "./FormSelect2";
import { FormSelect3 } from "./FormSelect3";
import { FormikProps } from "formik";
import { get, set } from "lodash";

interface FormEntitySelect3Props<T> extends FormSelect2Props<T> {
  options?: any;
  useQuery: any;
  multiple?: boolean;
  convertToNative?: boolean;
  queryDSL?: string;
  formEffect?: { form: FormikProps<any>; field: string };
  withPreloads?: string;
}

export function FormEntitySelect3<T>(
  props: FormEntitySelect3Props<T> & {
    labelFn?: (item: T) => void;
  }
) {
  const queryClient = useQueryClient();
  let [keyword, setKeyword] = useState<string>("");

  const { query } = props.useQuery({
    queryClient,
    query: {
      itemsPerPage: 20,
      query: (props.queryDSL || `name %?%`).replaceAll("?", keyword),
      withPreloads: props.withPreloads,
    },
    queryOptions: {
      refetchOnWindowFocus: false,
    },
  });

  const onInputChange = (keyword: string) => {
    setKeyword(keyword);
  };

  const affects = props.formEffect
    ? (value: any) => {
        if (props.formEffect) {
          const newValue = {
            ...props.formEffect.form.values,
          };

          set(newValue, props.formEffect.field, value);

          if (props.multiple) {
            newValue[props.formEffect.field + "ListId"] = value.map(
              (t: any) => t.uniqueId
            );
            set(
              newValue,
              props.formEffect.field + "ListId",
              (value || []).map((t: any) => t.uniqueId)
            );
          } else {
            // newValue[props.formEffect.field + "Id"] = value.uniqueId;
            set(newValue, props.formEffect.field + "Id", value.uniqueId);
          }

          props.formEffect?.form.setValues(newValue);
        }
      }
    : undefined;

  const convertToNative =
    (query.data?.data?.totalAvailableItems < 10 || props.convertToNative) &&
    !props.multiple;
  const defaultLabel = (m: any) => m.name;
  const value = props.formEffect
    ? get(props.formEffect?.form.values, props.formEffect?.field || "")
    : props.value;

  const keyExtractor =
    typeof value === "string" ? (x: string) => x : (t: any) => t["uniqueId"];

  return (
    <div>
      <FormSelect3
        {...props}
        value={value}
        errorMessage={
          (props.formEffect?.form.errors[
            props.formEffect?.field + "Id"
          ] as any) || ""
        }
        convertToNative={convertToNative}
        onChange={affects || (props.onChange as any)}
        fnLabelFormat={(props.labelFn as any) || defaultLabel}
        keyExtractor={keyExtractor}
        onInputChange={onInputChange}
        fnLoadOptions={null as any}
        label={props.label}
        options={query?.data?.data?.items || []}
      />
    </div>
  );
}
