import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { type EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailSenderDto } from "@/modules/fireback/sdk/messaging/EmailSenderDto";

export const EmailSenderEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<EmailSenderDto>>) => {
  const t = useT();
  const { values, setFieldValue, errors } = form;

  return (
    <>
      <FormText
        value={values.fromEmailAddress}
        onChange={(value) =>
          setFieldValue(EmailSenderDto.Fields.fromEmailAddress, value, false)
        }
        autoFocus={!isEditing}
        errorMessage={errors.fromEmailAddress}
        label={t.mailProvider.fromEmailAddress}
        hint={t.mailProvider.fromEmailAddressHint}
      />
      <FormText
        value={values.fromName}
        onChange={(value) =>
          setFieldValue(EmailSenderDto.Fields.fromName, value, false)
        }
        errorMessage={errors.fromName}
        label={t.mailProvider.fromName}
        hint={t.mailProvider.fromNameHint}
      />
      <FormText
        value={values.nickName}
        onChange={(value) =>
          setFieldValue(EmailSenderDto.Fields.nickName, value, false)
        }
        errorMessage={errors.nickName}
        label={t.mailProvider.nickName}
        hint={t.mailProvider.nickNameHint}
      />
      <FormText
        value={values.replyTo}
        onChange={(value) =>
          setFieldValue(EmailSenderDto.Fields.replyTo, value, false)
        }
        errorMessage={errors.replyTo}
        label={t.mailProvider.replyTo}
        hint={t.mailProvider.replyToHint}
      />
    </>
  );
};
