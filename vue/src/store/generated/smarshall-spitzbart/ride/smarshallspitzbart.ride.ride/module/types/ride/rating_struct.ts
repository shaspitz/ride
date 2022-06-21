/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "smarshallspitzbart.ride.ride";

export interface RatingStruct {
  index: string;
  rating: string;
}

const baseRatingStruct: object = { index: "", rating: "" };

export const RatingStruct = {
  encode(message: RatingStruct, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.rating !== "") {
      writer.uint32(18).string(message.rating);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RatingStruct {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRatingStruct } as RatingStruct;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.rating = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RatingStruct {
    const message = { ...baseRatingStruct } as RatingStruct;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.rating !== undefined && object.rating !== null) {
      message.rating = String(object.rating);
    } else {
      message.rating = "";
    }
    return message;
  },

  toJSON(message: RatingStruct): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.rating !== undefined && (obj.rating = message.rating);
    return obj;
  },

  fromPartial(object: DeepPartial<RatingStruct>): RatingStruct {
    const message = { ...baseRatingStruct } as RatingStruct;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.rating !== undefined && object.rating !== null) {
      message.rating = object.rating;
    } else {
      message.rating = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
