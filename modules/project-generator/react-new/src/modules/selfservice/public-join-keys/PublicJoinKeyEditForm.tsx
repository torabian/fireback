import { FormSelect } from "../../fireback-ui/components/forms/form-select/FormSelect";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { PublicJoinKeyDto } from "../../sdk/abac/PublicJoinKeyDto";
import { useRolesQuerySource } from "../../fireback-ui/hooks/useRolesQuerySource";
import { useContext } from "react";

export const PublicJoinKeyEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<PublicJoinKeyDto>>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const { options } = useContext(RemoteQueryContext);
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
