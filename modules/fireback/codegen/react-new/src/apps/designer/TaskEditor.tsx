import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { useEffect, useState } from "react";
import { Module3Task } from "./defs";
import { FieldsEditor } from "./FieldsEditor";

interface TaskEditorProps {
  task: Module3Task;
  onChange: (task: Module3Task) => void;
}

export function TaskEditor({ task, onChange }: TaskEditorProps) {
  const [data, setData$] = useState<Module3Task>({ name: "" });

  const setData = (params: Module3Task) => {
    setData$(params);
    onChange(params);
  };

  useEffect(() => {
    setData(task);
  }, [task]);

  return (
    <div className="record-item">
      <h2 className="item-title">
        <div className="row">
          <div className="form-group col-md-4">
            <FormText
              label="Task name"
              value={data.name}
              onChange={(v) => setData({ ...data, name: v })}
            />
          </div>
          <div className="form-group col-md-4">
            <FormText
              label="Task Input Dto"
              value={data.in?.dto}
              onChange={(v) => setData({ ...data, in: { ...data.in, dto: v } })}
            />
          </div>
          <div className="form-group col-md-4">
            <FormText
              label="Task Input Entity"
              value={data.in?.entity}
              onChange={(v) =>
                setData({ ...data, in: { ...data.in, entity: v } })
              }
            />
          </div>
        </div>
        <div>
          <FieldsEditor
            fields={data.in?.fields}
            onChange={(e) =>
              setData({ ...data, in: { ...(data.in || {}), fields: e } })
            }
          />
        </div>
      </h2>
    </div>
  );
}
