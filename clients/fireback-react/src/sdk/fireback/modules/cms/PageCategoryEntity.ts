import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PageCategoryEntityKeys =
  keyof typeof PageCategoryEntity.Fields;
export class PageCategoryEntity extends BaseEntity {
  public children?: PageCategoryEntity[] | null;
  public name?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/page-category/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/page-category/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/page-category/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/page-categories`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "page-category/edit/:uniqueId",
      Rcreate: "page-category/new",
      Rsingle: "page-category/:uniqueId",
      Rquery: "page-categories",
  };
  public static definition = {
  "name": "pageCategory",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "validate": "required",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
}
}