import { EmailSenderDto } from "@/modules/fireback/sdk/messaging/EmailSenderDto";
import { enTranslations } from "@/modules/fireback/translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: EmailSenderDto.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.fromName,
    title: t.mailProvider.fromName,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.fromEmailAddress,
    title: t.mailProvider.fromEmailAddress,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.nickName,
    title: t.mailProvider.nickName,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.replyTo,
    title: t.mailProvider.replyTo,
    width: 200,
  },
];
