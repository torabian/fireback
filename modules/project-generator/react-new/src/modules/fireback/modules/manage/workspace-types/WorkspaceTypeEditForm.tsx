import { FormText } from "../../../components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../hooks/useT";
import { WorkspaceTypeDto } from "../../../sdk/abac/WorkspaceTypeDto";

import { useContext } from "react";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { useRolesQuerySource } from "../../../hooks/useRolesQuerySource";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";

export const WorkspaceTypeEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceTypeDto>>) => {
  const { values, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <pre>{JSON.stringify(form.errors)}</pre>
      <FormText
        value={values.uniqueId}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeDto.Fields.uniqueId, value, false)
        }
        errorMessage={form.errors.uniqueId}
        label={t.wokspaces.workspaceTypeUniqueId}
        autoFocus={!isEditing}
        hint={t.wokspaces.workspaceTypeUniqueIdHint}
      />
      <FormText
        value={values.title}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeDto.Fields.title, value, false)
        }
        errorMessage={form.errors.title}
        label={t.wokspaces.workspaceTypeTitle}
        autoFocus={!isEditing}
        hint={t.wokspaces.workspaceTypeTitleHint}
      />
      <FormText
        value={values.slug}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeDto.Fields.slug, value, false)
        }
        errorMessage={form.errors.slug}
        label={t.wokspaces.workspaceTypeSlug}
        hint={t.wokspaces.workspaceTypeSlugHint}
      />
      <FormSelect
        label={t.wokspaces.invite.role}
        hint={t.wokspaces.invite.roleHint}
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
        label={t.wokspaces.typeDescription}
        hint={t.wokspaces.typeDescriptionHint}
      />
    </>
  );
};
