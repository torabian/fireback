import { FormText } from "../../../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../../fireback-ui/types/EntityManagement";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { WorkspaceDto } from "../../../sdk/abac/WorkspaceDto";

export const WorkspaceEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceDto>>) => {
  const { values, setFieldValue, errors } = form;
  const s = useS(strings);

  return (
    <>
      <FormText
        value={values.name}
        autoFocus={!isEditing}
        onChange={(value) =>
          setFieldValue(WorkspaceDto.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={s.workspaceName}
        hint={s.workspaceNameHint}
      />
    </>
  );
};
