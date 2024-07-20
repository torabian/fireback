import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
export class CustomerAddress extends BaseEntity {
  public zipCode?: string | null;
  public country?: string | null;
  public street?: string | null;
  public city?: string | null;
}
// Class body
export type CustomerEntityKeys =
  keyof typeof CustomerEntity.Fields;
export class CustomerEntity extends BaseEntity {
  public children?: CustomerEntity[] | null;
  public firstName?: string | null;
  public avatar?: string | null;
  public sex?: string | null;
  public subscriptionTier?: string | null;
  public birthday?: string | null;
  public lastName?: string | null;
  public address?: CustomerAddress | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/customer/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/customer/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/customer/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/customers`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "customer/edit/:uniqueId",
      Rcreate: "customer/new",
      Rsingle: "customer/:uniqueId",
      Rquery: "customers",
      rAddressCreate: "customer/:linkerId/address/new",
      rAddressEdit: "customer/:linkerId/address/edit/:uniqueId",
      editAddress(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/customer/${linkerId}/address/edit/${uniqueId}`;
      },
      createAddress(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/customer/${linkerId}/address/new`;
      },
  };
  public static definition = {
  "name": "customer",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "firstName",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "avatar",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "sex",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "subscriptionTier",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "birthday",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "lastName",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "linkedTo": "CustomerEntity",
      "name": "address",
      "type": "object",
      "computedType": "CustomerAddress",
      "gormMap": {},
      "fullName": "CustomerAddress",
      "fields": [
        {
          "name": "zipCode",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "country",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "street",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "city",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        }
      ]
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      firstName: 'firstName',
      avatar: 'avatar',
      sex: 'sex',
      subscriptionTier: 'subscriptionTier',
      birthday: 'birthday',
      lastName: 'lastName',
      address$: 'address',
      address: {
  ...BaseEntity.Fields,
      zipCode: 'zipCode',
      country: 'country',
      street: 'street',
      city: 'city',
      },
}
}