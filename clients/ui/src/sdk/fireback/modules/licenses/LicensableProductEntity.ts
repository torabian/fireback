import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type LicensableProductEntityKeys =
  keyof typeof LicensableProductEntity.Fields;
export class LicensableProductEntity extends BaseEntity {
  public children?: LicensableProductEntity[] | null;
  public name?: string | null;
  public privateKey?: string | null;
  public publicKey?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/licensable-product/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/licensable-product/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/licensable-product/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/licensable-products`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "licensable-product/edit/:uniqueId",
      Rcreate: "licensable-product/new",
      Rsingle: "licensable-product/:uniqueId",
      Rquery: "licensable-products",
  };
  public static definition = {
  "name": "licensableProduct",
  "queryScope": "public",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "validate": "required,omitempty,min=1,max=100",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "privateKey",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "publicKey",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliName": "product"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      privateKey: 'privateKey',
      publicKey: 'publicKey',
}
}
