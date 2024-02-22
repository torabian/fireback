import { FormSelect } from "@/components/forms/form-select/FormSelect";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";
import { EmailSenderEntity } from "src/sdk/fireback";
import { EmailSenderEntityFields } from "src/sdk/fireback/modules/workspaces/email-sender-fields";
import { FormikProps } from "formik";

export const EmailSenderEditForm = ({
  form,
  isEditing,
}: {
  form: FormikProps<Partial<EmailSenderEntity>>;
  isEditing?: boolean;
}) => {
  const t = useT();
  const { values, setFieldValue, errors } = form;

  return (
    <>
      <FormText
        value={values.fromEmailAddress}
        onChange={(value) =>
          setFieldValue(EmailSenderEntityFields.fromEmailAddress, value, false)
        }
        autoFocus={!isEditing}
        errorMessage={errors.fromEmailAddress}
        label={t.mailProvider.fromEmailAddress}
        hint={t.mailProvider.fromEmailAddressHint}
      />
      <FormText
        value={values.fromName}
        onChange={(value) =>
          setFieldValue(EmailSenderEntityFields.fromName, value, false)
        }
        errorMessage={errors.fromName}
        label={t.mailProvider.fromName}
        hint={t.mailProvider.fromNameHint}
      />
      <FormText
        value={values.nickName}
        onChange={(value) =>
          setFieldValue(EmailSenderEntityFields.nickName, value, false)
        }
        errorMessage={errors.nickName}
        label={t.mailProvider.nickName}
        hint={t.mailProvider.nickNameHint}
      />
      <FormText
        value={values.replyTo}
        onChange={(value) =>
          setFieldValue(EmailSenderEntityFields.replyTo, value, false)
        }
        errorMessage={errors.replyTo}
        label={t.mailProvider.replyTo}
        hint={t.mailProvider.replyToHint}
      />
    </>
  );
};
