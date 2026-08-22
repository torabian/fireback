import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletDto } from "./WalletDto";
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
 * Action to communicate with the action adminCreateWallet
 */
export type AdminCreateWalletActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type AdminCreateWalletActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AdminCreateWalletActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletDto;
  }>;
export const useAdminCreateWalletAction = (
  options?: AdminCreateWalletActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: AdminCreateWalletActionReq) => {
    setCompleteState(false);
    return AdminCreateWalletAction.Fetch(
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
 * AdminCreateWalletAction
 */
export class AdminCreateWalletAction {
  //
  static URL = "/wallet/admin-create";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(AdminCreateWalletAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<AdminCreateWalletActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletDto>, AdminCreateWalletActionReq, unknown>(
      overrideUrl ?? AdminCreateWalletAction.NewUrl(qs),
      {
        method: AdminCreateWalletAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<AdminCreateWalletActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletDto(item));
    const res = await AdminCreateWalletAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletDto>();
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
    name: "adminCreateWallet",
    cliShort: "admin-create",
    url: "/wallet/admin-create",
    method: "post",
    description:
      "Root-only: creates a wallet on behalf of any user or workspace, not just the caller - the admin/support counterpart to walletpublic's createWallet (which only ever creates for the caller themselves). Enforces the same walletConfig.maxWalletsPer{User,Workspace} [PerCurrency] limits and active-currency check as the owner-facing path (see wallet.CheckWalletLimit in WalletLimit.go, shared by both). Rejects if the target userId/workspaceId doesn't exist.",
    in: {
      fields: [
        {
          name: "ownerType",
          description:
            "Whether the new wallet belongs to a user or a workspace.",
          type: "string",
          tags: {
            validate: "required,oneof=user workspace",
          },
        },
        {
          name: "userId",
          description:
            'Unique id of the target user. Required when ownerType is "user" - must be an existing user.',
          type: "string?",
          tags: {
            validate: "required_if=ownerType user",
          },
        },
        {
          name: "workspaceId",
          description:
            'Unique id of the target workspace. Required when ownerType is "workspace" - must be an existing workspace. Unlike walletpublic\'s createWallet, the caller does not need to be a member of it.',
          type: "string?",
          tags: {
            validate: "required_if=ownerType workspace",
          },
        },
        {
          name: "currency",
          description:
            "Currency code for the new wallet - must match an active walletCurrency.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "label",
          description: "Optional nickname for the new wallet.",
          type: "string?",
        },
      ],
    },
    out: {
      envelope: "GResponse",
      dto: "WalletDto",
    },
  };
}
/**
 * The base class definition for adminCreateWalletActionReq
 **/
export class AdminCreateWalletActionReq {
  /**
   * Whether the new wallet belongs to a user or a workspace.
   * @type {string}
   **/
  #ownerType: string = "";
  /**
   * Whether the new wallet belongs to a user or a workspace.
   * @returns {string}
   **/
  get ownerType() {
    return this.#ownerType;
  }
  /**
   * Whether the new wallet belongs to a user or a workspace.
   * @type {string}
   **/
  set ownerType(value: string) {
    this.#ownerType = String(value);
  }
  setOwnerType(value: string) {
    this.ownerType = value;
    return this;
  }
  /**
   * Unique id of the target user. Required when ownerType is "user" - must be an existing user.
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * Unique id of the target user. Required when ownerType is "user" - must be an existing user.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * Unique id of the target user. Required when ownerType is "user" - must be an existing user.
   * @type {string}
   **/
  set userId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#userId = correctType ? value : String(value);
  }
  setUserId(value: string | null | undefined) {
    this.userId = value;
    return this;
  }
  /**
   * Unique id of the target workspace. Required when ownerType is "workspace" - must be an existing workspace. Unlike walletpublic's createWallet, the caller does not need to be a member of it.
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   * Unique id of the target workspace. Required when ownerType is "workspace" - must be an existing workspace. Unlike walletpublic's createWallet, the caller does not need to be a member of it.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * Unique id of the target workspace. Required when ownerType is "workspace" - must be an existing workspace. Unlike walletpublic's createWallet, the caller does not need to be a member of it.
   * @type {string}
   **/
  set workspaceId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#workspaceId = correctType ? value : String(value);
  }
  setWorkspaceId(value: string | null | undefined) {
    this.workspaceId = value;
    return this;
  }
  /**
   * Currency code for the new wallet - must match an active walletCurrency.
   * @type {string}
   **/
  #currency: string = "";
  /**
   * Currency code for the new wallet - must match an active walletCurrency.
   * @returns {string}
   **/
  get currency() {
    return this.#currency;
  }
  /**
   * Currency code for the new wallet - must match an active walletCurrency.
   * @type {string}
   **/
  set currency(value: string) {
    this.#currency = String(value);
  }
  setCurrency(value: string) {
    this.currency = value;
    return this;
  }
  /**
   * Optional nickname for the new wallet.
   * @type {string}
   **/
  #label?: string | null = undefined;
  /**
   * Optional nickname for the new wallet.
   * @returns {string}
   **/
  get label() {
    return this.#label;
  }
  /**
   * Optional nickname for the new wallet.
   * @type {string}
   **/
  set label(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#label = correctType ? value : String(value);
  }
  setLabel(value: string | null | undefined) {
    this.label = value;
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
    const d = data as Partial<AdminCreateWalletActionReq>;
    if (d.ownerType !== undefined) {
      this.ownerType = d.ownerType;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.currency !== undefined) {
      this.currency = d.currency;
    }
    if (d.label !== undefined) {
      this.label = d.label;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      ownerType: this.#ownerType,
      userId: this.#userId,
      workspaceId: this.#workspaceId,
      currency: this.#currency,
      label: this.#label,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      ownerType: "ownerType",
      userId: "userId",
      workspaceId: "workspaceId",
      currency: "currency",
      label: "label",
    };
  }
  /**
   * Creates an instance of AdminCreateWalletActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AdminCreateWalletActionReqType) {
    return new AdminCreateWalletActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of AdminCreateWalletActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AdminCreateWalletActionReqType>) {
    return new AdminCreateWalletActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AdminCreateWalletActionReqType>,
  ): InstanceType<typeof AdminCreateWalletActionReq> {
    return new AdminCreateWalletActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AdminCreateWalletActionReq> {
    return new AdminCreateWalletActionReq(this.toJSON());
  }
}
export abstract class AdminCreateWalletActionReqFactory {
  abstract create(data: unknown): AdminCreateWalletActionReq;
}
/**
 * The base type definition for adminCreateWalletActionReq
 **/
export type AdminCreateWalletActionReqType = {
  /**
   * Whether the new wallet belongs to a user or a workspace.
   * @type {string}
   **/
  ownerType: string;
  /**
   * Unique id of the target user. Required when ownerType is "user" - must be an existing user.
   * @type {string}
   **/
  userId?: string;
  /**
   * Unique id of the target workspace. Required when ownerType is "workspace" - must be an existing workspace. Unlike walletpublic's createWallet, the caller does not need to be a member of it.
   * @type {string}
   **/
  workspaceId?: string;
  /**
   * Currency code for the new wallet - must match an active walletCurrency.
   * @type {string}
   **/
  currency: string;
  /**
   * Optional nickname for the new wallet.
   * @type {string}
   **/
  label?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdminCreateWalletActionReqType {}
