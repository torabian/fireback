/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { IError, QueryFilterRequest, RemoveReply } from "../../core/common";

export const protobufPackage = "";

export interface CommonProfileCreateReply {
  data: CommonProfileEntity | undefined;
  error: IError | undefined;
}

export interface CommonProfileReply {
  data: CommonProfileEntity | undefined;
  error: IError | undefined;
}

export interface CommonProfileQueryReply {
  items: CommonProfileEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface CommonProfileEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: CommonProfileEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"firstName"  ) */
  firstName?: string | undefined;
  /** @tag(  yaml:"lastName"  ) */
  lastName?: string | undefined;
  /** @tag(  yaml:"phoneNumber"  ) */
  phoneNumber?: string | undefined;
  /** @tag(  yaml:"email"  ) */
  email?: string | undefined;
  /** @tag(  yaml:"company"  ) */
  company?: string | undefined;
  /** @tag(  yaml:"street"  ) */
  street?: string | undefined;
  /** @tag(  yaml:"houseNumber"  ) */
  houseNumber?: string | undefined;
  /** @tag(  yaml:"zipCode"  ) */
  zipCode?: string | undefined;
  /** @tag(  yaml:"city"  ) */
  city?: string | undefined;
  /** @tag(  yaml:"gender"  ) */
  gender?: string | undefined;
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

function createBaseCommonProfileCreateReply(): CommonProfileCreateReply {
  return { data: undefined, error: undefined };
}

export const CommonProfileCreateReply = {
  encode(
    message: CommonProfileCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      CommonProfileEntity.encode(
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
  ): CommonProfileCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCommonProfileCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = CommonProfileEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): CommonProfileCreateReply {
    return {
      data: isSet(object.data)
        ? CommonProfileEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CommonProfileCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? CommonProfileEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CommonProfileCreateReply>, I>>(
    base?: I
  ): CommonProfileCreateReply {
    return CommonProfileCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CommonProfileCreateReply>, I>>(
    object: I
  ): CommonProfileCreateReply {
    const message = createBaseCommonProfileCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? CommonProfileEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseCommonProfileReply(): CommonProfileReply {
  return { data: undefined, error: undefined };
}

export const CommonProfileReply = {
  encode(
    message: CommonProfileReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      CommonProfileEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CommonProfileReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCommonProfileReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = CommonProfileEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): CommonProfileReply {
    return {
      data: isSet(object.data)
        ? CommonProfileEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CommonProfileReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? CommonProfileEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CommonProfileReply>, I>>(
    base?: I
  ): CommonProfileReply {
    return CommonProfileReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CommonProfileReply>, I>>(
    object: I
  ): CommonProfileReply {
    const message = createBaseCommonProfileReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? CommonProfileEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseCommonProfileQueryReply(): CommonProfileQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const CommonProfileQueryReply = {
  encode(
    message: CommonProfileQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      CommonProfileEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): CommonProfileQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCommonProfileQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            CommonProfileEntity.decode(reader, reader.uint32())
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

  fromJSON(object: any): CommonProfileQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => CommonProfileEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CommonProfileQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? CommonProfileEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<CommonProfileQueryReply>, I>>(
    base?: I
  ): CommonProfileQueryReply {
    return CommonProfileQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CommonProfileQueryReply>, I>>(
    object: I
  ): CommonProfileQueryReply {
    const message = createBaseCommonProfileQueryReply();
    message.items =
      object.items?.map((e) => CommonProfileEntity.fromPartial(e)) || [];
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

function createBaseCommonProfileEntity(): CommonProfileEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    firstName: undefined,
    lastName: undefined,
    phoneNumber: undefined,
    email: undefined,
    company: undefined,
    street: undefined,
    houseNumber: undefined,
    zipCode: undefined,
    city: undefined,
    gender: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const CommonProfileEntity = {
  encode(
    message: CommonProfileEntity,
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
      CommonProfileEntity.encode(
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
    if (message.firstName !== undefined) {
      writer.uint32(74).string(message.firstName);
    }
    if (message.lastName !== undefined) {
      writer.uint32(82).string(message.lastName);
    }
    if (message.phoneNumber !== undefined) {
      writer.uint32(90).string(message.phoneNumber);
    }
    if (message.email !== undefined) {
      writer.uint32(98).string(message.email);
    }
    if (message.company !== undefined) {
      writer.uint32(106).string(message.company);
    }
    if (message.street !== undefined) {
      writer.uint32(114).string(message.street);
    }
    if (message.houseNumber !== undefined) {
      writer.uint32(122).string(message.houseNumber);
    }
    if (message.zipCode !== undefined) {
      writer.uint32(130).string(message.zipCode);
    }
    if (message.city !== undefined) {
      writer.uint32(138).string(message.city);
    }
    if (message.gender !== undefined) {
      writer.uint32(146).string(message.gender);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): CommonProfileEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCommonProfileEntity();
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
          message.parent = CommonProfileEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.firstName = reader.string();
          break;
        case 10:
          message.lastName = reader.string();
          break;
        case 11:
          message.phoneNumber = reader.string();
          break;
        case 12:
          message.email = reader.string();
          break;
        case 13:
          message.company = reader.string();
          break;
        case 14:
          message.street = reader.string();
          break;
        case 15:
          message.houseNumber = reader.string();
          break;
        case 16:
          message.zipCode = reader.string();
          break;
        case 17:
          message.city = reader.string();
          break;
        case 18:
          message.gender = reader.string();
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

  fromJSON(object: any): CommonProfileEntity {
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
        ? CommonProfileEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      firstName: isSet(object.firstName) ? String(object.firstName) : undefined,
      lastName: isSet(object.lastName) ? String(object.lastName) : undefined,
      phoneNumber: isSet(object.phoneNumber)
        ? String(object.phoneNumber)
        : undefined,
      email: isSet(object.email) ? String(object.email) : undefined,
      company: isSet(object.company) ? String(object.company) : undefined,
      street: isSet(object.street) ? String(object.street) : undefined,
      houseNumber: isSet(object.houseNumber)
        ? String(object.houseNumber)
        : undefined,
      zipCode: isSet(object.zipCode) ? String(object.zipCode) : undefined,
      city: isSet(object.city) ? String(object.city) : undefined,
      gender: isSet(object.gender) ? String(object.gender) : undefined,
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

  toJSON(message: CommonProfileEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? CommonProfileEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.phoneNumber !== undefined &&
      (obj.phoneNumber = message.phoneNumber);
    message.email !== undefined && (obj.email = message.email);
    message.company !== undefined && (obj.company = message.company);
    message.street !== undefined && (obj.street = message.street);
    message.houseNumber !== undefined &&
      (obj.houseNumber = message.houseNumber);
    message.zipCode !== undefined && (obj.zipCode = message.zipCode);
    message.city !== undefined && (obj.city = message.city);
    message.gender !== undefined && (obj.gender = message.gender);
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

  create<I extends Exact<DeepPartial<CommonProfileEntity>, I>>(
    base?: I
  ): CommonProfileEntity {
    return CommonProfileEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CommonProfileEntity>, I>>(
    object: I
  ): CommonProfileEntity {
    const message = createBaseCommonProfileEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? CommonProfileEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.firstName = object.firstName ?? undefined;
    message.lastName = object.lastName ?? undefined;
    message.phoneNumber = object.phoneNumber ?? undefined;
    message.email = object.email ?? undefined;
    message.company = object.company ?? undefined;
    message.street = object.street ?? undefined;
    message.houseNumber = object.houseNumber ?? undefined;
    message.zipCode = object.zipCode ?? undefined;
    message.city = object.city ?? undefined;
    message.gender = object.gender ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

export interface CommonProfiles {
  CommonProfileActionCreate(
    request: CommonProfileEntity
  ): Promise<CommonProfileCreateReply>;
  CommonProfileActionUpdate(
    request: CommonProfileEntity
  ): Promise<CommonProfileCreateReply>;
  CommonProfileActionQuery(
    request: QueryFilterRequest
  ): Promise<CommonProfileQueryReply>;
  CommonProfileActionGetOne(
    request: QueryFilterRequest
  ): Promise<CommonProfileReply>;
  CommonProfileActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class CommonProfilesClientImpl implements CommonProfiles {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "CommonProfiles";
    this.rpc = rpc;
    this.CommonProfileActionCreate = this.CommonProfileActionCreate.bind(this);
    this.CommonProfileActionUpdate = this.CommonProfileActionUpdate.bind(this);
    this.CommonProfileActionQuery = this.CommonProfileActionQuery.bind(this);
    this.CommonProfileActionGetOne = this.CommonProfileActionGetOne.bind(this);
    this.CommonProfileActionRemove = this.CommonProfileActionRemove.bind(this);
  }
  CommonProfileActionCreate(
    request: CommonProfileEntity
  ): Promise<CommonProfileCreateReply> {
    const data = CommonProfileEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CommonProfileActionCreate",
      data
    );
    return promise.then((data) =>
      CommonProfileCreateReply.decode(new _m0.Reader(data))
    );
  }

  CommonProfileActionUpdate(
    request: CommonProfileEntity
  ): Promise<CommonProfileCreateReply> {
    const data = CommonProfileEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CommonProfileActionUpdate",
      data
    );
    return promise.then((data) =>
      CommonProfileCreateReply.decode(new _m0.Reader(data))
    );
  }

  CommonProfileActionQuery(
    request: QueryFilterRequest
  ): Promise<CommonProfileQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CommonProfileActionQuery",
      data
    );
    return promise.then((data) =>
      CommonProfileQueryReply.decode(new _m0.Reader(data))
    );
  }

  CommonProfileActionGetOne(
    request: QueryFilterRequest
  ): Promise<CommonProfileReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CommonProfileActionGetOne",
      data
    );
    return promise.then((data) =>
      CommonProfileReply.decode(new _m0.Reader(data))
    );
  }

  CommonProfileActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CommonProfileActionRemove",
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
