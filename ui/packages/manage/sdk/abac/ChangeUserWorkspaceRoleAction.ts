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
 * Action to communicate with the action ChangeUserWorkspaceRole
 */
export type ChangeUserWorkspaceRoleActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type ChangeUserWorkspaceRoleActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ChangeUserWorkspaceRoleActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => ChangeUserWorkspaceRoleActionRes;
  }>;
export const useChangeUserWorkspaceRoleAction = (
  options?: ChangeUserWorkspaceRoleActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: ChangeUserWorkspaceRoleActionReq) => {
    setCompleteState(false);
    return ChangeUserWorkspaceRoleAction.Fetch(
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
 * ChangeUserWorkspaceRoleAction
 */
export class ChangeUserWorkspaceRoleAction {
  //
  static URL = "/workspace/change-user-role";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ChangeUserWorkspaceRoleAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<ChangeUserWorkspaceRoleActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<ChangeUserWorkspaceRoleActionRes>,
      ChangeUserWorkspaceRoleActionReq,
      unknown
    >(
      overrideUrl ?? ChangeUserWorkspaceRoleAction.NewUrl(qs),
      {
        method: ChangeUserWorkspaceRoleAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<ChangeUserWorkspaceRoleActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => ChangeUserWorkspaceRoleActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ChangeUserWorkspaceRoleActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new ChangeUserWorkspaceRoleActionRes(item));
    const res = await ChangeUserWorkspaceRoleAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<ChangeUserWorkspaceRoleActionRes>();
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
    name: "ChangeUserWorkspaceRole",
    cliName: "change-user-role",
    url: "/workspace/change-user-role",
    method: "post",
    description:
      "Replaces a user's role assignment(s) in a workspace with a single new role, root only. roleId must belong to workspaceId, not to root - list the real options with 'role browse' (or GET /role/browse) using a Workspace-Id header set to workspaceId, not root.",
    in: {
      fields: [
        {
          name: "userId",
          description:
            "UniqueId of the user whose role is changing. Must already be a member of workspaceId (see AddUserToWorkspace otherwise).",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "workspaceId",
          description: "UniqueId of the workspace the membership belongs to.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "roleId",
          description:
            "UniqueId of the new role (must belong to workspaceId) to assign to the user.",
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
          description: "UniqueId of the userWorkspace membership row.",
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
        {
          name: "roleId",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for changeUserWorkspaceRoleActionReq
 **/
export class ChangeUserWorkspaceRoleActionReq {
  /**
   * UniqueId of the user whose role is changing. Must already be a member of workspaceId (see AddUserToWorkspace otherwise).
   * @type {string}
   **/
  #userId: string = "";
  /**
   * UniqueId of the user whose role is changing. Must already be a member of workspaceId (see AddUserToWorkspace otherwise).
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * UniqueId of the user whose role is changing. Must already be a member of workspaceId (see AddUserToWorkspace otherwise).
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
   * UniqueId of the workspace the membership belongs to.
   * @type {string}
   **/
  #workspaceId: string = "";
  /**
   * UniqueId of the workspace the membership belongs to.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * UniqueId of the workspace the membership belongs to.
   * @type {string}
   **/
  set workspaceId(value: string) {
    this.#workspaceId = String(value);
  }
  setWorkspaceId(value: string) {
    this.workspaceId = value;
    return this;
  }
  /**
   * UniqueId of the new role (must belong to workspaceId) to assign to the user.
   * @type {string}
   **/
  #roleId: string = "";
  /**
   * UniqueId of the new role (must belong to workspaceId) to assign to the user.
   * @returns {string}
   **/
  get roleId() {
    return this.#roleId;
  }
  /**
   * UniqueId of the new role (must belong to workspaceId) to assign to the user.
   * @type {string}
   **/
  set roleId(value: string) {
    this.#roleId = String(value);
  }
  setRoleId(value: string) {
    this.roleId = value;
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
    const d = data as Partial<ChangeUserWorkspaceRoleActionReq>;
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.roleId !== undefined) {
      this.roleId = d.roleId;
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
      roleId: this.#roleId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      userId: "userId",
      workspaceId: "workspaceId",
      roleId: "roleId",
    };
  }
  /**
   * Creates an instance of ChangeUserWorkspaceRoleActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ChangeUserWorkspaceRoleActionReqType) {
    return new ChangeUserWorkspaceRoleActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ChangeUserWorkspaceRoleActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<ChangeUserWorkspaceRoleActionReqType>,
  ) {
    return new ChangeUserWorkspaceRoleActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ChangeUserWorkspaceRoleActionReqType>,
  ): InstanceType<typeof ChangeUserWorkspaceRoleActionReq> {
    return new ChangeUserWorkspaceRoleActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof ChangeUserWorkspaceRoleActionReq> {
    return new ChangeUserWorkspaceRoleActionReq(this.toJSON());
  }
}
export abstract class ChangeUserWorkspaceRoleActionReqFactory {
  abstract create(data: unknown): ChangeUserWorkspaceRoleActionReq;
}
/**
 * The base type definition for changeUserWorkspaceRoleActionReq
 **/
export type ChangeUserWorkspaceRoleActionReqType = {
  /**
   * UniqueId of the user whose role is changing. Must already be a member of workspaceId (see AddUserToWorkspace otherwise).
   * @type {string}
   **/
  userId: string;
  /**
   * UniqueId of the workspace the membership belongs to.
   * @type {string}
   **/
  workspaceId: string;
  /**
   * UniqueId of the new role (must belong to workspaceId) to assign to the user.
   * @type {string}
   **/
  roleId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ChangeUserWorkspaceRoleActionReqType {}
/**
 * The base class definition for changeUserWorkspaceRoleActionRes
 **/
export class ChangeUserWorkspaceRoleActionRes {
  /**
   * UniqueId of the userWorkspace membership row.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * UniqueId of the userWorkspace membership row.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * UniqueId of the userWorkspace membership row.
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
  /**
   *
   * @type {string}
   **/
  #roleId: string = "";
  /**
   *
   * @returns {string}
   **/
  get roleId() {
    return this.#roleId;
  }
  /**
   *
   * @type {string}
   **/
  set roleId(value: string) {
    this.#roleId = String(value);
  }
  setRoleId(value: string) {
    this.roleId = value;
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
    const d = data as Partial<ChangeUserWorkspaceRoleActionRes>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.roleId !== undefined) {
      this.roleId = d.roleId;
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
      roleId: this.#roleId,
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
      roleId: "roleId",
    };
  }
  /**
   * Creates an instance of ChangeUserWorkspaceRoleActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ChangeUserWorkspaceRoleActionResType) {
    return new ChangeUserWorkspaceRoleActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ChangeUserWorkspaceRoleActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<ChangeUserWorkspaceRoleActionResType>,
  ) {
    return new ChangeUserWorkspaceRoleActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ChangeUserWorkspaceRoleActionResType>,
  ): InstanceType<typeof ChangeUserWorkspaceRoleActionRes> {
    return new ChangeUserWorkspaceRoleActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof ChangeUserWorkspaceRoleActionRes> {
    return new ChangeUserWorkspaceRoleActionRes(this.toJSON());
  }
}
export abstract class ChangeUserWorkspaceRoleActionResFactory {
  abstract create(data: unknown): ChangeUserWorkspaceRoleActionRes;
}
/**
 * The base type definition for changeUserWorkspaceRoleActionRes
 **/
export type ChangeUserWorkspaceRoleActionResType = {
  /**
   * UniqueId of the userWorkspace membership row.
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
  /**
   *
   * @type {string}
   **/
  roleId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ChangeUserWorkspaceRoleActionResType {}
