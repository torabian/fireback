import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { WorkspaceEntity } from "@/sdk/fireback/modules/workspaces/WorkspaceEntity";

export const WorkspaceEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceEntity>>) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();

  return (
    <>
      <FormText
        value={values.name}
        autoFocus={!isEditing}
        onChange={(value) =>
          setFieldValue(WorkspaceEntity.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.wokspaces.workspaceName}
        hint={t.wokspaces.workspaceNameHint}
      />
    </>
  );
};
