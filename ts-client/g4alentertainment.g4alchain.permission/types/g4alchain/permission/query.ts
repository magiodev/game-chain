/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Administrator } from "./administrator";
import { Params } from "./params";

export const protobufPackage = "g4alentertainment.g4alchain.permission";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetAdministratorRequest {
  address: string;
}

export interface QueryGetAdministratorResponse {
  administrator: Administrator | undefined;
}

export interface QueryAllAdministratorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAdministratorResponse {
  administrator: Administrator[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetAdministratorRequest(): QueryGetAdministratorRequest {
  return { address: "" };
}

export const QueryGetAdministratorRequest = {
  encode(message: QueryGetAdministratorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAdministratorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAdministratorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAdministratorRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryGetAdministratorRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAdministratorRequest>, I>>(object: I): QueryGetAdministratorRequest {
    const message = createBaseQueryGetAdministratorRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryGetAdministratorResponse(): QueryGetAdministratorResponse {
  return { administrator: undefined };
}

export const QueryGetAdministratorResponse = {
  encode(message: QueryGetAdministratorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.administrator !== undefined) {
      Administrator.encode(message.administrator, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAdministratorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAdministratorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.administrator = Administrator.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAdministratorResponse {
    return { administrator: isSet(object.administrator) ? Administrator.fromJSON(object.administrator) : undefined };
  },

  toJSON(message: QueryGetAdministratorResponse): unknown {
    const obj: any = {};
    message.administrator !== undefined
      && (obj.administrator = message.administrator ? Administrator.toJSON(message.administrator) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAdministratorResponse>, I>>(
    object: I,
  ): QueryGetAdministratorResponse {
    const message = createBaseQueryGetAdministratorResponse();
    message.administrator = (object.administrator !== undefined && object.administrator !== null)
      ? Administrator.fromPartial(object.administrator)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAdministratorRequest(): QueryAllAdministratorRequest {
  return { pagination: undefined };
}

export const QueryAllAdministratorRequest = {
  encode(message: QueryAllAdministratorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAdministratorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAdministratorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAdministratorRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllAdministratorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllAdministratorRequest>, I>>(object: I): QueryAllAdministratorRequest {
    const message = createBaseQueryAllAdministratorRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAdministratorResponse(): QueryAllAdministratorResponse {
  return { administrator: [], pagination: undefined };
}

export const QueryAllAdministratorResponse = {
  encode(message: QueryAllAdministratorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.administrator) {
      Administrator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAdministratorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAdministratorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.administrator.push(Administrator.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAdministratorResponse {
    return {
      administrator: Array.isArray(object?.administrator)
        ? object.administrator.map((e: any) => Administrator.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllAdministratorResponse): unknown {
    const obj: any = {};
    if (message.administrator) {
      obj.administrator = message.administrator.map((e) => e ? Administrator.toJSON(e) : undefined);
    } else {
      obj.administrator = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllAdministratorResponse>, I>>(
    object: I,
  ): QueryAllAdministratorResponse {
    const message = createBaseQueryAllAdministratorResponse();
    message.administrator = object.administrator?.map((e) => Administrator.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Administrator by index. */
  Administrator(request: QueryGetAdministratorRequest): Promise<QueryGetAdministratorResponse>;
  /** Queries a list of Administrator items. */
  AdministratorAll(request: QueryAllAdministratorRequest): Promise<QueryAllAdministratorResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Administrator = this.Administrator.bind(this);
    this.AdministratorAll = this.AdministratorAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.permission.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Administrator(request: QueryGetAdministratorRequest): Promise<QueryGetAdministratorResponse> {
    const data = QueryGetAdministratorRequest.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.permission.Query", "Administrator", data);
    return promise.then((data) => QueryGetAdministratorResponse.decode(new _m0.Reader(data)));
  }

  AdministratorAll(request: QueryAllAdministratorRequest): Promise<QueryAllAdministratorResponse> {
    const data = QueryAllAdministratorRequest.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.permission.Query", "AdministratorAll", data);
    return promise.then((data) => QueryAllAdministratorResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
