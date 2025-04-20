import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";
import { enTranslations } from "../../../translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: EmailProviderEntity.Fields.uniqueId,
    title: t.table.uniqueId,
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
