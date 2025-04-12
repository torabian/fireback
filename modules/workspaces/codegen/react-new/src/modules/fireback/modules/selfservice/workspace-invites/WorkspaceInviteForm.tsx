import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { strings } from "./strings/translations";
import { WorkspaceInviteEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceInviteEntity";
import { useT } from "@/modules/fireback/hooks/useT";
import { useS } from "@/modules/fireback/hooks/useS";
import { createQuerySource } from "@/modules/fireback/hooks/useAsQuery";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";
import { useGetRoles } from "@/modules/fireback/sdk/modules/abac/useGetRoles";
import { FormCheckbox } from "@/modules/fireback/components/forms/form-switch/FormSwitch";
import { interfaceLanguages } from "../personal-settings/Langugages";

export const WorkspaceInviteForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceInviteEntity>>) => {
  const t = useT();
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);

  const languages = interfaceLanguages(t);
  const languagesQuerySource = createQuerySource(languages);

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
            keyExtractor={(item) => item.value}
            formEffect={{
              form,
              field: WorkspaceInviteEntity.Fields.targetUserLocale,
              beforeSet(item) {
                return item.value;
              },
            }}
            errorMessage={form.errors.targetUserLocale}
            querySource={languagesQuerySource}
            label={s.targetLocale}
            hint={s.targetLocaleHint}
          />
        </div>
        <div className="col-md-12">
          <FormRichText
            value={values.coverLetter}
            onChange={(value) =>
              setFieldValue(
                WorkspaceInviteEntity.Fields.coverLetter,
                value,
                false
              )
            }
            forceBasic
            errorMessage={errors.coverLetter}
            label={s.coverLetter}
            placeholder={s.coverLetterHint}
            hint={s.coverLetterHint}
          />
        </div>
        <div className="col-md-12">
          <FormSelect
            formEffect={{ field: WorkspaceInviteEntity.Fields.role$, form }}
            querySource={useGetRoles}
            label={t.wokspaces.invite.role}
            errorMessage={errors.roleId}
            fnLabelFormat={(item) => item.name}
            hint={t.wokspaces.invite.roleHint}
          />
        </div>
      </div>

      <div className="row">
        <div className="col-md-12">
          <FormText
            value={values.email}
            onChange={(value) =>
              setFieldValue(WorkspaceInviteEntity.Fields.email, value, false)
            }
            errorMessage={errors.email}
            label={t.wokspaces.invite.email}
            hint={t.wokspaces.invite.emailHint}
          />
        </div>
        <div className="col-md-12">
          <FormCheckbox
            value={values.forceEmailAddress}
            onChange={(value) =>
              setFieldValue(
                WorkspaceInviteEntity.Fields.forceEmailAddress,
                value
              )
            }
            errorMessage={errors.forceEmailAddress}
            label={s.forcedEmailAddress}
            hint={s.forcedEmailAddressHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values.phonenumber}
            onChange={(value) =>
              setFieldValue(
                WorkspaceInviteEntity.Fields.phonenumber,
                value,
                false
              )
            }
            errorMessage={errors.phonenumber}
            type="phonenumber"
            label={t.wokspaces.invite.phoneNumber}
            hint={t.wokspaces.invite.phoneNumberHint}
          />
        </div>
        <div className="col-md-12">
          <FormCheckbox
            value={values.forcePhoneNumber}
            onChange={(value) =>
              setFieldValue(
                WorkspaceInviteEntity.Fields.forcePhoneNumber,
                value
              )
            }
            errorMessage={errors.forcePhoneNumber}
            label={s.forcedPhone}
            hint={s.forcedPhoneHint}
          />
        </div>
      </div>
    </>
  );
};
