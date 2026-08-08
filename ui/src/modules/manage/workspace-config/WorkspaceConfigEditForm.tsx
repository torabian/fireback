import { FormCheckbox } from "../../fireback-ui/components/forms/form-switch/FormSwitch";
import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { useS } from "../../fireback-ui/hooks/useS";
import { WorkspaceConfigDto } from "../../sdk/abac/WorkspaceConfigDto";
import { strings } from "./strings/translations";
import { PageSection } from "../../fireback-ui/components/page-section/PageSection";
import { FormSelect } from "../../fireback-ui/components/forms/form-select/FormSelect";
import { useEmailProvidersQuerySource } from "../../fireback-ui/hooks/useEmailProvidersQuerySource";
import { useGsmProvidersQuerySource } from "../../fireback-ui/hooks/useGsmProvidersQuerySource";
import { useRegionalContentsQuerySource } from "../../fireback-ui/hooks/useRegionalContentsQuerySource";

export const WorkspaceConfigForm = ({
  form,
  isEditing,
}: EntityFormProps<WorkspaceConfigDto>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
      <PageSection
        title={s.workspaceConfigs.recaptchaSectionTitle}
        description={s.workspaceConfigs.recaptchaSectionDescription}
      >
        <FormCheckbox
          value={values.enableRecaptcha2}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigDto.Fields.enableRecaptcha2,
              value,
              false,
            )
          }
          errorMessage={errors.enableRecaptcha2}
          label={s.workspaceConfigs.enableRecaptcha2}
          hint={s.workspaceConfigs.enableRecaptcha2Hint}
        />

        <FormText
          value={values.recaptcha2ServerKey}
          disabled={!values.enableRecaptcha2}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigDto.Fields.recaptcha2ServerKey,
              value,
              false,
            )
          }
          errorMessage={errors.recaptcha2ServerKey}
          label={s.workspaceConfigs.recaptcha2ServerKey}
          hint={s.workspaceConfigs.recaptcha2ServerKeyHint}
        />
        <FormText
          value={values.recaptcha2ClientKey}
          disabled={!values.enableRecaptcha2}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigDto.Fields.recaptcha2ClientKey,
              value,
              false,
            )
          }
          errorMessage={errors.recaptcha2ClientKey}
          label={s.workspaceConfigs.recaptcha2ClientKey}
          hint={s.workspaceConfigs.recaptcha2ClientKeyHint}
        />
      </PageSection>

      <PageSection
        title={s.workspaceConfigs.otpSectionTitle}
        description={s.workspaceConfigs.otpSectionDescription}
      >
        <FormCheckbox
          value={values.enableOtp}
          onChange={(value) =>
            setFieldValue(WorkspaceConfigDto.Fields.enableOtp, value, false)
          }
          errorMessage={errors.enableOtp}
          label={s.workspaceConfigs.enableOtp}
          hint={s.workspaceConfigs.enableOtpHint}
        />

        <FormCheckbox
          value={values.requireOtpOnSignup}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigDto.Fields.requireOtpOnSignup,
              value,
              false,
            )
          }
          errorMessage={errors.requireOtpOnSignup}
          label={s.workspaceConfigs.requireOtpOnSignup}
          hint={s.workspaceConfigs.requireOtpOnSignupHint}
        />

        <FormCheckbox
          value={values.requireOtpOnSignin}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigDto.Fields.requireOtpOnSignin,
              value,
              false,
            )
          }
          errorMessage={errors.requireOtpOnSignin}
          label={s.workspaceConfigs.requireOtpOnSignin}
          hint={s.workspaceConfigs.requireOtpOnSigninHint}
        />
      </PageSection>
      <PageSection
        title={s.workspaceConfigs.totpSectionTitle}
        description={s.workspaceConfigs.totpSectionDescription}
      >
        <FormCheckbox
          value={values.enableTotp}
          onChange={(value) =>
            setFieldValue(WorkspaceConfigDto.Fields.enableTotp, value, false)
          }
          errorMessage={errors.enableTotp}
          label={s.workspaceConfigs.enableTotp}
          hint={s.workspaceConfigs.enableTotpHint}
        />

        <FormCheckbox
          value={values.forceTotp}
          onChange={(value) =>
            setFieldValue(WorkspaceConfigDto.Fields.forceTotp, value, false)
          }
          errorMessage={errors.forceTotp}
          label={s.workspaceConfigs.forceTotp}
          hint={s.workspaceConfigs.forceTotpHint}
        />
      </PageSection>
      <PageSection
        title={s.workspaceConfigs.passwordSectionTitle}
        description={s.workspaceConfigs.passwordSectionDescription}
      >
        <FormCheckbox
          value={values.forcePasswordOnPhone}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigDto.Fields.forcePasswordOnPhone,
              value,
              false,
            )
          }
          errorMessage={errors.forcePasswordOnPhone}
          label={s.workspaceConfigs.forcePasswordOnPhone}
          hint={s.workspaceConfigs.forcePasswordOnPhoneHint}
        />

        <FormCheckbox
          value={values.forcePersonNameOnPhone}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigDto.Fields.forcePersonNameOnPhone,
              value,
              false,
            )
          }
          errorMessage={errors.forcePersonNameOnPhone}
          label={s.workspaceConfigs.forcePersonNameOnPhone}
          hint={s.workspaceConfigs.forcePersonNameOnPhoneHint}
        />
      </PageSection>
      <PageSection
        title={s.workspaceConfigs.passwordSectionTitle}
        description={s.workspaceConfigs.passwordSectionDescription}
      >
        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigDto.Fields.generalEmailProviderId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={(e) => `${e.type} (${e.uniqueId})`}
          querySource={useEmailProvidersQuerySource}
          errorMessage={errors.generalEmailProviderId}
          label={s.workspaceConfigs.generalEmailProviderLabel}
          hint={s.workspaceConfigs.generalEmailProviderHint}
        />

        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigDto.Fields.generalGsmProviderId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={(e) => `${e.type} (${e.uniqueId})`}
          querySource={useGsmProvidersQuerySource}
          errorMessage={errors.generalGsmProviderId}
          label={s.workspaceConfigs.generalGsmProviderLabel}
          hint={s.workspaceConfigs.generalGsmProviderHint}
        />

        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigDto.Fields.inviteToWorkspaceContentId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={(e) => `${e.title})`}
          querySource={useRegionalContentsQuerySource}
          errorMessage={errors.inviteToWorkspaceContentId}
          label={s.workspaceConfigs.inviteToWorkspaceContentLabel}
          hint={s.workspaceConfigs.inviteToWorkspaceContentHint}
        />

        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigDto.Fields.emailOtpContentId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={(e) => `${e.title})`}
          querySource={useRegionalContentsQuerySource}
          errorMessage={errors.emailOtpContentId}
          label={s.workspaceConfigs.emailOtpContentLabel}
          hint={s.workspaceConfigs.emailOtpContentHint}
        />
        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigDto.Fields.smsOtpContentId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={(e) => `${e.title})`}
          querySource={useRegionalContentsQuerySource}
          errorMessage={errors.smsOtpContentId}
          label={s.workspaceConfigs.smsOtpContentLabel}
          hint={s.workspaceConfigs.smsOtpContentHint}
        />
      </PageSection>
    </>
  );
};
