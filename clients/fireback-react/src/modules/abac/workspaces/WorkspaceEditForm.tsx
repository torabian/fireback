import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";
import { WorkspaceEntity } from "src/sdk/fireback";
import { WorkspaceEntityFields } from "src/sdk/fireback/modules/workspaces/workspace-fields";
import { FormikProps } from "formik";

export const WorkspaceEditForm = ({
  form,
  isEditing,
}: {
  isEditing?: boolean;
  form: FormikProps<Partial<WorkspaceEntity>>;
}) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();

  return (
    <>
      <FormText
        value={values.name}
        autoFocus={!isEditing}
        onChange={(value) =>
          setFieldValue(WorkspaceEntityFields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.wokspaces.workspaceName}
        hint={t.wokspaces.workspaceNameHint}
      />
    </>
  );
};
