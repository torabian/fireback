import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
export class UserActivityActivities extends BaseDto {
  public uniqueId?: string | null;
  public activity?: number | null;
}
// Class body
export type UserActivityDtoKeys =
  keyof typeof UserActivityDto.Fields;
export class UserActivityDto extends BaseDto {
  public activities?: UserActivityActivities[] | null;
public static Fields = {
  ...BaseEntity.Fields,
      activities$: 'activities',
      activities: {
  ...BaseEntity.Fields,
      uniqueId: 'uniqueId',
      activity: 'activity',
      },
}
  public static definition = {
  "name": "userActivity",
  "fields": [
    {
      "linkedTo": "UserActivityDto",
      "name": "activities",
      "type": "array",
      "computedType": "UserActivityActivities[]",
      "gormMap": {},
      "fullName": "UserActivityActivities",
      "fields": [
        {
          "name": "uniqueId",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "activity",
          "type": "int64",
          "computedType": "number",
          "gormMap": {}
        }
      ]
    }
  ]
}
}
