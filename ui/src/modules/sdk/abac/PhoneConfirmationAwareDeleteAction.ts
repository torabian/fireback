import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "../sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action phoneConfirmationAwareDelete
 */
export type PhoneConfirmationAwareDeleteActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PhoneConfirmationAwareDeleteActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PhoneConfirmationAwareDeleteActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  };
export const usePhoneConfirmationAwareDeleteAction = (
  options?: PhoneConfirmationAwareDeleteActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PhoneConfirmationAwareDeleteActionReq) => {
    setCompleteState(false);
    return PhoneConfirmationAwareDeleteAction.Fetch(
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
 * PhoneConfirmationAwareDeleteAction
 */
export class PhoneConfirmationAwareDeleteAction {
  //
  static URL = "/phoneConfirmation/delete";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PhoneConfirmationAwareDeleteAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PhoneConfirmationAwareDeleteActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<unknown, PhoneConfirmationAwareDeleteActionReq, unknown>(
      overrideUrl ?? PhoneConfirmationAwareDeleteAction.NewUrl(qs),
      {
        method: PhoneConfirmationAwareDeleteAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PhoneConfirmationAwareDeleteActionReq, unknown>,
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
    const res = await PhoneConfirmationAwareDeleteAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "phoneConfirmationAwareDelete",
    cliName: "delete",
    cliShort: "phoneConfirmation-d",
    url: "/phoneConfirmation/delete",
    method: "post",
    description:
      'Deletes the given "phoneConfirmation" uniqueIds, along with everything phoneConfirmationAwareDeletePreview reports.',
    in: {
      fields: [
        {
          name: "uniqueIds",
          type: "slice",
          primitive: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for phoneConfirmationAwareDeleteActionReq
 **/
export class PhoneConfirmationAwareDeleteActionReq {
  /**
   *
   * @type {string[]}
   **/
  #uniqueIds: string[] = [];
  /**
   *
   * @returns {string[]}
   **/
  get uniqueIds() {
    return this.#uniqueIds;
  }
  /**
   *
   * @type {string[]}
   **/
  set uniqueIds(value: string[]) {
    this.#uniqueIds = value;
  }
  setUniqueIds(value: string[]) {
    this.uniqueIds = value;
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
    const d = data as Partial<PhoneConfirmationAwareDeleteActionReq>;
    if (d.uniqueIds !== undefined) {
      this.uniqueIds = d.uniqueIds;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueIds: this.#uniqueIds,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueIds$: "uniqueIds",
      get uniqueIds() {
        return "uniqueIds[:i]";
      },
    };
  }
  /**
   * Creates an instance of PhoneConfirmationAwareDeleteActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PhoneConfirmationAwareDeleteActionReqType) {
    return new PhoneConfirmationAwareDeleteActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of PhoneConfirmationAwareDeleteActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<PhoneConfirmationAwareDeleteActionReqType>,
  ) {
    return new PhoneConfirmationAwareDeleteActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PhoneConfirmationAwareDeleteActionReqType>,
  ): InstanceType<typeof PhoneConfirmationAwareDeleteActionReq> {
    return new PhoneConfirmationAwareDeleteActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof PhoneConfirmationAwareDeleteActionReq> {
    return new PhoneConfirmationAwareDeleteActionReq(this.toJSON());
  }
}
export abstract class PhoneConfirmationAwareDeleteActionReqFactory {
  abstract create(data: unknown): PhoneConfirmationAwareDeleteActionReq;
}
/**
 * The base type definition for phoneConfirmationAwareDeleteActionReq
 **/
export type PhoneConfirmationAwareDeleteActionReqType = {
  /**
   *
   * @type {string[]}
   **/
  uniqueIds: string[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PhoneConfirmationAwareDeleteActionReqType {}
