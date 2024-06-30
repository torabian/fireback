import { enTranslations } from "@/translations/en";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: ProductEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: ProductEntity.Fields.name,
    title: t.products.name,
    width: 100,
  },
  {
    name: ProductEntity.Fields.description,
    title: t.products.description,
    width: 100,
  },
  // {
  //   name: ProductEntity.Fields.uiSchema,
  //   title: t.products.uiSchema,
  //   width: 100,
  // },
  // {
  //   name: ProductEntity.Fields.jpsonSchema,
  //   title: t.products.jsonSchema,
  //   width: 100,
  // },
  // {
  //   name: ProductEntity.Fields.fields,
  //   title: t.products.fields,
  //   width: 100,
  // },
];
