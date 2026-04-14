import { GResponse } from "../../sdk/envelopes/index";
import { buildUrl } from "../../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action AcceptInvite
 */
export type AcceptInviteActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type AcceptInviteActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AcceptInviteActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => AcceptInviteActionRes;
  }>;
export const useAcceptInviteAction = (
  options?: AcceptInviteActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: AcceptInviteActionReq) => {
    setCompleteState(false);
    return AcceptInviteAction.Fetch(
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
 * AcceptInviteAction
 */
export class AcceptInviteAction {
  //
  static URL = "/user/invitation/accept";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(AcceptInviteAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<AcceptInviteActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<AcceptInviteActionRes>,
      AcceptInviteActionReq,
      unknown
    >(
      overrideUrl ?? AcceptInviteAction.NewUrl(qs),
      {
        method: AcceptInviteAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<AcceptInviteActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => AcceptInviteActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new AcceptInviteActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new AcceptInviteActionRes(item));
    const res = await AcceptInviteAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<AcceptInviteActionRes>();
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
    name: "AcceptInvite",
    url: "/user/invitation/accept",
    method: "post",
    description:
      "Use it when user accepts an invitation, and it will complete the joining process",
    in: {
      fields: [
        {
          name: "invitationUniqueId",
          description: "The invitation id which will be used to process",
          type: "string",
          tags: {
            validate: "required",
          },
        },
      ],
    },
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "accepted",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for acceptInviteActionReq
 **/
export class AcceptInviteActionReq {
  /**
   * The invitation id which will be used to process
   * @type {string}
   **/
  #invitationUniqueId: string = "";
  /**
   * The invitation id which will be used to process
   * @returns {string}
   **/
  get invitationUniqueId() {
    return this.#invitationUniqueId;
  }
  /**
   * The invitation id which will be used to process
   * @type {string}
   **/
  set invitationUniqueId(value: string) {
    this.#invitationUniqueId = String(value);
  }
  setInvitationUniqueId(value: string) {
    this.invitationUniqueId = value;
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
    const d = data as Partial<AcceptInviteActionReq>;
    if (d.invitationUniqueId !== undefined) {
      this.invitationUniqueId = d.invitationUniqueId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      invitationUniqueId: this.#invitationUniqueId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      invitationUniqueId: "invitationUniqueId",
    };
  }
  /**
   * Creates an instance of AcceptInviteActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AcceptInviteActionReqType) {
    return new AcceptInviteActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of AcceptInviteActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AcceptInviteActionReqType>) {
    return new AcceptInviteActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AcceptInviteActionReqType>,
  ): InstanceType<typeof AcceptInviteActionReq> {
    return new AcceptInviteActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AcceptInviteActionReq> {
    return new AcceptInviteActionReq(this.toJSON());
  }
}
export abstract class AcceptInviteActionReqFactory {
  abstract create(data: unknown): AcceptInviteActionReq;
}
/**
 * The base type definition for acceptInviteActionReq
 **/
export type AcceptInviteActionReqType = {
  /**
   * The invitation id which will be used to process
   * @type {string}
   **/
  invitationUniqueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AcceptInviteActionReqType {}
/**
 * The base class definition for acceptInviteActionRes
 **/
export class AcceptInviteActionRes {
  /**
   *
   * @type {boolean}
   **/
  #accepted!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get accepted() {
    return this.#accepted;
  }
  /**
   *
   * @type {boolean}
   **/
  set accepted(value: boolean) {
    this.#accepted = Boolean(value);
  }
  setAccepted(value: boolean) {
    this.accepted = value;
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
    const d = data as Partial<AcceptInviteActionRes>;
    if (d.accepted !== undefined) {
      this.accepted = d.accepted;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      accepted: this.#accepted,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      accepted: "accepted",
    };
  }
  /**
   * Creates an instance of AcceptInviteActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AcceptInviteActionResType) {
    return new AcceptInviteActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of AcceptInviteActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AcceptInviteActionResType>) {
    return new AcceptInviteActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AcceptInviteActionResType>,
  ): InstanceType<typeof AcceptInviteActionRes> {
    return new AcceptInviteActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AcceptInviteActionRes> {
    return new AcceptInviteActionRes(this.toJSON());
  }
}
export abstract class AcceptInviteActionResFactory {
  abstract create(data: unknown): AcceptInviteActionRes;
}
/**
 * The base type definition for acceptInviteActionRes
 **/
export type AcceptInviteActionResType = {
  /**
   *
   * @type {boolean}
   **/
  accepted: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AcceptInviteActionResType {}
