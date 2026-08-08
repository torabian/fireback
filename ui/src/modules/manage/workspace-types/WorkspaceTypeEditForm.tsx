import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { WorkspaceTypeDto } from "../../sdk/abac/WorkspaceTypeDto";

import { useContext } from "react";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { FormSelect } from "../../fireback-ui/components/forms/form-select/FormSelect";
import { useRolesQuerySource } from "../../fireback-ui/hooks/useRolesQuerySource";
import { FormRichText } from "../../fireback-ui/components/forms/form-richtext/FormRichText";

export const WorkspaceTypeEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceTypeDto>>) => {
  const { values, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const s = useS(strings);

  return (
    <>
      <FormText
        value={values.uniqueId}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeDto.Fields.uniqueId, value, false)
        }
        errorMessage={form.errors.uniqueId}
        label={s.workspaceTypeUniqueId}
        autoFocus={!isEditing}
        hint={s.workspaceTypeUniqueIdHint}
      />
      <FormText
        value={values.title}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeDto.Fields.title, value, false)
        }
        errorMessage={form.errors.title}
        label={s.workspaceTypeTitle}
        autoFocus={!isEditing}
        hint={s.workspaceTypeTitleHint}
      />
      <FormText
        value={values.slug}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeDto.Fields.slug, value, false)
        }
        errorMessage={form.errors.slug}
        label={s.workspaceTypeSlug}
        hint={s.workspaceTypeSlugHint}
      />
      <FormSelect
        label={s.roleFieldLabel}
        hint={s.roleFieldHint}
        fnLabelFormat={(role) => role.name}
        querySource={useRolesQuerySource}
        formEffect={{ form, field: "role" }}
        errorMessage={form.errors.roleId}
      />

      <FormRichText
        value={values.description}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeDto.Fields.description, value, false)
        }
        errorMessage={form.errors.description}
        label={s.typeDescription}
        hint={s.typeDescriptionHint}
      />
    </>
  );
};
