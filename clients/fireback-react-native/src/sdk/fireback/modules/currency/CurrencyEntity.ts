import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type CurrencyEntityKeys =
  keyof typeof CurrencyEntity.Fields;
export class CurrencyEntity extends BaseEntity {
  public children?: CurrencyEntity[] | null;
  public symbol?: string | null;
  public name?: string | null;
  public symbolNative?: string | null;
  public decimalDigits?: number | null;
  public rounding?: number | null;
  public code?: string | null;
  public namePlural?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/currency/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/currency/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/currency/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/currencies`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "currency/edit/:uniqueId",
      Rcreate: "currency/new",
      Rsingle: "currency/:uniqueId",
      Rquery: "currencies",
  };
  public static definition = {
  "name": "currency",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "symbol",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "name",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "symbolNative",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "decimalDigits",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "rounding",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "code",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "namePlural",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliShort": "curr",
  "cliDescription": "List of all famous currencies, both internal and user defined ones"
}
public static Fields = {
  ...BaseEntity.Fields,
      symbol: 'symbol',
      name: 'name',
      symbolNative: 'symbolNative',
      decimalDigits: 'decimalDigits',
      rounding: 'rounding',
      code: 'code',
      namePlural: 'namePlural',
}
}