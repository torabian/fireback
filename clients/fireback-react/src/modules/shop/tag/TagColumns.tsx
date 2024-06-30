import { enTranslations } from "@/translations/en";
import { TagEntity } from "src/sdk/fireback/modules/shop/TagEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: TagEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: TagEntity.Fields.name,
    title: t.tags.name,
    width: 100,
  },    
];