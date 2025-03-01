import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useS } from "@/modules/fireback/hooks/useS";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceConfigEntity";
import { useContext } from "react";
import { strings } from "./strings/translations";
import { FormCheckbox } from "@/modules/fireback/components/forms/form-switch/FormSwitch";

export const WorkspaceConfigForm = ({
  form,
  isEditing,
}: EntityFormProps<WorkspaceConfigEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
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

      <FormText
        value={values.recaptcha2ServerKey}
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
      {/*
          <FormText
            type="?"
            value={values.enableTotp }
            onChange={(value) => setFieldValue(WorkspaceConfigEntity.Fields.enableTotp, value, false)}
            errorMessage={errors.enableTotp }
            label={s.workspaceConfigs.enableTotp }
            hint={s.workspaceConfigs.enableTotpHint}
          />
         */}
      {/*
          <FormText
            type="?"
            value={values.forceTotp }
            onChange={(value) => setFieldValue(WorkspaceConfigEntity.Fields.forceTotp, value, false)}
            errorMessage={errors.forceTotp }
            label={s.workspaceConfigs.forceTotp }
            hint={s.workspaceConfigs.forceTotpHint}
          />
         */}
      {/*
          <FormText
            type="?"
            value={values.forcePasswordOnPhone }
            onChange={(value) => setFieldValue(WorkspaceConfigEntity.Fields.forcePasswordOnPhone, value, false)}
            errorMessage={errors.forcePasswordOnPhone }
            label={s.workspaceConfigs.forcePasswordOnPhone }
            hint={s.workspaceConfigs.forcePasswordOnPhoneHint}
          />
         */}
      {/*
          <FormText
            type="?"
            value={values.forcePersonNameOnPhone }
            onChange={(value) => setFieldValue(WorkspaceConfigEntity.Fields.forcePersonNameOnPhone, value, false)}
            errorMessage={errors.forcePersonNameOnPhone }
            label={s.workspaceConfigs.forcePersonNameOnPhone }
            hint={s.workspaceConfigs.forcePersonNameOnPhoneHint}
          />
         */}
    </>
  );
};
