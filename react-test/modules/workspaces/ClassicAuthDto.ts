import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type ClassicAuthDtoKeys =
  keyof typeof ClassicAuthDto.Fields;
export class ClassicAuthDto extends BaseDto {
  public value?: string | null;
  public password?: string | null;
  public firstName?: string | null;
  public lastName?: string | null;
  public inviteId?: string | null;
  public publicJoinKeyId?: string | null;
  public workspaceTypeId?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      value: 'value',
      password: 'password',
      firstName: 'firstName',
      lastName: 'lastName',
      inviteId: 'inviteId',
      publicJoinKeyId: 'publicJoinKeyId',
      workspaceTypeId: 'workspaceTypeId',
}
  public static definition = {
  "name": "classicAuth",
  "fields": [
    {
      "name": "value",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "password",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
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
      "name": "inviteId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "publicJoinKeyId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "workspaceTypeId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
