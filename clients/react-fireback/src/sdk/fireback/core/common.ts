/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export interface Timestamp {
  seconds: number;
  nanos: number;
}

export interface QueryFilter {
  query: string;
  startIndex: number;
  itemsPerPage: number;
  id: string;
  acceptLanguage: string;
  uniqueId: string;
}

export interface RemoveRequestData {
  rowsAffected: number;
}

export interface QueryFilterRequest {
  query: QueryFilter | undefined;
}

export interface DeleteRequest {
  list: string[];
  query: string;
  suspense: boolean;
}

export interface DeleteResponseData {
  DeleteRequest: number;
}

export interface DeleteResponse {
  data: DeleteResponseData | undefined;
}

export interface IErrorItem {
  location: string;
  message: string;
  messageTranslated: string;
  errorParam: string;
  type: string;
}

export interface EmptyRequest {}

export interface OkayResponseData {}

export interface OkayResponse {
  data: OkayResponseData | undefined;
}

export interface IError {
  message: string;
  messageTranslated: string;
  errors: IErrorItem[];
  httpCode: number;
}

export interface RemoveReply {
  rowsAffected: number;
  error: IError | undefined;
}

function createBaseTimestamp(): Timestamp {
  return { seconds: 0, nanos: 0 };
}

export const Timestamp = {
  encode(
    message: Timestamp,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.seconds !== 0) {
      writer.uint32(8).int64(message.seconds);
    }
    if (message.nanos !== 0) {
      writer.uint32(16).int32(message.nanos);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Timestamp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTimestamp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.seconds = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.nanos = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Timestamp {
    return {
      seconds: isSet(object.seconds) ? Number(object.seconds) : 0,
      nanos: isSet(object.nanos) ? Number(object.nanos) : 0,
    };
  },

  toJSON(message: Timestamp): unknown {
    const obj: any = {};
    message.seconds !== undefined &&
      (obj.seconds = Math.round(message.seconds));
    message.nanos !== undefined && (obj.nanos = Math.round(message.nanos));
    return obj;
  },

  create<I extends Exact<DeepPartial<Timestamp>, I>>(base?: I): Timestamp {
    return Timestamp.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Timestamp>, I>>(
    object: I
  ): Timestamp {
    const message = createBaseTimestamp();
    message.seconds = object.seconds ?? 0;
    message.nanos = object.nanos ?? 0;
    return message;
  },
};

function createBaseQueryFilter(): QueryFilter {
  return {
    query: "",
    startIndex: 0,
    itemsPerPage: 0,
    id: "",
    acceptLanguage: "",
    uniqueId: "",
  };
}

export const QueryFilter = {
  encode(
    message: QueryFilter,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.query !== "") {
      writer.uint32(10).string(message.query);
    }
    if (message.startIndex !== 0) {
      writer.uint32(16).int64(message.startIndex);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.id !== "") {
      writer.uint32(34).string(message.id);
    }
    if (message.acceptLanguage !== "") {
      writer.uint32(42).string(message.acceptLanguage);
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryFilter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFilter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.query = reader.string();
          break;
        case 2:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.id = reader.string();
          break;
        case 5:
          message.acceptLanguage = reader.string();
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilter {
    return {
      query: isSet(object.query) ? String(object.query) : "",
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      id: isSet(object.id) ? String(object.id) : "",
      acceptLanguage: isSet(object.acceptLanguage)
        ? String(object.acceptLanguage)
        : "",
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
    };
  },

  toJSON(message: QueryFilter): unknown {
    const obj: any = {};
    message.query !== undefined && (obj.query = message.query);
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.id !== undefined && (obj.id = message.id);
    message.acceptLanguage !== undefined &&
      (obj.acceptLanguage = message.acceptLanguage);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryFilter>, I>>(base?: I): QueryFilter {
    return QueryFilter.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<QueryFilter>, I>>(
    object: I
  ): QueryFilter {
    const message = createBaseQueryFilter();
    message.query = object.query ?? "";
    message.startIndex = object.startIndex ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.id = object.id ?? "";
    message.acceptLanguage = object.acceptLanguage ?? "";
    message.uniqueId = object.uniqueId ?? "";
    return message;
  },
};

function createBaseRemoveRequestData(): RemoveRequestData {
  return { rowsAffected: 0 };
}

export const RemoveRequestData = {
  encode(
    message: RemoveRequestData,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.rowsAffected !== 0) {
      writer.uint32(8).int64(message.rowsAffected);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RemoveRequestData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRemoveRequestData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rowsAffected = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RemoveRequestData {
    return {
      rowsAffected: isSet(object.rowsAffected)
        ? Number(object.rowsAffected)
        : 0,
    };
  },

  toJSON(message: RemoveRequestData): unknown {
    const obj: any = {};
    message.rowsAffected !== undefined &&
      (obj.rowsAffected = Math.round(message.rowsAffected));
    return obj;
  },

  create<I extends Exact<DeepPartial<RemoveRequestData>, I>>(
    base?: I
  ): RemoveRequestData {
    return RemoveRequestData.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RemoveRequestData>, I>>(
    object: I
  ): RemoveRequestData {
    const message = createBaseRemoveRequestData();
    message.rowsAffected = object.rowsAffected ?? 0;
    return message;
  },
};

function createBaseQueryFilterRequest(): QueryFilterRequest {
  return { query: undefined };
}

export const QueryFilterRequest = {
  encode(
    message: QueryFilterRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.query !== undefined) {
      QueryFilter.encode(message.query, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryFilterRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFilterRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.query = QueryFilter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilterRequest {
    return {
      query: isSet(object.query)
        ? QueryFilter.fromJSON(object.query)
        : undefined,
    };
  },

  toJSON(message: QueryFilterRequest): unknown {
    const obj: any = {};
    message.query !== undefined &&
      (obj.query = message.query
        ? QueryFilter.toJSON(message.query)
        : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryFilterRequest>, I>>(
    base?: I
  ): QueryFilterRequest {
    return QueryFilterRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<QueryFilterRequest>, I>>(
    object: I
  ): QueryFilterRequest {
    const message = createBaseQueryFilterRequest();
    message.query =
      object.query !== undefined && object.query !== null
        ? QueryFilter.fromPartial(object.query)
        : undefined;
    return message;
  },
};

function createBaseDeleteRequest(): DeleteRequest {
  return { list: [], query: "", suspense: false };
}

export const DeleteRequest = {
  encode(
    message: DeleteRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.list) {
      writer.uint32(10).string(v!);
    }
    if (message.query !== "") {
      writer.uint32(18).string(message.query);
    }
    if (message.suspense === true) {
      writer.uint32(24).bool(message.suspense);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.list.push(reader.string());
          break;
        case 2:
          message.query = reader.string();
          break;
        case 3:
          message.suspense = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteRequest {
    return {
      list: Array.isArray(object?.list)
        ? object.list.map((e: any) => String(e))
        : [],
      query: isSet(object.query) ? String(object.query) : "",
      suspense: isSet(object.suspense) ? Boolean(object.suspense) : false,
    };
  },

  toJSON(message: DeleteRequest): unknown {
    const obj: any = {};
    if (message.list) {
      obj.list = message.list.map((e) => e);
    } else {
      obj.list = [];
    }
    message.query !== undefined && (obj.query = message.query);
    message.suspense !== undefined && (obj.suspense = message.suspense);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteRequest>, I>>(
    base?: I
  ): DeleteRequest {
    return DeleteRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteRequest>, I>>(
    object: I
  ): DeleteRequest {
    const message = createBaseDeleteRequest();
    message.list = object.list?.map((e) => e) || [];
    message.query = object.query ?? "";
    message.suspense = object.suspense ?? false;
    return message;
  },
};

function createBaseDeleteResponseData(): DeleteResponseData {
  return { DeleteRequest: 0 };
}

export const DeleteResponseData = {
  encode(
    message: DeleteResponseData,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.DeleteRequest !== 0) {
      writer.uint32(8).int64(message.DeleteRequest);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteResponseData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteResponseData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.DeleteRequest = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteResponseData {
    return {
      DeleteRequest: isSet(object.DeleteRequest)
        ? Number(object.DeleteRequest)
        : 0,
    };
  },

  toJSON(message: DeleteResponseData): unknown {
    const obj: any = {};
    message.DeleteRequest !== undefined &&
      (obj.DeleteRequest = Math.round(message.DeleteRequest));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteResponseData>, I>>(
    base?: I
  ): DeleteResponseData {
    return DeleteResponseData.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteResponseData>, I>>(
    object: I
  ): DeleteResponseData {
    const message = createBaseDeleteResponseData();
    message.DeleteRequest = object.DeleteRequest ?? 0;
    return message;
  },
};

function createBaseDeleteResponse(): DeleteResponse {
  return { data: undefined };
}

export const DeleteResponse = {
  encode(
    message: DeleteResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      DeleteResponseData.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = DeleteResponseData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteResponse {
    return {
      data: isSet(object.data)
        ? DeleteResponseData.fromJSON(object.data)
        : undefined,
    };
  },

  toJSON(message: DeleteResponse): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? DeleteResponseData.toJSON(message.data)
        : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteResponse>, I>>(
    base?: I
  ): DeleteResponse {
    return DeleteResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteResponse>, I>>(
    object: I
  ): DeleteResponse {
    const message = createBaseDeleteResponse();
    message.data =
      object.data !== undefined && object.data !== null
        ? DeleteResponseData.fromPartial(object.data)
        : undefined;
    return message;
  },
};

function createBaseIErrorItem(): IErrorItem {
  return {
    location: "",
    message: "",
    messageTranslated: "",
    errorParam: "",
    type: "",
  };
}

export const IErrorItem = {
  encode(
    message: IErrorItem,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.location !== "") {
      writer.uint32(10).string(message.location);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.messageTranslated !== "") {
      writer.uint32(34).string(message.messageTranslated);
    }
    if (message.errorParam !== "") {
      writer.uint32(42).string(message.errorParam);
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IErrorItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIErrorItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.location = reader.string();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 4:
          message.messageTranslated = reader.string();
          break;
        case 5:
          message.errorParam = reader.string();
          break;
        case 3:
          message.type = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IErrorItem {
    return {
      location: isSet(object.location) ? String(object.location) : "",
      message: isSet(object.message) ? String(object.message) : "",
      messageTranslated: isSet(object.messageTranslated)
        ? String(object.messageTranslated)
        : "",
      errorParam: isSet(object.errorParam) ? String(object.errorParam) : "",
      type: isSet(object.type) ? String(object.type) : "",
    };
  },

  toJSON(message: IErrorItem): unknown {
    const obj: any = {};
    message.location !== undefined && (obj.location = message.location);
    message.message !== undefined && (obj.message = message.message);
    message.messageTranslated !== undefined &&
      (obj.messageTranslated = message.messageTranslated);
    message.errorParam !== undefined && (obj.errorParam = message.errorParam);
    message.type !== undefined && (obj.type = message.type);
    return obj;
  },

  create<I extends Exact<DeepPartial<IErrorItem>, I>>(base?: I): IErrorItem {
    return IErrorItem.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<IErrorItem>, I>>(
    object: I
  ): IErrorItem {
    const message = createBaseIErrorItem();
    message.location = object.location ?? "";
    message.message = object.message ?? "";
    message.messageTranslated = object.messageTranslated ?? "";
    message.errorParam = object.errorParam ?? "";
    message.type = object.type ?? "";
    return message;
  },
};

function createBaseEmptyRequest(): EmptyRequest {
  return {};
}

export const EmptyRequest = {
  encode(
    _: EmptyRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EmptyRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmptyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): EmptyRequest {
    return {};
  },

  toJSON(_: EmptyRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<EmptyRequest>, I>>(
    base?: I
  ): EmptyRequest {
    return EmptyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmptyRequest>, I>>(
    _: I
  ): EmptyRequest {
    const message = createBaseEmptyRequest();
    return message;
  },
};

function createBaseOkayResponseData(): OkayResponseData {
  return {};
}

export const OkayResponseData = {
  encode(
    _: OkayResponseData,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OkayResponseData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOkayResponseData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): OkayResponseData {
    return {};
  },

  toJSON(_: OkayResponseData): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<OkayResponseData>, I>>(
    base?: I
  ): OkayResponseData {
    return OkayResponseData.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OkayResponseData>, I>>(
    _: I
  ): OkayResponseData {
    const message = createBaseOkayResponseData();
    return message;
  },
};

function createBaseOkayResponse(): OkayResponse {
  return { data: undefined };
}

export const OkayResponse = {
  encode(
    message: OkayResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      OkayResponseData.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OkayResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOkayResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = OkayResponseData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OkayResponse {
    return {
      data: isSet(object.data)
        ? OkayResponseData.fromJSON(object.data)
        : undefined,
    };
  },

  toJSON(message: OkayResponse): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? OkayResponseData.toJSON(message.data)
        : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<OkayResponse>, I>>(
    base?: I
  ): OkayResponse {
    return OkayResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OkayResponse>, I>>(
    object: I
  ): OkayResponse {
    const message = createBaseOkayResponse();
    message.data =
      object.data !== undefined && object.data !== null
        ? OkayResponseData.fromPartial(object.data)
        : undefined;
    return message;
  },
};

function createBaseIError(): IError {
  return { message: "", messageTranslated: "", errors: [], httpCode: 0 };
}

export const IError = {
  encode(
    message: IError,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.messageTranslated !== "") {
      writer.uint32(34).string(message.messageTranslated);
    }
    for (const v of message.errors) {
      IErrorItem.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.httpCode !== 0) {
      writer.uint32(24).int32(message.httpCode);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IError {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIError();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        case 4:
          message.messageTranslated = reader.string();
          break;
        case 2:
          message.errors.push(IErrorItem.decode(reader, reader.uint32()));
          break;
        case 3:
          message.httpCode = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IError {
    return {
      message: isSet(object.message) ? String(object.message) : "",
      messageTranslated: isSet(object.messageTranslated)
        ? String(object.messageTranslated)
        : "",
      errors: Array.isArray(object?.errors)
        ? object.errors.map((e: any) => IErrorItem.fromJSON(e))
        : [],
      httpCode: isSet(object.httpCode) ? Number(object.httpCode) : 0,
    };
  },

  toJSON(message: IError): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    message.messageTranslated !== undefined &&
      (obj.messageTranslated = message.messageTranslated);
    if (message.errors) {
      obj.errors = message.errors.map((e) =>
        e ? IErrorItem.toJSON(e) : undefined
      );
    } else {
      obj.errors = [];
    }
    message.httpCode !== undefined &&
      (obj.httpCode = Math.round(message.httpCode));
    return obj;
  },

  create<I extends Exact<DeepPartial<IError>, I>>(base?: I): IError {
    return IError.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<IError>, I>>(object: I): IError {
    const message = createBaseIError();
    message.message = object.message ?? "";
    message.messageTranslated = object.messageTranslated ?? "";
    message.errors = object.errors?.map((e) => IErrorItem.fromPartial(e)) || [];
    message.httpCode = object.httpCode ?? 0;
    return message;
  },
};

function createBaseRemoveReply(): RemoveReply {
  return { rowsAffected: 0, error: undefined };
}

export const RemoveReply = {
  encode(
    message: RemoveReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.rowsAffected !== 0) {
      writer.uint32(8).int64(message.rowsAffected);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RemoveReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRemoveReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rowsAffected = longToNumber(reader.int64() as Long);
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

  fromJSON(object: any): RemoveReply {
    return {
      rowsAffected: isSet(object.rowsAffected)
        ? Number(object.rowsAffected)
        : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: RemoveReply): unknown {
    const obj: any = {};
    message.rowsAffected !== undefined &&
      (obj.rowsAffected = Math.round(message.rowsAffected));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<RemoveReply>, I>>(base?: I): RemoveReply {
    return RemoveReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RemoveReply>, I>>(
    object: I
  ): RemoveReply {
    const message = createBaseRemoveReply();
    message.rowsAffected = object.rowsAffected ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
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
