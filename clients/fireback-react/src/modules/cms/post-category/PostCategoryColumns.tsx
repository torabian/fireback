import { enTranslations } from "@/translations/en";
import { PostCategoryEntity } from "src/sdk/fireback/modules/cms/PostCategoryEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: PostCategoryEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: PostCategoryEntity.Fields.name,
    title: t.postcategories.name,
    width: 100,
  },    
];