import { FormText } from "../../../components/forms/form-text/FormText";
import { EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../hooks/useT";
import { UserEntity } from "../../../sdk/modules/abac/UserEntity";
import { FormikProps } from "formik";
import { useContext } from "react";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";

export const UserEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<UserEntity>>) => {
  const { values, setFieldValue, errors, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <FormText
            value={values?.firstName}
            onChange={(value) =>
              setFieldValue(UserEntity.Fields.firstName, value, false)
            }
            autoFocus={!isEditing}
            errorMessage={errors?.firstName}
            label={t.wokspaces.invite.firstName}
            hint={t.wokspaces.invite.firstNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.lastName}
            onChange={(value) =>
              setFieldValue(UserEntity.Fields.lastName, value, false)
            }
            errorMessage={errors?.lastName}
            label={t.wokspaces.invite.lastName}
            hint={t.wokspaces.invite.lastNameHint}
          />
        </div>
      </div>
    </>
  );
};
