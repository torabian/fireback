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
 * Action to communicate with the action AddUserToWorkspace
 */
export type AddUserToWorkspaceActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type AddUserToWorkspaceActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AddUserToWorkspaceActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => AddUserToWorkspaceActionRes;
  }>;
export const useAddUserToWorkspaceAction = (
  options?: AddUserToWorkspaceActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: AddUserToWorkspaceActionReq) => {
    setCompleteState(false);
    return AddUserToWorkspaceAction.Fetch(
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
 * AddUserToWorkspaceAction
 */
export class AddUserToWorkspaceAction {
  //
  static URL = "/workspace/add-user";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(AddUserToWorkspaceAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<AddUserToWorkspaceActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<AddUserToWorkspaceActionRes>,
      AddUserToWorkspaceActionReq,
      unknown
    >(
      overrideUrl ?? AddUserToWorkspaceAction.NewUrl(qs),
      {
        method: AddUserToWorkspaceAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<AddUserToWorkspaceActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => AddUserToWorkspaceActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new AddUserToWorkspaceActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new AddUserToWorkspaceActionRes(item));
    const res = await AddUserToWorkspaceAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<AddUserToWorkspaceActionRes>();
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
    name: "AddUserToWorkspace",
    cliName: "add-user",
    url: "/workspace/add-user",
    method: "post",
    description:
      "Adds an existing user directly to a workspace with one of that workspace's existing roles - no invitation/email involved, root only. Use InviteToWorkspace instead when the target person doesn't have an account yet.",
    in: {
      fields: [
        {
          name: "userId",
          description: "UniqueId of the existing user to add to the workspace.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "workspaceId",
          description: "UniqueId of the workspace to add the user to.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "roleId",
          description:
            "UniqueId of the role (must belong to workspaceId) to assign to the user.",
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
          description: "UniqueId of the created userWorkspace membership row.",
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
 * The base class definition for addUserToWorkspaceActionReq
 **/
export class AddUserToWorkspaceActionReq {
  /**
   * UniqueId of the existing user to add to the workspace.
   * @type {string}
   **/
  #userId: string = "";
  /**
   * UniqueId of the existing user to add to the workspace.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * UniqueId of the existing user to add to the workspace.
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
   * UniqueId of the workspace to add the user to.
   * @type {string}
   **/
  #workspaceId: string = "";
  /**
   * UniqueId of the workspace to add the user to.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * UniqueId of the workspace to add the user to.
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
   * UniqueId of the role (must belong to workspaceId) to assign to the user.
   * @type {string}
   **/
  #roleId: string = "";
  /**
   * UniqueId of the role (must belong to workspaceId) to assign to the user.
   * @returns {string}
   **/
  get roleId() {
    return this.#roleId;
  }
  /**
   * UniqueId of the role (must belong to workspaceId) to assign to the user.
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
    const d = data as Partial<AddUserToWorkspaceActionReq>;
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
   * Creates an instance of AddUserToWorkspaceActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AddUserToWorkspaceActionReqType) {
    return new AddUserToWorkspaceActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of AddUserToWorkspaceActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AddUserToWorkspaceActionReqType>) {
    return new AddUserToWorkspaceActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AddUserToWorkspaceActionReqType>,
  ): InstanceType<typeof AddUserToWorkspaceActionReq> {
    return new AddUserToWorkspaceActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AddUserToWorkspaceActionReq> {
    return new AddUserToWorkspaceActionReq(this.toJSON());
  }
}
export abstract class AddUserToWorkspaceActionReqFactory {
  abstract create(data: unknown): AddUserToWorkspaceActionReq;
}
/**
 * The base type definition for addUserToWorkspaceActionReq
 **/
export type AddUserToWorkspaceActionReqType = {
  /**
   * UniqueId of the existing user to add to the workspace.
   * @type {string}
   **/
  userId: string;
  /**
   * UniqueId of the workspace to add the user to.
   * @type {string}
   **/
  workspaceId: string;
  /**
   * UniqueId of the role (must belong to workspaceId) to assign to the user.
   * @type {string}
   **/
  roleId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AddUserToWorkspaceActionReqType {}
/**
 * The base class definition for addUserToWorkspaceActionRes
 **/
export class AddUserToWorkspaceActionRes {
  /**
   * UniqueId of the created userWorkspace membership row.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * UniqueId of the created userWorkspace membership row.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * UniqueId of the created userWorkspace membership row.
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
    const d = data as Partial<AddUserToWorkspaceActionRes>;
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
   * Creates an instance of AddUserToWorkspaceActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AddUserToWorkspaceActionResType) {
    return new AddUserToWorkspaceActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of AddUserToWorkspaceActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AddUserToWorkspaceActionResType>) {
    return new AddUserToWorkspaceActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AddUserToWorkspaceActionResType>,
  ): InstanceType<typeof AddUserToWorkspaceActionRes> {
    return new AddUserToWorkspaceActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AddUserToWorkspaceActionRes> {
    return new AddUserToWorkspaceActionRes(this.toJSON());
  }
}
export abstract class AddUserToWorkspaceActionResFactory {
  abstract create(data: unknown): AddUserToWorkspaceActionRes;
}
/**
 * The base type definition for addUserToWorkspaceActionRes
 **/
export type AddUserToWorkspaceActionResType = {
  /**
   * UniqueId of the created userWorkspace membership row.
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
export namespace AddUserToWorkspaceActionResType {}
