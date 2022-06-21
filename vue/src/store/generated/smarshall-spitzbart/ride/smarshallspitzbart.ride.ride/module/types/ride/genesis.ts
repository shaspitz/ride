/* eslint-disable */
import { Params } from "../ride/params";
import { NextRide } from "../ride/next_ride";
import { StoredRide } from "../ride/stored_ride";
import { RatingStruct } from "../ride/rating_struct";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "smarshallspitzbart.ride.ride";

/** GenesisState defines the ride module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  nextRide: NextRide | undefined;
  storedRideList: StoredRide[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  ratingStructList: RatingStruct[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextRide !== undefined) {
      NextRide.encode(message.nextRide, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.storedRideList) {
      StoredRide.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.ratingStructList) {
      RatingStruct.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.storedRideList = [];
    message.ratingStructList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.nextRide = NextRide.decode(reader, reader.uint32());
          break;
        case 3:
          message.storedRideList.push(
            StoredRide.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.ratingStructList.push(
            RatingStruct.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedRideList = [];
    message.ratingStructList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nextRide !== undefined && object.nextRide !== null) {
      message.nextRide = NextRide.fromJSON(object.nextRide);
    } else {
      message.nextRide = undefined;
    }
    if (object.storedRideList !== undefined && object.storedRideList !== null) {
      for (const e of object.storedRideList) {
        message.storedRideList.push(StoredRide.fromJSON(e));
      }
    }
    if (
      object.ratingStructList !== undefined &&
      object.ratingStructList !== null
    ) {
      for (const e of object.ratingStructList) {
        message.ratingStructList.push(RatingStruct.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.nextRide !== undefined &&
      (obj.nextRide = message.nextRide
        ? NextRide.toJSON(message.nextRide)
        : undefined);
    if (message.storedRideList) {
      obj.storedRideList = message.storedRideList.map((e) =>
        e ? StoredRide.toJSON(e) : undefined
      );
    } else {
      obj.storedRideList = [];
    }
    if (message.ratingStructList) {
      obj.ratingStructList = message.ratingStructList.map((e) =>
        e ? RatingStruct.toJSON(e) : undefined
      );
    } else {
      obj.ratingStructList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedRideList = [];
    message.ratingStructList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nextRide !== undefined && object.nextRide !== null) {
      message.nextRide = NextRide.fromPartial(object.nextRide);
    } else {
      message.nextRide = undefined;
    }
    if (object.storedRideList !== undefined && object.storedRideList !== null) {
      for (const e of object.storedRideList) {
        message.storedRideList.push(StoredRide.fromPartial(e));
      }
    }
    if (
      object.ratingStructList !== undefined &&
      object.ratingStructList !== null
    ) {
      for (const e of object.ratingStructList) {
        message.ratingStructList.push(RatingStruct.fromPartial(e));
      }
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
