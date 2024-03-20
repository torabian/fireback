import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { TagEntity } from "src/sdk/fireback/modules/shop/TagEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
export const TagForm = ({ form, isEditing }: EntityFormProps<TagEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) => setFieldValue(TagEntity.Fields.name, value, false)}
        errorMessage={errors.name}
        label={t.tags.name}
        autoFocus={!isEditing}
        hint={t.tags.nameHint}
      />
    </>
  );
};
