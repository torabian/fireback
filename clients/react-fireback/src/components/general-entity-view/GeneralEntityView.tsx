import { useT } from "@/hooks/useT";

export interface GeneralEntityField {
  label: string;
  elem: any;
}

export function GeneralEntityView({
  entity,
  fields,
}: {
  entity: any;
  fields?: Array<GeneralEntityField>;
}) {
  const d = entity;
  const t = useT();
  return (
    <div className="mt-4">
      <div className="general-entity-view ">
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{t.table.info}</div>
          <div className="field-value">{t.table.value}</div>
        </div>

        {entity?.uniqueId && (
          <div className="entity-view-row entity-view-body">
            <div className="field-info">{t.table.uniqueId}</div>
            <div className="field-value">{entity.uniqueId}</div>
          </div>
        )}
        {(fields || [])?.map((field, index) => {
          let value = field.elem === undefined ? "-" : field.elem;

          if (field.elem === true) {
            value = t.common.yes;
          }

          if (field.elem === false) {
            value = t.common.no;
          }

          return (
            <div key={index} className="entity-view-row entity-view-body">
              <div className="field-info">{field.label}</div>
              <div className="field-value">{value}</div>
            </div>
          );
        })}

        {entity?.createdFormatted && (
          <div className="entity-view-row entity-view-body">
            <div className="field-info">{t.table.created}</div>
            <div className="field-value">{entity.createdFormatted}</div>
          </div>
        )}
      </div>
    </div>
  );
}
