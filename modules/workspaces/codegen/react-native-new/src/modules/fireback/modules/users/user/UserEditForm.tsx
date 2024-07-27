import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { UserEntity } from "src/sdk/fireback/modules/workspaces/UserEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
export const UserForm = ({
  form,
  isEditing,
}: EntityFormProps<UserEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
        <FormEntitySelect3
          formEffect={ { form, field: UserEntity.Fields.person$ } }
          useQuery={useGetPeople}
          label={t.users.person }
          hint={t.users.personHint}
        />
        <FormText
          value={values.avatar }
          onChange={(value) => setFieldValue(UserEntity.Fields.avatar, value, false)}
          errorMessage={errors.avatar }
          label={t.users.avatar }
          hint={t.users.avatarHint}
        />
    </>
  );
};