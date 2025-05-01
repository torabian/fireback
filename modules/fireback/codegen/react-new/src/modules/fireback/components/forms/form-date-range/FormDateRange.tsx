import {
  FormDateTimeRange,
  FormDateTimeRangeProps,
} from "../form-datetime-range/FormDateTimeRange";

export const FormDateRange = (props: FormDateTimeRangeProps) => {
  const normalizeDate = (d?: Date) =>
    d ? new Date(d.getFullYear(), d.getMonth(), d.getDate()) : undefined;

  const value = {
    startDate: normalizeDate(new Date(props.value?.startDate)),
    endDate: normalizeDate(new Date(props.value?.endDate)),
  };

  const onChange = (value) => {
    props.onChange({
      startDate: normalizeDate(value.startDate),
      endDate: normalizeDate(value.endDate),
    });
  };

  return <FormDateTimeRange {...props} value={value} onChange={onChange} />;
};
