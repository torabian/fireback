import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletViewDto } from "./WalletViewDto";
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
 * Action to communicate with the action createWallet
 */
export type CreateWalletActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CreateWalletActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CreateWalletActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletViewDto;
  }>;
export const useCreateWalletAction = (
  options?: CreateWalletActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: CreateWalletActionReq) => {
    setCompleteState(false);
    return CreateWalletAction.Fetch(
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
 * CreateWalletAction
 */
export class CreateWalletAction {
  //
  static URL = "/wallet/create";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CreateWalletAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<CreateWalletActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletViewDto>, CreateWalletActionReq, unknown>(
      overrideUrl ?? CreateWalletAction.NewUrl(qs),
      {
        method: CreateWalletAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<CreateWalletActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletViewDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletViewDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletViewDto(item));
    const res = await CreateWalletAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletViewDto>();
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
    name: "createWallet",
    cliShort: "create",
    url: "/wallet/create",
    method: "post",
    description:
      'Creates a new wallet for the caller (ownerType "user") or for a workspace the caller belongs to (ownerType "workspace"), in the given currency, with a starting balance of "0". Rejected if it would exceed walletConfig\'s maxWalletsPer{User,Workspace}[PerCurrency] limits, or if currency has no active walletCurrency.',
    in: {
      fields: [
        {
          name: "ownerType",
          description:
            "Whether this wallet belongs to the caller directly or to a workspace.",
          type: "string",
          tags: {
            validate: "required,oneof=user workspace",
          },
        },
        {
          name: "workspaceId",
          description:
            'Unique id of the owning workspace. Required when ownerType is "workspace" - the caller must belong to this workspace.',
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
      dto: "WalletViewDto",
    },
  };
}
/**
 * The base class definition for createWalletActionReq
 **/
export class CreateWalletActionReq {
  /**
   * Whether this wallet belongs to the caller directly or to a workspace.
   * @type {string}
   **/
  #ownerType: string = "";
  /**
   * Whether this wallet belongs to the caller directly or to a workspace.
   * @returns {string}
   **/
  get ownerType() {
    return this.#ownerType;
  }
  /**
   * Whether this wallet belongs to the caller directly or to a workspace.
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
   * Unique id of the owning workspace. Required when ownerType is "workspace" - the caller must belong to this workspace.
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   * Unique id of the owning workspace. Required when ownerType is "workspace" - the caller must belong to this workspace.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * Unique id of the owning workspace. Required when ownerType is "workspace" - the caller must belong to this workspace.
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
    const d = data as Partial<CreateWalletActionReq>;
    if (d.ownerType !== undefined) {
      this.ownerType = d.ownerType;
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
      workspaceId: "workspaceId",
      currency: "currency",
      label: "label",
    };
  }
  /**
   * Creates an instance of CreateWalletActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CreateWalletActionReqType) {
    return new CreateWalletActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of CreateWalletActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<CreateWalletActionReqType>) {
    return new CreateWalletActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CreateWalletActionReqType>,
  ): InstanceType<typeof CreateWalletActionReq> {
    return new CreateWalletActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CreateWalletActionReq> {
    return new CreateWalletActionReq(this.toJSON());
  }
}
export abstract class CreateWalletActionReqFactory {
  abstract create(data: unknown): CreateWalletActionReq;
}
/**
 * The base type definition for createWalletActionReq
 **/
export type CreateWalletActionReqType = {
  /**
   * Whether this wallet belongs to the caller directly or to a workspace.
   * @type {string}
   **/
  ownerType: string;
  /**
   * Unique id of the owning workspace. Required when ownerType is "workspace" - the caller must belong to this workspace.
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
export namespace CreateWalletActionReqType {}
