import { enTranslations } from "@/translations/en";
import { BrandEntity } from "src/sdk/fireback/modules/shop/BrandEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: BrandEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: BrandEntity.Fields.name,
    title: t.brands.name,
    width: 100,
  },    
];