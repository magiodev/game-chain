/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "g4alentertainment.g4alchain.game";

export interface MsgCreateProject {
  creator: string;
  symbol: string;
  name: string;
  description: string;
  delegate: string[];
}

export interface MsgCreateProjectResponse {
}

export interface MsgUpdateProject {
  creator: string;
  symbol: string;
  name: string;
  description: string;
  delegate: string[];
}

export interface MsgUpdateProjectResponse {
}

function createBaseMsgCreateProject(): MsgCreateProject {
  return { creator: "", symbol: "", name: "", description: "", delegate: [] };
}

export const MsgCreateProject = {
  encode(message: MsgCreateProject, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.symbol !== "") {
      writer.uint32(18).string(message.symbol);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    for (const v of message.delegate) {
      writer.uint32(42).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateProject {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateProject();
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
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.delegate.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateProject {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      delegate: Array.isArray(object?.delegate) ? object.delegate.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgCreateProject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    if (message.delegate) {
      obj.delegate = message.delegate.map((e) => e);
    } else {
      obj.delegate = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateProject>, I>>(object: I): MsgCreateProject {
    const message = createBaseMsgCreateProject();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.delegate = object.delegate?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgCreateProjectResponse(): MsgCreateProjectResponse {
  return {};
}

export const MsgCreateProjectResponse = {
  encode(_: MsgCreateProjectResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateProjectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateProjectResponse();
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

  fromJSON(_: any): MsgCreateProjectResponse {
    return {};
  },

  toJSON(_: MsgCreateProjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateProjectResponse>, I>>(_: I): MsgCreateProjectResponse {
    const message = createBaseMsgCreateProjectResponse();
    return message;
  },
};

function createBaseMsgUpdateProject(): MsgUpdateProject {
  return { creator: "", symbol: "", name: "", description: "", delegate: [] };
}

export const MsgUpdateProject = {
  encode(message: MsgUpdateProject, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.symbol !== "") {
      writer.uint32(18).string(message.symbol);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    for (const v of message.delegate) {
      writer.uint32(42).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateProject {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateProject();
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
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.delegate.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateProject {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      delegate: Array.isArray(object?.delegate) ? object.delegate.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgUpdateProject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    if (message.delegate) {
      obj.delegate = message.delegate.map((e) => e);
    } else {
      obj.delegate = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateProject>, I>>(object: I): MsgUpdateProject {
    const message = createBaseMsgUpdateProject();
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.delegate = object.delegate?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgUpdateProjectResponse(): MsgUpdateProjectResponse {
  return {};
}

export const MsgUpdateProjectResponse = {
  encode(_: MsgUpdateProjectResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateProjectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateProjectResponse();
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

  fromJSON(_: any): MsgUpdateProjectResponse {
    return {};
  },

  toJSON(_: MsgUpdateProjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateProjectResponse>, I>>(_: I): MsgUpdateProjectResponse {
    const message = createBaseMsgUpdateProjectResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateProject(request: MsgCreateProject): Promise<MsgCreateProjectResponse>;
  UpdateProject(request: MsgUpdateProject): Promise<MsgUpdateProjectResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateProject = this.CreateProject.bind(this);
    this.UpdateProject = this.UpdateProject.bind(this);
  }
  CreateProject(request: MsgCreateProject): Promise<MsgCreateProjectResponse> {
    const data = MsgCreateProject.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.game.Msg", "CreateProject", data);
    return promise.then((data) => MsgCreateProjectResponse.decode(new _m0.Reader(data)));
  }

  UpdateProject(request: MsgUpdateProject): Promise<MsgUpdateProjectResponse> {
    const data = MsgUpdateProject.encode(request).finish();
    const promise = this.rpc.request("g4alentertainment.g4alchain.game.Msg", "UpdateProject", data);
    return promise.then((data) => MsgUpdateProjectResponse.decode(new _m0.Reader(data)));
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
