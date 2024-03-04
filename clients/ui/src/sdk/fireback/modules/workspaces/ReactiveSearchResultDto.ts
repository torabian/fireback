import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type ReactiveSearchResultDtoKeys =
  keyof typeof ReactiveSearchResultDto.Fields;
export class ReactiveSearchResultDto extends BaseDto {
  public uniqueId?: string | null;
  public phrase?: string | null;
  public icon?: string | null;
  public description?: string | null;
  public group?: string | null;
  public uiLocation?: string | null;
  public actionFn?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      uniqueId: 'uniqueId',
      phrase: 'phrase',
      icon: 'icon',
      description: 'description',
      group: 'group',
      uiLocation: 'uiLocation',
      actionFn: 'actionFn',
}
  public static definition = {
  "name": "reactiveSearchResult",
  "fields": [
    {
      "name": "uniqueId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "phrase",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "icon",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "description",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "group",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "uiLocation",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "actionFn",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
