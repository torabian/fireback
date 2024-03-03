import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PersonEntityKeys =
  keyof typeof PersonEntity.Fields;
export class PersonEntity extends BaseEntity {
  public firstName?: string | null;
  public lastName?: string | null;
  public photo?: string | null;
  public gender?: string | null;
  public title?: string | null;
  public birthDate?: Date | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/person/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/person/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/person/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/people`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "person/edit/:uniqueId",
      Rcreate: "person/new",
      Rsingle: "person/:uniqueId",
      Rquery: "people",
  };
  public static definition = {
  "name": "person",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "firstName",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "lastName",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "photo",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "gender",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "title",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "birthDate",
      "type": "date",
      "computedType": "Date",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      firstName: 'firstName',
      lastName: 'lastName',
      photo: 'photo',
      gender: 'gender',
      title: 'title',
      birthDate: 'birthDate',
}
}
