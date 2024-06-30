import { EmailSenderEntity } from "@/sdk/fireback/modules/workspaces/EmailSenderEntity";
import { enTranslations } from "@/translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: EmailSenderEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: EmailSenderEntity.Fields.fromName,
    title: t.mailProvider.fromName,
    width: 200,
  },
  {
    name: EmailSenderEntity.Fields.fromEmailAddress,
    title: t.mailProvider.fromEmailAddress,
    width: 200,
  },
  {
    name: EmailSenderEntity.Fields.nickName,
    title: t.mailProvider.nickName,
    width: 200,
  },
  {
    name: EmailSenderEntity.Fields.replyTo,
    title: t.mailProvider.replyTo,
    width: 200,
  },
];
