import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { type EntityFormProps } from "@fireback/ui-core/types/EntityManagement";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { EmailSenderDto } from "@fireback/messaging/sdk/messaging/EmailSenderDto";

export const EmailSenderEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<EmailSenderDto>>) => {
  const s = useS(strings);
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
        label={s.fromEmailAddress}
        hint={s.fromEmailAddressHint}
      />
      <FormText
        value={values.fromName}
        onChange={(value) =>
          setFieldValue(EmailSenderDto.Fields.fromName, value, false)
        }
        errorMessage={errors.fromName}
        label={s.fromName}
        hint={s.fromNameHint}
      />
      <FormText
        value={values.nickName}
        onChange={(value) =>
          setFieldValue(EmailSenderDto.Fields.nickName, value, false)
        }
        errorMessage={errors.nickName}
        label={s.nickName}
        hint={s.nickNameHint}
      />
      <FormText
        value={values.replyTo}
        onChange={(value) =>
          setFieldValue(EmailSenderDto.Fields.replyTo, value, false)
        }
        errorMessage={errors.replyTo}
        label={s.replyTo}
        hint={s.replyToHint}
      />
    </>
  );
};
