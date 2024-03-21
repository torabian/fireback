import { enTranslations } from "@/translations/en";
import { PageTagEntity } from "src/sdk/fireback/modules/cms/PageTagEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: PageTagEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: PageTagEntity.Fields.name,
    title: t.pagetags.name,
    width: 100,
  },    
];