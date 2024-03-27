import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { PostEntity } from "src/sdk/fireback/modules/cms/PostEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { useGetPostCategories } from "@/sdk/fireback/modules/cms/useGetPostCategories";
export const PostForm = ({ form, isEditing }: EntityFormProps<PostEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.title}
        onChange={(value) =>
          setFieldValue(PostEntity.Fields.title, value, false)
        }
        autoFocus={!isEditing}
        errorMessage={errors.title}
        label={t.posts.title}
        hint={t.posts.titleHint}
      />
      <FormEntitySelect3
        formEffect={{ form, field: PostEntity.Fields.category$ }}
        useQuery={useGetPostCategories}
        label={t.posts.category}
        hint={t.posts.categoryHint}
      />
    </>
  );
};
