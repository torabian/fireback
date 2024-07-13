import { useT } from "@/modules/fireback/hooks/useT";
import { get } from "lodash";

export const ChildForm = ({ form, Component, part }: any) => {
  const { setFieldValue } = form;

  // type Mv = keyof typeof TemperatureHmiComponentDto.Fields;

  const setChildValue = (field: any, value: any, shouldValidate: boolean) => {
    setFieldValue(part + "." + field, value, shouldValidate);
  };
  const data = get(form.values, part) || {};
  const errors = form.errors;
  const t = useT();

  return (
    <Component
      t={t}
      data={data}
      errors={errors}
      setFieldValue={setChildValue}
      form={form}
    />
  );
};

type ConvertKeysToString<T> = { [K in keyof T]: string };

export interface ParialFormProps<T, K> {
  setFieldValue: (field: K, value: any, validate?: boolean) => void;
  data: T;
  errors: ConvertKeysToString<T>;
  t: any;
}
