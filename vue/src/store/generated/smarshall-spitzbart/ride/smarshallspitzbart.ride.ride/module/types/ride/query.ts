/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../ride/params";
import { NextRide } from "../ride/next_ride";
import { StoredRide } from "../ride/stored_ride";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { RatingStruct } from "../ride/rating_struct";

export const protobufPackage = "smarshallspitzbart.ride.ride";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNextRideRequest {}

export interface QueryGetNextRideResponse {
  NextRide: NextRide | undefined;
}

export interface QueryGetStoredRideRequest {
  index: string;
}

export interface QueryGetStoredRideResponse {
  storedRide: StoredRide | undefined;
}

export interface QueryAllStoredRideRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredRideResponse {
  storedRide: StoredRide[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRatingStructRequest {
  index: string;
}

export interface QueryGetRatingStructResponse {
  ratingStruct: RatingStruct | undefined;
}

export interface QueryAllRatingStructRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRatingStructResponse {
  ratingStruct: RatingStruct[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetNextRideRequest: object = {};

export const QueryGetNextRideRequest = {
  encode(_: QueryGetNextRideRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNextRideRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNextRideRequest,
    } as QueryGetNextRideRequest;
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

  fromJSON(_: any): QueryGetNextRideRequest {
    const message = {
      ...baseQueryGetNextRideRequest,
    } as QueryGetNextRideRequest;
    return message;
  },

  toJSON(_: QueryGetNextRideRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetNextRideRequest>
  ): QueryGetNextRideRequest {
    const message = {
      ...baseQueryGetNextRideRequest,
    } as QueryGetNextRideRequest;
    return message;
  },
};

const baseQueryGetNextRideResponse: object = {};

export const QueryGetNextRideResponse = {
  encode(
    message: QueryGetNextRideResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.NextRide !== undefined) {
      NextRide.encode(message.NextRide, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNextRideResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNextRideResponse,
    } as QueryGetNextRideResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.NextRide = NextRide.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNextRideResponse {
    const message = {
      ...baseQueryGetNextRideResponse,
    } as QueryGetNextRideResponse;
    if (object.NextRide !== undefined && object.NextRide !== null) {
      message.NextRide = NextRide.fromJSON(object.NextRide);
    } else {
      message.NextRide = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNextRideResponse): unknown {
    const obj: any = {};
    message.NextRide !== undefined &&
      (obj.NextRide = message.NextRide
        ? NextRide.toJSON(message.NextRide)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNextRideResponse>
  ): QueryGetNextRideResponse {
    const message = {
      ...baseQueryGetNextRideResponse,
    } as QueryGetNextRideResponse;
    if (object.NextRide !== undefined && object.NextRide !== null) {
      message.NextRide = NextRide.fromPartial(object.NextRide);
    } else {
      message.NextRide = undefined;
    }
    return message;
  },
};

const baseQueryGetStoredRideRequest: object = { index: "" };

export const QueryGetStoredRideRequest = {
  encode(
    message: QueryGetStoredRideRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredRideRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredRideRequest,
    } as QueryGetStoredRideRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredRideRequest {
    const message = {
      ...baseQueryGetStoredRideRequest,
    } as QueryGetStoredRideRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetStoredRideRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredRideRequest>
  ): QueryGetStoredRideRequest {
    const message = {
      ...baseQueryGetStoredRideRequest,
    } as QueryGetStoredRideRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetStoredRideResponse: object = {};

export const QueryGetStoredRideResponse = {
  encode(
    message: QueryGetStoredRideResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storedRide !== undefined) {
      StoredRide.encode(message.storedRide, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredRideResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredRideResponse,
    } as QueryGetStoredRideResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedRide = StoredRide.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredRideResponse {
    const message = {
      ...baseQueryGetStoredRideResponse,
    } as QueryGetStoredRideResponse;
    if (object.storedRide !== undefined && object.storedRide !== null) {
      message.storedRide = StoredRide.fromJSON(object.storedRide);
    } else {
      message.storedRide = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStoredRideResponse): unknown {
    const obj: any = {};
    message.storedRide !== undefined &&
      (obj.storedRide = message.storedRide
        ? StoredRide.toJSON(message.storedRide)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredRideResponse>
  ): QueryGetStoredRideResponse {
    const message = {
      ...baseQueryGetStoredRideResponse,
    } as QueryGetStoredRideResponse;
    if (object.storedRide !== undefined && object.storedRide !== null) {
      message.storedRide = StoredRide.fromPartial(object.storedRide);
    } else {
      message.storedRide = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredRideRequest: object = {};

export const QueryAllStoredRideRequest = {
  encode(
    message: QueryAllStoredRideRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredRideRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredRideRequest,
    } as QueryAllStoredRideRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredRideRequest {
    const message = {
      ...baseQueryAllStoredRideRequest,
    } as QueryAllStoredRideRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredRideRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredRideRequest>
  ): QueryAllStoredRideRequest {
    const message = {
      ...baseQueryAllStoredRideRequest,
    } as QueryAllStoredRideRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredRideResponse: object = {};

export const QueryAllStoredRideResponse = {
  encode(
    message: QueryAllStoredRideResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.storedRide) {
      StoredRide.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredRideResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredRideResponse,
    } as QueryAllStoredRideResponse;
    message.storedRide = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedRide.push(StoredRide.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredRideResponse {
    const message = {
      ...baseQueryAllStoredRideResponse,
    } as QueryAllStoredRideResponse;
    message.storedRide = [];
    if (object.storedRide !== undefined && object.storedRide !== null) {
      for (const e of object.storedRide) {
        message.storedRide.push(StoredRide.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredRideResponse): unknown {
    const obj: any = {};
    if (message.storedRide) {
      obj.storedRide = message.storedRide.map((e) =>
        e ? StoredRide.toJSON(e) : undefined
      );
    } else {
      obj.storedRide = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredRideResponse>
  ): QueryAllStoredRideResponse {
    const message = {
      ...baseQueryAllStoredRideResponse,
    } as QueryAllStoredRideResponse;
    message.storedRide = [];
    if (object.storedRide !== undefined && object.storedRide !== null) {
      for (const e of object.storedRide) {
        message.storedRide.push(StoredRide.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetRatingStructRequest: object = { index: "" };

export const QueryGetRatingStructRequest = {
  encode(
    message: QueryGetRatingStructRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRatingStructRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRatingStructRequest,
    } as QueryGetRatingStructRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRatingStructRequest {
    const message = {
      ...baseQueryGetRatingStructRequest,
    } as QueryGetRatingStructRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetRatingStructRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRatingStructRequest>
  ): QueryGetRatingStructRequest {
    const message = {
      ...baseQueryGetRatingStructRequest,
    } as QueryGetRatingStructRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetRatingStructResponse: object = {};

export const QueryGetRatingStructResponse = {
  encode(
    message: QueryGetRatingStructResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ratingStruct !== undefined) {
      RatingStruct.encode(
        message.ratingStruct,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRatingStructResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRatingStructResponse,
    } as QueryGetRatingStructResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ratingStruct = RatingStruct.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRatingStructResponse {
    const message = {
      ...baseQueryGetRatingStructResponse,
    } as QueryGetRatingStructResponse;
    if (object.ratingStruct !== undefined && object.ratingStruct !== null) {
      message.ratingStruct = RatingStruct.fromJSON(object.ratingStruct);
    } else {
      message.ratingStruct = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRatingStructResponse): unknown {
    const obj: any = {};
    message.ratingStruct !== undefined &&
      (obj.ratingStruct = message.ratingStruct
        ? RatingStruct.toJSON(message.ratingStruct)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRatingStructResponse>
  ): QueryGetRatingStructResponse {
    const message = {
      ...baseQueryGetRatingStructResponse,
    } as QueryGetRatingStructResponse;
    if (object.ratingStruct !== undefined && object.ratingStruct !== null) {
      message.ratingStruct = RatingStruct.fromPartial(object.ratingStruct);
    } else {
      message.ratingStruct = undefined;
    }
    return message;
  },
};

const baseQueryAllRatingStructRequest: object = {};

export const QueryAllRatingStructRequest = {
  encode(
    message: QueryAllRatingStructRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRatingStructRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRatingStructRequest,
    } as QueryAllRatingStructRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRatingStructRequest {
    const message = {
      ...baseQueryAllRatingStructRequest,
    } as QueryAllRatingStructRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRatingStructRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRatingStructRequest>
  ): QueryAllRatingStructRequest {
    const message = {
      ...baseQueryAllRatingStructRequest,
    } as QueryAllRatingStructRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRatingStructResponse: object = {};

export const QueryAllRatingStructResponse = {
  encode(
    message: QueryAllRatingStructResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.ratingStruct) {
      RatingStruct.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRatingStructResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRatingStructResponse,
    } as QueryAllRatingStructResponse;
    message.ratingStruct = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ratingStruct.push(
            RatingStruct.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRatingStructResponse {
    const message = {
      ...baseQueryAllRatingStructResponse,
    } as QueryAllRatingStructResponse;
    message.ratingStruct = [];
    if (object.ratingStruct !== undefined && object.ratingStruct !== null) {
      for (const e of object.ratingStruct) {
        message.ratingStruct.push(RatingStruct.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRatingStructResponse): unknown {
    const obj: any = {};
    if (message.ratingStruct) {
      obj.ratingStruct = message.ratingStruct.map((e) =>
        e ? RatingStruct.toJSON(e) : undefined
      );
    } else {
      obj.ratingStruct = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRatingStructResponse>
  ): QueryAllRatingStructResponse {
    const message = {
      ...baseQueryAllRatingStructResponse,
    } as QueryAllRatingStructResponse;
    message.ratingStruct = [];
    if (object.ratingStruct !== undefined && object.ratingStruct !== null) {
      for (const e of object.ratingStruct) {
        message.ratingStruct.push(RatingStruct.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a NextRide by index. */
  NextRide(request: QueryGetNextRideRequest): Promise<QueryGetNextRideResponse>;
  /** Queries a StoredRide by index. */
  StoredRide(
    request: QueryGetStoredRideRequest
  ): Promise<QueryGetStoredRideResponse>;
  /** Queries a list of StoredRide items. */
  StoredRideAll(
    request: QueryAllStoredRideRequest
  ): Promise<QueryAllStoredRideResponse>;
  /** Queries a RatingStruct by index. */
  RatingStruct(
    request: QueryGetRatingStructRequest
  ): Promise<QueryGetRatingStructResponse>;
  /** Queries a list of RatingStruct items. */
  RatingStructAll(
    request: QueryAllRatingStructRequest
  ): Promise<QueryAllRatingStructResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  NextRide(
    request: QueryGetNextRideRequest
  ): Promise<QueryGetNextRideResponse> {
    const data = QueryGetNextRideRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Query",
      "NextRide",
      data
    );
    return promise.then((data) =>
      QueryGetNextRideResponse.decode(new Reader(data))
    );
  }

  StoredRide(
    request: QueryGetStoredRideRequest
  ): Promise<QueryGetStoredRideResponse> {
    const data = QueryGetStoredRideRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Query",
      "StoredRide",
      data
    );
    return promise.then((data) =>
      QueryGetStoredRideResponse.decode(new Reader(data))
    );
  }

  StoredRideAll(
    request: QueryAllStoredRideRequest
  ): Promise<QueryAllStoredRideResponse> {
    const data = QueryAllStoredRideRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Query",
      "StoredRideAll",
      data
    );
    return promise.then((data) =>
      QueryAllStoredRideResponse.decode(new Reader(data))
    );
  }

  RatingStruct(
    request: QueryGetRatingStructRequest
  ): Promise<QueryGetRatingStructResponse> {
    const data = QueryGetRatingStructRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Query",
      "RatingStruct",
      data
    );
    return promise.then((data) =>
      QueryGetRatingStructResponse.decode(new Reader(data))
    );
  }

  RatingStructAll(
    request: QueryAllRatingStructRequest
  ): Promise<QueryAllRatingStructResponse> {
    const data = QueryAllRatingStructRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smarshallspitzbart.ride.ride.Query",
      "RatingStructAll",
      data
    );
    return promise.then((data) =>
      QueryAllRatingStructResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
