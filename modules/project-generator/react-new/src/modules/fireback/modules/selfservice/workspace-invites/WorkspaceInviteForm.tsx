import { type EntityFormProps } from "../../../definitions/definitions";
import { strings } from "./strings/translations";
import { WorkspaceInviteDto } from "../../../sdk/abac/WorkspaceInviteDto";
import { useT } from "../../../hooks/useT";
import { useS } from "../../../hooks/useS";
import { createQuerySource } from "../../../hooks/useAsQuery";
import { FormText } from "../../../components/forms/form-text/FormText";
import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { FormRichText } from "../../../components/forms/form-richtext/FormRichText";
import { useRolesQuerySource } from "../../../hooks/useRolesQuerySource";
import { FormCheckbox } from "../../../components/forms/form-switch/FormSwitch";
import { interfaceLanguages } from "../personal-settings/Langugages";

export const WorkspaceInviteForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceInviteDto>>) => {
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
                WorkspaceInviteDto.Fields.firstName,
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
              setFieldValue(WorkspaceInviteDto.Fields.lastName, value, false)
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
              field: WorkspaceInviteDto.Fields.targetUserLocale,
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
                WorkspaceInviteDto.Fields.coverLetter,
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
            formEffect={{ field: WorkspaceInviteDto.Fields.role$, form }}
            querySource={useRolesQuerySource}
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
              setFieldValue(WorkspaceInviteDto.Fields.email, value, false)
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
                WorkspaceInviteDto.Fields.forceEmailAddress,
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
                WorkspaceInviteDto.Fields.phonenumber,
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
                WorkspaceInviteDto.Fields.forcePhoneNumber,
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
