import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { UserEntity } from "@/sdk/fireback/modules/workspaces/UserEntity";
import { FormikProps } from "formik";
import { useContext } from "react";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";

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
            value={values.person?.firstName}
            onChange={(value) =>
              setFieldValue("person.firstName", value, false)
            }
            autoFocus={!isEditing}
            errorMessage={errors?.person?.firstName}
            label={t.wokspaces.invite.firstName}
            hint={t.wokspaces.invite.firstNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values.person?.lastName}
            onChange={(value) => setFieldValue("person.lastName", value, false)}
            errorMessage={errors?.person?.lastName}
            label={t.wokspaces.invite.lastName}
            hint={t.wokspaces.invite.lastNameHint}
          />
        </div>
      </div>
    </>
  );
};
