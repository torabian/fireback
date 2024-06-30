import { enTranslations } from "@/translations/en";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: DiscountCodeEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: DiscountCodeEntity.Fields.series,
    title: t.discountCodes.series,
    width: 100,
  },
  {
    name: DiscountCodeEntity.Fields.limit,
    title: t.discountCodes.limit,
    width: 100,
  },
  {
    name: "validFrom",
    title: t.discountCodes.validFrom,
    width: 100,
    getCellValue: (entity: DiscountCodeEntity) => entity.valid?.startFormatted,
  },
  {
    name: "validUntil",
    title: t.discountCodes.validUntil,
    getCellValue: (entity: DiscountCodeEntity) => entity.valid?.endFormatted,
    width: 100,
  },
  {
    name: DiscountCodeEntity.Fields.appliedProducts$,
    title: t.discountCodes.appliedProducts,
    width: 100,
    getCellValue: (entity: DiscountCodeEntity) =>
      entity.appliedProducts?.length,
  },
  {
    name: DiscountCodeEntity.Fields.excludedProducts$,
    title: t.discountCodes.excludedProducts,
    width: 100,
    getCellValue: (entity: DiscountCodeEntity) =>
      entity.excludedProducts?.length,
  },
  {
    name: DiscountCodeEntity.Fields.appliedCategories$,
    title: t.discountCodes.appliedCategories,
    width: 100,
    getCellValue: (entity: DiscountCodeEntity) =>
      entity.appliedCategories?.length,
  },
  {
    name: DiscountCodeEntity.Fields.excludedCategories$,
    title: t.discountCodes.excludedCategories,
    width: 100,
    getCellValue: (entity: DiscountCodeEntity) =>
      entity.excludedCategories?.length,
  },
];
