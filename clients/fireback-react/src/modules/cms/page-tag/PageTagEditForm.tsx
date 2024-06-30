import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { PageTagEntity } from "src/sdk/fireback/modules/cms/PageTagEntity";
import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/fireback/components/forms/form-select/FormEntitySelect3";
export const PageTagForm = ({
  form,
  isEditing,
}: EntityFormProps<PageTagEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) =>
          setFieldValue(PageTagEntity.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.pagetags.name}
        hint={t.pagetags.nameHint}
      />
    </>
  );
};
