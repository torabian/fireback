import { useContext } from "react";
import { FormText } from "../../../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { UserDto } from "../../../sdk/abac/UserDto";
import { strings } from "./strings/translations";

export const UserEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<UserDto>>) => {
  const { values, setFieldValue, errors, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();
  const s = useS(strings);

  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <FormText
            value={values?.firstName}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.firstName, value, false)
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
              setFieldValue(UserDto.Fields.lastName, value, false)
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
              setFieldValue(UserDto.Fields.primaryAddress.city, value, false)
            }
            errorMessage={errors?.primaryAddress?.city}
            label={s.cityName}
            hint={s.cityNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.addressLine1}
            onChange={(value) =>
              setFieldValue(
                UserDto.Fields.primaryAddress.addressLine1,
                value,
                false,
              )
            }
            errorMessage={errors?.primaryAddress?.addressLine1}
            label={s.addressLine1}
            hint={s.addressLine1Hint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.addressLine2}
            onChange={(value) =>
              setFieldValue(
                UserDto.Fields.primaryAddress.addressLine2,
                value,
                false,
              )
            }
            errorMessage={errors?.primaryAddress?.addressLine2}
            label={s.addressLine2}
            hint={s.addressLine2Hint}
          />
        </div>
      </div>
    </>
  );
};
