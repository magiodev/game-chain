/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "g4alentertainment.g4alchain.denomfactory";

export interface Denom {
  symbol: string;
  project: string;
  maxSupply: number;
  canChangeMaxSupply: boolean;
  creator: string;
}

function createBaseDenom(): Denom {
  return { symbol: "", project: "", maxSupply: 0, canChangeMaxSupply: false, creator: "" };
}

export const Denom = {
  encode(message: Denom, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    if (message.project !== "") {
      writer.uint32(18).string(message.project);
    }
    if (message.maxSupply !== 0) {
      writer.uint32(24).int32(message.maxSupply);
    }
    if (message.canChangeMaxSupply === true) {
      writer.uint32(32).bool(message.canChangeMaxSupply);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Denom {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDenom();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        case 2:
          message.project = reader.string();
          break;
        case 3:
          message.maxSupply = reader.int32();
          break;
        case 4:
          message.canChangeMaxSupply = reader.bool();
          break;
        case 5:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Denom {
    return {
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      project: isSet(object.project) ? String(object.project) : "",
      maxSupply: isSet(object.maxSupply) ? Number(object.maxSupply) : 0,
      canChangeMaxSupply: isSet(object.canChangeMaxSupply) ? Boolean(object.canChangeMaxSupply) : false,
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: Denom): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.project !== undefined && (obj.project = message.project);
    message.maxSupply !== undefined && (obj.maxSupply = Math.round(message.maxSupply));
    message.canChangeMaxSupply !== undefined && (obj.canChangeMaxSupply = message.canChangeMaxSupply);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Denom>, I>>(object: I): Denom {
    const message = createBaseDenom();
    message.symbol = object.symbol ?? "";
    message.project = object.project ?? "";
    message.maxSupply = object.maxSupply ?? 0;
    message.canChangeMaxSupply = object.canChangeMaxSupply ?? false;
    message.creator = object.creator ?? "";
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
