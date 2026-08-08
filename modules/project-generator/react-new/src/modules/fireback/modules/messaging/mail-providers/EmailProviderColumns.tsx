import { EmailProviderDto } from "@/modules/fireback/sdk/messaging/EmailProviderDto";
import { enTranslations } from "../../../translations/en";
import { strings } from "./strings/translations";

export const columns = (t: typeof enTranslations, s: typeof strings) => [
  {
    name: EmailProviderDto.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: "title",
    title: s.emailProviders.title,
    width: 200,
  },
  {
    name: EmailProviderDto.Fields.type,
    title: t.mailProvider.type,
    width: 200,
  },
  {
    name: EmailProviderDto.Fields.apiKey,
    title: t.mailProvider.apiKey,
    width: 200,
  },
];
