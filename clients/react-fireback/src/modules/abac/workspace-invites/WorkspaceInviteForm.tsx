import { FormEntitySelect2 } from "@/components/forms/form-select/FormEntitySelect2";
import { FormCheckbox } from "@/components/forms/form-switch/FormSwitch";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";
import { RoleEntity, WorkspaceInviteEntity } from "src/sdk/fireback";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { RoleActions } from "src/sdk/fireback/modules/workspaces/role-actions";
import { FormikProps } from "formik";
import { useContext } from "react";

export const WorkspaceInviteForm = ({
  form,
  isEditing,
}: {
  form: FormikProps<Partial<WorkspaceInviteEntity>>;
  isEditing?: boolean;
}) => {
  const t = useT();
  const { values, setValues, setFieldValue, errors } = form;
  const { options } = useContext(RemoteQueryContext);

  return (
    <>
      <div className="row">
        <div className="col-md-6">
          <FormText
            value={values.firstName}
            onChange={(value) => setFieldValue("firstName", value, false)}
            errorMessage={errors.firstName}
            label={t.wokspaces.invite.firstName}
            autoFocus={!isEditing}
            hint={t.wokspaces.invite.firstNameHint}
          />
        </div>
        <div className="col-md-6">
          <FormText
            value={values.lastName}
            onChange={(value) => setFieldValue("lastName", value, false)}
            errorMessage={errors.lastName}
            label={t.wokspaces.invite.lastName}
            hint={t.wokspaces.invite.lastNameHint}
          />
        </div>
      </div>

      <div className="row">
        <div className="col-md-6">
          <FormEntitySelect2
            label={t.wokspaces.invite.role}
            hint={t.wokspaces.invite.roleHint}
            fnLoadOptions={async (keyword) => {
              return (
                (
                  await RoleActions.fn(options)
                    .query(`name %"${keyword}"%`)
                    .getRoles()
                ).data?.items || []
              );
            }}
            value={(values as any).role}
            onChange={(entity) => {
              setValues({
                ...values,
                role: entity,
                roleId: entity.uniqueId,
              });
            }}
            labelFn={(t: RoleEntity) => [t?.name].join(" ")}
            errorMessage={errors.roleId}
          />
        </div>
        {/* <div className="col-md-6">
          <FormSelect
            value={values.passportType}
            onChange={(value) => setFieldValue("passportType", value, false)}
            errorMessage={errors.passportType}
            options={getPassportOptions(t)}
            label="User primary passport"
            hint="You can define how the user can login into their account"
          />
        </div> */}
      </div>

      <div className="row">
        <div className="col-md-12">
          <FormCheckbox
            value={values.passportMode === "forced"}
            onChange={(value) =>
              setFieldValue("passportMode", value ? "forced" : "optional")
            }
            errorMessage={errors.email}
            label={t.wokspaces.invite.forcePassport}
          />
        </div>
        <div className="col-md-6">
          <FormText
            value={values.email}
            onChange={(value) => setFieldValue("email", value, false)}
            errorMessage={errors.email}
            label={t.wokspaces.invite.email}
            hint={t.wokspaces.invite.emailHint}
          />
        </div>
        <div className="col-md-6">
          <FormText
            value={values.phoneNumber}
            onChange={(value) => setFieldValue("phoneNumber", value, false)}
            errorMessage={errors.phoneNumber}
            type="phonenumber"
            label={t.wokspaces.invite.phoneNumber}
            hint={t.wokspaces.invite.phoneNumberHint}
          />
        </div>
      </div>

      {/* 
          <FormSelect
            value={values.passwordMethod}
            type="verbose"
            onChange={(value) => setFieldValue("passwordMethod", value, false)}
            errorMessage={errors.passwordMethod}
            options={getPasswordOptions(t)}
            name="passwordMethod"
            label="How to send password"
            hint="Determine how the user will get the password for first time."
          />
        </>
       */}
    </>
  );
};
