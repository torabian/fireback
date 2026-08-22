import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletPaymentAttemptViewDto } from "./WalletPaymentAttemptViewDto";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action cancelPaymentAttempt
 */
export type CancelPaymentAttemptActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CancelPaymentAttemptActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CancelPaymentAttemptActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletPaymentAttemptViewDto;
  }>;
export const useCancelPaymentAttemptAction = (
  options?: CancelPaymentAttemptActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: CancelPaymentAttemptActionReq) => {
    setCompleteState(false);
    return CancelPaymentAttemptAction.Fetch(
      {
        body,
        headers: options?.headers,
      },
      {
        creatorFn: options?.creatorFn,
        qs: options?.qs,
        ctx,
        onMessage: options?.onMessage,
        overrideUrl: options?.overrideUrl,
      },
    ).then((x) => {
      x.done.then(() => {
        setCompleteState(true);
      });
      setResponse(x.response);
      return x.response.result;
    });
  };
  const result = useMutation({
    mutationFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
/**
 * CancelPaymentAttemptAction
 */
export class CancelPaymentAttemptAction {
  //
  static URL = "/wallet/attempts/cancel";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CancelPaymentAttemptAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<CancelPaymentAttemptActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletPaymentAttemptViewDto>,
      CancelPaymentAttemptActionReq,
      unknown
    >(
      overrideUrl ?? CancelPaymentAttemptAction.NewUrl(qs),
      {
        method: CancelPaymentAttemptAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<CancelPaymentAttemptActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletPaymentAttemptViewDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletPaymentAttemptViewDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletPaymentAttemptViewDto(item));
    const res = await CancelPaymentAttemptAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletPaymentAttemptViewDto>();
        if (creatorFn) {
          resp.setCreator(creatorFn);
        }
        resp.inject(data);
        return resp;
      },
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "cancelPaymentAttempt",
    cliShort: "cancel",
    url: "/wallet/attempts/cancel",
    method: "post",
    description:
      'Cancels one of the caller\'s own payment attempts, only while it is still "pending".',
    in: {
      fields: [
        {
          name: "attemptId",
          description: "Unique id of the payment attempt to cancel.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
      ],
    },
    out: {
      envelope: "GResponse",
      dto: "WalletPaymentAttemptViewDto",
    },
  };
}
/**
 * The base class definition for cancelPaymentAttemptActionReq
 **/
export class CancelPaymentAttemptActionReq {
  /**
   * Unique id of the payment attempt to cancel.
   * @type {string}
   **/
  #attemptId: string = "";
  /**
   * Unique id of the payment attempt to cancel.
   * @returns {string}
   **/
  get attemptId() {
    return this.#attemptId;
  }
  /**
   * Unique id of the payment attempt to cancel.
   * @type {string}
   **/
  set attemptId(value: string) {
    this.#attemptId = String(value);
  }
  setAttemptId(value: string) {
    this.attemptId = value;
    return this;
  }
  constructor(data: unknown = undefined) {
    if (data === null || data === undefined) {
      return;
    }
    if (typeof data === "string") {
      this.applyFromObject(JSON.parse(data));
    } else if (this.#isJsonAppliable(data)) {
      this.applyFromObject(data);
    } else {
      throw new Error(
        "Instance cannot be created on an unknown value, check the content being passed. got: " +
          typeof data,
      );
    }
  }
  #isJsonAppliable(obj: unknown) {
    const g = globalThis as unknown as { Buffer: any; Blob: any };
    const isBuffer =
      typeof g.Buffer !== "undefined" &&
      typeof g.Buffer.isBuffer === "function" &&
      g.Buffer.isBuffer(obj);
    const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
    return (
      obj &&
      typeof obj === "object" &&
      !Array.isArray(obj) &&
      !isBuffer &&
      !(obj instanceof ArrayBuffer) &&
      !isBlob
    );
  }
  /**
   * casts the fields of a javascript object into the class properties one by one
   **/
  applyFromObject(data = {}) {
    const d = data as Partial<CancelPaymentAttemptActionReq>;
    if (d.attemptId !== undefined) {
      this.attemptId = d.attemptId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      attemptId: this.#attemptId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      attemptId: "attemptId",
    };
  }
  /**
   * Creates an instance of CancelPaymentAttemptActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CancelPaymentAttemptActionReqType) {
    return new CancelPaymentAttemptActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of CancelPaymentAttemptActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CancelPaymentAttemptActionReqType>,
  ) {
    return new CancelPaymentAttemptActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CancelPaymentAttemptActionReqType>,
  ): InstanceType<typeof CancelPaymentAttemptActionReq> {
    return new CancelPaymentAttemptActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CancelPaymentAttemptActionReq> {
    return new CancelPaymentAttemptActionReq(this.toJSON());
  }
}
export abstract class CancelPaymentAttemptActionReqFactory {
  abstract create(data: unknown): CancelPaymentAttemptActionReq;
}
/**
 * The base type definition for cancelPaymentAttemptActionReq
 **/
export type CancelPaymentAttemptActionReqType = {
  /**
   * Unique id of the payment attempt to cancel.
   * @type {string}
   **/
  attemptId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CancelPaymentAttemptActionReqType {}
