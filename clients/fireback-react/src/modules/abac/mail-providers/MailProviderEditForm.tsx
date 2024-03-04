import { FormSelect } from "@/components/forms/form-select/FormSelect";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";
import { EmailProviderEntity } from "@/sdk/fireback/modules/workspaces/EmailProviderEntity";
import { FormikProps } from "formik";

export const EmailProviderEditForm = ({
  form,
}: {
  form: FormikProps<Partial<EmailProviderEntity>>;
}) => {
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
