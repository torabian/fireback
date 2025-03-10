/*
*	Generated by fireback 1.2.0
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/
    import {
        BaseDto,
        BaseEntity,
    } from "../../core/definitions"
    import {
        SecurityModel,
    } from "./SecurityModel"
// In this section we have sub entities related to this object
// Class body
export type AuthContextDtoKeys =
  keyof typeof AuthContextDto.Fields;
export class AuthContextDto extends BaseDto {
  public skipWorkspaceId?: boolean | null;
  public workspaceId?: string | null;
  public token?: string | null;
  public security?: SecurityModel | null;
      securityId?: string | null;
  public capabilities?: unknown[] | null;
public static Fields = {
  ...BaseEntity.Fields,
      skipWorkspaceId: `skipWorkspaceId`,
      workspaceId: `workspaceId`,
      token: `token`,
          securityId: `securityId`,
      security$: `security`,
        security: SecurityModel.Fields,
      capabilities: `capabilities`,
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
      "name": "security",
      "type": "one",
      "target": "SecurityModel",
      "computedType": "SecurityModel",
      "gormMap": {}
    },
    {
      "name": "capabilities",
      "type": "arrayP",
      "primitive": "PermissionInfo",
      "computedType": "unknown[]",
      "gormMap": {}
    }
  ]
}
}
