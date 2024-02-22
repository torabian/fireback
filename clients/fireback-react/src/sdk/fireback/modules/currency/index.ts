/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { IError, QueryFilterRequest, RemoveReply } from "../../core/common";

export const protobufPackage = "";

export interface CurrencyCreateReply {
  data: CurrencyEntity | undefined;
  error: IError | undefined;
}

export interface CurrencyReply {
  data: CurrencyEntity | undefined;
  error: IError | undefined;
}

export interface CurrencyQueryReply {
  items: CurrencyEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface CurrencyEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: CurrencyEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: CurrencyEntityPolyglot[];
  /** @tag(  yaml:"symbol"  ) */
  symbol?: string | undefined;
  /** @tag(translate:"true"  yaml:"name"  ) */
  name?: string | undefined;
  /** @tag(  yaml:"symbolNative"  ) */
  symbolNative?: string | undefined;
  /** @tag(  yaml:"decimalDigits"  ) */
  decimalDigits?: number | undefined;
  /** @tag(  yaml:"rounding"  ) */
  rounding?: number | undefined;
  /** @tag(  yaml:"code"  ) */
  code?: string | undefined;
  /** @tag(  yaml:"namePlural"  ) */
  namePlural?: string | undefined;
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
export interface CurrencyEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"name" json:"name"); */
  name: string;
}

export interface PriceTagCreateReply {
  data: PriceTagEntity | undefined;
  error: IError | undefined;
}

export interface PriceTagReply {
  data: PriceTagEntity | undefined;
  error: IError | undefined;
}

export interface PriceTagQueryReply {
  items: PriceTagEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface PriceTagEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PriceTagEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** repeated PriceTagVariationEntity variations = 9; // @tag(gorm:"foreignKey:LinkerId;references:UniqueId" yaml:"variations") */
  variations: PriceTagVariationEntity[];
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

export interface PriceTagVariationEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PriceTagVariationEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  currencyId?: string | undefined;
  /** @tag(gorm:"foreignKey:CurrencyId;references:UniqueId" json:"currency" yaml:"currency" fbtype:"one") */
  currency: CurrencyEntity | undefined;
  /** @tag(  yaml:"amount"  ) */
  amount?: number | undefined;
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

function createBaseCurrencyCreateReply(): CurrencyCreateReply {
  return { data: undefined, error: undefined };
}

export const CurrencyCreateReply = {
  encode(
    message: CurrencyCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      CurrencyEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CurrencyCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCurrencyCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = CurrencyEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): CurrencyCreateReply {
    return {
      data: isSet(object.data)
        ? CurrencyEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CurrencyCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? CurrencyEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CurrencyCreateReply>, I>>(
    base?: I
  ): CurrencyCreateReply {
    return CurrencyCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CurrencyCreateReply>, I>>(
    object: I
  ): CurrencyCreateReply {
    const message = createBaseCurrencyCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? CurrencyEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseCurrencyReply(): CurrencyReply {
  return { data: undefined, error: undefined };
}

export const CurrencyReply = {
  encode(
    message: CurrencyReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      CurrencyEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CurrencyReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCurrencyReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = CurrencyEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): CurrencyReply {
    return {
      data: isSet(object.data)
        ? CurrencyEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CurrencyReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? CurrencyEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CurrencyReply>, I>>(
    base?: I
  ): CurrencyReply {
    return CurrencyReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CurrencyReply>, I>>(
    object: I
  ): CurrencyReply {
    const message = createBaseCurrencyReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? CurrencyEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseCurrencyQueryReply(): CurrencyQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const CurrencyQueryReply = {
  encode(
    message: CurrencyQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      CurrencyEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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

  decode(input: _m0.Reader | Uint8Array, length?: number): CurrencyQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCurrencyQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(CurrencyEntity.decode(reader, reader.uint32()));
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

  fromJSON(object: any): CurrencyQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => CurrencyEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CurrencyQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? CurrencyEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<CurrencyQueryReply>, I>>(
    base?: I
  ): CurrencyQueryReply {
    return CurrencyQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CurrencyQueryReply>, I>>(
    object: I
  ): CurrencyQueryReply {
    const message = createBaseCurrencyQueryReply();
    message.items =
      object.items?.map((e) => CurrencyEntity.fromPartial(e)) || [];
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

function createBaseCurrencyEntity(): CurrencyEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    symbol: undefined,
    name: undefined,
    symbolNative: undefined,
    decimalDigits: undefined,
    rounding: undefined,
    code: undefined,
    namePlural: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const CurrencyEntity = {
  encode(
    message: CurrencyEntity,
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
      CurrencyEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.translations) {
      CurrencyEntityPolyglot.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.symbol !== undefined) {
      writer.uint32(82).string(message.symbol);
    }
    if (message.name !== undefined) {
      writer.uint32(90).string(message.name);
    }
    if (message.symbolNative !== undefined) {
      writer.uint32(98).string(message.symbolNative);
    }
    if (message.decimalDigits !== undefined) {
      writer.uint32(104).int64(message.decimalDigits);
    }
    if (message.rounding !== undefined) {
      writer.uint32(112).int64(message.rounding);
    }
    if (message.code !== undefined) {
      writer.uint32(122).string(message.code);
    }
    if (message.namePlural !== undefined) {
      writer.uint32(130).string(message.namePlural);
    }
    if (message.rank !== 0) {
      writer.uint32(136).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(144).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(152).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(162).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(170).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CurrencyEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCurrencyEntity();
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
          message.parent = CurrencyEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            CurrencyEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.symbol = reader.string();
          break;
        case 11:
          message.name = reader.string();
          break;
        case 12:
          message.symbolNative = reader.string();
          break;
        case 13:
          message.decimalDigits = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.rounding = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.code = reader.string();
          break;
        case 16:
          message.namePlural = reader.string();
          break;
        case 17:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 19:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 20:
          message.createdFormatted = reader.string();
          break;
        case 21:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CurrencyEntity {
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
        ? CurrencyEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            CurrencyEntityPolyglot.fromJSON(e)
          )
        : [],
      symbol: isSet(object.symbol) ? String(object.symbol) : undefined,
      name: isSet(object.name) ? String(object.name) : undefined,
      symbolNative: isSet(object.symbolNative)
        ? String(object.symbolNative)
        : undefined,
      decimalDigits: isSet(object.decimalDigits)
        ? Number(object.decimalDigits)
        : undefined,
      rounding: isSet(object.rounding) ? Number(object.rounding) : undefined,
      code: isSet(object.code) ? String(object.code) : undefined,
      namePlural: isSet(object.namePlural)
        ? String(object.namePlural)
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

  toJSON(message: CurrencyEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? CurrencyEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? CurrencyEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.name !== undefined && (obj.name = message.name);
    message.symbolNative !== undefined &&
      (obj.symbolNative = message.symbolNative);
    message.decimalDigits !== undefined &&
      (obj.decimalDigits = Math.round(message.decimalDigits));
    message.rounding !== undefined &&
      (obj.rounding = Math.round(message.rounding));
    message.code !== undefined && (obj.code = message.code);
    message.namePlural !== undefined && (obj.namePlural = message.namePlural);
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

  create<I extends Exact<DeepPartial<CurrencyEntity>, I>>(
    base?: I
  ): CurrencyEntity {
    return CurrencyEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CurrencyEntity>, I>>(
    object: I
  ): CurrencyEntity {
    const message = createBaseCurrencyEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? CurrencyEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) => CurrencyEntityPolyglot.fromPartial(e)) ||
      [];
    message.symbol = object.symbol ?? undefined;
    message.name = object.name ?? undefined;
    message.symbolNative = object.symbolNative ?? undefined;
    message.decimalDigits = object.decimalDigits ?? undefined;
    message.rounding = object.rounding ?? undefined;
    message.code = object.code ?? undefined;
    message.namePlural = object.namePlural ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseCurrencyEntityPolyglot(): CurrencyEntityPolyglot {
  return { linkerId: "", languageId: "", name: "" };
}

export const CurrencyEntityPolyglot = {
  encode(
    message: CurrencyEntityPolyglot,
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
  ): CurrencyEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCurrencyEntityPolyglot();
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

  fromJSON(object: any): CurrencyEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: CurrencyEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<CurrencyEntityPolyglot>, I>>(
    base?: I
  ): CurrencyEntityPolyglot {
    return CurrencyEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CurrencyEntityPolyglot>, I>>(
    object: I
  ): CurrencyEntityPolyglot {
    const message = createBaseCurrencyEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBasePriceTagCreateReply(): PriceTagCreateReply {
  return { data: undefined, error: undefined };
}

export const PriceTagCreateReply = {
  encode(
    message: PriceTagCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PriceTagEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PriceTagCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriceTagCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PriceTagEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): PriceTagCreateReply {
    return {
      data: isSet(object.data)
        ? PriceTagEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PriceTagCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PriceTagEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PriceTagCreateReply>, I>>(
    base?: I
  ): PriceTagCreateReply {
    return PriceTagCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PriceTagCreateReply>, I>>(
    object: I
  ): PriceTagCreateReply {
    const message = createBasePriceTagCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PriceTagEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePriceTagReply(): PriceTagReply {
  return { data: undefined, error: undefined };
}

export const PriceTagReply = {
  encode(
    message: PriceTagReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PriceTagEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PriceTagReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriceTagReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PriceTagEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): PriceTagReply {
    return {
      data: isSet(object.data)
        ? PriceTagEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PriceTagReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PriceTagEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PriceTagReply>, I>>(
    base?: I
  ): PriceTagReply {
    return PriceTagReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PriceTagReply>, I>>(
    object: I
  ): PriceTagReply {
    const message = createBasePriceTagReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PriceTagEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePriceTagQueryReply(): PriceTagQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const PriceTagQueryReply = {
  encode(
    message: PriceTagQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      PriceTagEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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

  decode(input: _m0.Reader | Uint8Array, length?: number): PriceTagQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriceTagQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(PriceTagEntity.decode(reader, reader.uint32()));
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

  fromJSON(object: any): PriceTagQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => PriceTagEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PriceTagQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? PriceTagEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<PriceTagQueryReply>, I>>(
    base?: I
  ): PriceTagQueryReply {
    return PriceTagQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PriceTagQueryReply>, I>>(
    object: I
  ): PriceTagQueryReply {
    const message = createBasePriceTagQueryReply();
    message.items =
      object.items?.map((e) => PriceTagEntity.fromPartial(e)) || [];
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

function createBasePriceTagEntity(): PriceTagEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    variations: [],
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PriceTagEntity = {
  encode(
    message: PriceTagEntity,
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
      PriceTagEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.variations) {
      PriceTagVariationEntity.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(80).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(88).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(96).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(106).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(114).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PriceTagEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriceTagEntity();
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
          message.parent = PriceTagEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.variations.push(
            PriceTagVariationEntity.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 11:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.createdFormatted = reader.string();
          break;
        case 14:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PriceTagEntity {
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
        ? PriceTagEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      variations: Array.isArray(object?.variations)
        ? object.variations.map((e: any) => PriceTagVariationEntity.fromJSON(e))
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

  toJSON(message: PriceTagEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PriceTagEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.variations) {
      obj.variations = message.variations.map((e) =>
        e ? PriceTagVariationEntity.toJSON(e) : undefined
      );
    } else {
      obj.variations = [];
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

  create<I extends Exact<DeepPartial<PriceTagEntity>, I>>(
    base?: I
  ): PriceTagEntity {
    return PriceTagEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PriceTagEntity>, I>>(
    object: I
  ): PriceTagEntity {
    const message = createBasePriceTagEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PriceTagEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.variations =
      object.variations?.map((e) => PriceTagVariationEntity.fromPartial(e)) ||
      [];
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBasePriceTagVariationEntity(): PriceTagVariationEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    currencyId: undefined,
    currency: undefined,
    amount: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PriceTagVariationEntity = {
  encode(
    message: PriceTagVariationEntity,
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
      PriceTagVariationEntity.encode(
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
    if (message.currencyId !== undefined) {
      writer.uint32(82).string(message.currencyId);
    }
    if (message.currency !== undefined) {
      CurrencyEntity.encode(
        message.currency,
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.amount !== undefined) {
      writer.uint32(97).double(message.amount);
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

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PriceTagVariationEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePriceTagVariationEntity();
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
          message.parent = PriceTagVariationEntity.decode(
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
        case 10:
          message.currencyId = reader.string();
          break;
        case 11:
          message.currency = CurrencyEntity.decode(reader, reader.uint32());
          break;
        case 12:
          message.amount = reader.double();
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

  fromJSON(object: any): PriceTagVariationEntity {
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
        ? PriceTagVariationEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      currencyId: isSet(object.currencyId)
        ? String(object.currencyId)
        : undefined,
      currency: isSet(object.currency)
        ? CurrencyEntity.fromJSON(object.currency)
        : undefined,
      amount: isSet(object.amount) ? Number(object.amount) : undefined,
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

  toJSON(message: PriceTagVariationEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PriceTagVariationEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.currencyId !== undefined && (obj.currencyId = message.currencyId);
    message.currency !== undefined &&
      (obj.currency = message.currency
        ? CurrencyEntity.toJSON(message.currency)
        : undefined);
    message.amount !== undefined && (obj.amount = message.amount);
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

  create<I extends Exact<DeepPartial<PriceTagVariationEntity>, I>>(
    base?: I
  ): PriceTagVariationEntity {
    return PriceTagVariationEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PriceTagVariationEntity>, I>>(
    object: I
  ): PriceTagVariationEntity {
    const message = createBasePriceTagVariationEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PriceTagVariationEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.currencyId = object.currencyId ?? undefined;
    message.currency =
      object.currency !== undefined && object.currency !== null
        ? CurrencyEntity.fromPartial(object.currency)
        : undefined;
    message.amount = object.amount ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

export interface Currencys {
  CurrencyActionCreate(request: CurrencyEntity): Promise<CurrencyCreateReply>;
  CurrencyActionUpdate(request: CurrencyEntity): Promise<CurrencyCreateReply>;
  CurrencyActionQuery(request: QueryFilterRequest): Promise<CurrencyQueryReply>;
  CurrencyActionGetOne(request: QueryFilterRequest): Promise<CurrencyReply>;
  CurrencyActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class CurrencysClientImpl implements Currencys {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Currencys";
    this.rpc = rpc;
    this.CurrencyActionCreate = this.CurrencyActionCreate.bind(this);
    this.CurrencyActionUpdate = this.CurrencyActionUpdate.bind(this);
    this.CurrencyActionQuery = this.CurrencyActionQuery.bind(this);
    this.CurrencyActionGetOne = this.CurrencyActionGetOne.bind(this);
    this.CurrencyActionRemove = this.CurrencyActionRemove.bind(this);
  }
  CurrencyActionCreate(request: CurrencyEntity): Promise<CurrencyCreateReply> {
    const data = CurrencyEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CurrencyActionCreate",
      data
    );
    return promise.then((data) =>
      CurrencyCreateReply.decode(new _m0.Reader(data))
    );
  }

  CurrencyActionUpdate(request: CurrencyEntity): Promise<CurrencyCreateReply> {
    const data = CurrencyEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CurrencyActionUpdate",
      data
    );
    return promise.then((data) =>
      CurrencyCreateReply.decode(new _m0.Reader(data))
    );
  }

  CurrencyActionQuery(
    request: QueryFilterRequest
  ): Promise<CurrencyQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CurrencyActionQuery", data);
    return promise.then((data) =>
      CurrencyQueryReply.decode(new _m0.Reader(data))
    );
  }

  CurrencyActionGetOne(request: QueryFilterRequest): Promise<CurrencyReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CurrencyActionGetOne",
      data
    );
    return promise.then((data) => CurrencyReply.decode(new _m0.Reader(data)));
  }

  CurrencyActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CurrencyActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface PriceTags {
  PriceTagActionCreate(request: PriceTagEntity): Promise<PriceTagCreateReply>;
  PriceTagActionUpdate(request: PriceTagEntity): Promise<PriceTagCreateReply>;
  PriceTagActionQuery(request: QueryFilterRequest): Promise<PriceTagQueryReply>;
  PriceTagActionGetOne(request: QueryFilterRequest): Promise<PriceTagReply>;
  PriceTagActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class PriceTagsClientImpl implements PriceTags {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "PriceTags";
    this.rpc = rpc;
    this.PriceTagActionCreate = this.PriceTagActionCreate.bind(this);
    this.PriceTagActionUpdate = this.PriceTagActionUpdate.bind(this);
    this.PriceTagActionQuery = this.PriceTagActionQuery.bind(this);
    this.PriceTagActionGetOne = this.PriceTagActionGetOne.bind(this);
    this.PriceTagActionRemove = this.PriceTagActionRemove.bind(this);
  }
  PriceTagActionCreate(request: PriceTagEntity): Promise<PriceTagCreateReply> {
    const data = PriceTagEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PriceTagActionCreate",
      data
    );
    return promise.then((data) =>
      PriceTagCreateReply.decode(new _m0.Reader(data))
    );
  }

  PriceTagActionUpdate(request: PriceTagEntity): Promise<PriceTagCreateReply> {
    const data = PriceTagEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PriceTagActionUpdate",
      data
    );
    return promise.then((data) =>
      PriceTagCreateReply.decode(new _m0.Reader(data))
    );
  }

  PriceTagActionQuery(
    request: QueryFilterRequest
  ): Promise<PriceTagQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "PriceTagActionQuery", data);
    return promise.then((data) =>
      PriceTagQueryReply.decode(new _m0.Reader(data))
    );
  }

  PriceTagActionGetOne(request: QueryFilterRequest): Promise<PriceTagReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PriceTagActionGetOne",
      data
    );
    return promise.then((data) => PriceTagReply.decode(new _m0.Reader(data)));
  }

  PriceTagActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PriceTagActionRemove",
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
