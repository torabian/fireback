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
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action MyNotifications
 */
export type MyNotificationsActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type MyNotificationsActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<MyNotificationsActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  MyNotificationsActionOptions &
  Partial<{
    creatorFn: (item: unknown) => MyNotificationsActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useMyNotificationsActionQuery = (
  options: MyNotificationsActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return MyNotificationsAction.Fetch(
      {
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
  const result = useQuery({
    queryKey: [MyNotificationsAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type MyNotificationsActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  MyNotificationsActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => MyNotificationsActionRes;
  }>;
export const useMyNotificationsAction = (
  options?: MyNotificationsActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return MyNotificationsAction.Fetch(
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
 * MyNotificationsAction
 */
export class MyNotificationsAction {
  //
  static URL = "/notification/mine";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(MyNotificationsAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<MyNotificationsActionRes>, unknown, unknown>(
      overrideUrl ?? MyNotificationsAction.NewUrl(qs),
      {
        method: MyNotificationsAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => MyNotificationsActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new MyNotificationsActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new MyNotificationsActionRes(item));
    const res = await MyNotificationsAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<MyNotificationsActionRes>();
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
    name: "MyNotifications",
    cliName: "mine",
    url: "/notification/mine",
    method: "get",
    description:
      "Self-service. Lists the calling user's own notifications (whichever workspace they're currently in - notifications aren't workspace-scoped from the recipient's side), most recent first. See SendNotificationAction for how they're created and MarkNotificationReadAction to mark one read.",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "uniqueId",
          type: "string",
        },
        {
          name: "title",
          type: "string",
        },
        {
          name: "body",
          type: "string",
        },
        {
          name: "isRead",
          type: "bool",
        },
        {
          name: "createdAt",
          description: "RFC3339 timestamp this notification was sent at.",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for myNotificationsActionRes
 **/
export class MyNotificationsActionRes {
  /**
   *
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   *
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   *
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
  #title: string = "";
  /**
   *
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   *
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
   *
   * @type {string}
   **/
  #body: string = "";
  /**
   *
   * @returns {string}
   **/
  get body() {
    return this.#body;
  }
  /**
   *
   * @type {string}
   **/
  set body(value: string) {
    this.#body = String(value);
  }
  setBody(value: string) {
    this.body = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #isRead!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get isRead() {
    return this.#isRead;
  }
  /**
   *
   * @type {boolean}
   **/
  set isRead(value: boolean) {
    this.#isRead = Boolean(value);
  }
  setIsRead(value: boolean) {
    this.isRead = value;
    return this;
  }
  /**
   * RFC3339 timestamp this notification was sent at.
   * @type {string}
   **/
  #createdAt: string = "";
  /**
   * RFC3339 timestamp this notification was sent at.
   * @returns {string}
   **/
  get createdAt() {
    return this.#createdAt;
  }
  /**
   * RFC3339 timestamp this notification was sent at.
   * @type {string}
   **/
  set createdAt(value: string) {
    this.#createdAt = String(value);
  }
  setCreatedAt(value: string) {
    this.createdAt = value;
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
    const d = data as Partial<MyNotificationsActionRes>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.body !== undefined) {
      this.body = d.body;
    }
    if (d.isRead !== undefined) {
      this.isRead = d.isRead;
    }
    if (d.createdAt !== undefined) {
      this.createdAt = d.createdAt;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      title: this.#title,
      body: this.#body,
      isRead: this.#isRead,
      createdAt: this.#createdAt,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      title: "title",
      body: "body",
      isRead: "isRead",
      createdAt: "createdAt",
    };
  }
  /**
   * Creates an instance of MyNotificationsActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: MyNotificationsActionResType) {
    return new MyNotificationsActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of MyNotificationsActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<MyNotificationsActionResType>) {
    return new MyNotificationsActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<MyNotificationsActionResType>,
  ): InstanceType<typeof MyNotificationsActionRes> {
    return new MyNotificationsActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof MyNotificationsActionRes> {
    return new MyNotificationsActionRes(this.toJSON());
  }
}
export abstract class MyNotificationsActionResFactory {
  abstract create(data: unknown): MyNotificationsActionRes;
}
/**
 * The base type definition for myNotificationsActionRes
 **/
export type MyNotificationsActionResType = {
  /**
   *
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {string}
   **/
  title: string;
  /**
   *
   * @type {string}
   **/
  body: string;
  /**
   *
   * @type {boolean}
   **/
  isRead: boolean;
  /**
   * RFC3339 timestamp this notification was sent at.
   * @type {string}
   **/
  createdAt: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MyNotificationsActionResType {}
