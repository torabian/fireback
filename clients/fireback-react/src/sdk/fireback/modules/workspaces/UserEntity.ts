import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    PersonEntity,
} from "./PersonEntity"
// In this section we have sub entities related to this object
// Class body
export type UserEntityKeys =
  keyof typeof UserEntity.Fields;
export class UserEntity extends BaseEntity {
  public person?: PersonEntity | null;
      personId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/user/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/user/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/user/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/users`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "user/edit/:uniqueId",
      Rcreate: "user/new",
      Rsingle: "user/:uniqueId",
      Rquery: "users",
  };
  public static definition = {
  "name": "user",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "person",
      "type": "one",
      "target": "PersonEntity",
      "allowCreate": true,
      "computedType": "PersonEntity",
      "gormMap": {}
    }
  ],
  "cliDescription": "Manage the users who are in the current app (root only)"
}
public static Fields = {
  ...BaseEntity.Fields,
          personId: 'personId',
      person$: 'person',
      person: PersonEntity.Fields,
}
}
