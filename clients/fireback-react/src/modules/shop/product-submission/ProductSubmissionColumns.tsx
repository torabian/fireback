import { BrandEntity } from "@/sdk/fireback/modules/shop/BrandEntity";
import { enTranslations } from "@/translations/en";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";

function withEssentials(
  t: typeof enTranslations,
  items: Array<any>
): Array<any> {
  return [
    {
      name: ProductSubmissionEntity.Fields.uniqueId,
      title: t.table.uniqueId,
      width: 200,
    },
    ...items,
    {
      name: "createdFormatted",
      title: t.table.created,
      width: 200,
    },
    {
      name: "updatedFormatted",
      title: t.table.updated,
      width: 200,
    },
  ];
}

export const columns = (t: typeof enTranslations) =>
  withEssentials(t, [
    {
      name: ProductSubmissionEntity.Fields.name,
      title: t.productsubmissions.name,
      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.product,
      title: t.productsubmissions.product,
      width: 100,
    },
    // {
    //   name: ProductSubmissionEntity.Fields.data,
    //   title: t.productsubmissions.data,
    //   width: 100,
    // },
    // {
    //   name: ProductSubmissionEntity.Fields.values,
    //   title: t.productsubmissions.values,
    //   width: 100,
    // },
    {
      name: ProductSubmissionEntity.Fields.price,
      title: t.productsubmissions.price,
      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.description,
      title: t.productsubmissions.description,
      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.sku,
      title: t.productsubmissions.sku,
      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.brand$,
      title: t.productsubmissions.brand,
      getCellValue: (entity: ProductSubmissionEntity) => entity?.brand?.name,

      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.category$,
      title: t.productsubmissions.category,
      getCellValue: (entity: ProductSubmissionEntity) => entity.category?.name,
      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.tags$,
      title: t.productsubmissions.tags,
      getCellValue: (entity: ProductSubmissionEntity) =>
        (entity.tags || [])?.map((t) => t.name).join(", "),
      width: 100,
    },
  ]);
