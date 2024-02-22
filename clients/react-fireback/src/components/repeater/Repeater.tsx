import { uuidv4 } from "@/helpers/api";
import { source } from "@/helpers/source";
import { useT } from "@/hooks/useT";
import { FormikProps } from "formik";
import { osResources } from "../mulittarget/multitarget-resource";

export function Repeater({
  value,
  onChange,
  label,
  Component,
  form,
  disabled,
}: {
  value?: any[];
  onChange?: (value: Partial<any>[]) => void;
  label: string;
  Component: any;
  form: FormikProps<any>;
  disabled?: boolean;
}) {
  const t = useT();
  const readOnly = !!onChange;

  if (!Component) {
    return null;
  }

  return (
    <div>
      {(value || [])?.map((item, index) => (
        <div className="repeater-item">
          <div className="repeater-actions">
            <button
              disabled={disabled}
              onClick={(v) => {
                value = value?.filter((v, i) => i !== index);
                onChange && onChange([...(value || [])]);
              }}
              type="button"
              className="delete-btn"
            >
              <img src={source(osResources.delete)} />
            </button>
          </div>

          <div className="row repeater-element" key={item.uniqueId}>
            <Component disabled={disabled} form={form} index={index} />
          </div>
        </div>
      ))}
      <div className="repeater-end-actions">
        <button
          className="btn btn-primary"
          type="button"
          disabled={disabled}
          onClick={() =>
            onChange &&
            onChange([
              ...(value || []),
              {
                uniqueId: uuidv4(),
              },
            ])
          }
        >
          {label}
        </button>
      </div>
    </div>
  );
}
