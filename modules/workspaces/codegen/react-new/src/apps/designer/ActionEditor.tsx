import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { useEffect, useState } from "react";
import { Module3Action } from "./defs";

interface ActionEditorProps {
  action: Module3Action;
  onChange: (action: Module3Action) => void;
}

export function ActionEditor({ action, onChange }: ActionEditorProps) {
  const [data, setData$] = useState<Module3Action>({});

  const setData = (params: Module3Action) => {
    setData$(params);
    onChange(params);
  };

  useEffect(() => {
    setData(action);
  }, [action]);

  return (
    <div className="record-item">
      <h2 className="item-title">
        <div className="row">
          <div className="form-group col-md-4">
            <FormText
              label="Action name"
              value={data.name}
              onChange={(v) => setData({ ...data, name: v })}
            />
          </div>
          <div className="form-group col-md-4">
            <FormText
              label="Cli name"
              placeholder={data.name}
              value={data.cliName}
              onChange={(v) => setData({ ...data, cliName: v })}
            />
          </div>
        </div>
      </h2>
    </div>
  );
}
