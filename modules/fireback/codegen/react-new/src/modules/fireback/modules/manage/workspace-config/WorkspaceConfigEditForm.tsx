import { FormCheckbox } from "@/modules/fireback/components/forms/form-switch/FormSwitch";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { type EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useS } from "@/modules/fireback/hooks/useS";
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceConfigEntity";
import { strings } from "./strings/translations";
import { PageSection } from "@/modules/fireback/components/page-section/PageSection";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { useGetEmailProviders } from "@/modules/fireback/sdk/modules/abac/useGetEmailProviders";
import { useGetGsmProviders } from "@/modules/fireback/sdk/modules/abac/useGetGsmProviders";
import { useGetRegionalContents } from "@/modules/fireback/sdk/modules/abac/useGetRegionalContents";

export const WorkspaceConfigForm = ({
  form,
  isEditing,
}: EntityFormProps<WorkspaceConfigEntity>) => {
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
              WorkspaceConfigEntity.Fields.enableRecaptcha2,
              value,
              false
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
              WorkspaceConfigEntity.Fields.recaptcha2ServerKey,
              value,
              false
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
              WorkspaceConfigEntity.Fields.recaptcha2ClientKey,
              value,
              false
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
            setFieldValue(WorkspaceConfigEntity.Fields.enableOtp, value, false)
          }
          errorMessage={errors.enableOtp}
          label={s.workspaceConfigs.enableOtp}
          hint={s.workspaceConfigs.enableOtpHint}
        />

        <FormCheckbox
          value={values.requireOtpOnSignup}
          onChange={(value) =>
            setFieldValue(
              WorkspaceConfigEntity.Fields.requireOtpOnSignup,
              value,
              false
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
              WorkspaceConfigEntity.Fields.requireOtpOnSignin,
              value,
              false
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
            setFieldValue(WorkspaceConfigEntity.Fields.enableTotp, value, false)
          }
          errorMessage={errors.enableTotp}
          label={s.workspaceConfigs.enableTotp}
          hint={s.workspaceConfigs.enableTotpHint}
        />

        <FormCheckbox
          value={values.forceTotp}
          onChange={(value) =>
            setFieldValue(WorkspaceConfigEntity.Fields.forceTotp, value, false)
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
              WorkspaceConfigEntity.Fields.forcePasswordOnPhone,
              value,
              false
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
              WorkspaceConfigEntity.Fields.forcePersonNameOnPhone,
              value,
              false
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
            field: WorkspaceConfigEntity.Fields.generalEmailProviderId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={e => `${e.type} (${e.uniqueId})`}
          querySource={useGetEmailProviders}
          errorMessage={errors.generalEmailProviderId}
          label={s.workspaceConfigs.generalEmailProviderLabel}
          hint={s.workspaceConfigs.generalEmailProviderHint}
        />

        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigEntity.Fields.generalGsmProviderId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={e => `${e.type} (${e.uniqueId})`}
          querySource={useGetGsmProviders}
          errorMessage={errors.generalGsmProviderId}
          label={s.workspaceConfigs.generalGsmProviderLabel}
          hint={s.workspaceConfigs.generalGsmProviderHint}
        />

        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigEntity.Fields.inviteToWorkspaceContentId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={e => `${e.title})`}
          querySource={useGetRegionalContents}
          errorMessage={errors.inviteToWorkspaceContentId}
          label={s.workspaceConfigs.inviteToWorkspaceContentLabel}
          hint={s.workspaceConfigs.inviteToWorkspaceContentHint}
        />

        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigEntity.Fields.emailOtpContentId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={e => `${e.title})`}
          querySource={useGetRegionalContents}
          errorMessage={errors.emailOtpContentId}
          label={s.workspaceConfigs.emailOtpContentLabel}
          hint={s.workspaceConfigs.emailOtpContentHint}
        />
        <FormSelect
          keyExtractor={(t) => t.uniqueId}
          formEffect={{
            form,
            field: WorkspaceConfigEntity.Fields.smsOtpContentId,
            beforeSet(item) {
              return item.uniqueId;
            },
          }}
          fnLabelFormat={e => `${e.title})`}
          querySource={useGetRegionalContents}
          errorMessage={errors.smsOtpContentId}
          label={s.workspaceConfigs.smsOtpContentLabel}
          hint={s.workspaceConfigs.smsOtpContentHint}
        />

      </PageSection>
    </>
  );
};
