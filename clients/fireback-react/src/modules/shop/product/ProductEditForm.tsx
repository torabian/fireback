import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext, useState } from "react";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";
import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/fireback/components/forms/form-select/FormEntitySelect3";
import FormBuilder from "@/thirdparty/form-builder/FormBuilder";

export const ProductForm = ({
  form,
  isEditing,
}: EntityFormProps<ProductEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();

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
        autoFocus={!isEditing}
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
        schema={JSON.stringify(values.jsonSchema || {})}
        uischema={JSON.stringify(values.uiSchema || {})}
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
