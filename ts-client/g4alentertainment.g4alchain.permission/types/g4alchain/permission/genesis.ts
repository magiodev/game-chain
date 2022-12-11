/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Administrator } from "./administrator";
import { Developer } from "./developer";
import { Params } from "./params";

export const protobufPackage = "g4alentertainment.g4alchain.permission";

/** GenesisState defines the permission module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  administratorList: Administrator[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  developerList: Developer[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, administratorList: [], developerList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.administratorList) {
      Administrator.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.developerList) {
      Developer.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.administratorList.push(Administrator.decode(reader, reader.uint32()));
          break;
        case 3:
          message.developerList.push(Developer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      administratorList: Array.isArray(object?.administratorList)
        ? object.administratorList.map((e: any) => Administrator.fromJSON(e))
        : [],
      developerList: Array.isArray(object?.developerList)
        ? object.developerList.map((e: any) => Developer.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.administratorList) {
      obj.administratorList = message.administratorList.map((e) => e ? Administrator.toJSON(e) : undefined);
    } else {
      obj.administratorList = [];
    }
    if (message.developerList) {
      obj.developerList = message.developerList.map((e) => e ? Developer.toJSON(e) : undefined);
    } else {
      obj.developerList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.administratorList = object.administratorList?.map((e) => Administrator.fromPartial(e)) || [];
    message.developerList = object.developerList?.map((e) => Developer.fromPartial(e)) || [];
    return message;
  },
};

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