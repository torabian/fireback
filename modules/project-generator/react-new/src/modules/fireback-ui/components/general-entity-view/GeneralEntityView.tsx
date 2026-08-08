import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import { CopyCell } from "../entity-manager/CopyCell";

export interface GeneralEntityField {
  label: string;
  copyableContent?: string;
  elem: any;
}

export function GeneralEntityView({
  entity,
  fields,
  title,
  description,
}: {
  entity: any;
  fields?: Array<GeneralEntityField>;
  title?: string;
  description?: string;
}) {
  const d = entity;
  const s = useS(strings);
  return (
    <div className="mt-4">
      <div className="general-entity-view ">
        {title ? <h1>{title}</h1> : null}
        {description ? <p>{description}</p> : null}
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{s.table.info}</div>
          <div className="field-value">{s.table.value}</div>
        </div>

        {entity?.uniqueId && (
          <div className="entity-view-row entity-view-body">
            <div className="field-info">{s.table.uniqueId}</div>
            <div className="field-value">{entity.uniqueId}</div>
          </div>
        )}
        {(fields || [])?.map((field, index) => {
          let value = field.elem === undefined ? "-" : field.elem;

          if (field.elem === true) {
            value = s.common.yes;
          }

          if (field.elem === false) {
            value = s.common.no;
          }

          if (field.elem === null) {
            value = (
              <i>
                <b>{s.common.isNUll}</b>
              </i>
            );
          }

          return (
            <div key={index} className="entity-view-row entity-view-body">
              <div className="field-info">{field.label}</div>
              <div
                className="field-value"
                data-test-id={field.label?.toString() || ""}
              >
                {value} <CopyCell value={field.copyableContent || value} />
              </div>
            </div>
          );
        })}

        {entity?.createdFormatted && (
          <div className="entity-view-row entity-view-body">
            <div className="field-info">{s.table.created}</div>
            <div className="field-value">{entity.createdFormatted}</div>
          </div>
        )}
      </div>
    </div>
  );
}
