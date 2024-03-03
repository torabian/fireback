import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    UserEntity,
} from "./UserEntity"
// In this section we have sub entities related to this object
// Class body
export type TokenEntityKeys =
  keyof typeof TokenEntity.Fields;
export class TokenEntity extends BaseEntity {
  public user?: UserEntity | null;
  public validUntil?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/token/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/token/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/token/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/tokens`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "token/edit/:uniqueId",
      Rcreate: "token/new",
      Rsingle: "token/:uniqueId",
      Rquery: "tokens",
  };
  public static definition = {
  "name": "token",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/UserDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "user",
      "type": "one",
      "target": "UserEntity",
      "computedType": "UserEntity",
      "gormMap": {}
    },
    {
      "name": "validUntil",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      user$: 'user',
      user: UserEntity.Fields,
      validUntil: 'validUntil',
}
}
