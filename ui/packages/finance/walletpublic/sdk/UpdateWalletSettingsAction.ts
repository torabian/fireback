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
 * Action to communicate with the action updateWalletSettings
 */
export type UpdateWalletSettingsActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type UpdateWalletSettingsActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UpdateWalletSettingsActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletViewDto;
  }>;
export const useUpdateWalletSettingsAction = (
  options?: UpdateWalletSettingsActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: UpdateWalletSettingsActionReq) => {
    setCompleteState(false);
    return UpdateWalletSettingsAction.Fetch(
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
 * UpdateWalletSettingsAction
 */
export class UpdateWalletSettingsAction {
  //
  static URL = "/wallet/settings";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(UpdateWalletSettingsAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<UpdateWalletSettingsActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletViewDto>,
      UpdateWalletSettingsActionReq,
      unknown
    >(
      overrideUrl ?? UpdateWalletSettingsAction.NewUrl(qs),
      {
        method: UpdateWalletSettingsAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<UpdateWalletSettingsActionReq, unknown>,
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
    const res = await UpdateWalletSettingsAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
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
    name: "updateWalletSettings",
    cliShort: "settings",
    url: "/wallet/settings",
    method: "post",
    description:
      "Updates a wallet's non-financial fields only (label/status/isDefault) - balance can never be set through this action. The caller must be the wallet's owning user, or a member of its owning workspace.",
    in: {
      fields: [
        {
          name: "walletId",
          description: "Unique id of the wallet to update.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "label",
          description: "New label. Omit to leave unchanged.",
          type: "string?",
        },
        {
          name: "status",
          description:
            "New status. Only active/frozen may be set here - closing a wallet is a separate, root-only operation (see wallet module's adjustBalance/admin tooling).",
          type: "string?",
          tags: {
            validate: "omitempty,oneof=active frozen",
          },
        },
        {
          name: "isDefault",
          description: "New isDefault flag. Omit to leave unchanged.",
          type: "bool?",
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
 * The base class definition for updateWalletSettingsActionReq
 **/
export class UpdateWalletSettingsActionReq {
  /**
   * Unique id of the wallet to update.
   * @type {string}
   **/
  #walletId: string = "";
  /**
   * Unique id of the wallet to update.
   * @returns {string}
   **/
  get walletId() {
    return this.#walletId;
  }
  /**
   * Unique id of the wallet to update.
   * @type {string}
   **/
  set walletId(value: string) {
    this.#walletId = String(value);
  }
  setWalletId(value: string) {
    this.walletId = value;
    return this;
  }
  /**
   * New label. Omit to leave unchanged.
   * @type {string}
   **/
  #label?: string | null = undefined;
  /**
   * New label. Omit to leave unchanged.
   * @returns {string}
   **/
  get label() {
    return this.#label;
  }
  /**
   * New label. Omit to leave unchanged.
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
  /**
   * New status. Only active/frozen may be set here - closing a wallet is a separate, root-only operation (see wallet module's adjustBalance/admin tooling).
   * @type {string}
   **/
  #status?: string | null = undefined;
  /**
   * New status. Only active/frozen may be set here - closing a wallet is a separate, root-only operation (see wallet module's adjustBalance/admin tooling).
   * @returns {string}
   **/
  get status() {
    return this.#status;
  }
  /**
   * New status. Only active/frozen may be set here - closing a wallet is a separate, root-only operation (see wallet module's adjustBalance/admin tooling).
   * @type {string}
   **/
  set status(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#status = correctType ? value : String(value);
  }
  setStatus(value: string | null | undefined) {
    this.status = value;
    return this;
  }
  /**
   * New isDefault flag. Omit to leave unchanged.
   * @type {boolean}
   **/
  #isDefault?: boolean | null = undefined;
  /**
   * New isDefault flag. Omit to leave unchanged.
   * @returns {boolean}
   **/
  get isDefault() {
    return this.#isDefault;
  }
  /**
   * New isDefault flag. Omit to leave unchanged.
   * @type {boolean}
   **/
  set isDefault(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isDefault = correctType ? value : Boolean(value);
  }
  setIsDefault(value: boolean | null | undefined) {
    this.isDefault = value;
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
    const d = data as Partial<UpdateWalletSettingsActionReq>;
    if (d.walletId !== undefined) {
      this.walletId = d.walletId;
    }
    if (d.label !== undefined) {
      this.label = d.label;
    }
    if (d.status !== undefined) {
      this.status = d.status;
    }
    if (d.isDefault !== undefined) {
      this.isDefault = d.isDefault;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      walletId: this.#walletId,
      label: this.#label,
      status: this.#status,
      isDefault: this.#isDefault,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      walletId: "walletId",
      label: "label",
      status: "status",
      isDefault: "isDefault",
    };
  }
  /**
   * Creates an instance of UpdateWalletSettingsActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UpdateWalletSettingsActionReqType) {
    return new UpdateWalletSettingsActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of UpdateWalletSettingsActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<UpdateWalletSettingsActionReqType>,
  ) {
    return new UpdateWalletSettingsActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UpdateWalletSettingsActionReqType>,
  ): InstanceType<typeof UpdateWalletSettingsActionReq> {
    return new UpdateWalletSettingsActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UpdateWalletSettingsActionReq> {
    return new UpdateWalletSettingsActionReq(this.toJSON());
  }
}
export abstract class UpdateWalletSettingsActionReqFactory {
  abstract create(data: unknown): UpdateWalletSettingsActionReq;
}
/**
 * The base type definition for updateWalletSettingsActionReq
 **/
export type UpdateWalletSettingsActionReqType = {
  /**
   * Unique id of the wallet to update.
   * @type {string}
   **/
  walletId: string;
  /**
   * New label. Omit to leave unchanged.
   * @type {string}
   **/
  label?: string;
  /**
   * New status. Only active/frozen may be set here - closing a wallet is a separate, root-only operation (see wallet module's adjustBalance/admin tooling).
   * @type {string}
   **/
  status?: string;
  /**
   * New isDefault flag. Omit to leave unchanged.
   * @type {boolean}
   **/
  isDefault?: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UpdateWalletSettingsActionReqType {}
