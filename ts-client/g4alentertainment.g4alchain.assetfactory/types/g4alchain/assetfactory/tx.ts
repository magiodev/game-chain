/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "g4alentertainment.g4alchain.assetfactory";

export interface MsgCreateClass {
  creator: string;
  symbol: string;
  project: string;
  maxSupply: number;
  canChangeMaxSupply: boolean;
}

export interface MsgCreateClassResponse {
}

export interface MsgUpdateClass {
  creator: string;
  symbol: string;
  project: string;
  maxSupply: number;
  canChangeMaxSupply: boolean;
}

export interface MsgUpdateClassResponse {
}

function createBaseMsgCreateClass(): MsgCreateClass {
  return { creator: "", symbol: "", project: "", maxSupply: 0, canChangeMaxSupply: false };
}

export const MsgCreateClass = {
  encode(message: MsgCreateClass, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.symbol !== "") {
      writer.uint32(18).string(message.symbol);
    }
    if (message.project !== "") {
      writer.uint32(26).string(message.project);
    }
    if (message.maxSupply !== 0) {
      writer.uint32(32).int32(message.maxSupply);
    }
    if (message.canChangeMaxSupply === true) {
      writer.uint32(40).bool(message.canChangeMaxSupply);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateClass {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateClass();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.symbol = reader.string();
          break;
        case 3:
          message.project = reader.string();
          break;
        case 4:
          message.maxSupply = reader.int32();
          break;
        case 5:
          message.canChangeMaxSupply = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateClass {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      project: isSet(object.project) ? String(object.project) : "",
      maxSupply: isSet(object.maxSupply) ? Number(object.maxSupply) : 0,
      canChangeMaxSupply: isSet(object.canChangeMaxSupply) ? Boolean(object.canChangeMaxSupply) : false,
    };
  },

  toJSON(message: MsgCreateClass): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.project !== undefined && (obj.project = message.project);
    message.maxSupply !== undefined && (obj.maxSupply = Math.round(message.maxSupply));
    message.canChangeMaxSupply !== undefined && (obj.canChangeMaxSupply = message.canChangeMaxSupply);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateClass>, I>>(object: I): MsgCreateClass {
    const message = createBaseMsgCreateClass();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.project = object.project ?? "";
    message.maxSupply = object.maxSupply ?? 0;
    message.canChangeMaxSupply = object.canChangeMaxSupply ?? false;
    return message;
  },
};

function createBaseMsgCreateClassResponse(): MsgCreateClassResponse {
  return {};
}

export const MsgCreateClassResponse = {
  encode(_: MsgCreateClassResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateClassResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateClassResponse();
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

  fromJSON(_: any): MsgCreateClassResponse {
    return {};
  },

  toJSON(_: MsgCreateClassResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateClassResponse>, I>>(_: I): MsgCreateClassResponse {
    const message = createBaseMsgCreateClassResponse();
    return message;
  },
};

function createBaseMsgUpdateClass(): MsgUpdateClass {
  return { creator: "", symbol: "", project: "", maxSupply: 0, canChangeMaxSupply: false };
}

export const MsgUpdateClass = {
  encode(message: MsgUpdateClass, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.symbol !== "") {
      writer.uint32(18).string(message.symbol);
    }
    if (message.project !== "") {
      writer.uint32(26).string(message.project);
    }
    if (message.maxSupply !== 0) {
      writer.uint32(32).int32(message.maxSupply);
    }
    if (message.canChangeMaxSupply === true) {
      writer.uint32(40).bool(message.canChangeMaxSupply);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateClass {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateClass();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.symbol = reader.string();
          break;
        case 3:
          message.project = reader.string();
          break;
        case 4:
          message.maxSupply = reader.int32();
          break;
        case 5:
          message.canChangeMaxSupply = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateClass {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      project: isSet(object.project) ? String(object.project) : "",
      maxSupply: isSet(object.maxSupply) ? Number(object.maxSupply) : 0,
      canChangeMaxSupply: isSet(object.canChangeMaxSupply) ? Boolean(object.canChangeMaxSupply) : false,
    };
  },

  toJSON(message: MsgUpdateClass): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.project !== undefined && (obj.project = message.project);
    message.maxSupply !== undefined && (obj.maxSupply = Math.round(message.maxSupply));
    message.canChangeMaxSupply !== undefined && (obj.canChangeMaxSupply = message.canChangeMaxSupply);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateClass>, I>>(object: I): MsgUpdateClass {
    const message = createBaseMsgUpdateClass();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.project = object.project ?? "";
    message.maxSupply = object.maxSupply ?? 0;
    message.canChangeMaxSupply = object.canChangeMaxSupply ?? false;
    return message;
  },
};

function createBaseMsgUpdateClassResponse(): MsgUpdateClassResponse {
  return {};
}

export const MsgUpdateClassResponse = {
  encode(_: MsgUpdateClassResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateClassResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateClassResponse();
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

  fromJSON(_: any): MsgUpdateClassResponse {
    return {};
  },

  toJSON(_: MsgUpdateClassResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateClassResponse>, I>>(_: I): MsgUpdateClassResponse {
    const message = createBaseMsgUpdateClassResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateClass(request: MsgCreateClass): Promise<MsgCreateClassResponse>;
  UpdateClass(request: MsgUpdateClass): Promise<MsgUpdateClassResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateClass = this.CreateClass.bind(this);
    this.UpdateClass = this.UpdateClass.bind(this);
  }
  CreateClass(request: MsgCreateClass): Promise<MsgCreateClassResponse> {
    const data = MsgCreateClass.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.assetfactory.Msg", "CreateClass", data);
    return promise.then((data) => MsgCreateClassResponse.decode(new _m0.Reader(data)));
  }

  UpdateClass(request: MsgUpdateClass): Promise<MsgUpdateClassResponse> {
    const data = MsgUpdateClass.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.assetfactory.Msg", "UpdateClass", data);
    return promise.then((data) => MsgUpdateClassResponse.decode(new _m0.Reader(data)));
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
