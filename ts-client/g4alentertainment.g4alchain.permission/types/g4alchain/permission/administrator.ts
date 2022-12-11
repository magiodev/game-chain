/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "g4alentertainment.g4alchain.permission";

export interface Administrator {
  address: string;
  createdAt: number;
  updatedAt: number;
  blocked: boolean;
  creator: string;
}

function createBaseAdministrator(): Administrator {
  return { address: "", createdAt: 0, updatedAt: 0, blocked: false, creator: "" };
}

export const Administrator = {
  encode(message: Administrator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.createdAt !== 0) {
      writer.uint32(16).int32(message.createdAt);
    }
    if (message.updatedAt !== 0) {
      writer.uint32(24).int32(message.updatedAt);
    }
    if (message.blocked === true) {
      writer.uint32(32).bool(message.blocked);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Administrator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAdministrator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.createdAt = reader.int32();
          break;
        case 3:
          message.updatedAt = reader.int32();
          break;
        case 4:
          message.blocked = reader.bool();
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

  fromJSON(object: any): Administrator {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      updatedAt: isSet(object.updatedAt) ? Number(object.updatedAt) : 0,
      blocked: isSet(object.blocked) ? Boolean(object.blocked) : false,
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: Administrator): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.createdAt !== undefined && (obj.createdAt = Math.round(message.createdAt));
    message.updatedAt !== undefined && (obj.updatedAt = Math.round(message.updatedAt));
    message.blocked !== undefined && (obj.blocked = message.blocked);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Administrator>, I>>(object: I): Administrator {
    const message = createBaseAdministrator();
    message.address = object.address ?? "";
    message.createdAt = object.createdAt ?? 0;
    message.updatedAt = object.updatedAt ?? 0;
    message.blocked = object.blocked ?? false;
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
