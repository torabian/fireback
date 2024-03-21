import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { PageTagEntity } from "src/sdk/fireback/modules/cms/PageTagEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
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
          value={values.name }
          onChange={(value) => setFieldValue(PageTagEntity.Fields.name, value, false)}
          errorMessage={errors.name }
          label={t.pagetags.name }
          hint={t.pagetags.nameHint}
        />
    </>
  );
};