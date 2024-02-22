import { enTranslations } from "@/translations/en";
import { EmailSenderEntityFields } from "src/sdk/fireback/modules/workspaces/email-sender-fields";

export const columns = (t: typeof enTranslations) => [
  {
    name: EmailSenderEntityFields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: EmailSenderEntityFields.fromName,
    title: t.mailProvider.fromName,
    width: 200,
  },
  {
    name: EmailSenderEntityFields.fromEmailAddress,
    title: t.mailProvider.fromEmailAddress,
    width: 200,
  },
  {
    name: EmailSenderEntityFields.nickName,
    title: t.mailProvider.nickName,
    width: 200,
  },
  {
    name: EmailSenderEntityFields.replyTo,
    title: t.mailProvider.replyTo,
    width: 200,
  },
];
