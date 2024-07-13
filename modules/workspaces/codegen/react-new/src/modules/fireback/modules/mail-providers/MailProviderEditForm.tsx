import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailProviderEntity } from "../../sdk/modules/workspaces/EmailProviderEntity";
import { FormikProps } from "formik";

export const EmailProviderEditForm = ({
  form,
  isEditing,
}: EntityFormProps<EmailProviderEntity>) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormSelect
        value={values.type}
        onChange={(value) =>
          setFieldValue(EmailProviderEntity.Fields.type, value, false)
        }
        options={[{ label: "Sendgrid", value: "sendgrid" }]}
        errorMessage={errors.type}
        label={t.mailProvider.type}
        hint={t.mailProvider.typeHint}
      />

      <FormText
        value={values.apiKey}
        autoFocus={!isEditing}
        onChange={(value) =>
          setFieldValue(EmailProviderEntity.Fields.apiKey, value, false)
        }
        dir="ltr"
        errorMessage={errors.apiKey}
        label={t.mailProvider.apiKey}
        hint={t.mailProvider.apiKeyHint}
      />
    </>
  );
};
