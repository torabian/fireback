/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { IError, QueryFilterRequest, RemoveReply } from "../../core/common";

export const protobufPackage = "";

export interface KeyboardShortcutCreateReply {
  data: KeyboardShortcutEntity | undefined;
  error: IError | undefined;
}

export interface KeyboardShortcutReply {
  data: KeyboardShortcutEntity | undefined;
  error: IError | undefined;
}

export interface KeyboardShortcutQueryReply {
  items: KeyboardShortcutEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface KeyboardShortcutEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: KeyboardShortcutEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: KeyboardShortcutEntityPolyglot[];
  /** @tag(  yaml:"os"  ) */
  os?: string | undefined;
  /** @tag(  yaml:"host"  ) */
  host?: string | undefined;
  /** This is an object, another entity needs to be created for */
  defaultCombination: KeyboardShortcutDefaultCombinationEntity | undefined;
  /** This is an object, another entity needs to be created for */
  userCombination: KeyboardShortcutUserCombinationEntity | undefined;
  /** @tag(translate:"true"  yaml:"action"  ) */
  action?: string | undefined;
  /** @tag(  yaml:"actionKey"  ) */
  actionKey?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

/** Because it has translation field, we need a translation table for this */
export interface KeyboardShortcutEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"action" json:"action"); */
  action: string;
}

export interface KeyboardShortcutDefaultCombinationEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: KeyboardShortcutDefaultCombinationEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"altKey") */
  altKey?: boolean | undefined;
  /** @tag(  yaml:"key"  ) */
  key?: string | undefined;
  /** @tag(  yaml:"metaKey") */
  metaKey?: boolean | undefined;
  /** @tag(  yaml:"shiftKey") */
  shiftKey?: boolean | undefined;
  /** @tag(  yaml:"ctrlKey") */
  ctrlKey?: boolean | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface KeyboardShortcutUserCombinationEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: KeyboardShortcutUserCombinationEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"altKey") */
  altKey?: boolean | undefined;
  /** @tag(  yaml:"key"  ) */
  key?: string | undefined;
  /** @tag(  yaml:"metaKey") */
  metaKey?: boolean | undefined;
  /** @tag(  yaml:"shiftKey") */
  shiftKey?: boolean | undefined;
  /** @tag(  yaml:"ctrlKey") */
  ctrlKey?: boolean | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

function createBaseKeyboardShortcutCreateReply(): KeyboardShortcutCreateReply {
  return { data: undefined, error: undefined };
}

export const KeyboardShortcutCreateReply = {
  encode(
    message: KeyboardShortcutCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      KeyboardShortcutEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): KeyboardShortcutCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyboardShortcutCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = KeyboardShortcutEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeyboardShortcutCreateReply {
    return {
      data: isSet(object.data)
        ? KeyboardShortcutEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: KeyboardShortcutCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? KeyboardShortcutEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<KeyboardShortcutCreateReply>, I>>(
    base?: I
  ): KeyboardShortcutCreateReply {
    return KeyboardShortcutCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<KeyboardShortcutCreateReply>, I>>(
    object: I
  ): KeyboardShortcutCreateReply {
    const message = createBaseKeyboardShortcutCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? KeyboardShortcutEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseKeyboardShortcutReply(): KeyboardShortcutReply {
  return { data: undefined, error: undefined };
}

export const KeyboardShortcutReply = {
  encode(
    message: KeyboardShortcutReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      KeyboardShortcutEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): KeyboardShortcutReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyboardShortcutReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = KeyboardShortcutEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeyboardShortcutReply {
    return {
      data: isSet(object.data)
        ? KeyboardShortcutEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: KeyboardShortcutReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? KeyboardShortcutEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<KeyboardShortcutReply>, I>>(
    base?: I
  ): KeyboardShortcutReply {
    return KeyboardShortcutReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<KeyboardShortcutReply>, I>>(
    object: I
  ): KeyboardShortcutReply {
    const message = createBaseKeyboardShortcutReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? KeyboardShortcutEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseKeyboardShortcutQueryReply(): KeyboardShortcutQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const KeyboardShortcutQueryReply = {
  encode(
    message: KeyboardShortcutQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      KeyboardShortcutEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): KeyboardShortcutQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyboardShortcutQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            KeyboardShortcutEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeyboardShortcutQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => KeyboardShortcutEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: KeyboardShortcutQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? KeyboardShortcutEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<KeyboardShortcutQueryReply>, I>>(
    base?: I
  ): KeyboardShortcutQueryReply {
    return KeyboardShortcutQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<KeyboardShortcutQueryReply>, I>>(
    object: I
  ): KeyboardShortcutQueryReply {
    const message = createBaseKeyboardShortcutQueryReply();
    message.items =
      object.items?.map((e) => KeyboardShortcutEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseKeyboardShortcutEntity(): KeyboardShortcutEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    os: undefined,
    host: undefined,
    defaultCombination: undefined,
    userCombination: undefined,
    action: undefined,
    actionKey: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const KeyboardShortcutEntity = {
  encode(
    message: KeyboardShortcutEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      KeyboardShortcutEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.translations) {
      KeyboardShortcutEntityPolyglot.encode(
        v!,
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.os !== undefined) {
      writer.uint32(82).string(message.os);
    }
    if (message.host !== undefined) {
      writer.uint32(90).string(message.host);
    }
    if (message.defaultCombination !== undefined) {
      KeyboardShortcutDefaultCombinationEntity.encode(
        message.defaultCombination,
        writer.uint32(98).fork()
      ).ldelim();
    }
    if (message.userCombination !== undefined) {
      KeyboardShortcutUserCombinationEntity.encode(
        message.userCombination,
        writer.uint32(106).fork()
      ).ldelim();
    }
    if (message.action !== undefined) {
      writer.uint32(114).string(message.action);
    }
    if (message.actionKey !== undefined) {
      writer.uint32(122).string(message.actionKey);
    }
    if (message.rank !== 0) {
      writer.uint32(128).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(136).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(144).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(154).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(162).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): KeyboardShortcutEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyboardShortcutEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = KeyboardShortcutEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            KeyboardShortcutEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.os = reader.string();
          break;
        case 11:
          message.host = reader.string();
          break;
        case 12:
          message.defaultCombination =
            KeyboardShortcutDefaultCombinationEntity.decode(
              reader,
              reader.uint32()
            );
          break;
        case 13:
          message.userCombination =
            KeyboardShortcutUserCombinationEntity.decode(
              reader,
              reader.uint32()
            );
          break;
        case 14:
          message.action = reader.string();
          break;
        case 15:
          message.actionKey = reader.string();
          break;
        case 16:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 19:
          message.createdFormatted = reader.string();
          break;
        case 20:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeyboardShortcutEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? KeyboardShortcutEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            KeyboardShortcutEntityPolyglot.fromJSON(e)
          )
        : [],
      os: isSet(object.os) ? String(object.os) : undefined,
      host: isSet(object.host) ? String(object.host) : undefined,
      defaultCombination: isSet(object.defaultCombination)
        ? KeyboardShortcutDefaultCombinationEntity.fromJSON(
            object.defaultCombination
          )
        : undefined,
      userCombination: isSet(object.userCombination)
        ? KeyboardShortcutUserCombinationEntity.fromJSON(object.userCombination)
        : undefined,
      action: isSet(object.action) ? String(object.action) : undefined,
      actionKey: isSet(object.actionKey) ? String(object.actionKey) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: KeyboardShortcutEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? KeyboardShortcutEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? KeyboardShortcutEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.os !== undefined && (obj.os = message.os);
    message.host !== undefined && (obj.host = message.host);
    message.defaultCombination !== undefined &&
      (obj.defaultCombination = message.defaultCombination
        ? KeyboardShortcutDefaultCombinationEntity.toJSON(
            message.defaultCombination
          )
        : undefined);
    message.userCombination !== undefined &&
      (obj.userCombination = message.userCombination
        ? KeyboardShortcutUserCombinationEntity.toJSON(message.userCombination)
        : undefined);
    message.action !== undefined && (obj.action = message.action);
    message.actionKey !== undefined && (obj.actionKey = message.actionKey);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<KeyboardShortcutEntity>, I>>(
    base?: I
  ): KeyboardShortcutEntity {
    return KeyboardShortcutEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<KeyboardShortcutEntity>, I>>(
    object: I
  ): KeyboardShortcutEntity {
    const message = createBaseKeyboardShortcutEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? KeyboardShortcutEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) =>
        KeyboardShortcutEntityPolyglot.fromPartial(e)
      ) || [];
    message.os = object.os ?? undefined;
    message.host = object.host ?? undefined;
    message.defaultCombination =
      object.defaultCombination !== undefined &&
      object.defaultCombination !== null
        ? KeyboardShortcutDefaultCombinationEntity.fromPartial(
            object.defaultCombination
          )
        : undefined;
    message.userCombination =
      object.userCombination !== undefined && object.userCombination !== null
        ? KeyboardShortcutUserCombinationEntity.fromPartial(
            object.userCombination
          )
        : undefined;
    message.action = object.action ?? undefined;
    message.actionKey = object.actionKey ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseKeyboardShortcutEntityPolyglot(): KeyboardShortcutEntityPolyglot {
  return { linkerId: "", languageId: "", action: "" };
}

export const KeyboardShortcutEntityPolyglot = {
  encode(
    message: KeyboardShortcutEntityPolyglot,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.linkerId !== "") {
      writer.uint32(10).string(message.linkerId);
    }
    if (message.languageId !== "") {
      writer.uint32(18).string(message.languageId);
    }
    if (message.action !== "") {
      writer.uint32(26).string(message.action);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): KeyboardShortcutEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyboardShortcutEntityPolyglot();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.linkerId = reader.string();
          break;
        case 2:
          message.languageId = reader.string();
          break;
        case 3:
          message.action = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeyboardShortcutEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      action: isSet(object.action) ? String(object.action) : "",
    };
  },

  toJSON(message: KeyboardShortcutEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.action !== undefined && (obj.action = message.action);
    return obj;
  },

  create<I extends Exact<DeepPartial<KeyboardShortcutEntityPolyglot>, I>>(
    base?: I
  ): KeyboardShortcutEntityPolyglot {
    return KeyboardShortcutEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<KeyboardShortcutEntityPolyglot>, I>>(
    object: I
  ): KeyboardShortcutEntityPolyglot {
    const message = createBaseKeyboardShortcutEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.action = object.action ?? "";
    return message;
  },
};

function createBaseKeyboardShortcutDefaultCombinationEntity(): KeyboardShortcutDefaultCombinationEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    altKey: undefined,
    key: undefined,
    metaKey: undefined,
    shiftKey: undefined,
    ctrlKey: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const KeyboardShortcutDefaultCombinationEntity = {
  encode(
    message: KeyboardShortcutDefaultCombinationEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      KeyboardShortcutDefaultCombinationEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.altKey !== undefined) {
      writer.uint32(72).bool(message.altKey);
    }
    if (message.key !== undefined) {
      writer.uint32(82).string(message.key);
    }
    if (message.metaKey !== undefined) {
      writer.uint32(88).bool(message.metaKey);
    }
    if (message.shiftKey !== undefined) {
      writer.uint32(96).bool(message.shiftKey);
    }
    if (message.ctrlKey !== undefined) {
      writer.uint32(104).bool(message.ctrlKey);
    }
    if (message.rank !== 0) {
      writer.uint32(112).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(120).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(128).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(138).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(146).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): KeyboardShortcutDefaultCombinationEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyboardShortcutDefaultCombinationEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = KeyboardShortcutDefaultCombinationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.altKey = reader.bool();
          break;
        case 10:
          message.key = reader.string();
          break;
        case 11:
          message.metaKey = reader.bool();
          break;
        case 12:
          message.shiftKey = reader.bool();
          break;
        case 13:
          message.ctrlKey = reader.bool();
          break;
        case 14:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.createdFormatted = reader.string();
          break;
        case 18:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeyboardShortcutDefaultCombinationEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? KeyboardShortcutDefaultCombinationEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      altKey: isSet(object.altKey) ? Boolean(object.altKey) : undefined,
      key: isSet(object.key) ? String(object.key) : undefined,
      metaKey: isSet(object.metaKey) ? Boolean(object.metaKey) : undefined,
      shiftKey: isSet(object.shiftKey) ? Boolean(object.shiftKey) : undefined,
      ctrlKey: isSet(object.ctrlKey) ? Boolean(object.ctrlKey) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: KeyboardShortcutDefaultCombinationEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? KeyboardShortcutDefaultCombinationEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.altKey !== undefined && (obj.altKey = message.altKey);
    message.key !== undefined && (obj.key = message.key);
    message.metaKey !== undefined && (obj.metaKey = message.metaKey);
    message.shiftKey !== undefined && (obj.shiftKey = message.shiftKey);
    message.ctrlKey !== undefined && (obj.ctrlKey = message.ctrlKey);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<
    I extends Exact<DeepPartial<KeyboardShortcutDefaultCombinationEntity>, I>
  >(base?: I): KeyboardShortcutDefaultCombinationEntity {
    return KeyboardShortcutDefaultCombinationEntity.fromPartial(base ?? {});
  },

  fromPartial<
    I extends Exact<DeepPartial<KeyboardShortcutDefaultCombinationEntity>, I>
  >(object: I): KeyboardShortcutDefaultCombinationEntity {
    const message = createBaseKeyboardShortcutDefaultCombinationEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? KeyboardShortcutDefaultCombinationEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.altKey = object.altKey ?? undefined;
    message.key = object.key ?? undefined;
    message.metaKey = object.metaKey ?? undefined;
    message.shiftKey = object.shiftKey ?? undefined;
    message.ctrlKey = object.ctrlKey ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseKeyboardShortcutUserCombinationEntity(): KeyboardShortcutUserCombinationEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    altKey: undefined,
    key: undefined,
    metaKey: undefined,
    shiftKey: undefined,
    ctrlKey: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const KeyboardShortcutUserCombinationEntity = {
  encode(
    message: KeyboardShortcutUserCombinationEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      KeyboardShortcutUserCombinationEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.altKey !== undefined) {
      writer.uint32(72).bool(message.altKey);
    }
    if (message.key !== undefined) {
      writer.uint32(82).string(message.key);
    }
    if (message.metaKey !== undefined) {
      writer.uint32(88).bool(message.metaKey);
    }
    if (message.shiftKey !== undefined) {
      writer.uint32(96).bool(message.shiftKey);
    }
    if (message.ctrlKey !== undefined) {
      writer.uint32(104).bool(message.ctrlKey);
    }
    if (message.rank !== 0) {
      writer.uint32(112).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(120).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(128).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(138).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(146).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): KeyboardShortcutUserCombinationEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyboardShortcutUserCombinationEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = KeyboardShortcutUserCombinationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.altKey = reader.bool();
          break;
        case 10:
          message.key = reader.string();
          break;
        case 11:
          message.metaKey = reader.bool();
          break;
        case 12:
          message.shiftKey = reader.bool();
          break;
        case 13:
          message.ctrlKey = reader.bool();
          break;
        case 14:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.createdFormatted = reader.string();
          break;
        case 18:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): KeyboardShortcutUserCombinationEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? KeyboardShortcutUserCombinationEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      altKey: isSet(object.altKey) ? Boolean(object.altKey) : undefined,
      key: isSet(object.key) ? String(object.key) : undefined,
      metaKey: isSet(object.metaKey) ? Boolean(object.metaKey) : undefined,
      shiftKey: isSet(object.shiftKey) ? Boolean(object.shiftKey) : undefined,
      ctrlKey: isSet(object.ctrlKey) ? Boolean(object.ctrlKey) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: KeyboardShortcutUserCombinationEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? KeyboardShortcutUserCombinationEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.altKey !== undefined && (obj.altKey = message.altKey);
    message.key !== undefined && (obj.key = message.key);
    message.metaKey !== undefined && (obj.metaKey = message.metaKey);
    message.shiftKey !== undefined && (obj.shiftKey = message.shiftKey);
    message.ctrlKey !== undefined && (obj.ctrlKey = message.ctrlKey);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<
    I extends Exact<DeepPartial<KeyboardShortcutUserCombinationEntity>, I>
  >(base?: I): KeyboardShortcutUserCombinationEntity {
    return KeyboardShortcutUserCombinationEntity.fromPartial(base ?? {});
  },

  fromPartial<
    I extends Exact<DeepPartial<KeyboardShortcutUserCombinationEntity>, I>
  >(object: I): KeyboardShortcutUserCombinationEntity {
    const message = createBaseKeyboardShortcutUserCombinationEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? KeyboardShortcutUserCombinationEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.altKey = object.altKey ?? undefined;
    message.key = object.key ?? undefined;
    message.metaKey = object.metaKey ?? undefined;
    message.shiftKey = object.shiftKey ?? undefined;
    message.ctrlKey = object.ctrlKey ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

export interface KeyboardShortcuts {
  KeyboardShortcutActionCreate(
    request: KeyboardShortcutEntity
  ): Promise<KeyboardShortcutCreateReply>;
  KeyboardShortcutActionUpdate(
    request: KeyboardShortcutEntity
  ): Promise<KeyboardShortcutCreateReply>;
  KeyboardShortcutActionQuery(
    request: QueryFilterRequest
  ): Promise<KeyboardShortcutQueryReply>;
  KeyboardShortcutActionGetOne(
    request: QueryFilterRequest
  ): Promise<KeyboardShortcutReply>;
  KeyboardShortcutActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class KeyboardShortcutsClientImpl implements KeyboardShortcuts {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "KeyboardShortcuts";
    this.rpc = rpc;
    this.KeyboardShortcutActionCreate =
      this.KeyboardShortcutActionCreate.bind(this);
    this.KeyboardShortcutActionUpdate =
      this.KeyboardShortcutActionUpdate.bind(this);
    this.KeyboardShortcutActionQuery =
      this.KeyboardShortcutActionQuery.bind(this);
    this.KeyboardShortcutActionGetOne =
      this.KeyboardShortcutActionGetOne.bind(this);
    this.KeyboardShortcutActionRemove =
      this.KeyboardShortcutActionRemove.bind(this);
  }
  KeyboardShortcutActionCreate(
    request: KeyboardShortcutEntity
  ): Promise<KeyboardShortcutCreateReply> {
    const data = KeyboardShortcutEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "KeyboardShortcutActionCreate",
      data
    );
    return promise.then((data) =>
      KeyboardShortcutCreateReply.decode(new _m0.Reader(data))
    );
  }

  KeyboardShortcutActionUpdate(
    request: KeyboardShortcutEntity
  ): Promise<KeyboardShortcutCreateReply> {
    const data = KeyboardShortcutEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "KeyboardShortcutActionUpdate",
      data
    );
    return promise.then((data) =>
      KeyboardShortcutCreateReply.decode(new _m0.Reader(data))
    );
  }

  KeyboardShortcutActionQuery(
    request: QueryFilterRequest
  ): Promise<KeyboardShortcutQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "KeyboardShortcutActionQuery",
      data
    );
    return promise.then((data) =>
      KeyboardShortcutQueryReply.decode(new _m0.Reader(data))
    );
  }

  KeyboardShortcutActionGetOne(
    request: QueryFilterRequest
  ): Promise<KeyboardShortcutReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "KeyboardShortcutActionGetOne",
      data
    );
    return promise.then((data) =>
      KeyboardShortcutReply.decode(new _m0.Reader(data))
    );
  }

  KeyboardShortcutActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "KeyboardShortcutActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & {
      [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
    };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error(
      "Value is larger than Number.MAX_SAFE_INTEGER"
    );
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
