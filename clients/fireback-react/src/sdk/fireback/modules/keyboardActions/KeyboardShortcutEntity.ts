import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
export class KeyboardShortcutDefaultCombination extends BaseEntity {
  public altKey?: boolean | null;
  public key?: string | null;
  public metaKey?: boolean | null;
  public shiftKey?: boolean | null;
  public ctrlKey?: boolean | null;
}
export class KeyboardShortcutUserCombination extends BaseEntity {
  public altKey?: boolean | null;
  public key?: string | null;
  public metaKey?: boolean | null;
  public shiftKey?: boolean | null;
  public ctrlKey?: boolean | null;
}
// Class body
export type KeyboardShortcutEntityKeys =
  keyof typeof KeyboardShortcutEntity.Fields;
export class KeyboardShortcutEntity extends BaseEntity {
  public os?: string | null;
  public host?: string | null;
  public defaultCombination?: KeyboardShortcutDefaultCombination | null;
  public userCombination?: KeyboardShortcutUserCombination | null;
  public action?: string | null;
  public actionKey?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcut/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcut/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcut/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcuts`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "keyboard-shortcut/edit/:uniqueId",
      Rcreate: "keyboard-shortcut/new",
      Rsingle: "keyboard-shortcut/:uniqueId",
      Rquery: "keyboard-shortcuts",
      rDefaultCombinationCreate: "keyboard-shortcut/:linkerId/default_combination/new",
      rDefaultCombinationEdit: "keyboard-shortcut/:linkerId/default_combination/edit/:uniqueId",
      editDefaultCombination(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcut/${linkerId}/default_combination/edit/${uniqueId}`;
      },
      createDefaultCombination(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcut/${linkerId}/default_combination/new`;
      },
      rUserCombinationCreate: "keyboard-shortcut/:linkerId/user_combination/new",
      rUserCombinationEdit: "keyboard-shortcut/:linkerId/user_combination/edit/:uniqueId",
      editUserCombination(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcut/${linkerId}/user_combination/edit/${uniqueId}`;
      },
      createUserCombination(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/keyboard-shortcut/${linkerId}/user_combination/new`;
      },
  };
  public static definition = {
  "name": "keyboardShortcut",
  "queryScope": "public",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "os",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "host",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "linkedTo": "KeyboardShortcutEntity",
      "name": "defaultCombination",
      "type": "object",
      "computedType": "KeyboardShortcutDefaultCombination",
      "gormMap": {},
      "fullName": "KeyboardShortcutDefaultCombination",
      "fields": [
        {
          "name": "altKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        },
        {
          "name": "key",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "metaKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        },
        {
          "name": "shiftKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        },
        {
          "name": "ctrlKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        }
      ]
    },
    {
      "linkedTo": "KeyboardShortcutEntity",
      "name": "userCombination",
      "type": "object",
      "computedType": "KeyboardShortcutUserCombination",
      "gormMap": {},
      "fullName": "KeyboardShortcutUserCombination",
      "fields": [
        {
          "name": "altKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        },
        {
          "name": "key",
          "type": "string",
          "computedType": "string",
          "gormMap": {}
        },
        {
          "name": "metaKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        },
        {
          "name": "shiftKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        },
        {
          "name": "ctrlKey",
          "type": "bool",
          "computedType": "boolean",
          "gormMap": {}
        }
      ]
    },
    {
      "name": "action",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "actionKey",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliShort": "kbshort",
  "cliDescription": "Manage the keyboard shortcuts in web and desktop apps (accessibility)"
}
public static Fields = {
  ...BaseEntity.Fields,
      os: 'os',
      host: 'host',
      defaultCombination$: 'defaultCombination',
      defaultCombination: {
  ...BaseEntity.Fields,
      altKey: 'altKey',
      key: 'key',
      metaKey: 'metaKey',
      shiftKey: 'shiftKey',
      ctrlKey: 'ctrlKey',
      },
      userCombination$: 'userCombination',
      userCombination: {
  ...BaseEntity.Fields,
      altKey: 'altKey',
      key: 'key',
      metaKey: 'metaKey',
      shiftKey: 'shiftKey',
      ctrlKey: 'ctrlKey',
      },
      action: 'action',
      actionKey: 'actionKey',
}
}
