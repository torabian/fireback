/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { UserEntity, WorkspaceEntity } from "../workspaces/index";

export const protobufPackage = "";

export interface FileEntity {
  name: string;
  diskPath: string;
  size: number;
  virtualPath: string;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  type: string;
  /** @tag(gorm:"foreignKey:WorkspaceId;references:UniqueId" json:"-") */
  workspace: WorkspaceEntity | undefined;
  /** @tag(json:"workspaceId" gorm:"size:100;") */
  workspaceId?: string | undefined;
  /** @tag(gorm:"foreignKey:UserId;references:UniqueId" json:"-") */
  user: UserEntity | undefined;
  /** @tag(json:"userId" gorm:"size:100;") */
  userId: string;
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
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

function createBaseFileEntity(): FileEntity {
  return {
    name: "",
    diskPath: "",
    size: 0,
    virtualPath: "",
    uniqueId: "",
    type: "",
    workspace: undefined,
    workspaceId: undefined,
    user: undefined,
    userId: "",
    visibility: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const FileEntity = {
  encode(
    message: FileEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.diskPath !== "") {
      writer.uint32(18).string(message.diskPath);
    }
    if (message.size !== 0) {
      writer.uint32(24).int64(message.size);
    }
    if (message.virtualPath !== "") {
      writer.uint32(34).string(message.virtualPath);
    }
    if (message.uniqueId !== "") {
      writer.uint32(42).string(message.uniqueId);
    }
    if (message.type !== "") {
      writer.uint32(50).string(message.type);
    }
    if (message.workspace !== undefined) {
      WorkspaceEntity.encode(
        message.workspace,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(66).string(message.workspaceId);
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(74).fork()).ldelim();
    }
    if (message.userId !== "") {
      writer.uint32(82).string(message.userId);
    }
    if (message.visibility !== undefined) {
      writer.uint32(90).string(message.visibility);
    }
    if (message.rank !== 0) {
      writer.uint32(120).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(128).int64(message.updated);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): FileEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFileEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.diskPath = reader.string();
          break;
        case 3:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.virtualPath = reader.string();
          break;
        case 5:
          message.uniqueId = reader.string();
          break;
        case 6:
          message.type = reader.string();
          break;
        case 7:
          message.workspace = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 8:
          message.workspaceId = reader.string();
          break;
        case 9:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        case 10:
          message.userId = reader.string();
          break;
        case 11:
          message.visibility = reader.string();
          break;
        case 15:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 16:
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

  fromJSON(object: any): FileEntity {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      diskPath: isSet(object.diskPath) ? String(object.diskPath) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
      virtualPath: isSet(object.virtualPath) ? String(object.virtualPath) : "",
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      type: isSet(object.type) ? String(object.type) : "",
      workspace: isSet(object.workspace)
        ? WorkspaceEntity.fromJSON(object.workspace)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
      userId: isSet(object.userId) ? String(object.userId) : "",
      visibility: isSet(object.visibility)
        ? String(object.visibility)
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

  toJSON(message: FileEntity): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.diskPath !== undefined && (obj.diskPath = message.diskPath);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.virtualPath !== undefined &&
      (obj.virtualPath = message.virtualPath);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.type !== undefined && (obj.type = message.type);
    message.workspace !== undefined &&
      (obj.workspace = message.workspace
        ? WorkspaceEntity.toJSON(message.workspace)
        : undefined);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    message.userId !== undefined && (obj.userId = message.userId);
    message.visibility !== undefined && (obj.visibility = message.visibility);
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

  create<I extends Exact<DeepPartial<FileEntity>, I>>(base?: I): FileEntity {
    return FileEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FileEntity>, I>>(
    object: I
  ): FileEntity {
    const message = createBaseFileEntity();
    message.name = object.name ?? "";
    message.diskPath = object.diskPath ?? "";
    message.size = object.size ?? 0;
    message.virtualPath = object.virtualPath ?? "";
    message.uniqueId = object.uniqueId ?? "";
    message.type = object.type ?? "";
    message.workspace =
      object.workspace !== undefined && object.workspace !== null
        ? WorkspaceEntity.fromPartial(object.workspace)
        : undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    message.userId = object.userId ?? "";
    message.visibility = object.visibility ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

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
