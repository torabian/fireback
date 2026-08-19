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
 * Action to communicate with the action SendNotification
 */
export type SendNotificationActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type SendNotificationActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  SendNotificationActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => SendNotificationActionRes;
  }>;
export const useSendNotificationAction = (
  options?: SendNotificationActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: SendNotificationActionReq) => {
    setCompleteState(false);
    return SendNotificationAction.Fetch(
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
 * SendNotificationAction
 */
export class SendNotificationAction {
  //
  static URL = "/notification/send";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(SendNotificationAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<SendNotificationActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<SendNotificationActionRes>,
      SendNotificationActionReq,
      unknown
    >(
      overrideUrl ?? SendNotificationAction.NewUrl(qs),
      {
        method: SendNotificationAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<SendNotificationActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => SendNotificationActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new SendNotificationActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new SendNotificationActionRes(item));
    const res = await SendNotificationAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<SendNotificationActionRes>();
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
    name: "SendNotification",
    cliName: "send",
    url: "/notification/send",
    method: "post",
    description:
      "Root-only. Sends a notification to one or more existing users at once - creates one NotificationEntity row per userId (all sharing the same title/body/senderId), rather than requiring one call per recipient. Any userId that doesn't correspond to an existing user is skipped (not a hard failure for the whole batch) and reported back in skippedUserIds.",
    in: {
      fields: [
        {
          name: "userIds",
          description: "UniqueIds of the existing users to notify.",
          type: "slice",
          primitive: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "title",
          description: "Short notification title.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "body",
          description: "Notification message body.",
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
          name: "sentCount",
          description: "Number of notifications actually created.",
          type: "int64",
        },
        {
          name: "skippedUserIds",
          description: "userIds skipped because no matching user exists.",
          type: "slice",
          primitive: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for sendNotificationActionReq
 **/
export class SendNotificationActionReq {
  /**
   * UniqueIds of the existing users to notify.
   * @type {string[]}
   **/
  #userIds: string[] = [];
  /**
   * UniqueIds of the existing users to notify.
   * @returns {string[]}
   **/
  get userIds() {
    return this.#userIds;
  }
  /**
   * UniqueIds of the existing users to notify.
   * @type {string[]}
   **/
  set userIds(value: string[]) {
    this.#userIds = value;
  }
  setUserIds(value: string[]) {
    this.userIds = value;
    return this;
  }
  /**
   * Short notification title.
   * @type {string}
   **/
  #title: string = "";
  /**
   * Short notification title.
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   * Short notification title.
   * @type {string}
   **/
  set title(value: string) {
    this.#title = String(value);
  }
  setTitle(value: string) {
    this.title = value;
    return this;
  }
  /**
   * Notification message body.
   * @type {string}
   **/
  #body: string = "";
  /**
   * Notification message body.
   * @returns {string}
   **/
  get body() {
    return this.#body;
  }
  /**
   * Notification message body.
   * @type {string}
   **/
  set body(value: string) {
    this.#body = String(value);
  }
  setBody(value: string) {
    this.body = value;
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
    const d = data as Partial<SendNotificationActionReq>;
    if (d.userIds !== undefined) {
      this.userIds = d.userIds;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.body !== undefined) {
      this.body = d.body;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      userIds: this.#userIds,
      title: this.#title,
      body: this.#body,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      userIds$: "userIds",
      get userIds() {
        return "userIds[:i]";
      },
      title: "title",
      body: "body",
    };
  }
  /**
   * Creates an instance of SendNotificationActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendNotificationActionReqType) {
    return new SendNotificationActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of SendNotificationActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SendNotificationActionReqType>) {
    return new SendNotificationActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendNotificationActionReqType>,
  ): InstanceType<typeof SendNotificationActionReq> {
    return new SendNotificationActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SendNotificationActionReq> {
    return new SendNotificationActionReq(this.toJSON());
  }
}
export abstract class SendNotificationActionReqFactory {
  abstract create(data: unknown): SendNotificationActionReq;
}
/**
 * The base type definition for sendNotificationActionReq
 **/
export type SendNotificationActionReqType = {
  /**
   * UniqueIds of the existing users to notify.
   * @type {string[]}
   **/
  userIds: string[];
  /**
   * Short notification title.
   * @type {string}
   **/
  title: string;
  /**
   * Notification message body.
   * @type {string}
   **/
  body: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SendNotificationActionReqType {}
/**
 * The base class definition for sendNotificationActionRes
 **/
export class SendNotificationActionRes {
  /**
   * Number of notifications actually created.
   * @type {number}
   **/
  #sentCount: number = 0;
  /**
   * Number of notifications actually created.
   * @returns {number}
   **/
  get sentCount() {
    return this.#sentCount;
  }
  /**
   * Number of notifications actually created.
   * @type {number}
   **/
  set sentCount(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#sentCount = parsedValue;
    }
  }
  setSentCount(value: number) {
    this.sentCount = value;
    return this;
  }
  /**
   * userIds skipped because no matching user exists.
   * @type {string[]}
   **/
  #skippedUserIds: string[] = [];
  /**
   * userIds skipped because no matching user exists.
   * @returns {string[]}
   **/
  get skippedUserIds() {
    return this.#skippedUserIds;
  }
  /**
   * userIds skipped because no matching user exists.
   * @type {string[]}
   **/
  set skippedUserIds(value: string[]) {
    this.#skippedUserIds = value;
  }
  setSkippedUserIds(value: string[]) {
    this.skippedUserIds = value;
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
    const d = data as Partial<SendNotificationActionRes>;
    if (d.sentCount !== undefined) {
      this.sentCount = d.sentCount;
    }
    if (d.skippedUserIds !== undefined) {
      this.skippedUserIds = d.skippedUserIds;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      sentCount: this.#sentCount,
      skippedUserIds: this.#skippedUserIds,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      sentCount: "sentCount",
      skippedUserIds$: "skippedUserIds",
      get skippedUserIds() {
        return "skippedUserIds[:i]";
      },
    };
  }
  /**
   * Creates an instance of SendNotificationActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendNotificationActionResType) {
    return new SendNotificationActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of SendNotificationActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SendNotificationActionResType>) {
    return new SendNotificationActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendNotificationActionResType>,
  ): InstanceType<typeof SendNotificationActionRes> {
    return new SendNotificationActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SendNotificationActionRes> {
    return new SendNotificationActionRes(this.toJSON());
  }
}
export abstract class SendNotificationActionResFactory {
  abstract create(data: unknown): SendNotificationActionRes;
}
/**
 * The base type definition for sendNotificationActionRes
 **/
export type SendNotificationActionResType = {
  /**
   * Number of notifications actually created.
   * @type {number}
   **/
  sentCount: number;
  /**
   * userIds skipped because no matching user exists.
   * @type {string[]}
   **/
  skippedUserIds: string[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SendNotificationActionResType {}
