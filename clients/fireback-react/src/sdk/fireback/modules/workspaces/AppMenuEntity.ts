import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CapabilityEntity,
} from "./CapabilityEntity"
// In this section we have sub entities related to this object
// Class body
export type AppMenuEntityKeys =
  keyof typeof AppMenuEntity.Fields;
export class AppMenuEntity extends BaseEntity {
  public href?: string | null;
  public icon?: string | null;
  public label?: string | null;
  public activeMatcher?: string | null;
  public applyType?: string | null;
  public capability?: CapabilityEntity | null;
      capabilityId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/app-menu/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/app-menu/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/app-menu/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/app-menus`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "app-menu/edit/:uniqueId",
      Rcreate: "app-menu/new",
      Rsingle: "app-menu/:uniqueId",
      Rquery: "app-menus",
  };
  public static definition = {
  "name": "appMenu",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/CapabilityDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "href",
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
      "name": "label",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "activeMatcher",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "applyType",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "capability",
      "type": "one",
      "target": "CapabilityEntity",
      "computedType": "CapabilityEntity",
      "gormMap": {}
    }
  ],
  "cliDescription": "Manages the menus in the app, (for example tab views, sidebar items, etc.)",
  "cte": true
}
public static Fields = {
  ...BaseEntity.Fields,
      href: 'href',
      icon: 'icon',
      label: 'label',
      activeMatcher: 'activeMatcher',
      applyType: 'applyType',
          capabilityId: 'capabilityId',
      capability$: 'capability',
      capability: CapabilityEntity.Fields,
}
}
