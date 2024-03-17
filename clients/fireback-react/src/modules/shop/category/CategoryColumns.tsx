import { enTranslations } from "@/translations/en";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: CategoryEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: CategoryEntity.Fields.name,
    title: t.categories.name,
    width: 100,
  },    
];