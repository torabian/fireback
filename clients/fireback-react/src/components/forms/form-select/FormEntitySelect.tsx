/**
 * This is a select input, which can query specific type of entity
 */

import { UseQueryResult } from "react-query";
import { FormSelect, FormSelectProps } from "./FormSelect";

interface FormEntitySelectProps extends FormSelectProps {
  options?: any;
}

export function FormEntitySelect(
  props: FormEntitySelectProps & {
    query: UseQueryResult<any, unknown>;
    labelFn?: (item: any) => void;
  }
) {
  // @todo : not always it's name, make a mechanism in backend to have the generalLabel
  const options = (props.query.data?.data?.items || []).map((t: any) => ({
    label: props.labelFn ? props.labelFn(t) : t.name,
    value: t.uniqueId,
  }));

  return (
    <div>
      <FormSelect {...props} options={options} />
    </div>
  );
}
