import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CurrencyEntity,
} from "./CurrencyEntity"
// In this section we have sub entities related to this object
export class PriceTagVariations extends BaseEntity {
  public currency?: CurrencyEntity | null;
      currencyId?: string | null;
  public amount?: number | null;
}
// Class body
export type PriceTagEntityKeys =
  keyof typeof PriceTagEntity.Fields;
export class PriceTagEntity extends BaseEntity {
  public children?: PriceTagEntity[] | null;
  public variations?: PriceTagVariations[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/price-tag/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/price-tag/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/price-tag/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/price-tags`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "price-tag/edit/:uniqueId",
      Rcreate: "price-tag/new",
      Rsingle: "price-tag/:uniqueId",
      Rquery: "price-tags",
      rVariationsCreate: "price-tag/:linkerId/variations/new",
      rVariationsEdit: "price-tag/:linkerId/variations/edit/:uniqueId",
      editVariations(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/price-tag/${linkerId}/variations/edit/${uniqueId}`;
      },
      createVariations(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/price-tag/${linkerId}/variations/new`;
      },
  };
  public static definition = {
  "name": "priceTag",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "linkedTo": "PriceTagEntity",
      "name": "variations",
      "type": "array",
      "computedType": "PriceTagVariations[]",
      "gormMap": {},
      "fullName": "PriceTagVariations",
      "fields": [
        {
          "name": "currency",
          "type": "one",
          "target": "CurrencyEntity",
          "computedType": "CurrencyEntity",
          "gormMap": {}
        },
        {
          "name": "amount",
          "type": "float64",
          "computedType": "number",
          "gormMap": {}
        }
      ]
    }
  ],
  "cliDescription": "Price tag is a definition of a price, in different currencies or regions"
}
public static Fields = {
  ...BaseEntity.Fields,
      variations$: 'variations',
      variations: {
  ...BaseEntity.Fields,
          currencyId: 'currencyId',
      currency$: 'currency',
        currency: CurrencyEntity.Fields,
      amount: 'amount',
      },
}
}