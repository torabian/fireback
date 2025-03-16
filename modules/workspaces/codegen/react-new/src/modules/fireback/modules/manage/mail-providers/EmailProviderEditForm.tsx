import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { createQuerySource } from "@/modules/fireback/hooks/useAsQuery";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/workspaces/EmailProviderEntity";

export const EmailProviderEditForm = ({
  form,
  isEditing,
}: EntityFormProps<EmailProviderEntity>) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();

  const emailProviders = [{ label: "Sendgrid", value: "sendgrid" }];
  const querySource = createQuerySource(emailProviders);

  return (
    <>
      <FormSelect
        formEffect={{
          form,
          field: EmailProviderEntity.Fields.type,
          beforeSet(item) {
            return item.value;
          },
        }}
        querySource={querySource}
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
