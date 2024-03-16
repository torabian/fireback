import { FormText } from "@/components/forms/form-text/FormText";
import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { useContext, useState } from "react";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";
import { FormBuilder } from "@ginkgo-bioworks/react-json-schema-form-builder";

export const CategoryForm = ({
  form,
  isEditing,
}: EntityFormProps<CategoryEntity>) => {
  const [data, setData] = useState({
    schema:
      '{"type":"object","properties":{"newInput1":{"title":"New Input 1","type":"string"},"newInput3":{"title":"New Input 3","type":"object"},"newInput2":{"title":"New Input 2","type":"string"}},"dependencies":{},"required":[]}',
    uischema: '{"ui:order":["newInput1","newInput3","newInput2"]}',
  });
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormBuilder
        schema={data.schema}
        uischema={data.uischema}
        onChange={(newSchema: string, newUiSchema: string) => {
          setData({
            schema: newSchema,
            uischema: newUiSchema,
          });
        }}
      />
      <pre>{JSON.stringify(data, null, 2)}</pre>
      <FormText
        value={values.name}
        onChange={(value) =>
          setFieldValue(CategoryEntity.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.categories.name}
        autoFocus={!isEditing}
        hint={t.categories.nameHint}
      />
    </>
  );
};
