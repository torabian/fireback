import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { type EntityFormProps } from "@fireback/ui-core/types/EntityManagement";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings as coreStrings } from "@fireback/ui-core/components/strings/translations";
import { strings } from "./strings/translations";
import { WorkspaceDto } from "@fireback/manage/sdk/abac/WorkspaceDto";
import { useWorkspacesQuerySource } from "./useWorkspacesQuerySource";

export const WorkspaceEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceDto>>) => {
  const { values, setFieldValue, errors } = form;
  const s = useS(strings);
  const cs = useS(coreStrings);
  const workspacesQuerySource = useWorkspacesQuerySource;

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

      <FormSelect
        querySource={workspacesQuerySource}
        formEffect={{
          form,
          field: WorkspaceDto.Fields.parentId,
          beforeSet(item) {
            return item?.uniqueId ?? null;
          },
        }}
        fnLabelFormat={(item) => item?.name}
        errorMessage={errors.parentId as string | undefined}
        label={cs.common.parent}
        hint={cs.common.parentHint}
      />

      {/* The unique id can only ever be chosen once, at creation - editing it
          afterwards would silently break every existing reference to this
          workspace (WorkspaceUpdateAction, in fact, never even reads a uniqueId
          out of its request body - see WorkspaceActions.go). Left entirely out of
          the form once isEditing is true, rather than just disabled, so there's
          nothing on screen suggesting it's still changeable. */}
      {!isEditing && (
        <FormText
          value={values.uniqueId ?? ""}
          onChange={(value) =>
            setFieldValue(WorkspaceDto.Fields.uniqueId, value, false)
          }
          errorMessage={errors.uniqueId as string | undefined}
          label={s.workspaceUniqueId}
          hint={s.workspaceUniqueIdHint}
        />
      )}
    </>
  );
};
