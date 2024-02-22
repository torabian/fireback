import { enTranslations } from "@/translations/en";
import { EmailProviderEntityFields } from "src/sdk/fireback/modules/workspaces/email-provider-fields";

export const columns = (t: typeof enTranslations) => [
  {
    name: EmailProviderEntityFields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: EmailProviderEntityFields.type,
    title: t.mailProvider.type,
    width: 200,
  },
  {
    name: EmailProviderEntityFields.apiKey,
    title: t.mailProvider.apiKey,
    width: 200,
  },
];
