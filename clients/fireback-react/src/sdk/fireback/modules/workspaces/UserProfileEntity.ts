import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type UserProfileEntityKeys =
  keyof typeof UserProfileEntity.Fields;
export class UserProfileEntity extends BaseEntity {
  public firstName?: string | null;
  public lastName?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/user-profile/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/user-profile/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/user-profile/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/user-profiles`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "user-profile/edit/:uniqueId",
      Rcreate: "user-profile/new",
      Rsingle: "user-profile/:uniqueId",
      Rquery: "user-profiles",
  };
  public static definition = {
  "name": "userProfile",
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
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      firstName: 'firstName',
      lastName: 'lastName',
}
}
