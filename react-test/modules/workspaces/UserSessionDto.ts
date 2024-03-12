import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    PassportEntity,
} from "./PassportEntity"
import {
    UserEntity,
} from "./UserEntity"
import {
    UserWorkspaceEntity,
} from "./UserWorkspaceEntity"
// In this section we have sub entities related to this object
// Class body
export type UserSessionDtoKeys =
  keyof typeof UserSessionDto.Fields;
export class UserSessionDto extends BaseDto {
  public passport?: PassportEntity | null;
      passportId?: string | null;
  public token?: string | null;
  public exchangeKey?: string | null;
  public userWorkspaces?: UserWorkspaceEntity[] | null;
    userWorkspacesListId?: string[] | null;
  public user?: UserEntity | null;
  public userId?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
          passportId: 'passportId',
      passport$: 'passport',
      passport: PassportEntity.Fields,
      token: 'token',
      exchangeKey: 'exchangeKey',
        userWorkspacesListId: 'userWorkspacesListId',
      userWorkspaces$: 'userWorkspaces',
      userWorkspaces: UserWorkspaceEntity.Fields,
      user$: 'user',
      user: UserEntity.Fields,
      userId: 'userId',
}
  public static definition = {
  "name": "userSession",
  "fields": [
    {
      "name": "passport",
      "type": "one",
      "target": "PassportEntity",
      "computedType": "PassportEntity",
      "gormMap": {}
    },
    {
      "name": "token",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "exchangeKey",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "userWorkspaces",
      "type": "many2many",
      "target": "UserWorkspaceEntity",
      "computedType": "UserWorkspaceEntity[]",
      "gormMap": {}
    },
    {
      "name": "user",
      "type": "one",
      "target": "UserEntity",
      "computedType": "UserEntity",
      "gormMap": {}
    },
    {
      "name": "userId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
