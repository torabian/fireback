import { useEffect, useState } from "react";
import { Module2Entity } from "./defs";
import { FieldsEditor } from "./FieldsEditor";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";

interface EntityEditorProps {
  entity: Module2Entity;
  onChange: (entity: Module2Entity) => void;
}

export function EntityEditor({ entity, onChange }: EntityEditorProps) {
  const [data, setData$] = useState<Module2Entity>({});

  const setData = (params: Module2Entity) => {
    setData$(params);
    onChange(params);
  };

  useEffect(() => {
    setData(entity);
  }, [entity]);

  return (
    <div className="record-item">
      <h2 className="item-title">
        <div className="row">
          <div className="form-group col-md-4">
            <FormText
              label="Entity name"
              value={data.name}
              onChange={(v) => setData({ ...data, name: v })}
            />
          </div>
          <div className="form-group col-md-4">
            <FormText
              label="Cli name"
              value={data.cliName}
              placeholder={data.name}
              onChange={(v) => setData({ ...data, cliName: v })}
            />
          </div>
          <div className="form-group col-md-12">
            <FormRichText
              label="Description"
              onChange={(e) => setData({ ...data, description: e })}
              placeholder="Cli Description visible"
              value={data.description}
            ></FormRichText>
          </div>
        </div>
        <div>
          <FieldsEditor
            fields={data.fields}
            onChange={(e) => setData({ ...data, fields: e })}
          />
        </div>
      </h2>
    </div>
  );
}
