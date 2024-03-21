import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { PageEntity } from "src/sdk/fireback/modules/cms/PageEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { useGetPageCategories } from "@/sdk/fireback/modules/cms/useGetPageCategories";
export const PageForm = ({ form, isEditing }: EntityFormProps<PageEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.title}
        onChange={(value) =>
          setFieldValue(PageEntity.Fields.title, value, false)
        }
        errorMessage={errors.title}
        label={t.pages.title}
        hint={t.pages.titleHint}
      />
      <FormEntitySelect3
        formEffect={{ form, field: PageEntity.Fields.category$ }}
        useQuery={useGetPageCategories}
        label={t.pages.category}
        hint={t.pages.categoryHint}
      />
    </>
  );
};
