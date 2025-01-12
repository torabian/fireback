import { useEffect, useState } from "react";
import { Module3Dto } from "./defs";
import { FieldsEditor } from "./FieldsEditor";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";

interface DtoEditorProps {
  dto: Module3Dto;
  onChange: (dto: Module3Dto) => void;
}

export function DtoEditor({ dto, onChange }: DtoEditorProps) {
  const [data, setData$] = useState<Module3Dto>({});

  const setData = (params: Module3Dto) => {
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
