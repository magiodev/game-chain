/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "g4alentertainment.g4alchain.permission";

export interface MsgCreateAdministrator {
  creator: string;
  address: string;
  createdAt: number;
  updatedAt: number;
  blocked: boolean;
}

export interface MsgCreateAdministratorResponse {
}

export interface MsgUpdateAdministrator {
  creator: string;
  address: string;
  createdAt: number;
  updatedAt: number;
  blocked: boolean;
}

export interface MsgUpdateAdministratorResponse {
}

export interface MsgDeleteAdministrator {
  creator: string;
  address: string;
}

export interface MsgDeleteAdministratorResponse {
}

function createBaseMsgCreateAdministrator(): MsgCreateAdministrator {
  return { creator: "", address: "", createdAt: 0, updatedAt: 0, blocked: false };
}

export const MsgCreateAdministrator = {
  encode(message: MsgCreateAdministrator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.createdAt !== 0) {
      writer.uint32(24).int32(message.createdAt);
    }
    if (message.updatedAt !== 0) {
      writer.uint32(32).int32(message.updatedAt);
    }
    if (message.blocked === true) {
      writer.uint32(40).bool(message.blocked);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAdministrator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAdministrator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.createdAt = reader.int32();
          break;
        case 4:
          message.updatedAt = reader.int32();
          break;
        case 5:
          message.blocked = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAdministrator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      updatedAt: isSet(object.updatedAt) ? Number(object.updatedAt) : 0,
      blocked: isSet(object.blocked) ? Boolean(object.blocked) : false,
    };
  },

  toJSON(message: MsgCreateAdministrator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.createdAt !== undefined && (obj.createdAt = Math.round(message.createdAt));
    message.updatedAt !== undefined && (obj.updatedAt = Math.round(message.updatedAt));
    message.blocked !== undefined && (obj.blocked = message.blocked);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAdministrator>, I>>(object: I): MsgCreateAdministrator {
    const message = createBaseMsgCreateAdministrator();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.createdAt = object.createdAt ?? 0;
    message.updatedAt = object.updatedAt ?? 0;
    message.blocked = object.blocked ?? false;
    return message;
  },
};

function createBaseMsgCreateAdministratorResponse(): MsgCreateAdministratorResponse {
  return {};
}

export const MsgCreateAdministratorResponse = {
  encode(_: MsgCreateAdministratorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAdministratorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAdministratorResponse();
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

  fromJSON(_: any): MsgCreateAdministratorResponse {
    return {};
  },

  toJSON(_: MsgCreateAdministratorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAdministratorResponse>, I>>(_: I): MsgCreateAdministratorResponse {
    const message = createBaseMsgCreateAdministratorResponse();
    return message;
  },
};

function createBaseMsgUpdateAdministrator(): MsgUpdateAdministrator {
  return { creator: "", address: "", createdAt: 0, updatedAt: 0, blocked: false };
}

export const MsgUpdateAdministrator = {
  encode(message: MsgUpdateAdministrator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.createdAt !== 0) {
      writer.uint32(24).int32(message.createdAt);
    }
    if (message.updatedAt !== 0) {
      writer.uint32(32).int32(message.updatedAt);
    }
    if (message.blocked === true) {
      writer.uint32(40).bool(message.blocked);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAdministrator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAdministrator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.createdAt = reader.int32();
          break;
        case 4:
          message.updatedAt = reader.int32();
          break;
        case 5:
          message.blocked = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateAdministrator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      updatedAt: isSet(object.updatedAt) ? Number(object.updatedAt) : 0,
      blocked: isSet(object.blocked) ? Boolean(object.blocked) : false,
    };
  },

  toJSON(message: MsgUpdateAdministrator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.createdAt !== undefined && (obj.createdAt = Math.round(message.createdAt));
    message.updatedAt !== undefined && (obj.updatedAt = Math.round(message.updatedAt));
    message.blocked !== undefined && (obj.blocked = message.blocked);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateAdministrator>, I>>(object: I): MsgUpdateAdministrator {
    const message = createBaseMsgUpdateAdministrator();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.createdAt = object.createdAt ?? 0;
    message.updatedAt = object.updatedAt ?? 0;
    message.blocked = object.blocked ?? false;
    return message;
  },
};

function createBaseMsgUpdateAdministratorResponse(): MsgUpdateAdministratorResponse {
  return {};
}

export const MsgUpdateAdministratorResponse = {
  encode(_: MsgUpdateAdministratorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAdministratorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAdministratorResponse();
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

  fromJSON(_: any): MsgUpdateAdministratorResponse {
    return {};
  },

  toJSON(_: MsgUpdateAdministratorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateAdministratorResponse>, I>>(_: I): MsgUpdateAdministratorResponse {
    const message = createBaseMsgUpdateAdministratorResponse();
    return message;
  },
};

function createBaseMsgDeleteAdministrator(): MsgDeleteAdministrator {
  return { creator: "", address: "" };
}

export const MsgDeleteAdministrator = {
  encode(message: MsgDeleteAdministrator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAdministrator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAdministrator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteAdministrator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgDeleteAdministrator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAdministrator>, I>>(object: I): MsgDeleteAdministrator {
    const message = createBaseMsgDeleteAdministrator();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgDeleteAdministratorResponse(): MsgDeleteAdministratorResponse {
  return {};
}

export const MsgDeleteAdministratorResponse = {
  encode(_: MsgDeleteAdministratorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAdministratorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAdministratorResponse();
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

  fromJSON(_: any): MsgDeleteAdministratorResponse {
    return {};
  },

  toJSON(_: MsgDeleteAdministratorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAdministratorResponse>, I>>(_: I): MsgDeleteAdministratorResponse {
    const message = createBaseMsgDeleteAdministratorResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateAdministrator(request: MsgCreateAdministrator): Promise<MsgCreateAdministratorResponse>;
  UpdateAdministrator(request: MsgUpdateAdministrator): Promise<MsgUpdateAdministratorResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteAdministrator(request: MsgDeleteAdministrator): Promise<MsgDeleteAdministratorResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateAdministrator = this.CreateAdministrator.bind(this);
    this.UpdateAdministrator = this.UpdateAdministrator.bind(this);
    this.DeleteAdministrator = this.DeleteAdministrator.bind(this);
  }
  CreateAdministrator(request: MsgCreateAdministrator): Promise<MsgCreateAdministratorResponse> {
    const data = MsgCreateAdministrator.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.permission.Msg", "CreateAdministrator", data);
    return promise.then((data) => MsgCreateAdministratorResponse.decode(new _m0.Reader(data)));
  }

  UpdateAdministrator(request: MsgUpdateAdministrator): Promise<MsgUpdateAdministratorResponse> {
    const data = MsgUpdateAdministrator.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.permission.Msg", "UpdateAdministrator", data);
    return promise.then((data) => MsgUpdateAdministratorResponse.decode(new _m0.Reader(data)));
  }

  DeleteAdministrator(request: MsgDeleteAdministrator): Promise<MsgDeleteAdministratorResponse> {
    const data = MsgDeleteAdministrator.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.permission.Msg", "DeleteAdministrator", data);
    return promise.then((data) => MsgDeleteAdministratorResponse.decode(new _m0.Reader(data)));
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
