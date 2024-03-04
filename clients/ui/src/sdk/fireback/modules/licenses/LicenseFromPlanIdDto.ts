import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type LicenseFromPlanIdDtoKeys =
  keyof typeof LicenseFromPlanIdDto.Fields;
export class LicenseFromPlanIdDto extends BaseDto {
  public machineId?: string | null;
  public email?: string | null;
  public owner?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      machineId: 'machineId',
      email: 'email',
      owner: 'owner',
}
  public static definition = {
  "name": "licenseFromPlanId",
  "fields": [
    {
      "name": "machineId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "email",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "owner",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
