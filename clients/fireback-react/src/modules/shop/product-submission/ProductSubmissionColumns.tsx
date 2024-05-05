import { withEssentials } from "@/fireback/hooks/columns";
import { enTranslations } from "@/translations/en";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";

export const columns = (t: typeof enTranslations) =>
  withEssentials(t, [
    {
      name: ProductSubmissionEntity.Fields.name,
      title: t.productsubmissions.name,
      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.product$,
      title: t.productsubmissions.product,
      getCellValue: (entity: ProductSubmissionEntity) =>
        entity?.product?.name || entity?.productId,
      width: 100,
    },
    {
      name: ProductSubmissionEntity.Fields.price,
      title: t.productsubmissions.price,
      width: 100,
    },
    {
      name: "descriptionExcerpt",
      // name: ProductSubmissionEntity.Fields.descriptionExcerpt,
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
