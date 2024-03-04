import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type AuthContextDtoKeys =
  keyof typeof AuthContextDto.Fields;
export class AuthContextDto extends BaseDto {
  public skipWorkspaceId?: boolean | null;
  public workspaceId?: string | null;
  public token?: string | null;
  public capabilities?: string[] | null;
public static Fields = {
  ...BaseEntity.Fields,
      skipWorkspaceId: 'skipWorkspaceId',
      workspaceId: 'workspaceId',
      token: 'token',
      capabilities: 'capabilities',
}
  public static definition = {
  "name": "authContext",
  "fields": [
    {
      "name": "skipWorkspaceId",
      "type": "bool",
      "computedType": "boolean",
      "gormMap": {}
    },
    {
      "name": "workspaceId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "token",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "capabilities",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {}
    }
  ]
}
}
