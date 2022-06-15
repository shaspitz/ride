/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "smarshallspitzbart.ride.ride";

export interface StoredRide {
  index: string;
  destination: string;
  driver: string;
  passenger: string;
  acceptanceTime: string;
  finishTime: string;
  finishLocation: string;
  mutualStake: number;
  payPerHour: number;
  distanceTip: number;
  beforeId: string;
  afterId: string;
  deadline: string;
}

const baseStoredRide: object = {
  index: "",
  destination: "",
  driver: "",
  passenger: "",
  acceptanceTime: "",
  finishTime: "",
  finishLocation: "",
  mutualStake: 0,
  payPerHour: 0,
  distanceTip: 0,
  beforeId: "",
  afterId: "",
  deadline: "",
};

export const StoredRide = {
  encode(message: StoredRide, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.destination !== "") {
      writer.uint32(18).string(message.destination);
    }
    if (message.driver !== "") {
      writer.uint32(26).string(message.driver);
    }
    if (message.passenger !== "") {
      writer.uint32(34).string(message.passenger);
    }
    if (message.acceptanceTime !== "") {
      writer.uint32(42).string(message.acceptanceTime);
    }
    if (message.finishTime !== "") {
      writer.uint32(50).string(message.finishTime);
    }
    if (message.finishLocation !== "") {
      writer.uint32(58).string(message.finishLocation);
    }
    if (message.mutualStake !== 0) {
      writer.uint32(64).uint64(message.mutualStake);
    }
    if (message.payPerHour !== 0) {
      writer.uint32(72).uint64(message.payPerHour);
    }
    if (message.distanceTip !== 0) {
      writer.uint32(80).uint64(message.distanceTip);
    }
    if (message.beforeId !== "") {
      writer.uint32(90).string(message.beforeId);
    }
    if (message.afterId !== "") {
      writer.uint32(98).string(message.afterId);
    }
    if (message.deadline !== "") {
      writer.uint32(106).string(message.deadline);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredRide {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredRide } as StoredRide;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.destination = reader.string();
          break;
        case 3:
          message.driver = reader.string();
          break;
        case 4:
          message.passenger = reader.string();
          break;
        case 5:
          message.acceptanceTime = reader.string();
          break;
        case 6:
          message.finishTime = reader.string();
          break;
        case 7:
          message.finishLocation = reader.string();
          break;
        case 8:
          message.mutualStake = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.payPerHour = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.distanceTip = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.beforeId = reader.string();
          break;
        case 12:
          message.afterId = reader.string();
          break;
        case 13:
          message.deadline = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredRide {
    const message = { ...baseStoredRide } as StoredRide;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = String(object.destination);
    } else {
      message.destination = "";
    }
    if (object.driver !== undefined && object.driver !== null) {
      message.driver = String(object.driver);
    } else {
      message.driver = "";
    }
    if (object.passenger !== undefined && object.passenger !== null) {
      message.passenger = String(object.passenger);
    } else {
      message.passenger = "";
    }
    if (object.acceptanceTime !== undefined && object.acceptanceTime !== null) {
      message.acceptanceTime = String(object.acceptanceTime);
    } else {
      message.acceptanceTime = "";
    }
    if (object.finishTime !== undefined && object.finishTime !== null) {
      message.finishTime = String(object.finishTime);
    } else {
      message.finishTime = "";
    }
    if (object.finishLocation !== undefined && object.finishLocation !== null) {
      message.finishLocation = String(object.finishLocation);
    } else {
      message.finishLocation = "";
    }
    if (object.mutualStake !== undefined && object.mutualStake !== null) {
      message.mutualStake = Number(object.mutualStake);
    } else {
      message.mutualStake = 0;
    }
    if (object.payPerHour !== undefined && object.payPerHour !== null) {
      message.payPerHour = Number(object.payPerHour);
    } else {
      message.payPerHour = 0;
    }
    if (object.distanceTip !== undefined && object.distanceTip !== null) {
      message.distanceTip = Number(object.distanceTip);
    } else {
      message.distanceTip = 0;
    }
    if (object.beforeId !== undefined && object.beforeId !== null) {
      message.beforeId = String(object.beforeId);
    } else {
      message.beforeId = "";
    }
    if (object.afterId !== undefined && object.afterId !== null) {
      message.afterId = String(object.afterId);
    } else {
      message.afterId = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = String(object.deadline);
    } else {
      message.deadline = "";
    }
    return message;
  },

  toJSON(message: StoredRide): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.destination !== undefined &&
      (obj.destination = message.destination);
    message.driver !== undefined && (obj.driver = message.driver);
    message.passenger !== undefined && (obj.passenger = message.passenger);
    message.acceptanceTime !== undefined &&
      (obj.acceptanceTime = message.acceptanceTime);
    message.finishTime !== undefined && (obj.finishTime = message.finishTime);
    message.finishLocation !== undefined &&
      (obj.finishLocation = message.finishLocation);
    message.mutualStake !== undefined &&
      (obj.mutualStake = message.mutualStake);
    message.payPerHour !== undefined && (obj.payPerHour = message.payPerHour);
    message.distanceTip !== undefined &&
      (obj.distanceTip = message.distanceTip);
    message.beforeId !== undefined && (obj.beforeId = message.beforeId);
    message.afterId !== undefined && (obj.afterId = message.afterId);
    message.deadline !== undefined && (obj.deadline = message.deadline);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredRide>): StoredRide {
    const message = { ...baseStoredRide } as StoredRide;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = object.destination;
    } else {
      message.destination = "";
    }
    if (object.driver !== undefined && object.driver !== null) {
      message.driver = object.driver;
    } else {
      message.driver = "";
    }
    if (object.passenger !== undefined && object.passenger !== null) {
      message.passenger = object.passenger;
    } else {
      message.passenger = "";
    }
    if (object.acceptanceTime !== undefined && object.acceptanceTime !== null) {
      message.acceptanceTime = object.acceptanceTime;
    } else {
      message.acceptanceTime = "";
    }
    if (object.finishTime !== undefined && object.finishTime !== null) {
      message.finishTime = object.finishTime;
    } else {
      message.finishTime = "";
    }
    if (object.finishLocation !== undefined && object.finishLocation !== null) {
      message.finishLocation = object.finishLocation;
    } else {
      message.finishLocation = "";
    }
    if (object.mutualStake !== undefined && object.mutualStake !== null) {
      message.mutualStake = object.mutualStake;
    } else {
      message.mutualStake = 0;
    }
    if (object.payPerHour !== undefined && object.payPerHour !== null) {
      message.payPerHour = object.payPerHour;
    } else {
      message.payPerHour = 0;
    }
    if (object.distanceTip !== undefined && object.distanceTip !== null) {
      message.distanceTip = object.distanceTip;
    } else {
      message.distanceTip = 0;
    }
    if (object.beforeId !== undefined && object.beforeId !== null) {
      message.beforeId = object.beforeId;
    } else {
      message.beforeId = "";
    }
    if (object.afterId !== undefined && object.afterId !== null) {
      message.afterId = object.afterId;
    } else {
      message.afterId = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = object.deadline;
    } else {
      message.deadline = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
