import { useEffect, useState } from "react";
import { Module3Config } from "./defs";
import { FieldsEditor } from "./FieldsEditor";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";

interface ConfigEditorProps {
  config: Module3Config;
  onChange: (config: Module3Config) => void;
}

export function ConfigEditor({ config, onChange }: ConfigEditorProps) {
  const [data, setData$] = useState<Module3Config>({});

  const setData = (params: Module3Config) => {
    setData$(params);
    onChange(params);
  };

  useEffect(() => {
    setData(config);
  }, [config]);

  return (
    <div className="record-item">
      <h2 className="item-title">
        <div className="row">
          <div className="form-group col-md-4">
            <FormText
              label="Config name"
              value={data.name}
              onChange={(v) => setData({ ...data, name: v })}
            />
          </div>
          <div className="form-group col-md-12">
            <FormRichText
              label="Description"
              onChange={(e) => setData({ ...data, description: e })}
              placeholder="Description of purpose of the config"
              value={data.description}
            ></FormRichText>
          </div>
          <div className="col-md-2">
            <div className="form-group">
              <label>Type</label>
              <select
                className="form-control"
                value={data.type}
                onChange={(e) => setData({ ...data, type: e.target.value })}
              >
                <option value="string">string</option>
                <option value="int64">int64</option>
                <option value="float64">float64</option>
                <option value="bool">bool</option>
                <option value="html">html</option>
              </select>
            </div>
          </div>
        </div>
      </h2>
    </div>
  );
}
