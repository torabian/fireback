import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { strings } from "./strings/translations";
import { strings as personalSettingsStrings } from "../personal-settings/strings/translations";
import { WorkspaceInviteDto } from "../../sdk/abac/WorkspaceInviteDto";
import { useS } from "../../fireback-ui/hooks/useS";
import { createQuerySource } from "../../fireback-ui/hooks/useAsQuery";
import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { FormSelect } from "../../fireback-ui/components/forms/form-select/FormSelect";
import { FormRichText } from "../../fireback-ui/components/forms/form-richtext/FormRichText";
import { useRolesQuerySource } from "../../fireback-ui/hooks/useRolesQuerySource";
import { FormCheckbox } from "../../fireback-ui/components/forms/form-switch/FormSwitch";
import { interfaceLanguages } from "../personal-settings/Langugages";

export const WorkspaceInviteForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceInviteDto>>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  const personalSettingsS = useS(personalSettingsStrings);

  const languages = interfaceLanguages(personalSettingsS);
  const languagesQuerySource = createQuerySource(languages);

  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <FormText
            value={values.firstName}
            onChange={(value) =>
              setFieldValue(WorkspaceInviteDto.Fields.firstName, value, false)
            }
            errorMessage={errors.firstName}
            label={s.firstName}
            autoFocus={!isEditing}
            hint={s.firstNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values.lastName}
            onChange={(value) =>
              setFieldValue(WorkspaceInviteDto.Fields.lastName, value, false)
            }
            errorMessage={errors.lastName}
            label={s.lastName}
            hint={s.lastNameHint}
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
              setFieldValue(WorkspaceInviteDto.Fields.coverLetter, value, false)
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
            label={s.roleLabel}
            errorMessage={errors.roleId}
            fnLabelFormat={(item) => item.name}
            hint={s.roleHint}
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
            label={s.email}
            hint={s.emailHint}
          />
        </div>
        <div className="col-md-12">
          <FormCheckbox
            value={values.forceEmailAddress}
            onChange={(value) =>
              setFieldValue(WorkspaceInviteDto.Fields.forceEmailAddress, value)
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
              setFieldValue(WorkspaceInviteDto.Fields.phonenumber, value, false)
            }
            errorMessage={errors.phonenumber}
            type="phonenumber"
            label={s.phoneNumber}
            hint={s.phoneNumberHint}
          />
        </div>
        <div className="col-md-12">
          <FormCheckbox
            value={values.forcePhoneNumber}
            onChange={(value) =>
              setFieldValue(WorkspaceInviteDto.Fields.forcePhoneNumber, value)
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
