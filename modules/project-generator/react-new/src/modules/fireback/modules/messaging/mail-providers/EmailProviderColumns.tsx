import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";
import { enTranslations } from "../../../translations/en";
import { strings } from "./strings/translations";

export const columns = (t: typeof enTranslations, s: typeof strings) => [
  {
    name: EmailProviderEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: "title",
    title: s.emailProviders.title,
    width: 200,
  },
  {
    name: EmailProviderEntity.Fields.type,
    title: t.mailProvider.type,
    width: 200,
  },
  {
    name: EmailProviderEntity.Fields.apiKey,
    title: t.mailProvider.apiKey,
    width: 200,
  },
];
