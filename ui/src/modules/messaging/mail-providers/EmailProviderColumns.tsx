import { EmailProviderDto } from "../../sdk/messaging/EmailProviderDto";
import { type strings as uiStrings } from "../../fireback-ui/components/strings/translations";
import { strings } from "./strings/translations";

export const columns = (s: typeof strings, uiS: typeof uiStrings) => [
  {
    name: EmailProviderDto.Fields.uniqueId,
    title: uiS.table.uniqueId,
    width: 200,
  },
  {
    name: "title",
    title: s.emailProviders.title,
    width: 200,
  },
  {
    name: EmailProviderDto.Fields.type,
    title: s.type,
    width: 200,
  },
  {
    name: EmailProviderDto.Fields.apiKey,
    title: s.apiKey,
    width: 200,
  },
];
