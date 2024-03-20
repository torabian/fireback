import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type AcceptInviteDtoKeys =
  keyof typeof AcceptInviteDto.Fields;
export class AcceptInviteDto extends BaseDto {
  public inviteUniqueId?: string | null;
  public visibility?: string | null;
  public updated?: number | null;
  public created?: number | null;
public static Fields = {
  ...BaseEntity.Fields,
      inviteUniqueId: 'inviteUniqueId',
      visibility: 'visibility',
      updated: 'updated',
      created: 'created',
}
  public static definition = {
  "name": "acceptInvite",
  "fields": [
    {
      "name": "inviteUniqueId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "visibility",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "updated",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "created",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    }
  ]
}
}
