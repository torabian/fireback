import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type UserRoleWorkspaceDtoKeys =
  keyof typeof UserRoleWorkspaceDto.Fields;
export class UserRoleWorkspaceDto extends BaseDto {
  public roleId?: string | null;
  public capabilities?: string[] | null;
public static Fields = {
  ...BaseEntity.Fields,
      roleId: 'roleId',
      capabilities: 'capabilities',
}
  public static definition = {
  "name": "userRoleWorkspace",
  "fields": [
    {
      "name": "roleId",
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
