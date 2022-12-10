/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "g4alentertainment.g4alchain.denomfactory";

export interface MsgCreateDenom {
  creator: string;
  symbol: string;
  project: string;
  maxSupply: number;
  canChangeMaxSupply: boolean;
}

export interface MsgCreateDenomResponse {
}

export interface MsgUpdateDenom {
  creator: string;
  symbol: string;
  project: string;
  maxSupply: number;
  canChangeMaxSupply: boolean;
}

export interface MsgUpdateDenomResponse {
}

export interface MsgDeleteDenom {
  creator: string;
  symbol: string;
}

export interface MsgDeleteDenomResponse {
}

function createBaseMsgCreateDenom(): MsgCreateDenom {
  return { creator: "", symbol: "", project: "", maxSupply: 0, canChangeMaxSupply: false };
}

export const MsgCreateDenom = {
  encode(message: MsgCreateDenom, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDenom {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDenom();
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

  fromJSON(object: any): MsgCreateDenom {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      project: isSet(object.project) ? String(object.project) : "",
      maxSupply: isSet(object.maxSupply) ? Number(object.maxSupply) : 0,
      canChangeMaxSupply: isSet(object.canChangeMaxSupply) ? Boolean(object.canChangeMaxSupply) : false,
    };
  },

  toJSON(message: MsgCreateDenom): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.project !== undefined && (obj.project = message.project);
    message.maxSupply !== undefined && (obj.maxSupply = Math.round(message.maxSupply));
    message.canChangeMaxSupply !== undefined && (obj.canChangeMaxSupply = message.canChangeMaxSupply);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDenom>, I>>(object: I): MsgCreateDenom {
    const message = createBaseMsgCreateDenom();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.project = object.project ?? "";
    message.maxSupply = object.maxSupply ?? 0;
    message.canChangeMaxSupply = object.canChangeMaxSupply ?? false;
    return message;
  },
};

function createBaseMsgCreateDenomResponse(): MsgCreateDenomResponse {
  return {};
}

export const MsgCreateDenomResponse = {
  encode(_: MsgCreateDenomResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDenomResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDenomResponse();
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

  fromJSON(_: any): MsgCreateDenomResponse {
    return {};
  },

  toJSON(_: MsgCreateDenomResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDenomResponse>, I>>(_: I): MsgCreateDenomResponse {
    const message = createBaseMsgCreateDenomResponse();
    return message;
  },
};

function createBaseMsgUpdateDenom(): MsgUpdateDenom {
  return { creator: "", symbol: "", project: "", maxSupply: 0, canChangeMaxSupply: false };
}

export const MsgUpdateDenom = {
  encode(message: MsgUpdateDenom, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDenom {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDenom();
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

  fromJSON(object: any): MsgUpdateDenom {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      project: isSet(object.project) ? String(object.project) : "",
      maxSupply: isSet(object.maxSupply) ? Number(object.maxSupply) : 0,
      canChangeMaxSupply: isSet(object.canChangeMaxSupply) ? Boolean(object.canChangeMaxSupply) : false,
    };
  },

  toJSON(message: MsgUpdateDenom): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.project !== undefined && (obj.project = message.project);
    message.maxSupply !== undefined && (obj.maxSupply = Math.round(message.maxSupply));
    message.canChangeMaxSupply !== undefined && (obj.canChangeMaxSupply = message.canChangeMaxSupply);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDenom>, I>>(object: I): MsgUpdateDenom {
    const message = createBaseMsgUpdateDenom();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.project = object.project ?? "";
    message.maxSupply = object.maxSupply ?? 0;
    message.canChangeMaxSupply = object.canChangeMaxSupply ?? false;
    return message;
  },
};

function createBaseMsgUpdateDenomResponse(): MsgUpdateDenomResponse {
  return {};
}

export const MsgUpdateDenomResponse = {
  encode(_: MsgUpdateDenomResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDenomResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDenomResponse();
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

  fromJSON(_: any): MsgUpdateDenomResponse {
    return {};
  },

  toJSON(_: MsgUpdateDenomResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDenomResponse>, I>>(_: I): MsgUpdateDenomResponse {
    const message = createBaseMsgUpdateDenomResponse();
    return message;
  },
};

function createBaseMsgDeleteDenom(): MsgDeleteDenom {
  return { creator: "", symbol: "" };
}

export const MsgDeleteDenom = {
  encode(message: MsgDeleteDenom, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.symbol !== "") {
      writer.uint32(18).string(message.symbol);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDenom {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDenom();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.symbol = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDenom {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
    };
  },

  toJSON(message: MsgDeleteDenom): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDenom>, I>>(object: I): MsgDeleteDenom {
    const message = createBaseMsgDeleteDenom();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    return message;
  },
};

function createBaseMsgDeleteDenomResponse(): MsgDeleteDenomResponse {
  return {};
}

export const MsgDeleteDenomResponse = {
  encode(_: MsgDeleteDenomResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDenomResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDenomResponse();
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

  fromJSON(_: any): MsgDeleteDenomResponse {
    return {};
  },

  toJSON(_: MsgDeleteDenomResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDenomResponse>, I>>(_: I): MsgDeleteDenomResponse {
    const message = createBaseMsgDeleteDenomResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDenom(request: MsgCreateDenom): Promise<MsgCreateDenomResponse>;
  UpdateDenom(request: MsgUpdateDenom): Promise<MsgUpdateDenomResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteDenom(request: MsgDeleteDenom): Promise<MsgDeleteDenomResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateDenom = this.CreateDenom.bind(this);
    this.UpdateDenom = this.UpdateDenom.bind(this);
    this.DeleteDenom = this.DeleteDenom.bind(this);
  }
  CreateDenom(request: MsgCreateDenom): Promise<MsgCreateDenomResponse> {
    const data = MsgCreateDenom.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.denomfactory.Msg", "CreateDenom", data);
    return promise.then((data) => MsgCreateDenomResponse.decode(new _m0.Reader(data)));
  }

  UpdateDenom(request: MsgUpdateDenom): Promise<MsgUpdateDenomResponse> {
    const data = MsgUpdateDenom.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.denomfactory.Msg", "UpdateDenom", data);
    return promise.then((data) => MsgUpdateDenomResponse.decode(new _m0.Reader(data)));
  }

  DeleteDenom(request: MsgDeleteDenom): Promise<MsgDeleteDenomResponse> {
    const data = MsgDeleteDenom.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.denomfactory.Msg", "DeleteDenom", data);
    return promise.then((data) => MsgDeleteDenomResponse.decode(new _m0.Reader(data)));
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
