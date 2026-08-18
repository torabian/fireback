import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
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
 * Action to communicate with the action RemoveUserFromWorkspace
 */
export type RemoveUserFromWorkspaceActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type RemoveUserFromWorkspaceActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  RemoveUserFromWorkspaceActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => RemoveUserFromWorkspaceActionRes;
  }>;
export const useRemoveUserFromWorkspaceAction = (
  options?: RemoveUserFromWorkspaceActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: RemoveUserFromWorkspaceActionReq) => {
    setCompleteState(false);
    return RemoveUserFromWorkspaceAction.Fetch(
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
 * RemoveUserFromWorkspaceAction
 */
export class RemoveUserFromWorkspaceAction {
  //
  static URL = "/workspace/remove-user";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(RemoveUserFromWorkspaceAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<RemoveUserFromWorkspaceActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<RemoveUserFromWorkspaceActionRes>,
      RemoveUserFromWorkspaceActionReq,
      unknown
    >(
      overrideUrl ?? RemoveUserFromWorkspaceAction.NewUrl(qs),
      {
        method: RemoveUserFromWorkspaceAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<RemoveUserFromWorkspaceActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => RemoveUserFromWorkspaceActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new RemoveUserFromWorkspaceActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new RemoveUserFromWorkspaceActionRes(item));
    const res = await RemoveUserFromWorkspaceAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<RemoveUserFromWorkspaceActionRes>();
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
    name: "RemoveUserFromWorkspace",
    cliName: "remove-user",
    url: "/workspace/remove-user",
    method: "post",
    description:
      "Removes an existing user's membership - and every role assignment they have in it - from a workspace, root only. The counterpart to AddUserToWorkspace.",
    in: {
      fields: [
        {
          name: "userId",
          description: "UniqueId of the user to remove.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "workspaceId",
          description: "UniqueId of the workspace to remove the user from.",
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
          name: "uniqueId",
          description: "UniqueId of the removed userWorkspace membership row.",
          type: "string",
        },
        {
          name: "userId",
          type: "string",
        },
        {
          name: "workspaceId",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for removeUserFromWorkspaceActionReq
 **/
export class RemoveUserFromWorkspaceActionReq {
  /**
   * UniqueId of the user to remove.
   * @type {string}
   **/
  #userId: string = "";
  /**
   * UniqueId of the user to remove.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * UniqueId of the user to remove.
   * @type {string}
   **/
  set userId(value: string) {
    this.#userId = String(value);
  }
  setUserId(value: string) {
    this.userId = value;
    return this;
  }
  /**
   * UniqueId of the workspace to remove the user from.
   * @type {string}
   **/
  #workspaceId: string = "";
  /**
   * UniqueId of the workspace to remove the user from.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * UniqueId of the workspace to remove the user from.
   * @type {string}
   **/
  set workspaceId(value: string) {
    this.#workspaceId = String(value);
  }
  setWorkspaceId(value: string) {
    this.workspaceId = value;
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
    const d = data as Partial<RemoveUserFromWorkspaceActionReq>;
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      userId: this.#userId,
      workspaceId: this.#workspaceId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      userId: "userId",
      workspaceId: "workspaceId",
    };
  }
  /**
   * Creates an instance of RemoveUserFromWorkspaceActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: RemoveUserFromWorkspaceActionReqType) {
    return new RemoveUserFromWorkspaceActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of RemoveUserFromWorkspaceActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<RemoveUserFromWorkspaceActionReqType>,
  ) {
    return new RemoveUserFromWorkspaceActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<RemoveUserFromWorkspaceActionReqType>,
  ): InstanceType<typeof RemoveUserFromWorkspaceActionReq> {
    return new RemoveUserFromWorkspaceActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof RemoveUserFromWorkspaceActionReq> {
    return new RemoveUserFromWorkspaceActionReq(this.toJSON());
  }
}
export abstract class RemoveUserFromWorkspaceActionReqFactory {
  abstract create(data: unknown): RemoveUserFromWorkspaceActionReq;
}
/**
 * The base type definition for removeUserFromWorkspaceActionReq
 **/
export type RemoveUserFromWorkspaceActionReqType = {
  /**
   * UniqueId of the user to remove.
   * @type {string}
   **/
  userId: string;
  /**
   * UniqueId of the workspace to remove the user from.
   * @type {string}
   **/
  workspaceId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace RemoveUserFromWorkspaceActionReqType {}
/**
 * The base class definition for removeUserFromWorkspaceActionRes
 **/
export class RemoveUserFromWorkspaceActionRes {
  /**
   * UniqueId of the removed userWorkspace membership row.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * UniqueId of the removed userWorkspace membership row.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * UniqueId of the removed userWorkspace membership row.
   * @type {string}
   **/
  set uniqueId(value: string) {
    this.#uniqueId = String(value);
  }
  setUniqueId(value: string) {
    this.uniqueId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #userId: string = "";
  /**
   *
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   *
   * @type {string}
   **/
  set userId(value: string) {
    this.#userId = String(value);
  }
  setUserId(value: string) {
    this.userId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #workspaceId: string = "";
  /**
   *
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   *
   * @type {string}
   **/
  set workspaceId(value: string) {
    this.#workspaceId = String(value);
  }
  setWorkspaceId(value: string) {
    this.workspaceId = value;
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
    const d = data as Partial<RemoveUserFromWorkspaceActionRes>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      userId: this.#userId,
      workspaceId: this.#workspaceId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      userId: "userId",
      workspaceId: "workspaceId",
    };
  }
  /**
   * Creates an instance of RemoveUserFromWorkspaceActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: RemoveUserFromWorkspaceActionResType) {
    return new RemoveUserFromWorkspaceActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of RemoveUserFromWorkspaceActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<RemoveUserFromWorkspaceActionResType>,
  ) {
    return new RemoveUserFromWorkspaceActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<RemoveUserFromWorkspaceActionResType>,
  ): InstanceType<typeof RemoveUserFromWorkspaceActionRes> {
    return new RemoveUserFromWorkspaceActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof RemoveUserFromWorkspaceActionRes> {
    return new RemoveUserFromWorkspaceActionRes(this.toJSON());
  }
}
export abstract class RemoveUserFromWorkspaceActionResFactory {
  abstract create(data: unknown): RemoveUserFromWorkspaceActionRes;
}
/**
 * The base type definition for removeUserFromWorkspaceActionRes
 **/
export type RemoveUserFromWorkspaceActionResType = {
  /**
   * UniqueId of the removed userWorkspace membership row.
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {string}
   **/
  userId: string;
  /**
   *
   * @type {string}
   **/
  workspaceId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace RemoveUserFromWorkspaceActionResType {}
