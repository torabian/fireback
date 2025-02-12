import { FormSelect } from "../../components/forms/form-select/FormSelect";
import { FormText } from "../../components/forms/form-text/FormText";
import { EntityFormProps } from "../../definitions/definitions";
import { createQuerySource } from "../../hooks/useAsQuery";
import { useT } from "../../hooks/useT";
import { EmailProviderEntity } from "../../sdk/modules/workspaces/EmailProviderEntity";

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
