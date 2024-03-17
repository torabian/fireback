import { enTranslations } from "@/translations/en";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: ProductSubmissionEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: ProductSubmissionEntity.Fields.product,
    title: t.productsubmissions.product,
    width: 100,
  },    
  {
    name: ProductSubmissionEntity.Fields.data,
    title: t.productsubmissions.data,
    width: 100,
  },    
  {
    name: ProductSubmissionEntity.Fields.values,
    title: t.productsubmissions.values,
    width: 100,
  },    
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
    name: ProductSubmissionEntity.Fields.brand,
    title: t.productsubmissions.brand,
    width: 100,
  },    
  {
    name: ProductSubmissionEntity.Fields.category,
    title: t.productsubmissions.category,
    width: 100,
  },    
  {
    name: ProductSubmissionEntity.Fields.tags,
    title: t.productsubmissions.tags,
    width: 100,
  },    
];