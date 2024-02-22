import { FormMultiEntitySelect } from "@/components/forms/form-select/FormMultiEntitySelect";
import { FormSelect } from "@/components/forms/form-select/FormSelect";
import { FormText } from "@/components/forms/form-text/FormText";
import { RoleEntity, UserEntity } from "src/sdk/fireback";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { RoleActions } from "src/sdk/fireback/modules/workspaces/role-actions";
import { UserActions } from "src/sdk/fireback/modules/workspaces/user-actions";
import { UserEntityFields } from "src/sdk/fireback/modules/workspaces/user-fields";
import { FormikProps } from "formik";
import { useContext } from "react";
import {
  getPassportOptions,
  getPasswordOptions,
} from "../passports/PassportCommon";
import { useT } from "@/hooks/useT";

export const UserEditForm = ({
  form,
  isEditing,
}: {
  form: FormikProps<Partial<UserEntity>>;
  isEditing?: boolean;
}) => {
  const { values, setFieldValue, errors, setValues } = form as any;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <div className="row">
        <div className="col-md-6">
          <FormText
            value={values.firstName}
            onChange={(value) =>
              setFieldValue(UserEntityFields.firstName, value, false)
            }
            autoFocus={!isEditing}
            errorMessage={errors.firstName}
            label={t.wokspaces.invite.firstName}
            hint={t.wokspaces.invite.firstNameHint}
          />
        </div>
        <div className="col-md-6">
          <FormText
            value={values.lastName}
            onChange={(value) =>
              setFieldValue(UserEntityFields.lastName, value, false)
            }
            errorMessage={errors.lastName}
            label={t.wokspaces.invite.lastName}
            hint={t.wokspaces.invite.lastNameHint}
          />
        </div>
      </div>
    </>
  );
};
