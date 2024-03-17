import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext, useState } from "react";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import FormBuilder from "@/thirdparty/form-builder/FormBuilder";

export const ProductForm = ({
  form,
  isEditing,
}: EntityFormProps<ProductEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();

  const [data, setData] = useState({
    schema:
      '{"type":"object","properties":{"newInput1":{"title":"New Input 1","type":"string"},"newInput3":{"title":"New Input 3","type":"object"},"newInput2":{"title":"New Input 2","type":"string"}},"dependencies":{},"required":[]}',
    uischema: '{"ui:order":["newInput1","newInput3","newInput2"]}',
  });

  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) =>
          setFieldValue(ProductEntity.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.products.name}
        hint={t.products.nameHint}
      />
      <FormText
        value={values.description}
        onChange={(value) =>
          setFieldValue(ProductEntity.Fields.description, value, false)
        }
        errorMessage={errors.description}
        label={t.products.description}
        hint={t.products.descriptionHint}
      />
      <FormBuilder
        mods={{ showFormHead: false }}
        schema={JSON.stringify(values.jsonSchema)}
        uischema={JSON.stringify(values.uiSchema)}
        onChange={(newSchema: string, newUiSchema: string) => {
          setValues({
            ...values,
            jsonSchema: newSchema ? JSON.parse(newSchema) : {},
            uiSchema: newUiSchema ? JSON.parse(newUiSchema) : {},
          });
        }}
      />
    </>
  );
};
