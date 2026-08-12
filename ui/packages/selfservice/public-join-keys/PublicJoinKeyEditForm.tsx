import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { type EntityFormProps } from "@fireback/ui-core/types/EntityManagement";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { PublicJoinKeyDto } from "@fireback/selfservice/sdk/abac/PublicJoinKeyDto";
import { useRolesQuerySource } from "@fireback/ui-core/hooks/useRolesQuerySource";

export const PublicJoinKeyEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<PublicJoinKeyDto>>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);

  return (
    <>
      <FormSelect
        formEffect={{ field: PublicJoinKeyDto.Fields.role$, form }}
        querySource={useRolesQuerySource}
        label={s.roleFieldLabel}
        errorMessage={errors.roleId}
        fnLabelFormat={(item) => item.name}
        hint={s.roleFieldHint}
      />
    </>
  );
};
