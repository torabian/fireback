import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    PriceTagEntity,
} from "../currency/PriceTagEntity"
import {
    CapabilityEntity,
} from "../workspaces/CapabilityEntity"
import {
    LicensableProductEntity,
} from "./LicensableProductEntity"
// In this section we have sub entities related to this object
export class ProductPlanPermissions extends BaseEntity {
  public capability?: CapabilityEntity | null;
      capabilityId?: string | null;
}
// Class body
export type ProductPlanEntityKeys =
  keyof typeof ProductPlanEntity.Fields;
export class ProductPlanEntity extends BaseEntity {
  public name?: string | null;
  public duration?: number | null;
  public product?: LicensableProductEntity | null;
      productId?: string | null;
  public priceTag?: PriceTagEntity | null;
      priceTagId?: string | null;
  public permissions?: ProductPlanPermissions[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-plan/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/product-plan/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-plan/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-plans`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "product-plan/edit/:uniqueId",
      Rcreate: "product-plan/new",
      Rsingle: "product-plan/:uniqueId",
      Rquery: "product-plans",
      rPermissionsCreate: "product-plan/:linkerId/permissions/new",
      rPermissionsEdit: "product-plan/:linkerId/permissions/edit/:uniqueId",
      editPermissions(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-plan/${linkerId}/permissions/edit/${uniqueId}`;
      },
      createPermissions(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/product-plan/${linkerId}/permissions/new`;
      },
  };
  public static definition = {
  "name": "productPlan",
  "queryScope": "public",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/licenses/LicensableProductDefinitions.dyno.proto",
    "modules/currency/PriceTagDefinitions.dyno.proto",
    "modules/workspaces/CapabilityDefinitions.dyno.proto"
  ],
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
      "name": "duration",
      "type": "int64",
      "validate": "required",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "product",
      "type": "one",
      "target": "LicensableProductEntity",
      "validate": "required",
      "computedType": "LicensableProductEntity",
      "gormMap": {}
    },
    {
      "name": "priceTag",
      "type": "one",
      "target": "PriceTagEntity",
      "module": "currency",
      "computedType": "PriceTagEntity",
      "gormMap": {}
    },
    {
      "linkedTo": "ProductPlanEntity",
      "name": "permissions",
      "type": "array",
      "computedType": "ProductPlanPermissions[]",
      "gormMap": {},
      "fullName": "ProductPlanPermissions",
      "fields": [
        {
          "name": "capability",
          "type": "one",
          "target": "CapabilityEntity",
          "module": "workspaces",
          "computedType": "CapabilityEntity",
          "gormMap": {}
        }
      ]
    }
  ],
  "cliName": "plan"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      duration: 'duration',
          productId: 'productId',
      product$: 'product',
      product: LicensableProductEntity.Fields,
          priceTagId: 'priceTagId',
      priceTag$: 'priceTag',
      priceTag: PriceTagEntity.Fields,
      permissions$: 'permissions',
      permissions: {
  ...BaseEntity.Fields,
          capabilityId: 'capabilityId',
      capability$: 'capability',
      capability: CapabilityEntity.Fields,
      },
}
}
