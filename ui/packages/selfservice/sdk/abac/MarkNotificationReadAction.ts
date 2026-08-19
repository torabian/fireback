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
 * Action to communicate with the action MarkNotificationRead
 */
export type MarkNotificationReadActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type MarkNotificationReadActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  MarkNotificationReadActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  };
export const useMarkNotificationReadAction = (
  options?: MarkNotificationReadActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: MarkNotificationReadActionReq) => {
    setCompleteState(false);
    return MarkNotificationReadAction.Fetch(
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
 * MarkNotificationReadAction
 */
export class MarkNotificationReadAction {
  //
  static URL = "/notification/mark-read";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(MarkNotificationReadAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<MarkNotificationReadActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<unknown, MarkNotificationReadActionReq, unknown>(
      overrideUrl ?? MarkNotificationReadAction.NewUrl(qs),
      {
        method: MarkNotificationReadAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<MarkNotificationReadActionReq, unknown>,
    {
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {},
  ) => {
    const res = await MarkNotificationReadAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "MarkNotificationRead",
    cliName: "mark-read",
    url: "/notification/mark-read",
    method: "post",
    description:
      "Self-service. Marks one of the calling user's own notifications as read. 404s (ResourceNotFound, not a permission error) if the given uniqueId doesn't belong to the caller - never lets a user learn whether some other uniqueId exists at all, let alone mark it read.",
    in: {
      fields: [
        {
          name: "uniqueId",
          description: "The notification uniqueId to mark as read.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
      ],
    },
    out: {
      envelope: "GResponse",
    },
  };
}
/**
 * The base class definition for markNotificationReadActionReq
 **/
export class MarkNotificationReadActionReq {
  /**
   * The notification uniqueId to mark as read.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * The notification uniqueId to mark as read.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * The notification uniqueId to mark as read.
   * @type {string}
   **/
  set uniqueId(value: string) {
    this.#uniqueId = String(value);
  }
  setUniqueId(value: string) {
    this.uniqueId = value;
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
    const d = data as Partial<MarkNotificationReadActionReq>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
    };
  }
  /**
   * Creates an instance of MarkNotificationReadActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: MarkNotificationReadActionReqType) {
    return new MarkNotificationReadActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of MarkNotificationReadActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<MarkNotificationReadActionReqType>,
  ) {
    return new MarkNotificationReadActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<MarkNotificationReadActionReqType>,
  ): InstanceType<typeof MarkNotificationReadActionReq> {
    return new MarkNotificationReadActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof MarkNotificationReadActionReq> {
    return new MarkNotificationReadActionReq(this.toJSON());
  }
}
export abstract class MarkNotificationReadActionReqFactory {
  abstract create(data: unknown): MarkNotificationReadActionReq;
}
/**
 * The base type definition for markNotificationReadActionReq
 **/
export type MarkNotificationReadActionReqType = {
  /**
   * The notification uniqueId to mark as read.
   * @type {string}
   **/
  uniqueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MarkNotificationReadActionReqType {}
