import { useEffect, useState } from "react";
import { Module2Field } from "./defs";
import { FormCheckbox } from "@/modules/fireback/components/forms/form-switch/FormSwitch";

interface FieldsEditorProps {
  fields?: Module2Field[];
  onChange: (Field: Module2Field[]) => void;
}

export function FieldsEditor({ fields, onChange }: FieldsEditorProps) {
  const [data, setData$] = useState<Module2Field[]>([]);

  const setData = (params: Module2Field[]) => {
    setData$(params);
    onChange(params);
  };

  const addField = () => {
    const n = [...data, { name: "", type: "string" }];
    setData(n);
    onChange(n);
  };

  const setFieldData = (index: number, fieldData: Module2Field) => {
    const n = data.map((x, ind) => {
      if (ind === index) {
        return { ...x, ...fieldData };
      }
      return x;
    });

    setData(n);
    onChange(n);
  };

  const removeField = (index: number) => {
    const n = data.filter((x, ind) => ind !== index);

    setData(n);
    onChange(n);
  };

  useEffect(() => {
    setData(fields || []);
  }, [fields]);

  return (
    <div>
      {data.map((field, index) => {
        return (
          <div className="form-editor-group">
            <div className="row">
              <div className="col-md-4">
                <div className="form-group">
                  <label>Field name</label>
                  <input
                    type="text"
                    className="form-control"
                    placeholder="Field name"
                    value={field.name}
                    onChange={(e) =>
                      setFieldData(index, { name: e.target.value })
                    }
                  />
                </div>
              </div>

              <div className="col-md-2">
                <div className="form-group">
                  <label>Type</label>
                  <select
                    className="form-control"
                    value={field.type}
                    onChange={(e) =>
                      setFieldData(index, { type: e.target.value })
                    }
                  >
                    <option value="string">string</option>
                    <option value="int64">int64</option>
                    <option value="float64">float64</option>
                    <option value="bool">bool</option>
                    <option value="one">one</option>
                    <option value="many2many">many2many</option>
                    <option value="array">array</option>
                    <option value="object">object</option>
                    <option value="html">html</option>
                  </select>
                </div>
              </div>

              {field.type === "string" || field.type === "html" ? (
                <div className="col-md-1">
                  <FormCheckbox
                    label="Translate"
                    value={field.translate}
                    onChange={(v) => setFieldData(index, { translate: v })}
                  ></FormCheckbox>
                </div>
              ) : null}

              {field.type === "many2many" || field.type === "one" ? (
                <div className="col-md-4">
                  <div className="form-group">
                    <label>Module</label>
                    <input
                      type="text"
                      className="form-control"
                      placeholder="Module (if external)"
                      value={field.module}
                      onChange={(e) =>
                        setFieldData(index, { module: e.target.value })
                      }
                    />
                  </div>
                  <div className="form-group">
                    <label>Target Entity</label>
                    <input
                      type="text"
                      className="form-control"
                      placeholder="Target entity"
                      value={field.module}
                      onChange={(e) =>
                        setFieldData(index, { target: e.target.value })
                      }
                    />
                  </div>
                </div>
              ) : null}

              <div className="col-md-1">
                <button
                  className="btn btn-danger"
                  onClick={() => removeField(index)}
                >
                  Del
                </button>
              </div>
            </div>
            {field.type === "array" || field.type === "one" ? (
              <div className="child-fields">
                <FieldsEditor
                  onChange={(fields) => setFieldData(index, { fields })}
                  fields={field.fields}
                />
              </div>
            ) : null}
          </div>
        );
      })}
      <button className="m-2 btn btn-sm btn-primary" onClick={() => addField()}>
        Add field
      </button>
    </div>
  );
}
