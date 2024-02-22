/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { IError, QueryFilterRequest, RemoveReply } from "../../core/common";

export const protobufPackage = "";

export interface WidgetAreaCreateReply {
  data: WidgetAreaEntity | undefined;
  error: IError | undefined;
}

export interface WidgetAreaReply {
  data: WidgetAreaEntity | undefined;
  error: IError | undefined;
}

export interface WidgetAreaQueryReply {
  items: WidgetAreaEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface WidgetAreaEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: WidgetAreaEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: WidgetAreaEntityPolyglot[];
  /** @tag(translate:"true"  yaml:"name"  ) */
  name?: string | undefined;
  /** @tag(  yaml:"layouts"  ) */
  layouts?: string | undefined;
  /** repeated WidgetAreaWidgetsEntity widgets = 12; // @tag(gorm:"foreignKey:LinkerId;references:UniqueId" yaml:"widgets") */
  widgets: WidgetAreaWidgetsEntity[];
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
export interface WidgetAreaEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"name" json:"name"); */
  name: string;
}

export interface WidgetAreaWidgetsEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: WidgetAreaWidgetsEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: WidgetAreaWidgetsEntityPolyglot[];
  /** @tag(translate:"true"  yaml:"title"  ) */
  title?: string | undefined;
  /** One 2 one to external entity */
  widgetId?: string | undefined;
  /** @tag(gorm:"foreignKey:WidgetId;references:UniqueId" json:"widget" yaml:"widget" fbtype:"one") */
  widget: WidgetEntity | undefined;
  /** @tag(  yaml:"x"  ) */
  x?: number | undefined;
  /** @tag(  yaml:"y"  ) */
  y?: number | undefined;
  /** @tag(  yaml:"w"  ) */
  w?: number | undefined;
  /** @tag(  yaml:"h"  ) */
  h?: number | undefined;
  /** @tag(  yaml:"data"  ) */
  data?: string | undefined;
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
export interface WidgetAreaWidgetsEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"title" json:"title"); */
  title: string;
}

export interface WidgetCreateReply {
  data: WidgetEntity | undefined;
  error: IError | undefined;
}

export interface WidgetReply {
  data: WidgetEntity | undefined;
  error: IError | undefined;
}

export interface WidgetQueryReply {
  items: WidgetEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface WidgetEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: WidgetEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: WidgetEntityPolyglot[];
  /** @tag(translate:"true"  yaml:"name"  ) */
  name?: string | undefined;
  /** @tag(  yaml:"family"  ) */
  family?: string | undefined;
  /** @tag(  yaml:"providerKey"  ) */
  providerKey?: string | undefined;
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
export interface WidgetEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"name" json:"name"); */
  name: string;
}

function createBaseWidgetAreaCreateReply(): WidgetAreaCreateReply {
  return { data: undefined, error: undefined };
}

export const WidgetAreaCreateReply = {
  encode(
    message: WidgetAreaCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WidgetAreaEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WidgetAreaCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetAreaCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WidgetAreaEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): WidgetAreaCreateReply {
    return {
      data: isSet(object.data)
        ? WidgetAreaEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WidgetAreaCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WidgetAreaEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WidgetAreaCreateReply>, I>>(
    base?: I
  ): WidgetAreaCreateReply {
    return WidgetAreaCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetAreaCreateReply>, I>>(
    object: I
  ): WidgetAreaCreateReply {
    const message = createBaseWidgetAreaCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WidgetAreaEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWidgetAreaReply(): WidgetAreaReply {
  return { data: undefined, error: undefined };
}

export const WidgetAreaReply = {
  encode(
    message: WidgetAreaReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WidgetAreaEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WidgetAreaReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetAreaReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WidgetAreaEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): WidgetAreaReply {
    return {
      data: isSet(object.data)
        ? WidgetAreaEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WidgetAreaReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WidgetAreaEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WidgetAreaReply>, I>>(
    base?: I
  ): WidgetAreaReply {
    return WidgetAreaReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetAreaReply>, I>>(
    object: I
  ): WidgetAreaReply {
    const message = createBaseWidgetAreaReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WidgetAreaEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWidgetAreaQueryReply(): WidgetAreaQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const WidgetAreaQueryReply = {
  encode(
    message: WidgetAreaQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      WidgetAreaEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): WidgetAreaQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetAreaQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(WidgetAreaEntity.decode(reader, reader.uint32()));
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

  fromJSON(object: any): WidgetAreaQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => WidgetAreaEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WidgetAreaQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? WidgetAreaEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<WidgetAreaQueryReply>, I>>(
    base?: I
  ): WidgetAreaQueryReply {
    return WidgetAreaQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetAreaQueryReply>, I>>(
    object: I
  ): WidgetAreaQueryReply {
    const message = createBaseWidgetAreaQueryReply();
    message.items =
      object.items?.map((e) => WidgetAreaEntity.fromPartial(e)) || [];
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

function createBaseWidgetAreaEntity(): WidgetAreaEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    name: undefined,
    layouts: undefined,
    widgets: [],
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const WidgetAreaEntity = {
  encode(
    message: WidgetAreaEntity,
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
      WidgetAreaEntity.encode(
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
      WidgetAreaEntityPolyglot.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.name !== undefined) {
      writer.uint32(82).string(message.name);
    }
    if (message.layouts !== undefined) {
      writer.uint32(90).string(message.layouts);
    }
    for (const v of message.widgets) {
      WidgetAreaWidgetsEntity.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(104).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(112).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(120).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(130).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(138).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WidgetAreaEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetAreaEntity();
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
          message.parent = WidgetAreaEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            WidgetAreaEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.name = reader.string();
          break;
        case 11:
          message.layouts = reader.string();
          break;
        case 12:
          message.widgets.push(
            WidgetAreaWidgetsEntity.decode(reader, reader.uint32())
          );
          break;
        case 13:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.createdFormatted = reader.string();
          break;
        case 17:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WidgetAreaEntity {
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
        ? WidgetAreaEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            WidgetAreaEntityPolyglot.fromJSON(e)
          )
        : [],
      name: isSet(object.name) ? String(object.name) : undefined,
      layouts: isSet(object.layouts) ? String(object.layouts) : undefined,
      widgets: Array.isArray(object?.widgets)
        ? object.widgets.map((e: any) => WidgetAreaWidgetsEntity.fromJSON(e))
        : [],
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

  toJSON(message: WidgetAreaEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? WidgetAreaEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? WidgetAreaEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.name !== undefined && (obj.name = message.name);
    message.layouts !== undefined && (obj.layouts = message.layouts);
    if (message.widgets) {
      obj.widgets = message.widgets.map((e) =>
        e ? WidgetAreaWidgetsEntity.toJSON(e) : undefined
      );
    } else {
      obj.widgets = [];
    }
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

  create<I extends Exact<DeepPartial<WidgetAreaEntity>, I>>(
    base?: I
  ): WidgetAreaEntity {
    return WidgetAreaEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetAreaEntity>, I>>(
    object: I
  ): WidgetAreaEntity {
    const message = createBaseWidgetAreaEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? WidgetAreaEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) =>
        WidgetAreaEntityPolyglot.fromPartial(e)
      ) || [];
    message.name = object.name ?? undefined;
    message.layouts = object.layouts ?? undefined;
    message.widgets =
      object.widgets?.map((e) => WidgetAreaWidgetsEntity.fromPartial(e)) || [];
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseWidgetAreaEntityPolyglot(): WidgetAreaEntityPolyglot {
  return { linkerId: "", languageId: "", name: "" };
}

export const WidgetAreaEntityPolyglot = {
  encode(
    message: WidgetAreaEntityPolyglot,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.linkerId !== "") {
      writer.uint32(10).string(message.linkerId);
    }
    if (message.languageId !== "") {
      writer.uint32(18).string(message.languageId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WidgetAreaEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetAreaEntityPolyglot();
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
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WidgetAreaEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: WidgetAreaEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<WidgetAreaEntityPolyglot>, I>>(
    base?: I
  ): WidgetAreaEntityPolyglot {
    return WidgetAreaEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetAreaEntityPolyglot>, I>>(
    object: I
  ): WidgetAreaEntityPolyglot {
    const message = createBaseWidgetAreaEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseWidgetAreaWidgetsEntity(): WidgetAreaWidgetsEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    title: undefined,
    widgetId: undefined,
    widget: undefined,
    x: undefined,
    y: undefined,
    w: undefined,
    h: undefined,
    data: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const WidgetAreaWidgetsEntity = {
  encode(
    message: WidgetAreaWidgetsEntity,
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
      WidgetAreaWidgetsEntity.encode(
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
      WidgetAreaWidgetsEntityPolyglot.encode(
        v!,
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.title !== undefined) {
      writer.uint32(82).string(message.title);
    }
    if (message.widgetId !== undefined) {
      writer.uint32(98).string(message.widgetId);
    }
    if (message.widget !== undefined) {
      WidgetEntity.encode(message.widget, writer.uint32(106).fork()).ldelim();
    }
    if (message.x !== undefined) {
      writer.uint32(112).int64(message.x);
    }
    if (message.y !== undefined) {
      writer.uint32(120).int64(message.y);
    }
    if (message.w !== undefined) {
      writer.uint32(128).int64(message.w);
    }
    if (message.h !== undefined) {
      writer.uint32(136).int64(message.h);
    }
    if (message.data !== undefined) {
      writer.uint32(146).string(message.data);
    }
    if (message.rank !== 0) {
      writer.uint32(152).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(160).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(168).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(178).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(186).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WidgetAreaWidgetsEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetAreaWidgetsEntity();
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
          message.parent = WidgetAreaWidgetsEntity.decode(
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
            WidgetAreaWidgetsEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.title = reader.string();
          break;
        case 12:
          message.widgetId = reader.string();
          break;
        case 13:
          message.widget = WidgetEntity.decode(reader, reader.uint32());
          break;
        case 14:
          message.x = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.y = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.w = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.h = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.data = reader.string();
          break;
        case 19:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 20:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 21:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 22:
          message.createdFormatted = reader.string();
          break;
        case 23:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WidgetAreaWidgetsEntity {
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
        ? WidgetAreaWidgetsEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            WidgetAreaWidgetsEntityPolyglot.fromJSON(e)
          )
        : [],
      title: isSet(object.title) ? String(object.title) : undefined,
      widgetId: isSet(object.widgetId) ? String(object.widgetId) : undefined,
      widget: isSet(object.widget)
        ? WidgetEntity.fromJSON(object.widget)
        : undefined,
      x: isSet(object.x) ? Number(object.x) : undefined,
      y: isSet(object.y) ? Number(object.y) : undefined,
      w: isSet(object.w) ? Number(object.w) : undefined,
      h: isSet(object.h) ? Number(object.h) : undefined,
      data: isSet(object.data) ? String(object.data) : undefined,
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

  toJSON(message: WidgetAreaWidgetsEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? WidgetAreaWidgetsEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? WidgetAreaWidgetsEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.title !== undefined && (obj.title = message.title);
    message.widgetId !== undefined && (obj.widgetId = message.widgetId);
    message.widget !== undefined &&
      (obj.widget = message.widget
        ? WidgetEntity.toJSON(message.widget)
        : undefined);
    message.x !== undefined && (obj.x = Math.round(message.x));
    message.y !== undefined && (obj.y = Math.round(message.y));
    message.w !== undefined && (obj.w = Math.round(message.w));
    message.h !== undefined && (obj.h = Math.round(message.h));
    message.data !== undefined && (obj.data = message.data);
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

  create<I extends Exact<DeepPartial<WidgetAreaWidgetsEntity>, I>>(
    base?: I
  ): WidgetAreaWidgetsEntity {
    return WidgetAreaWidgetsEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetAreaWidgetsEntity>, I>>(
    object: I
  ): WidgetAreaWidgetsEntity {
    const message = createBaseWidgetAreaWidgetsEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? WidgetAreaWidgetsEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) =>
        WidgetAreaWidgetsEntityPolyglot.fromPartial(e)
      ) || [];
    message.title = object.title ?? undefined;
    message.widgetId = object.widgetId ?? undefined;
    message.widget =
      object.widget !== undefined && object.widget !== null
        ? WidgetEntity.fromPartial(object.widget)
        : undefined;
    message.x = object.x ?? undefined;
    message.y = object.y ?? undefined;
    message.w = object.w ?? undefined;
    message.h = object.h ?? undefined;
    message.data = object.data ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseWidgetAreaWidgetsEntityPolyglot(): WidgetAreaWidgetsEntityPolyglot {
  return { linkerId: "", languageId: "", title: "" };
}

export const WidgetAreaWidgetsEntityPolyglot = {
  encode(
    message: WidgetAreaWidgetsEntityPolyglot,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.linkerId !== "") {
      writer.uint32(10).string(message.linkerId);
    }
    if (message.languageId !== "") {
      writer.uint32(18).string(message.languageId);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WidgetAreaWidgetsEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetAreaWidgetsEntityPolyglot();
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
          message.title = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WidgetAreaWidgetsEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      title: isSet(object.title) ? String(object.title) : "",
    };
  },

  toJSON(message: WidgetAreaWidgetsEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.title !== undefined && (obj.title = message.title);
    return obj;
  },

  create<I extends Exact<DeepPartial<WidgetAreaWidgetsEntityPolyglot>, I>>(
    base?: I
  ): WidgetAreaWidgetsEntityPolyglot {
    return WidgetAreaWidgetsEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetAreaWidgetsEntityPolyglot>, I>>(
    object: I
  ): WidgetAreaWidgetsEntityPolyglot {
    const message = createBaseWidgetAreaWidgetsEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.title = object.title ?? "";
    return message;
  },
};

function createBaseWidgetCreateReply(): WidgetCreateReply {
  return { data: undefined, error: undefined };
}

export const WidgetCreateReply = {
  encode(
    message: WidgetCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WidgetEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WidgetCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WidgetEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): WidgetCreateReply {
    return {
      data: isSet(object.data) ? WidgetEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WidgetCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? WidgetEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WidgetCreateReply>, I>>(
    base?: I
  ): WidgetCreateReply {
    return WidgetCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetCreateReply>, I>>(
    object: I
  ): WidgetCreateReply {
    const message = createBaseWidgetCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WidgetEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWidgetReply(): WidgetReply {
  return { data: undefined, error: undefined };
}

export const WidgetReply = {
  encode(
    message: WidgetReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WidgetEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WidgetReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WidgetEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): WidgetReply {
    return {
      data: isSet(object.data) ? WidgetEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WidgetReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? WidgetEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WidgetReply>, I>>(base?: I): WidgetReply {
    return WidgetReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetReply>, I>>(
    object: I
  ): WidgetReply {
    const message = createBaseWidgetReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WidgetEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWidgetQueryReply(): WidgetQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const WidgetQueryReply = {
  encode(
    message: WidgetQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      WidgetEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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

  decode(input: _m0.Reader | Uint8Array, length?: number): WidgetQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(WidgetEntity.decode(reader, reader.uint32()));
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

  fromJSON(object: any): WidgetQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => WidgetEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WidgetQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? WidgetEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<WidgetQueryReply>, I>>(
    base?: I
  ): WidgetQueryReply {
    return WidgetQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetQueryReply>, I>>(
    object: I
  ): WidgetQueryReply {
    const message = createBaseWidgetQueryReply();
    message.items = object.items?.map((e) => WidgetEntity.fromPartial(e)) || [];
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

function createBaseWidgetEntity(): WidgetEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    name: undefined,
    family: undefined,
    providerKey: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const WidgetEntity = {
  encode(
    message: WidgetEntity,
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
      WidgetEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.translations) {
      WidgetEntityPolyglot.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.name !== undefined) {
      writer.uint32(82).string(message.name);
    }
    if (message.family !== undefined) {
      writer.uint32(90).string(message.family);
    }
    if (message.providerKey !== undefined) {
      writer.uint32(98).string(message.providerKey);
    }
    if (message.rank !== 0) {
      writer.uint32(104).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(112).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(120).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(130).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(138).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WidgetEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetEntity();
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
          message.parent = WidgetEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            WidgetEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.name = reader.string();
          break;
        case 11:
          message.family = reader.string();
          break;
        case 12:
          message.providerKey = reader.string();
          break;
        case 13:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.createdFormatted = reader.string();
          break;
        case 17:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WidgetEntity {
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
        ? WidgetEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) => WidgetEntityPolyglot.fromJSON(e))
        : [],
      name: isSet(object.name) ? String(object.name) : undefined,
      family: isSet(object.family) ? String(object.family) : undefined,
      providerKey: isSet(object.providerKey)
        ? String(object.providerKey)
        : undefined,
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

  toJSON(message: WidgetEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? WidgetEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? WidgetEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.name !== undefined && (obj.name = message.name);
    message.family !== undefined && (obj.family = message.family);
    message.providerKey !== undefined &&
      (obj.providerKey = message.providerKey);
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

  create<I extends Exact<DeepPartial<WidgetEntity>, I>>(
    base?: I
  ): WidgetEntity {
    return WidgetEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetEntity>, I>>(
    object: I
  ): WidgetEntity {
    const message = createBaseWidgetEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? WidgetEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) => WidgetEntityPolyglot.fromPartial(e)) ||
      [];
    message.name = object.name ?? undefined;
    message.family = object.family ?? undefined;
    message.providerKey = object.providerKey ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseWidgetEntityPolyglot(): WidgetEntityPolyglot {
  return { linkerId: "", languageId: "", name: "" };
}

export const WidgetEntityPolyglot = {
  encode(
    message: WidgetEntityPolyglot,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.linkerId !== "") {
      writer.uint32(10).string(message.linkerId);
    }
    if (message.languageId !== "") {
      writer.uint32(18).string(message.languageId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WidgetEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWidgetEntityPolyglot();
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
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WidgetEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: WidgetEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<WidgetEntityPolyglot>, I>>(
    base?: I
  ): WidgetEntityPolyglot {
    return WidgetEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WidgetEntityPolyglot>, I>>(
    object: I
  ): WidgetEntityPolyglot {
    const message = createBaseWidgetEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

export interface WidgetAreas {
  WidgetAreaActionCreate(
    request: WidgetAreaEntity
  ): Promise<WidgetAreaCreateReply>;
  WidgetAreaActionUpdate(
    request: WidgetAreaEntity
  ): Promise<WidgetAreaCreateReply>;
  WidgetAreaActionQuery(
    request: QueryFilterRequest
  ): Promise<WidgetAreaQueryReply>;
  WidgetAreaActionGetOne(request: QueryFilterRequest): Promise<WidgetAreaReply>;
  WidgetAreaActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class WidgetAreasClientImpl implements WidgetAreas {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "WidgetAreas";
    this.rpc = rpc;
    this.WidgetAreaActionCreate = this.WidgetAreaActionCreate.bind(this);
    this.WidgetAreaActionUpdate = this.WidgetAreaActionUpdate.bind(this);
    this.WidgetAreaActionQuery = this.WidgetAreaActionQuery.bind(this);
    this.WidgetAreaActionGetOne = this.WidgetAreaActionGetOne.bind(this);
    this.WidgetAreaActionRemove = this.WidgetAreaActionRemove.bind(this);
  }
  WidgetAreaActionCreate(
    request: WidgetAreaEntity
  ): Promise<WidgetAreaCreateReply> {
    const data = WidgetAreaEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WidgetAreaActionCreate",
      data
    );
    return promise.then((data) =>
      WidgetAreaCreateReply.decode(new _m0.Reader(data))
    );
  }

  WidgetAreaActionUpdate(
    request: WidgetAreaEntity
  ): Promise<WidgetAreaCreateReply> {
    const data = WidgetAreaEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WidgetAreaActionUpdate",
      data
    );
    return promise.then((data) =>
      WidgetAreaCreateReply.decode(new _m0.Reader(data))
    );
  }

  WidgetAreaActionQuery(
    request: QueryFilterRequest
  ): Promise<WidgetAreaQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WidgetAreaActionQuery",
      data
    );
    return promise.then((data) =>
      WidgetAreaQueryReply.decode(new _m0.Reader(data))
    );
  }

  WidgetAreaActionGetOne(
    request: QueryFilterRequest
  ): Promise<WidgetAreaReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WidgetAreaActionGetOne",
      data
    );
    return promise.then((data) => WidgetAreaReply.decode(new _m0.Reader(data)));
  }

  WidgetAreaActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WidgetAreaActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Widgets {
  WidgetActionCreate(request: WidgetEntity): Promise<WidgetCreateReply>;
  WidgetActionUpdate(request: WidgetEntity): Promise<WidgetCreateReply>;
  WidgetActionQuery(request: QueryFilterRequest): Promise<WidgetQueryReply>;
  WidgetActionGetOne(request: QueryFilterRequest): Promise<WidgetReply>;
  WidgetActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class WidgetsClientImpl implements Widgets {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Widgets";
    this.rpc = rpc;
    this.WidgetActionCreate = this.WidgetActionCreate.bind(this);
    this.WidgetActionUpdate = this.WidgetActionUpdate.bind(this);
    this.WidgetActionQuery = this.WidgetActionQuery.bind(this);
    this.WidgetActionGetOne = this.WidgetActionGetOne.bind(this);
    this.WidgetActionRemove = this.WidgetActionRemove.bind(this);
  }
  WidgetActionCreate(request: WidgetEntity): Promise<WidgetCreateReply> {
    const data = WidgetEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "WidgetActionCreate", data);
    return promise.then((data) =>
      WidgetCreateReply.decode(new _m0.Reader(data))
    );
  }

  WidgetActionUpdate(request: WidgetEntity): Promise<WidgetCreateReply> {
    const data = WidgetEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "WidgetActionUpdate", data);
    return promise.then((data) =>
      WidgetCreateReply.decode(new _m0.Reader(data))
    );
  }

  WidgetActionQuery(request: QueryFilterRequest): Promise<WidgetQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "WidgetActionQuery", data);
    return promise.then((data) =>
      WidgetQueryReply.decode(new _m0.Reader(data))
    );
  }

  WidgetActionGetOne(request: QueryFilterRequest): Promise<WidgetReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "WidgetActionGetOne", data);
    return promise.then((data) => WidgetReply.decode(new _m0.Reader(data)));
  }

  WidgetActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "WidgetActionRemove", data);
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
