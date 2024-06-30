import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { PostCategoryEntity } from "src/sdk/fireback/modules/cms/PostCategoryEntity";
import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/fireback/components/forms/form-select/FormEntitySelect3";
export const PostCategoryForm = ({
  form,
  isEditing,
}: EntityFormProps<PostCategoryEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) =>
          setFieldValue(PostCategoryEntity.Fields.name, value, false)
        }
        errorMessage={errors.name}
        autoFocus={!isEditing}
        label={t.postcategories.name}
        hint={t.postcategories.nameHint}
      />
    </>
  );
};
