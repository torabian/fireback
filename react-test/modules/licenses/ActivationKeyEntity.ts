import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    ProductPlanEntity,
} from "./ProductPlanEntity"
// In this section we have sub entities related to this object
// Class body
export type ActivationKeyEntityKeys =
  keyof typeof ActivationKeyEntity.Fields;
export class ActivationKeyEntity extends BaseEntity {
  public children?: ActivationKeyEntity[] | null;
  public series?: string | null;
  public used?: number | null;
  public plan?: ProductPlanEntity | null;
      planId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/activation-key/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/activation-key/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/activation-key/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/activation-keys`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "activation-key/edit/:uniqueId",
      Rcreate: "activation-key/new",
      Rsingle: "activation-key/:uniqueId",
      Rquery: "activation-keys",
  };
  public static definition = {
  "name": "activationKey",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/licenses/ProductPlanDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "series",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "used",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "plan",
      "type": "one",
      "target": "ProductPlanEntity",
      "computedType": "ProductPlanEntity",
      "gormMap": {}
    }
  ],
  "cliName": "key"
}
public static Fields = {
  ...BaseEntity.Fields,
      series: 'series',
      used: 'used',
          planId: 'planId',
      plan$: 'plan',
      plan: ProductPlanEntity.Fields,
}
}