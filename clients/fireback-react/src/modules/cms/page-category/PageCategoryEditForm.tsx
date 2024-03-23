import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { PageCategoryEntity } from "src/sdk/fireback/modules/cms/PageCategoryEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
export const PageCategoryForm = ({
  form,
  isEditing,
}: EntityFormProps<PageCategoryEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
        <FormText
          value={values.name }
          onChange={(value) => setFieldValue(PageCategoryEntity.Fields.name, value, false)}
          errorMessage={errors.name }
          label={t.pagecategories.name }
          hint={t.pagecategories.nameHint}
        />
    </>
  );
};