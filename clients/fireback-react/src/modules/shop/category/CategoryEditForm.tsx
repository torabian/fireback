import { FormText } from "@/components/forms/form-text/FormText";
import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { useContext, useState } from "react";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";

export const CategoryForm = ({
  form,
  isEditing,
}: EntityFormProps<CategoryEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
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
