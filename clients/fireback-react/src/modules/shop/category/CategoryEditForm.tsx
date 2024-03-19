import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { FormText } from "@/components/forms/form-text/FormText";
import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { useGetCategories } from "@/sdk/fireback/modules/shop/useGetCategories";
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
      <FormEntitySelect3
        formEffect={{ form, field: "parent" }}
        useQuery={useGetCategories}
        label={t.categories.parent}
        hint={t.categories.parentHint}
      />

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
