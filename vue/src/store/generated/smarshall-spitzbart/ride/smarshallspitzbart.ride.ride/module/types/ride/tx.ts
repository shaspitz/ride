/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "smarshallspitzbart.ride.ride";

export interface MsgRequestRide {
  creator: string;
  startLocation: string;
  destination: string;
  mutualStake: number;
  hourlyPay: number;
  distanceTip: number;
}

export interface MsgRequestRideResponse {
  idValue: string;
}

export interface MsgAccept {
  creator: string;
  idValue: string;
}

export interface MsgAcceptResponse {
  success: boolean;
}

const baseMsgRequestRide: object = {
  creator: "",
  startLocation: "",
  destination: "",
  mutualStake: 0,
  hourlyPay: 0,
  distanceTip: 0,
};

export const MsgRequestRide = {
  encode(message: MsgRequestRide, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.startLocation !== "") {
      writer.uint32(18).string(message.startLocation);
    }
    if (message.destination !== "") {
      writer.uint32(26).string(message.destination);
    }
    if (message.mutualStake !== 0) {
      writer.uint32(32).uint64(message.mutualStake);
    }
    if (message.hourlyPay !== 0) {
      writer.uint32(40).uint64(message.hourlyPay);
    }
    if (message.distanceTip !== 0) {
      writer.uint32(48).uint64(message.distanceTip);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestRide {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestRide } as MsgRequestRide;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.startLocation = reader.string();
          break;
        case 3:
          message.destination = reader.string();
          break;
        case 4:
          message.mutualStake = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.hourlyPay = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.distanceTip = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestRide {
    const message = { ...baseMsgRequestRide } as MsgRequestRide;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.startLocation !== undefined && object.startLocation !== null) {
      message.startLocation = String(object.startLocation);
    } else {
      message.startLocation = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = String(object.destination);
    } else {
      message.destination = "";
    }
    if (object.mutualStake !== undefined && object.mutualStake !== null) {
      message.mutualStake = Number(object.mutualStake);
    } else {
      message.mutualStake = 0;
    }
    if (object.hourlyPay !== undefined && object.hourlyPay !== null) {
      message.hourlyPay = Number(object.hourlyPay);
    } else {
      message.hourlyPay = 0;
    }
    if (object.distanceTip !== undefined && object.distanceTip !== null) {
      message.distanceTip = Number(object.distanceTip);
    } else {
      message.distanceTip = 0;
    }
    return message;
  },

  toJSON(message: MsgRequestRide): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.startLocation !== undefined &&
      (obj.startLocation = message.startLocation);
    message.destination !== undefined &&
      (obj.destination = message.destination);
    message.mutualStake !== undefined &&
      (obj.mutualStake = message.mutualStake);
    message.hourlyPay !== undefined && (obj.hourlyPay = message.hourlyPay);
    message.distanceTip !== undefined &&
      (obj.distanceTip = message.distanceTip);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRequestRide>): MsgRequestRide {
    const message = { ...baseMsgRequestRide } as MsgRequestRide;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.startLocation !== undefined && object.startLocation !== null) {
      message.startLocation = object.startLocation;
    } else {
      message.startLocation = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = object.destination;
    } else {
      message.destination = "";
    }
    if (object.mutualStake !== undefined && object.mutualStake !== null) {
      message.mutualStake = object.mutualStake;
    } else {
      message.mutualStake = 0;
    }
    if (object.hourlyPay !== undefined && object.hourlyPay !== null) {
      message.hourlyPay = object.hourlyPay;
    } else {
      message.hourlyPay = 0;
    }
    if (object.distanceTip !== undefined && object.distanceTip !== null) {
      message.distanceTip = object.distanceTip;
    } else {
      message.distanceTip = 0;
    }
    return message;
  },
};

const baseMsgRequestRideResponse: object = { idValue: "" };

export const MsgRequestRideResponse = {
  encode(
    message: MsgRequestRideResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.idValue !== "") {
      writer.uint32(10).string(message.idValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestRideResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestRideResponse } as MsgRequestRideResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestRideResponse {
    const message = { ...baseMsgRequestRideResponse } as MsgRequestRideResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    return message;
  },

  toJSON(message: MsgRequestRideResponse): unknown {
    const obj: any = {};
    message.idValue !== undefined && (obj.idValue = message.idValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRequestRideResponse>
  ): MsgRequestRideResponse {
    const message = { ...baseMsgRequestRideResponse } as MsgRequestRideResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    return message;
  },
};

const baseMsgAccept: object = { creator: "", idValue: "" };

export const MsgAccept = {
  encode(message: MsgAccept, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.idValue !== "") {
      writer.uint32(18).string(message.idValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAccept {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAccept } as MsgAccept;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.idValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAccept {
    const message = { ...baseMsgAccept } as MsgAccept;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    return message;
  },

  toJSON(message: MsgAccept): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.idValue !== undefined && (obj.idValue = message.idValue);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAccept>): MsgAccept {
    const message = { ...baseMsgAccept } as MsgAccept;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    return message;
  },
};

const baseMsgAcceptResponse: object = { success: false };

export const MsgAcceptResponse = {
  encode(message: MsgAcceptResponse, writer: Writer = Writer.create()): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAcceptResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAcceptResponse } as MsgAcceptResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAcceptResponse {
    const message = { ...baseMsgAcceptResponse } as MsgAcceptResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgAcceptResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAcceptResponse>): MsgAcceptResponse {
    const message = { ...baseMsgAcceptResponse } as MsgAcceptResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RequestRide(request: MsgRequestRide): Promise<MsgRequestRideResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Accept(request: MsgAccept): Promise<MsgAcceptResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RequestRide(request: MsgRequestRide): Promise<MsgRequestRideResponse> {
    const data = MsgRequestRide.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Msg",
      "RequestRide",
      data
    );
    return promise.then((data) =>
      MsgRequestRideResponse.decode(new Reader(data))
    );
  }

  Accept(request: MsgAccept): Promise<MsgAcceptResponse> {
    const data = MsgAccept.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Msg",
      "Accept",
      data
    );
    return promise.then((data) => MsgAcceptResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
