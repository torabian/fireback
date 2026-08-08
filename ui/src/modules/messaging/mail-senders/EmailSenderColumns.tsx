import { EmailSenderDto } from "../../sdk/messaging/EmailSenderDto";
import { type strings as uiStrings } from "../../fireback-ui/components/strings/translations";
import { type strings } from "./strings/translations";

export const columns = (s: typeof strings, uiS: typeof uiStrings) => [
  {
    name: EmailSenderDto.Fields.uniqueId,
    title: uiS.table.uniqueId,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.fromName,
    title: s.fromName,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.fromEmailAddress,
    title: s.fromEmailAddress,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.nickName,
    title: s.nickName,
    width: 200,
  },
  {
    name: EmailSenderDto.Fields.replyTo,
    title: s.replyTo,
    width: 200,
  },
];
