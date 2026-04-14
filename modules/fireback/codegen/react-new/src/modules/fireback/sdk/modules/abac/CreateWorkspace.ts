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
 * Action to communicate with the action CreateWorkspace
 */
export type CreateWorkspaceActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CreateWorkspaceActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CreateWorkspaceActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => CreateWorkspaceActionRes;
  }>;
export const useCreateWorkspaceAction = (
  options?: CreateWorkspaceActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: CreateWorkspaceActionReq) => {
    setCompleteState(false);
    return CreateWorkspaceAction.Fetch(
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
 * CreateWorkspaceAction
 */
export class CreateWorkspaceAction {
  //
  static URL = "/workspaces/create";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CreateWorkspaceAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<CreateWorkspaceActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<CreateWorkspaceActionRes>,
      CreateWorkspaceActionReq,
      unknown
    >(
      overrideUrl ?? CreateWorkspaceAction.NewUrl(qs),
      {
        method: CreateWorkspaceAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<CreateWorkspaceActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => CreateWorkspaceActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new CreateWorkspaceActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new CreateWorkspaceActionRes(item));
    const res = await CreateWorkspaceAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<CreateWorkspaceActionRes>();
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
    name: "CreateWorkspace",
    url: "/workspaces/create",
    method: "post",
    in: {
      fields: [
        {
          name: "name",
          type: "string",
        },
      ],
    },
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "workspaceId",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for createWorkspaceActionReq
 **/
export class CreateWorkspaceActionReq {
  /**
   *
   * @type {string}
   **/
  #name: string = "";
  /**
   *
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   *
   * @type {string}
   **/
  set name(value: string) {
    this.#name = String(value);
  }
  setName(value: string) {
    this.name = value;
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
    const d = data as Partial<CreateWorkspaceActionReq>;
    if (d.name !== undefined) {
      this.name = d.name;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      name: this.#name,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      name: "name",
    };
  }
  /**
   * Creates an instance of CreateWorkspaceActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CreateWorkspaceActionReqType) {
    return new CreateWorkspaceActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of CreateWorkspaceActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<CreateWorkspaceActionReqType>) {
    return new CreateWorkspaceActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CreateWorkspaceActionReqType>,
  ): InstanceType<typeof CreateWorkspaceActionReq> {
    return new CreateWorkspaceActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CreateWorkspaceActionReq> {
    return new CreateWorkspaceActionReq(this.toJSON());
  }
}
export abstract class CreateWorkspaceActionReqFactory {
  abstract create(data: unknown): CreateWorkspaceActionReq;
}
/**
 * The base type definition for createWorkspaceActionReq
 **/
export type CreateWorkspaceActionReqType = {
  /**
   *
   * @type {string}
   **/
  name: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CreateWorkspaceActionReqType {}
/**
 * The base class definition for createWorkspaceActionRes
 **/
export class CreateWorkspaceActionRes {
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
    const d = data as Partial<CreateWorkspaceActionRes>;
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
      workspaceId: this.#workspaceId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      workspaceId: "workspaceId",
    };
  }
  /**
   * Creates an instance of CreateWorkspaceActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CreateWorkspaceActionResType) {
    return new CreateWorkspaceActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of CreateWorkspaceActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<CreateWorkspaceActionResType>) {
    return new CreateWorkspaceActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CreateWorkspaceActionResType>,
  ): InstanceType<typeof CreateWorkspaceActionRes> {
    return new CreateWorkspaceActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CreateWorkspaceActionRes> {
    return new CreateWorkspaceActionRes(this.toJSON());
  }
}
export abstract class CreateWorkspaceActionResFactory {
  abstract create(data: unknown): CreateWorkspaceActionRes;
}
/**
 * The base type definition for createWorkspaceActionRes
 **/
export type CreateWorkspaceActionResType = {
  /**
   *
   * @type {string}
   **/
  workspaceId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CreateWorkspaceActionResType {}
