import { useEffect, useState } from "react";
import { Module2Dto } from "./defs";
import { FieldsEditor } from "./FieldsEditor";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";

interface DtoEditorProps {
  dto: Module2Dto;
  onChange: (dto: Module2Dto) => void;
}

export function DtoEditor({ dto, onChange }: DtoEditorProps) {
  const [data, setData$] = useState<Module2Dto>({});

  const setData = (params: Module2Dto) => {
    setData$(params);
    onChange(params);
  };

  useEffect(() => {
    setData(dto);
  }, [dto]);

  return (
    <div className="record-item">
      <h2 className="item-title">
        <div className="row">
          <div className="form-group col-md-4">
            <FormText
              label="Dto name"
              value={data.name}
              onChange={(v) => setData({ ...data, name: v })}
            />
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
