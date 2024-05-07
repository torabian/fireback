import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { PageEntity } from "src/sdk/fireback/modules/cms/PageEntity";
import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/fireback/components/forms/form-select/FormEntitySelect3";
import { useGetPageCategories } from "@/sdk/fireback/modules/cms/useGetPageCategories";
import { FormRichText } from "@/fireback/components/forms/form-richtext/FormRichText";
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
        autoFocus={!isEditing}
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
      <FormRichText
        forceRich
        label={t.pages.content}
        hint={t.pages.content}
        onChange={(val) => form.setFieldValue(PageEntity.Fields.content, val)}
        value={form.values.content}
      />
    </>
  );
};
