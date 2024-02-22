/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import {
  IError,
  QueryFilterRequest,
  RemoveReply,
  Timestamp,
} from "../../core/common";
import { PriceTagEntity } from "../currency/index";
import { CapabilityEntity } from "../workspaces/index";

export const protobufPackage = "";

export interface ActivationKeyCreateReply {
  data: ActivationKeyEntity | undefined;
  error: IError | undefined;
}

export interface ActivationKeyReply {
  data: ActivationKeyEntity | undefined;
  error: IError | undefined;
}

export interface ActivationKeyQueryReply {
  items: ActivationKeyEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface ActivationKeyEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: ActivationKeyEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"series"  ) */
  series?: string | undefined;
  /** @tag(  yaml:"used"  ) */
  used?: number | undefined;
  /** One 2 one to external entity */
  planId?: string | undefined;
  /** @tag(gorm:"foreignKey:PlanId;references:UniqueId" json:"plan" yaml:"plan" fbtype:"one") */
  plan: ProductPlanEntity | undefined;
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

export interface LicensableProductCreateReply {
  data: LicensableProductEntity | undefined;
  error: IError | undefined;
}

export interface LicensableProductReply {
  data: LicensableProductEntity | undefined;
  error: IError | undefined;
}

export interface LicensableProductQueryReply {
  items: LicensableProductEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface LicensableProductEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: LicensableProductEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: LicensableProductEntityPolyglot[];
  /** @tag(translate:"true" validate:"required,omitempty,min=1,max=100" yaml:"name"  ) */
  name?: string | undefined;
  /** @tag(  yaml:"privateKey"  ) */
  privateKey?: string | undefined;
  /** @tag(  yaml:"publicKey"  ) */
  publicKey?: string | undefined;
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
export interface LicensableProductEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"name" json:"name"); */
  name: string;
}

export interface LicenseCreateReply {
  data: LicenseEntity | undefined;
  error: IError | undefined;
}

export interface LicenseReply {
  data: LicenseEntity | undefined;
  error: IError | undefined;
}

export interface LicenseQueryReply {
  items: LicenseEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface LicenseEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: LicenseEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"name"  ) */
  name?: string | undefined;
  /** @tag(  yaml:"signedLicense"  ) */
  signedLicense?: string | undefined;
  /** @tag( yaml:"validityStartDate") */
  validityStartDate: Timestamp | undefined;
  /** @tag( yaml:"validityEndDate") */
  validityEndDate: Timestamp | undefined;
  /** repeated LicensePermissionEntity permissions = 13; // @tag(gorm:"foreignKey:LinkerId;references:UniqueId" yaml:"permissions") */
  permissions: LicensePermissionEntity[];
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

export interface LicensePermissionEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: LicensePermissionEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  capabilityId?: string | undefined;
  /** @tag(gorm:"foreignKey:CapabilityId;references:UniqueId" json:"capability" yaml:"capability" fbtype:"one") */
  capability: CapabilityEntity | undefined;
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

export interface LicenseFromActivationKeyDto {
  activationKeyId: string;
  machineId: string;
}

export interface LicenseFromPlanIdDto {
  machineId: string;
  email: string;
  owner: string;
}

export interface ProductPlanCreateReply {
  data: ProductPlanEntity | undefined;
  error: IError | undefined;
}

export interface ProductPlanReply {
  data: ProductPlanEntity | undefined;
  error: IError | undefined;
}

export interface ProductPlanQueryReply {
  items: ProductPlanEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface ProductPlanEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: ProductPlanEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: ProductPlanEntityPolyglot[];
  /** @tag(translate:"true" validate:"required,omitempty,min=1,max=100" yaml:"name"  ) */
  name?: string | undefined;
  /** @tag( validate:"required" yaml:"duration"  ) */
  duration?: number | undefined;
  /** One 2 one to external entity */
  productId?: string | undefined;
  /** @tag(gorm:"foreignKey:ProductId;references:UniqueId" json:"product" yaml:"product" fbtype:"one") */
  product: LicensableProductEntity | undefined;
  /** One 2 one to external entity */
  priceTagId?: string | undefined;
  /** @tag(gorm:"foreignKey:PriceTagId;references:UniqueId" json:"priceTag" yaml:"priceTag" fbtype:"one") */
  priceTag: PriceTagEntity | undefined;
  /** repeated ProductPlanPermissionEntity permissions = 18; // @tag(gorm:"foreignKey:LinkerId;references:UniqueId" yaml:"permissions") */
  permissions: ProductPlanPermissionEntity[];
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
export interface ProductPlanEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"name" json:"name"); */
  name: string;
}

export interface ProductPlanPermissionEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: ProductPlanPermissionEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  capabilityId?: string | undefined;
  /** @tag(gorm:"foreignKey:CapabilityId;references:UniqueId" json:"capability" yaml:"capability" fbtype:"one") */
  capability: CapabilityEntity | undefined;
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

function createBaseActivationKeyCreateReply(): ActivationKeyCreateReply {
  return { data: undefined, error: undefined };
}

export const ActivationKeyCreateReply = {
  encode(
    message: ActivationKeyCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      ActivationKeyEntity.encode(
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
  ): ActivationKeyCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActivationKeyCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = ActivationKeyEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): ActivationKeyCreateReply {
    return {
      data: isSet(object.data)
        ? ActivationKeyEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ActivationKeyCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? ActivationKeyEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ActivationKeyCreateReply>, I>>(
    base?: I
  ): ActivationKeyCreateReply {
    return ActivationKeyCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ActivationKeyCreateReply>, I>>(
    object: I
  ): ActivationKeyCreateReply {
    const message = createBaseActivationKeyCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? ActivationKeyEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseActivationKeyReply(): ActivationKeyReply {
  return { data: undefined, error: undefined };
}

export const ActivationKeyReply = {
  encode(
    message: ActivationKeyReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      ActivationKeyEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ActivationKeyReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActivationKeyReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = ActivationKeyEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): ActivationKeyReply {
    return {
      data: isSet(object.data)
        ? ActivationKeyEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ActivationKeyReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? ActivationKeyEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ActivationKeyReply>, I>>(
    base?: I
  ): ActivationKeyReply {
    return ActivationKeyReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ActivationKeyReply>, I>>(
    object: I
  ): ActivationKeyReply {
    const message = createBaseActivationKeyReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? ActivationKeyEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseActivationKeyQueryReply(): ActivationKeyQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const ActivationKeyQueryReply = {
  encode(
    message: ActivationKeyQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      ActivationKeyEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): ActivationKeyQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActivationKeyQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            ActivationKeyEntity.decode(reader, reader.uint32())
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

  fromJSON(object: any): ActivationKeyQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => ActivationKeyEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ActivationKeyQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? ActivationKeyEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<ActivationKeyQueryReply>, I>>(
    base?: I
  ): ActivationKeyQueryReply {
    return ActivationKeyQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ActivationKeyQueryReply>, I>>(
    object: I
  ): ActivationKeyQueryReply {
    const message = createBaseActivationKeyQueryReply();
    message.items =
      object.items?.map((e) => ActivationKeyEntity.fromPartial(e)) || [];
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

function createBaseActivationKeyEntity(): ActivationKeyEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    series: undefined,
    used: undefined,
    planId: undefined,
    plan: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const ActivationKeyEntity = {
  encode(
    message: ActivationKeyEntity,
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
      ActivationKeyEntity.encode(
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
    if (message.series !== undefined) {
      writer.uint32(74).string(message.series);
    }
    if (message.used !== undefined) {
      writer.uint32(80).int64(message.used);
    }
    if (message.planId !== undefined) {
      writer.uint32(98).string(message.planId);
    }
    if (message.plan !== undefined) {
      ProductPlanEntity.encode(
        message.plan,
        writer.uint32(106).fork()
      ).ldelim();
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

  decode(input: _m0.Reader | Uint8Array, length?: number): ActivationKeyEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActivationKeyEntity();
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
          message.parent = ActivationKeyEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.series = reader.string();
          break;
        case 10:
          message.used = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.planId = reader.string();
          break;
        case 13:
          message.plan = ProductPlanEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): ActivationKeyEntity {
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
        ? ActivationKeyEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      series: isSet(object.series) ? String(object.series) : undefined,
      used: isSet(object.used) ? Number(object.used) : undefined,
      planId: isSet(object.planId) ? String(object.planId) : undefined,
      plan: isSet(object.plan)
        ? ProductPlanEntity.fromJSON(object.plan)
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

  toJSON(message: ActivationKeyEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? ActivationKeyEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.series !== undefined && (obj.series = message.series);
    message.used !== undefined && (obj.used = Math.round(message.used));
    message.planId !== undefined && (obj.planId = message.planId);
    message.plan !== undefined &&
      (obj.plan = message.plan
        ? ProductPlanEntity.toJSON(message.plan)
        : undefined);
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

  create<I extends Exact<DeepPartial<ActivationKeyEntity>, I>>(
    base?: I
  ): ActivationKeyEntity {
    return ActivationKeyEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ActivationKeyEntity>, I>>(
    object: I
  ): ActivationKeyEntity {
    const message = createBaseActivationKeyEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? ActivationKeyEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.series = object.series ?? undefined;
    message.used = object.used ?? undefined;
    message.planId = object.planId ?? undefined;
    message.plan =
      object.plan !== undefined && object.plan !== null
        ? ProductPlanEntity.fromPartial(object.plan)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseLicensableProductCreateReply(): LicensableProductCreateReply {
  return { data: undefined, error: undefined };
}

export const LicensableProductCreateReply = {
  encode(
    message: LicensableProductCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      LicensableProductEntity.encode(
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
  ): LicensableProductCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicensableProductCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = LicensableProductEntity.decode(
            reader,
            reader.uint32()
          );
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

  fromJSON(object: any): LicensableProductCreateReply {
    return {
      data: isSet(object.data)
        ? LicensableProductEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: LicensableProductCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? LicensableProductEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<LicensableProductCreateReply>, I>>(
    base?: I
  ): LicensableProductCreateReply {
    return LicensableProductCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicensableProductCreateReply>, I>>(
    object: I
  ): LicensableProductCreateReply {
    const message = createBaseLicensableProductCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? LicensableProductEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseLicensableProductReply(): LicensableProductReply {
  return { data: undefined, error: undefined };
}

export const LicensableProductReply = {
  encode(
    message: LicensableProductReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      LicensableProductEntity.encode(
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
  ): LicensableProductReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicensableProductReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = LicensableProductEntity.decode(
            reader,
            reader.uint32()
          );
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

  fromJSON(object: any): LicensableProductReply {
    return {
      data: isSet(object.data)
        ? LicensableProductEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: LicensableProductReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? LicensableProductEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<LicensableProductReply>, I>>(
    base?: I
  ): LicensableProductReply {
    return LicensableProductReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicensableProductReply>, I>>(
    object: I
  ): LicensableProductReply {
    const message = createBaseLicensableProductReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? LicensableProductEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseLicensableProductQueryReply(): LicensableProductQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const LicensableProductQueryReply = {
  encode(
    message: LicensableProductQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      LicensableProductEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): LicensableProductQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicensableProductQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            LicensableProductEntity.decode(reader, reader.uint32())
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

  fromJSON(object: any): LicensableProductQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => LicensableProductEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: LicensableProductQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? LicensableProductEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<LicensableProductQueryReply>, I>>(
    base?: I
  ): LicensableProductQueryReply {
    return LicensableProductQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicensableProductQueryReply>, I>>(
    object: I
  ): LicensableProductQueryReply {
    const message = createBaseLicensableProductQueryReply();
    message.items =
      object.items?.map((e) => LicensableProductEntity.fromPartial(e)) || [];
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

function createBaseLicensableProductEntity(): LicensableProductEntity {
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
    privateKey: undefined,
    publicKey: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const LicensableProductEntity = {
  encode(
    message: LicensableProductEntity,
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
      LicensableProductEntity.encode(
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
      LicensableProductEntityPolyglot.encode(
        v!,
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.name !== undefined) {
      writer.uint32(82).string(message.name);
    }
    if (message.privateKey !== undefined) {
      writer.uint32(90).string(message.privateKey);
    }
    if (message.publicKey !== undefined) {
      writer.uint32(98).string(message.publicKey);
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
  ): LicensableProductEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicensableProductEntity();
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
          message.parent = LicensableProductEntity.decode(
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
            LicensableProductEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.name = reader.string();
          break;
        case 11:
          message.privateKey = reader.string();
          break;
        case 12:
          message.publicKey = reader.string();
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

  fromJSON(object: any): LicensableProductEntity {
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
        ? LicensableProductEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            LicensableProductEntityPolyglot.fromJSON(e)
          )
        : [],
      name: isSet(object.name) ? String(object.name) : undefined,
      privateKey: isSet(object.privateKey)
        ? String(object.privateKey)
        : undefined,
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : undefined,
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

  toJSON(message: LicensableProductEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? LicensableProductEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? LicensableProductEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.name !== undefined && (obj.name = message.name);
    message.privateKey !== undefined && (obj.privateKey = message.privateKey);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
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

  create<I extends Exact<DeepPartial<LicensableProductEntity>, I>>(
    base?: I
  ): LicensableProductEntity {
    return LicensableProductEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicensableProductEntity>, I>>(
    object: I
  ): LicensableProductEntity {
    const message = createBaseLicensableProductEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? LicensableProductEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) =>
        LicensableProductEntityPolyglot.fromPartial(e)
      ) || [];
    message.name = object.name ?? undefined;
    message.privateKey = object.privateKey ?? undefined;
    message.publicKey = object.publicKey ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseLicensableProductEntityPolyglot(): LicensableProductEntityPolyglot {
  return { linkerId: "", languageId: "", name: "" };
}

export const LicensableProductEntityPolyglot = {
  encode(
    message: LicensableProductEntityPolyglot,
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
  ): LicensableProductEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicensableProductEntityPolyglot();
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

  fromJSON(object: any): LicensableProductEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: LicensableProductEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<LicensableProductEntityPolyglot>, I>>(
    base?: I
  ): LicensableProductEntityPolyglot {
    return LicensableProductEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicensableProductEntityPolyglot>, I>>(
    object: I
  ): LicensableProductEntityPolyglot {
    const message = createBaseLicensableProductEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseLicenseCreateReply(): LicenseCreateReply {
  return { data: undefined, error: undefined };
}

export const LicenseCreateReply = {
  encode(
    message: LicenseCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      LicenseEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LicenseCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicenseCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = LicenseEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): LicenseCreateReply {
    return {
      data: isSet(object.data)
        ? LicenseEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: LicenseCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? LicenseEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<LicenseCreateReply>, I>>(
    base?: I
  ): LicenseCreateReply {
    return LicenseCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicenseCreateReply>, I>>(
    object: I
  ): LicenseCreateReply {
    const message = createBaseLicenseCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? LicenseEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseLicenseReply(): LicenseReply {
  return { data: undefined, error: undefined };
}

export const LicenseReply = {
  encode(
    message: LicenseReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      LicenseEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LicenseReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicenseReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = LicenseEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): LicenseReply {
    return {
      data: isSet(object.data)
        ? LicenseEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: LicenseReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? LicenseEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<LicenseReply>, I>>(
    base?: I
  ): LicenseReply {
    return LicenseReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicenseReply>, I>>(
    object: I
  ): LicenseReply {
    const message = createBaseLicenseReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? LicenseEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseLicenseQueryReply(): LicenseQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const LicenseQueryReply = {
  encode(
    message: LicenseQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      LicenseEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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

  decode(input: _m0.Reader | Uint8Array, length?: number): LicenseQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicenseQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(LicenseEntity.decode(reader, reader.uint32()));
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

  fromJSON(object: any): LicenseQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => LicenseEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: LicenseQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? LicenseEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<LicenseQueryReply>, I>>(
    base?: I
  ): LicenseQueryReply {
    return LicenseQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicenseQueryReply>, I>>(
    object: I
  ): LicenseQueryReply {
    const message = createBaseLicenseQueryReply();
    message.items =
      object.items?.map((e) => LicenseEntity.fromPartial(e)) || [];
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

function createBaseLicenseEntity(): LicenseEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    name: undefined,
    signedLicense: undefined,
    validityStartDate: undefined,
    validityEndDate: undefined,
    permissions: [],
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const LicenseEntity = {
  encode(
    message: LicenseEntity,
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
      LicenseEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.name !== undefined) {
      writer.uint32(74).string(message.name);
    }
    if (message.signedLicense !== undefined) {
      writer.uint32(82).string(message.signedLicense);
    }
    if (message.validityStartDate !== undefined) {
      Timestamp.encode(
        message.validityStartDate,
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.validityEndDate !== undefined) {
      Timestamp.encode(
        message.validityEndDate,
        writer.uint32(98).fork()
      ).ldelim();
    }
    for (const v of message.permissions) {
      LicensePermissionEntity.encode(v!, writer.uint32(106).fork()).ldelim();
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

  decode(input: _m0.Reader | Uint8Array, length?: number): LicenseEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicenseEntity();
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
          message.parent = LicenseEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.name = reader.string();
          break;
        case 10:
          message.signedLicense = reader.string();
          break;
        case 11:
          message.validityStartDate = Timestamp.decode(reader, reader.uint32());
          break;
        case 12:
          message.validityEndDate = Timestamp.decode(reader, reader.uint32());
          break;
        case 13:
          message.permissions.push(
            LicensePermissionEntity.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): LicenseEntity {
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
        ? LicenseEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      name: isSet(object.name) ? String(object.name) : undefined,
      signedLicense: isSet(object.signedLicense)
        ? String(object.signedLicense)
        : undefined,
      validityStartDate: isSet(object.validityStartDate)
        ? Timestamp.fromJSON(object.validityStartDate)
        : undefined,
      validityEndDate: isSet(object.validityEndDate)
        ? Timestamp.fromJSON(object.validityEndDate)
        : undefined,
      permissions: Array.isArray(object?.permissions)
        ? object.permissions.map((e: any) =>
            LicensePermissionEntity.fromJSON(e)
          )
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

  toJSON(message: LicenseEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? LicenseEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.name !== undefined && (obj.name = message.name);
    message.signedLicense !== undefined &&
      (obj.signedLicense = message.signedLicense);
    message.validityStartDate !== undefined &&
      (obj.validityStartDate = message.validityStartDate
        ? Timestamp.toJSON(message.validityStartDate)
        : undefined);
    message.validityEndDate !== undefined &&
      (obj.validityEndDate = message.validityEndDate
        ? Timestamp.toJSON(message.validityEndDate)
        : undefined);
    if (message.permissions) {
      obj.permissions = message.permissions.map((e) =>
        e ? LicensePermissionEntity.toJSON(e) : undefined
      );
    } else {
      obj.permissions = [];
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

  create<I extends Exact<DeepPartial<LicenseEntity>, I>>(
    base?: I
  ): LicenseEntity {
    return LicenseEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicenseEntity>, I>>(
    object: I
  ): LicenseEntity {
    const message = createBaseLicenseEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? LicenseEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.name = object.name ?? undefined;
    message.signedLicense = object.signedLicense ?? undefined;
    message.validityStartDate =
      object.validityStartDate !== undefined &&
      object.validityStartDate !== null
        ? Timestamp.fromPartial(object.validityStartDate)
        : undefined;
    message.validityEndDate =
      object.validityEndDate !== undefined && object.validityEndDate !== null
        ? Timestamp.fromPartial(object.validityEndDate)
        : undefined;
    message.permissions =
      object.permissions?.map((e) => LicensePermissionEntity.fromPartial(e)) ||
      [];
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseLicensePermissionEntity(): LicensePermissionEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    capabilityId: undefined,
    capability: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const LicensePermissionEntity = {
  encode(
    message: LicensePermissionEntity,
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
      LicensePermissionEntity.encode(
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
    if (message.capabilityId !== undefined) {
      writer.uint32(82).string(message.capabilityId);
    }
    if (message.capability !== undefined) {
      CapabilityEntity.encode(
        message.capability,
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(96).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(104).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(112).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(122).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(130).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): LicensePermissionEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicensePermissionEntity();
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
          message.parent = LicensePermissionEntity.decode(
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
          message.capabilityId = reader.string();
          break;
        case 11:
          message.capability = CapabilityEntity.decode(reader, reader.uint32());
          break;
        case 12:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.createdFormatted = reader.string();
          break;
        case 16:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LicensePermissionEntity {
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
        ? LicensePermissionEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      capabilityId: isSet(object.capabilityId)
        ? String(object.capabilityId)
        : undefined,
      capability: isSet(object.capability)
        ? CapabilityEntity.fromJSON(object.capability)
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

  toJSON(message: LicensePermissionEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? LicensePermissionEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.capabilityId !== undefined &&
      (obj.capabilityId = message.capabilityId);
    message.capability !== undefined &&
      (obj.capability = message.capability
        ? CapabilityEntity.toJSON(message.capability)
        : undefined);
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

  create<I extends Exact<DeepPartial<LicensePermissionEntity>, I>>(
    base?: I
  ): LicensePermissionEntity {
    return LicensePermissionEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicensePermissionEntity>, I>>(
    object: I
  ): LicensePermissionEntity {
    const message = createBaseLicensePermissionEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? LicensePermissionEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.capabilityId = object.capabilityId ?? undefined;
    message.capability =
      object.capability !== undefined && object.capability !== null
        ? CapabilityEntity.fromPartial(object.capability)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseLicenseFromActivationKeyDto(): LicenseFromActivationKeyDto {
  return { activationKeyId: "", machineId: "" };
}

export const LicenseFromActivationKeyDto = {
  encode(
    message: LicenseFromActivationKeyDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.activationKeyId !== "") {
      writer.uint32(10).string(message.activationKeyId);
    }
    if (message.machineId !== "") {
      writer.uint32(18).string(message.machineId);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): LicenseFromActivationKeyDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicenseFromActivationKeyDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.activationKeyId = reader.string();
          break;
        case 2:
          message.machineId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LicenseFromActivationKeyDto {
    return {
      activationKeyId: isSet(object.activationKeyId)
        ? String(object.activationKeyId)
        : "",
      machineId: isSet(object.machineId) ? String(object.machineId) : "",
    };
  },

  toJSON(message: LicenseFromActivationKeyDto): unknown {
    const obj: any = {};
    message.activationKeyId !== undefined &&
      (obj.activationKeyId = message.activationKeyId);
    message.machineId !== undefined && (obj.machineId = message.machineId);
    return obj;
  },

  create<I extends Exact<DeepPartial<LicenseFromActivationKeyDto>, I>>(
    base?: I
  ): LicenseFromActivationKeyDto {
    return LicenseFromActivationKeyDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicenseFromActivationKeyDto>, I>>(
    object: I
  ): LicenseFromActivationKeyDto {
    const message = createBaseLicenseFromActivationKeyDto();
    message.activationKeyId = object.activationKeyId ?? "";
    message.machineId = object.machineId ?? "";
    return message;
  },
};

function createBaseLicenseFromPlanIdDto(): LicenseFromPlanIdDto {
  return { machineId: "", email: "", owner: "" };
}

export const LicenseFromPlanIdDto = {
  encode(
    message: LicenseFromPlanIdDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.machineId !== "") {
      writer.uint32(10).string(message.machineId);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): LicenseFromPlanIdDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLicenseFromPlanIdDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.machineId = reader.string();
          break;
        case 2:
          message.email = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LicenseFromPlanIdDto {
    return {
      machineId: isSet(object.machineId) ? String(object.machineId) : "",
      email: isSet(object.email) ? String(object.email) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
    };
  },

  toJSON(message: LicenseFromPlanIdDto): unknown {
    const obj: any = {};
    message.machineId !== undefined && (obj.machineId = message.machineId);
    message.email !== undefined && (obj.email = message.email);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  create<I extends Exact<DeepPartial<LicenseFromPlanIdDto>, I>>(
    base?: I
  ): LicenseFromPlanIdDto {
    return LicenseFromPlanIdDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LicenseFromPlanIdDto>, I>>(
    object: I
  ): LicenseFromPlanIdDto {
    const message = createBaseLicenseFromPlanIdDto();
    message.machineId = object.machineId ?? "";
    message.email = object.email ?? "";
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseProductPlanCreateReply(): ProductPlanCreateReply {
  return { data: undefined, error: undefined };
}

export const ProductPlanCreateReply = {
  encode(
    message: ProductPlanCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      ProductPlanEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ProductPlanCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProductPlanCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = ProductPlanEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): ProductPlanCreateReply {
    return {
      data: isSet(object.data)
        ? ProductPlanEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ProductPlanCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? ProductPlanEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ProductPlanCreateReply>, I>>(
    base?: I
  ): ProductPlanCreateReply {
    return ProductPlanCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProductPlanCreateReply>, I>>(
    object: I
  ): ProductPlanCreateReply {
    const message = createBaseProductPlanCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? ProductPlanEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseProductPlanReply(): ProductPlanReply {
  return { data: undefined, error: undefined };
}

export const ProductPlanReply = {
  encode(
    message: ProductPlanReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      ProductPlanEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProductPlanReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProductPlanReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = ProductPlanEntity.decode(reader, reader.uint32());
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

  fromJSON(object: any): ProductPlanReply {
    return {
      data: isSet(object.data)
        ? ProductPlanEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ProductPlanReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? ProductPlanEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ProductPlanReply>, I>>(
    base?: I
  ): ProductPlanReply {
    return ProductPlanReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProductPlanReply>, I>>(
    object: I
  ): ProductPlanReply {
    const message = createBaseProductPlanReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? ProductPlanEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseProductPlanQueryReply(): ProductPlanQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const ProductPlanQueryReply = {
  encode(
    message: ProductPlanQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      ProductPlanEntity.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): ProductPlanQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProductPlanQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(ProductPlanEntity.decode(reader, reader.uint32()));
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

  fromJSON(object: any): ProductPlanQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => ProductPlanEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ProductPlanQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? ProductPlanEntity.toJSON(e) : undefined
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

  create<I extends Exact<DeepPartial<ProductPlanQueryReply>, I>>(
    base?: I
  ): ProductPlanQueryReply {
    return ProductPlanQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProductPlanQueryReply>, I>>(
    object: I
  ): ProductPlanQueryReply {
    const message = createBaseProductPlanQueryReply();
    message.items =
      object.items?.map((e) => ProductPlanEntity.fromPartial(e)) || [];
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

function createBaseProductPlanEntity(): ProductPlanEntity {
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
    duration: undefined,
    productId: undefined,
    product: undefined,
    priceTagId: undefined,
    priceTag: undefined,
    permissions: [],
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const ProductPlanEntity = {
  encode(
    message: ProductPlanEntity,
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
      ProductPlanEntity.encode(
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
      ProductPlanEntityPolyglot.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.name !== undefined) {
      writer.uint32(82).string(message.name);
    }
    if (message.duration !== undefined) {
      writer.uint32(88).int64(message.duration);
    }
    if (message.productId !== undefined) {
      writer.uint32(106).string(message.productId);
    }
    if (message.product !== undefined) {
      LicensableProductEntity.encode(
        message.product,
        writer.uint32(114).fork()
      ).ldelim();
    }
    if (message.priceTagId !== undefined) {
      writer.uint32(130).string(message.priceTagId);
    }
    if (message.priceTag !== undefined) {
      PriceTagEntity.encode(
        message.priceTag,
        writer.uint32(138).fork()
      ).ldelim();
    }
    for (const v of message.permissions) {
      ProductPlanPermissionEntity.encode(
        v!,
        writer.uint32(146).fork()
      ).ldelim();
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

  decode(input: _m0.Reader | Uint8Array, length?: number): ProductPlanEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProductPlanEntity();
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
          message.parent = ProductPlanEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            ProductPlanEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.name = reader.string();
          break;
        case 11:
          message.duration = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.productId = reader.string();
          break;
        case 14:
          message.product = LicensableProductEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 16:
          message.priceTagId = reader.string();
          break;
        case 17:
          message.priceTag = PriceTagEntity.decode(reader, reader.uint32());
          break;
        case 18:
          message.permissions.push(
            ProductPlanPermissionEntity.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): ProductPlanEntity {
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
        ? ProductPlanEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            ProductPlanEntityPolyglot.fromJSON(e)
          )
        : [],
      name: isSet(object.name) ? String(object.name) : undefined,
      duration: isSet(object.duration) ? Number(object.duration) : undefined,
      productId: isSet(object.productId) ? String(object.productId) : undefined,
      product: isSet(object.product)
        ? LicensableProductEntity.fromJSON(object.product)
        : undefined,
      priceTagId: isSet(object.priceTagId)
        ? String(object.priceTagId)
        : undefined,
      priceTag: isSet(object.priceTag)
        ? PriceTagEntity.fromJSON(object.priceTag)
        : undefined,
      permissions: Array.isArray(object?.permissions)
        ? object.permissions.map((e: any) =>
            ProductPlanPermissionEntity.fromJSON(e)
          )
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

  toJSON(message: ProductPlanEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? ProductPlanEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? ProductPlanEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.name !== undefined && (obj.name = message.name);
    message.duration !== undefined &&
      (obj.duration = Math.round(message.duration));
    message.productId !== undefined && (obj.productId = message.productId);
    message.product !== undefined &&
      (obj.product = message.product
        ? LicensableProductEntity.toJSON(message.product)
        : undefined);
    message.priceTagId !== undefined && (obj.priceTagId = message.priceTagId);
    message.priceTag !== undefined &&
      (obj.priceTag = message.priceTag
        ? PriceTagEntity.toJSON(message.priceTag)
        : undefined);
    if (message.permissions) {
      obj.permissions = message.permissions.map((e) =>
        e ? ProductPlanPermissionEntity.toJSON(e) : undefined
      );
    } else {
      obj.permissions = [];
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

  create<I extends Exact<DeepPartial<ProductPlanEntity>, I>>(
    base?: I
  ): ProductPlanEntity {
    return ProductPlanEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProductPlanEntity>, I>>(
    object: I
  ): ProductPlanEntity {
    const message = createBaseProductPlanEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? ProductPlanEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) =>
        ProductPlanEntityPolyglot.fromPartial(e)
      ) || [];
    message.name = object.name ?? undefined;
    message.duration = object.duration ?? undefined;
    message.productId = object.productId ?? undefined;
    message.product =
      object.product !== undefined && object.product !== null
        ? LicensableProductEntity.fromPartial(object.product)
        : undefined;
    message.priceTagId = object.priceTagId ?? undefined;
    message.priceTag =
      object.priceTag !== undefined && object.priceTag !== null
        ? PriceTagEntity.fromPartial(object.priceTag)
        : undefined;
    message.permissions =
      object.permissions?.map((e) =>
        ProductPlanPermissionEntity.fromPartial(e)
      ) || [];
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseProductPlanEntityPolyglot(): ProductPlanEntityPolyglot {
  return { linkerId: "", languageId: "", name: "" };
}

export const ProductPlanEntityPolyglot = {
  encode(
    message: ProductPlanEntityPolyglot,
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
  ): ProductPlanEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProductPlanEntityPolyglot();
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

  fromJSON(object: any): ProductPlanEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: ProductPlanEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<ProductPlanEntityPolyglot>, I>>(
    base?: I
  ): ProductPlanEntityPolyglot {
    return ProductPlanEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProductPlanEntityPolyglot>, I>>(
    object: I
  ): ProductPlanEntityPolyglot {
    const message = createBaseProductPlanEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseProductPlanPermissionEntity(): ProductPlanPermissionEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    capabilityId: undefined,
    capability: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const ProductPlanPermissionEntity = {
  encode(
    message: ProductPlanPermissionEntity,
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
      ProductPlanPermissionEntity.encode(
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
    if (message.capabilityId !== undefined) {
      writer.uint32(82).string(message.capabilityId);
    }
    if (message.capability !== undefined) {
      CapabilityEntity.encode(
        message.capability,
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(96).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(104).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(112).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(122).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(130).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ProductPlanPermissionEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProductPlanPermissionEntity();
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
          message.parent = ProductPlanPermissionEntity.decode(
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
          message.capabilityId = reader.string();
          break;
        case 11:
          message.capability = CapabilityEntity.decode(reader, reader.uint32());
          break;
        case 12:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.createdFormatted = reader.string();
          break;
        case 16:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ProductPlanPermissionEntity {
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
        ? ProductPlanPermissionEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      capabilityId: isSet(object.capabilityId)
        ? String(object.capabilityId)
        : undefined,
      capability: isSet(object.capability)
        ? CapabilityEntity.fromJSON(object.capability)
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

  toJSON(message: ProductPlanPermissionEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? ProductPlanPermissionEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.capabilityId !== undefined &&
      (obj.capabilityId = message.capabilityId);
    message.capability !== undefined &&
      (obj.capability = message.capability
        ? CapabilityEntity.toJSON(message.capability)
        : undefined);
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

  create<I extends Exact<DeepPartial<ProductPlanPermissionEntity>, I>>(
    base?: I
  ): ProductPlanPermissionEntity {
    return ProductPlanPermissionEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProductPlanPermissionEntity>, I>>(
    object: I
  ): ProductPlanPermissionEntity {
    const message = createBaseProductPlanPermissionEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? ProductPlanPermissionEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.capabilityId = object.capabilityId ?? undefined;
    message.capability =
      object.capability !== undefined && object.capability !== null
        ? CapabilityEntity.fromPartial(object.capability)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

export interface ActivationKeys {
  ActivationKeyActionCreate(
    request: ActivationKeyEntity
  ): Promise<ActivationKeyCreateReply>;
  ActivationKeyActionUpdate(
    request: ActivationKeyEntity
  ): Promise<ActivationKeyCreateReply>;
  ActivationKeyActionQuery(
    request: QueryFilterRequest
  ): Promise<ActivationKeyQueryReply>;
  ActivationKeyActionGetOne(
    request: QueryFilterRequest
  ): Promise<ActivationKeyReply>;
  ActivationKeyActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class ActivationKeysClientImpl implements ActivationKeys {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "ActivationKeys";
    this.rpc = rpc;
    this.ActivationKeyActionCreate = this.ActivationKeyActionCreate.bind(this);
    this.ActivationKeyActionUpdate = this.ActivationKeyActionUpdate.bind(this);
    this.ActivationKeyActionQuery = this.ActivationKeyActionQuery.bind(this);
    this.ActivationKeyActionGetOne = this.ActivationKeyActionGetOne.bind(this);
    this.ActivationKeyActionRemove = this.ActivationKeyActionRemove.bind(this);
  }
  ActivationKeyActionCreate(
    request: ActivationKeyEntity
  ): Promise<ActivationKeyCreateReply> {
    const data = ActivationKeyEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ActivationKeyActionCreate",
      data
    );
    return promise.then((data) =>
      ActivationKeyCreateReply.decode(new _m0.Reader(data))
    );
  }

  ActivationKeyActionUpdate(
    request: ActivationKeyEntity
  ): Promise<ActivationKeyCreateReply> {
    const data = ActivationKeyEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ActivationKeyActionUpdate",
      data
    );
    return promise.then((data) =>
      ActivationKeyCreateReply.decode(new _m0.Reader(data))
    );
  }

  ActivationKeyActionQuery(
    request: QueryFilterRequest
  ): Promise<ActivationKeyQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ActivationKeyActionQuery",
      data
    );
    return promise.then((data) =>
      ActivationKeyQueryReply.decode(new _m0.Reader(data))
    );
  }

  ActivationKeyActionGetOne(
    request: QueryFilterRequest
  ): Promise<ActivationKeyReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ActivationKeyActionGetOne",
      data
    );
    return promise.then((data) =>
      ActivationKeyReply.decode(new _m0.Reader(data))
    );
  }

  ActivationKeyActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ActivationKeyActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface LicensableProducts {
  LicensableProductActionCreate(
    request: LicensableProductEntity
  ): Promise<LicensableProductCreateReply>;
  LicensableProductActionUpdate(
    request: LicensableProductEntity
  ): Promise<LicensableProductCreateReply>;
  LicensableProductActionQuery(
    request: QueryFilterRequest
  ): Promise<LicensableProductQueryReply>;
  LicensableProductActionGetOne(
    request: QueryFilterRequest
  ): Promise<LicensableProductReply>;
  LicensableProductActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class LicensableProductsClientImpl implements LicensableProducts {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "LicensableProducts";
    this.rpc = rpc;
    this.LicensableProductActionCreate =
      this.LicensableProductActionCreate.bind(this);
    this.LicensableProductActionUpdate =
      this.LicensableProductActionUpdate.bind(this);
    this.LicensableProductActionQuery =
      this.LicensableProductActionQuery.bind(this);
    this.LicensableProductActionGetOne =
      this.LicensableProductActionGetOne.bind(this);
    this.LicensableProductActionRemove =
      this.LicensableProductActionRemove.bind(this);
  }
  LicensableProductActionCreate(
    request: LicensableProductEntity
  ): Promise<LicensableProductCreateReply> {
    const data = LicensableProductEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "LicensableProductActionCreate",
      data
    );
    return promise.then((data) =>
      LicensableProductCreateReply.decode(new _m0.Reader(data))
    );
  }

  LicensableProductActionUpdate(
    request: LicensableProductEntity
  ): Promise<LicensableProductCreateReply> {
    const data = LicensableProductEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "LicensableProductActionUpdate",
      data
    );
    return promise.then((data) =>
      LicensableProductCreateReply.decode(new _m0.Reader(data))
    );
  }

  LicensableProductActionQuery(
    request: QueryFilterRequest
  ): Promise<LicensableProductQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "LicensableProductActionQuery",
      data
    );
    return promise.then((data) =>
      LicensableProductQueryReply.decode(new _m0.Reader(data))
    );
  }

  LicensableProductActionGetOne(
    request: QueryFilterRequest
  ): Promise<LicensableProductReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "LicensableProductActionGetOne",
      data
    );
    return promise.then((data) =>
      LicensableProductReply.decode(new _m0.Reader(data))
    );
  }

  LicensableProductActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "LicensableProductActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Licenses {
  LicenseActionCreate(request: LicenseEntity): Promise<LicenseCreateReply>;
  LicenseActionUpdate(request: LicenseEntity): Promise<LicenseCreateReply>;
  LicenseActionQuery(request: QueryFilterRequest): Promise<LicenseQueryReply>;
  LicenseActionGetOne(request: QueryFilterRequest): Promise<LicenseReply>;
  LicenseActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class LicensesClientImpl implements Licenses {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Licenses";
    this.rpc = rpc;
    this.LicenseActionCreate = this.LicenseActionCreate.bind(this);
    this.LicenseActionUpdate = this.LicenseActionUpdate.bind(this);
    this.LicenseActionQuery = this.LicenseActionQuery.bind(this);
    this.LicenseActionGetOne = this.LicenseActionGetOne.bind(this);
    this.LicenseActionRemove = this.LicenseActionRemove.bind(this);
  }
  LicenseActionCreate(request: LicenseEntity): Promise<LicenseCreateReply> {
    const data = LicenseEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "LicenseActionCreate", data);
    return promise.then((data) =>
      LicenseCreateReply.decode(new _m0.Reader(data))
    );
  }

  LicenseActionUpdate(request: LicenseEntity): Promise<LicenseCreateReply> {
    const data = LicenseEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "LicenseActionUpdate", data);
    return promise.then((data) =>
      LicenseCreateReply.decode(new _m0.Reader(data))
    );
  }

  LicenseActionQuery(request: QueryFilterRequest): Promise<LicenseQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "LicenseActionQuery", data);
    return promise.then((data) =>
      LicenseQueryReply.decode(new _m0.Reader(data))
    );
  }

  LicenseActionGetOne(request: QueryFilterRequest): Promise<LicenseReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "LicenseActionGetOne", data);
    return promise.then((data) => LicenseReply.decode(new _m0.Reader(data)));
  }

  LicenseActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "LicenseActionRemove", data);
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface ProductPlans {
  ProductPlanActionCreate(
    request: ProductPlanEntity
  ): Promise<ProductPlanCreateReply>;
  ProductPlanActionUpdate(
    request: ProductPlanEntity
  ): Promise<ProductPlanCreateReply>;
  ProductPlanActionQuery(
    request: QueryFilterRequest
  ): Promise<ProductPlanQueryReply>;
  ProductPlanActionGetOne(
    request: QueryFilterRequest
  ): Promise<ProductPlanReply>;
  ProductPlanActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class ProductPlansClientImpl implements ProductPlans {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "ProductPlans";
    this.rpc = rpc;
    this.ProductPlanActionCreate = this.ProductPlanActionCreate.bind(this);
    this.ProductPlanActionUpdate = this.ProductPlanActionUpdate.bind(this);
    this.ProductPlanActionQuery = this.ProductPlanActionQuery.bind(this);
    this.ProductPlanActionGetOne = this.ProductPlanActionGetOne.bind(this);
    this.ProductPlanActionRemove = this.ProductPlanActionRemove.bind(this);
  }
  ProductPlanActionCreate(
    request: ProductPlanEntity
  ): Promise<ProductPlanCreateReply> {
    const data = ProductPlanEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ProductPlanActionCreate",
      data
    );
    return promise.then((data) =>
      ProductPlanCreateReply.decode(new _m0.Reader(data))
    );
  }

  ProductPlanActionUpdate(
    request: ProductPlanEntity
  ): Promise<ProductPlanCreateReply> {
    const data = ProductPlanEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ProductPlanActionUpdate",
      data
    );
    return promise.then((data) =>
      ProductPlanCreateReply.decode(new _m0.Reader(data))
    );
  }

  ProductPlanActionQuery(
    request: QueryFilterRequest
  ): Promise<ProductPlanQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ProductPlanActionQuery",
      data
    );
    return promise.then((data) =>
      ProductPlanQueryReply.decode(new _m0.Reader(data))
    );
  }

  ProductPlanActionGetOne(
    request: QueryFilterRequest
  ): Promise<ProductPlanReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ProductPlanActionGetOne",
      data
    );
    return promise.then((data) =>
      ProductPlanReply.decode(new _m0.Reader(data))
    );
  }

  ProductPlanActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ProductPlanActionRemove",
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
