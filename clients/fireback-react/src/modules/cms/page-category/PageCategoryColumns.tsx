import { enTranslations } from "@/translations/en";
import { PageCategoryEntity } from "src/sdk/fireback/modules/cms/PageCategoryEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: PageCategoryEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: PageCategoryEntity.Fields.name,
    title: t.pagecategories.name,
    width: 100,
  },    
];