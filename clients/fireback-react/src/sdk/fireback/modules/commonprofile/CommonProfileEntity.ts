import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type CommonProfileEntityKeys =
  keyof typeof CommonProfileEntity.Fields;
export class CommonProfileEntity extends BaseEntity {
  public firstName?: string | null;
  public lastName?: string | null;
  public phoneNumber?: string | null;
  public email?: string | null;
  public company?: string | null;
  public street?: string | null;
  public houseNumber?: string | null;
  public zipCode?: string | null;
  public city?: string | null;
  public gender?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/common-profile/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/common-profile/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/common-profile/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/common-profiles`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "common-profile/edit/:uniqueId",
      Rcreate: "common-profile/new",
      Rsingle: "common-profile/:uniqueId",
      Rquery: "common-profiles",
  };
  public static definition = {
  "name": "commonProfile",
  "distinctBy": "user",
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
      "name": "lastName",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "phoneNumber",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "email",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "company",
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
      "name": "houseNumber",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "zipCode",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "city",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "gender",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliDescription": "A common profile issues for every user (Set the living address, etc)"
}
public static Fields = {
  ...BaseEntity.Fields,
      firstName: 'firstName',
      lastName: 'lastName',
      phoneNumber: 'phoneNumber',
      email: 'email',
      company: 'company',
      street: 'street',
      houseNumber: 'houseNumber',
      zipCode: 'zipCode',
      city: 'city',
      gender: 'gender',
}
}
