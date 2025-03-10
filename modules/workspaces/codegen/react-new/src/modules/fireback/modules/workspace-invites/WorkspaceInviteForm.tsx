import { FormSelect } from "../../components/forms/form-select/FormSelect";
import { FormText } from "../../components/forms/form-text/FormText";
import { EntityFormProps } from "../../definitions/definitions";
import { useT } from "../../hooks/useT";
import { WorkspaceInviteEntity } from "../../sdk/modules/workspaces/WorkspaceInviteEntity";
import { useGetRoles } from "../../sdk/modules/workspaces/useGetRoles";

export const WorkspaceInviteForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceInviteEntity>>) => {
  const t = useT();
  const { values, setValues, setFieldValue, errors } = form;

  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <FormText
            value={values.firstName}
            onChange={(value) =>
              setFieldValue(
                WorkspaceInviteEntity.Fields.firstName,
                value,
                false
              )
            }
            errorMessage={errors.firstName}
            label={t.wokspaces.invite.firstName}
            autoFocus={!isEditing}
            hint={t.wokspaces.invite.firstNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values.lastName}
            onChange={(value) =>
              setFieldValue(WorkspaceInviteEntity.Fields.lastName, value, false)
            }
            errorMessage={errors.lastName}
            label={t.wokspaces.invite.lastName}
            hint={t.wokspaces.invite.lastNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormSelect
            formEffect={{ field: WorkspaceInviteEntity.Fields.role$, form }}
            querySource={useGetRoles}
            label={t.wokspaces.invite.role}
            fnLabelFormat={(item) => item.name}
            hint={t.wokspaces.invite.roleHint}
          />
        </div>
      </div>

      <div className="row">
        <div className="col-md-12">
          {/* <FormEntitySelect
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
          /> */}
        </div>
        {/* <div className="col-md-12">
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
        {/* <div className="col-md-12">
          <FormCheckbox
            value={values.passportMode === "forced"}
            onChange={(value) =>
              setFieldValue("passportMode", value ? "forced" : "optional")
            }
            errorMessage={errors.email}
            label={t.wokspaces.invite.forcePassport}
          />
        </div> */}
        <div className="col-md-12">
          <FormText
            value={values.value}
            onChange={(value) => setFieldValue("value", value, false)}
            errorMessage={errors.value}
            label={t.wokspaces.invite.email}
            hint={t.wokspaces.invite.emailHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values.value}
            onChange={(value) => setFieldValue("phoneNumber", value, false)}
            errorMessage={errors.value}
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
