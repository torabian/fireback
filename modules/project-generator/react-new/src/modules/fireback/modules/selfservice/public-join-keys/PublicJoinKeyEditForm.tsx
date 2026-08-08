import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { type EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../hooks/useT";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { PublicJoinKeyDto } from "../../../sdk/abac/PublicJoinKeyDto";
import { useRolesQuerySource } from "../../../hooks/useRolesQuerySource";
import { useContext } from "react";

export const PublicJoinKeyEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<PublicJoinKeyDto>>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <FormSelect
        formEffect={{ field: PublicJoinKeyDto.Fields.role$, form }}
        querySource={useRolesQuerySource}
        label={t.wokspaces.invite.role}
        errorMessage={errors.roleId}
        fnLabelFormat={(item) => item.name}
        hint={t.wokspaces.invite.roleHint}
      />
    </>
  );
};
