import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PostCategoryEntityKeys =
  keyof typeof PostCategoryEntity.Fields;
export class PostCategoryEntity extends BaseEntity {
  public children?: PostCategoryEntity[] | null;
  public name?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/post-category/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/post-category/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/post-category/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/post-categories`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "post-category/edit/:uniqueId",
      Rcreate: "post-category/new",
      Rsingle: "post-category/:uniqueId",
      Rquery: "post-categories",
  };
  public static definition = {
  "name": "postCategory",
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