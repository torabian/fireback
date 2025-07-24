import { FormText } from "../../../components/forms/form-text/FormText";
import { EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../hooks/useT";
import { WorkspaceTypeEntity } from "../../../sdk/modules/abac/WorkspaceTypeEntity";

import { useContext } from "react";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { useGetRoles } from "../../../sdk/modules/abac/useGetRoles";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";

export const WorkspaceTypeEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceTypeEntity>>) => {
  const { values, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <FormText
        value={values.uniqueId}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeEntity.Fields.uniqueId, value, false)
        }
        errorMessage={form.errors.uniqueId}
        label={t.wokspaces.workspaceTypeUniqueId}
        autoFocus={!isEditing}
        hint={t.wokspaces.workspaceTypeUniqueIdHint}
      />
      <FormText
        value={values.title}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeEntity.Fields.title, value, false)
        }
        errorMessage={form.errors.title}
        label={t.wokspaces.workspaceTypeTitle}
        autoFocus={!isEditing}
        hint={t.wokspaces.workspaceTypeTitleHint}
      />
      <FormText
        value={values.slug}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeEntity.Fields.slug, value, false)
        }
        errorMessage={form.errors.slug}
        label={t.wokspaces.workspaceTypeSlug}
        hint={t.wokspaces.workspaceTypeSlugHint}
      />
      <FormSelect
        label={t.wokspaces.invite.role}
        hint={t.wokspaces.invite.roleHint}
        fnLabelFormat={(role) => role.name}
        querySource={useGetRoles}
        formEffect={{ form, field: WorkspaceTypeEntity.Fields.role$ }}
        errorMessage={form.errors.roleId}
      />

      <FormRichText
        value={values.description}
        onChange={(value) =>
          form.setFieldValue(
            WorkspaceTypeEntity.Fields.description,
            value,
            false
          )
        }
        errorMessage={form.errors.description}
        label={t.wokspaces.typeDescription}
        hint={t.wokspaces.typeDescriptionHint}
      />
    </>
  );
};
