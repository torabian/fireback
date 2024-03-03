import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type GsmProviderEntityKeys =
  keyof typeof GsmProviderEntity.Fields;
export class GsmProviderEntity extends BaseEntity {
  public apiKey?: string | null;
  public mainSenderNumber?: string | null;
  public type?: "url" | "terminal" | "mediana" | null;
  public invokeUrl?: string | null;
  public invokeBody?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/gsm-provider/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/gsm-provider/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/gsm-provider/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/gsm-providers`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "gsm-provider/edit/:uniqueId",
      Rcreate: "gsm-provider/new",
      Rsingle: "gsm-provider/:uniqueId",
      Rquery: "gsm-providers",
  };
  public static definition = {
  "name": "gsmProvider",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "apiKey",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "mainSenderNumber",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "type",
      "type": "enum",
      "validate": "required",
      "of": [
        {
          "k": "url"
        },
        {
          "k": "terminal"
        },
        {
          "k": "mediana"
        }
      ],
      "computedType": "\"url\" | \"terminal\" | \"mediana\"",
      "gormMap": {}
    },
    {
      "name": "invokeUrl",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "invokeBody",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      apiKey: 'apiKey',
      mainSenderNumber: 'mainSenderNumber',
      type: 'type',
      invokeUrl: 'invokeUrl',
      invokeBody: 'invokeBody',
}
}
