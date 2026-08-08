import { FormText } from "../../../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { WorkspaceDto } from "../../../sdk/abac/WorkspaceDto";

export const WorkspaceEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceDto>>) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();

  return (
    <>
      <FormText
        value={values.name}
        autoFocus={!isEditing}
        onChange={(value) =>
          setFieldValue(WorkspaceDto.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.wokspaces.workspaceName}
        hint={t.wokspaces.workspaceNameHint}
      />
    </>
  );
};
