import { useContext } from "react";
import { FormText } from "../../../components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../hooks/useT";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { UserEntity } from "../../../sdk/modules/abac/UserEntity";

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
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.city}
            onChange={(value) =>
              setFieldValue(UserEntity.Fields.primaryAddress.city, value, false)
            }
            errorMessage={errors?.primaryAddress?.city}
            label={"City name"}
            hint={"The city address associated with user."}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.addressLine1}
            onChange={(value) =>
              setFieldValue(
                UserEntity.Fields.primaryAddress.addressLine1,
                value,
                false,
              )
            }
            errorMessage={errors?.primaryAddress?.addressLine1}
            label={"Address Line 1"}
            hint={"Address line 1 associated with user."}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.addressLine2}
            onChange={(value) =>
              setFieldValue(
                UserEntity.Fields.primaryAddress.addressLine2,
                value,
                false,
              )
            }
            errorMessage={errors?.primaryAddress?.addressLine2}
            label={"Address Line 2"}
            hint={"Address line 2 associated with user."}
          />
        </div>
      </div>
    </>
  );
};
