import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletConfigDto } from "./WalletConfigDto";
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
 * Action to communicate with the action updateWalletConfig
 */
export type UpdateWalletConfigActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type UpdateWalletConfigActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UpdateWalletConfigActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletConfigDto;
  }>;
export const useUpdateWalletConfigAction = (
  options?: UpdateWalletConfigActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: UpdateWalletConfigActionReq) => {
    setCompleteState(false);
    return UpdateWalletConfigAction.Fetch(
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
 * UpdateWalletConfigAction
 */
export class UpdateWalletConfigAction {
  //
  static URL = "/wallet/config";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(UpdateWalletConfigAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<UpdateWalletConfigActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletConfigDto>,
      UpdateWalletConfigActionReq,
      unknown
    >(
      overrideUrl ?? UpdateWalletConfigAction.NewUrl(qs),
      {
        method: UpdateWalletConfigAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<UpdateWalletConfigActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletConfigDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletConfigDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletConfigDto(item));
    const res = await UpdateWalletConfigAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletConfigDto>();
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
    name: "updateWalletConfig",
    cliShort: "config-set",
    url: "/wallet/config",
    method: "post",
    description:
      "Root-only: updates the wallet limits configuration. Unset fields are left unchanged (partial update).",
    in: {
      fields: [
        {
          name: "maxWalletsPerUser",
          description: "New maximum wallets per user, across all currencies.",
          type: "int64?",
        },
        {
          name: "maxWalletsPerWorkspace",
          description:
            "New maximum wallets per workspace, across all currencies.",
          type: "int64?",
        },
        {
          name: "maxWalletsPerUserPerCurrency",
          description: "New maximum wallets per user per single currency.",
          type: "int64?",
        },
        {
          name: "maxWalletsPerWorkspacePerCurrency",
          description: "New maximum wallets per workspace per single currency.",
          type: "int64?",
        },
      ],
    },
    out: {
      envelope: "GResponse",
      dto: "WalletConfigDto",
    },
  };
}
/**
 * The base class definition for updateWalletConfigActionReq
 **/
export class UpdateWalletConfigActionReq {
  /**
   * New maximum wallets per user, across all currencies.
   * @type {number}
   **/
  #maxWalletsPerUser?: number | null = undefined;
  /**
   * New maximum wallets per user, across all currencies.
   * @returns {number}
   **/
  get maxWalletsPerUser() {
    return this.#maxWalletsPerUser;
  }
  /**
   * New maximum wallets per user, across all currencies.
   * @type {number}
   **/
  set maxWalletsPerUser(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerUser = parsedValue;
    }
  }
  setMaxWalletsPerUser(value: number | null | undefined) {
    this.maxWalletsPerUser = value;
    return this;
  }
  /**
   * New maximum wallets per workspace, across all currencies.
   * @type {number}
   **/
  #maxWalletsPerWorkspace?: number | null = undefined;
  /**
   * New maximum wallets per workspace, across all currencies.
   * @returns {number}
   **/
  get maxWalletsPerWorkspace() {
    return this.#maxWalletsPerWorkspace;
  }
  /**
   * New maximum wallets per workspace, across all currencies.
   * @type {number}
   **/
  set maxWalletsPerWorkspace(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerWorkspace = parsedValue;
    }
  }
  setMaxWalletsPerWorkspace(value: number | null | undefined) {
    this.maxWalletsPerWorkspace = value;
    return this;
  }
  /**
   * New maximum wallets per user per single currency.
   * @type {number}
   **/
  #maxWalletsPerUserPerCurrency?: number | null = undefined;
  /**
   * New maximum wallets per user per single currency.
   * @returns {number}
   **/
  get maxWalletsPerUserPerCurrency() {
    return this.#maxWalletsPerUserPerCurrency;
  }
  /**
   * New maximum wallets per user per single currency.
   * @type {number}
   **/
  set maxWalletsPerUserPerCurrency(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerUserPerCurrency = parsedValue;
    }
  }
  setMaxWalletsPerUserPerCurrency(value: number | null | undefined) {
    this.maxWalletsPerUserPerCurrency = value;
    return this;
  }
  /**
   * New maximum wallets per workspace per single currency.
   * @type {number}
   **/
  #maxWalletsPerWorkspacePerCurrency?: number | null = undefined;
  /**
   * New maximum wallets per workspace per single currency.
   * @returns {number}
   **/
  get maxWalletsPerWorkspacePerCurrency() {
    return this.#maxWalletsPerWorkspacePerCurrency;
  }
  /**
   * New maximum wallets per workspace per single currency.
   * @type {number}
   **/
  set maxWalletsPerWorkspacePerCurrency(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerWorkspacePerCurrency = parsedValue;
    }
  }
  setMaxWalletsPerWorkspacePerCurrency(value: number | null | undefined) {
    this.maxWalletsPerWorkspacePerCurrency = value;
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
    const d = data as Partial<UpdateWalletConfigActionReq>;
    if (d.maxWalletsPerUser !== undefined) {
      this.maxWalletsPerUser = d.maxWalletsPerUser;
    }
    if (d.maxWalletsPerWorkspace !== undefined) {
      this.maxWalletsPerWorkspace = d.maxWalletsPerWorkspace;
    }
    if (d.maxWalletsPerUserPerCurrency !== undefined) {
      this.maxWalletsPerUserPerCurrency = d.maxWalletsPerUserPerCurrency;
    }
    if (d.maxWalletsPerWorkspacePerCurrency !== undefined) {
      this.maxWalletsPerWorkspacePerCurrency =
        d.maxWalletsPerWorkspacePerCurrency;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      maxWalletsPerUser: this.#maxWalletsPerUser,
      maxWalletsPerWorkspace: this.#maxWalletsPerWorkspace,
      maxWalletsPerUserPerCurrency: this.#maxWalletsPerUserPerCurrency,
      maxWalletsPerWorkspacePerCurrency:
        this.#maxWalletsPerWorkspacePerCurrency,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      maxWalletsPerUser: "maxWalletsPerUser",
      maxWalletsPerWorkspace: "maxWalletsPerWorkspace",
      maxWalletsPerUserPerCurrency: "maxWalletsPerUserPerCurrency",
      maxWalletsPerWorkspacePerCurrency: "maxWalletsPerWorkspacePerCurrency",
    };
  }
  /**
   * Creates an instance of UpdateWalletConfigActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UpdateWalletConfigActionReqType) {
    return new UpdateWalletConfigActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of UpdateWalletConfigActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UpdateWalletConfigActionReqType>) {
    return new UpdateWalletConfigActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UpdateWalletConfigActionReqType>,
  ): InstanceType<typeof UpdateWalletConfigActionReq> {
    return new UpdateWalletConfigActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UpdateWalletConfigActionReq> {
    return new UpdateWalletConfigActionReq(this.toJSON());
  }
}
export abstract class UpdateWalletConfigActionReqFactory {
  abstract create(data: unknown): UpdateWalletConfigActionReq;
}
/**
 * The base type definition for updateWalletConfigActionReq
 **/
export type UpdateWalletConfigActionReqType = {
  /**
   * New maximum wallets per user, across all currencies.
   * @type {number}
   **/
  maxWalletsPerUser?: number;
  /**
   * New maximum wallets per workspace, across all currencies.
   * @type {number}
   **/
  maxWalletsPerWorkspace?: number;
  /**
   * New maximum wallets per user per single currency.
   * @type {number}
   **/
  maxWalletsPerUserPerCurrency?: number;
  /**
   * New maximum wallets per workspace per single currency.
   * @type {number}
   **/
  maxWalletsPerWorkspacePerCurrency?: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UpdateWalletConfigActionReqType {}
