import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useT } from "@/modules/fireback/hooks/useT";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/workspaces/PublicJoinKeyEntity";
import { useGetRoles } from "@/modules/fireback/sdk/modules/workspaces/useGetRoles";
import { useContext } from "react";

export const PublicJoinKeyEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<PublicJoinKeyEntity>>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <FormSelect
        formEffect={{ field: PublicJoinKeyEntity.Fields.role$, form }}
        querySource={useGetRoles}
        label={t.wokspaces.invite.role}
        errorMessage={errors.roleId}
        fnLabelFormat={(item) => item.name}
        hint={t.wokspaces.invite.roleHint}
      />
    </>
  );
};
