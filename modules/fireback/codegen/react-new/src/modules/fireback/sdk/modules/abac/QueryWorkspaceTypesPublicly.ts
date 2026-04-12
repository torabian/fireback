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
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action QueryWorkspaceTypesPublicly
 */
export type QueryWorkspaceTypesPubliclyActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type QueryWorkspaceTypesPubliclyActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<QueryWorkspaceTypesPubliclyActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  QueryWorkspaceTypesPubliclyActionOptions &
  Partial<{
    creatorFn: (item: unknown) => QueryWorkspaceTypesPubliclyActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useQueryWorkspaceTypesPubliclyActionQuery = (
  options: QueryWorkspaceTypesPubliclyActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return QueryWorkspaceTypesPubliclyAction.Fetch(
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
    queryKey: [QueryWorkspaceTypesPubliclyAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type QueryWorkspaceTypesPubliclyActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  QueryWorkspaceTypesPubliclyActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => QueryWorkspaceTypesPubliclyActionRes;
  }>;
export const useQueryWorkspaceTypesPubliclyAction = (
  options?: QueryWorkspaceTypesPubliclyActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return QueryWorkspaceTypesPubliclyAction.Fetch(
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
 * QueryWorkspaceTypesPubliclyAction
 */
export class QueryWorkspaceTypesPubliclyAction {
  //
  static URL = "/workspace/public/types";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(QueryWorkspaceTypesPubliclyAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<QueryWorkspaceTypesPubliclyActionRes>,
      unknown,
      unknown
    >(
      overrideUrl ?? QueryWorkspaceTypesPubliclyAction.NewUrl(qs),
      {
        method: QueryWorkspaceTypesPubliclyAction.Method,
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
      creatorFn?:
        | ((item: unknown) => QueryWorkspaceTypesPubliclyActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new QueryWorkspaceTypesPubliclyActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new QueryWorkspaceTypesPubliclyActionRes(item));
    const res = await QueryWorkspaceTypesPubliclyAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<QueryWorkspaceTypesPubliclyActionRes>();
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
    name: "QueryWorkspaceTypesPublicly",
    cliName: "public-types",
    url: "/workspace/public/types",
    method: "get",
    description:
      "Returns the workspaces types available in the project publicly without authentication, and the value could be used upon signup to go different route.",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "title",
          type: "string",
        },
        {
          name: "description",
          type: "string",
        },
        {
          name: "uniqueId",
          type: "string",
        },
        {
          name: "slug",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for queryWorkspaceTypesPubliclyActionRes
 **/
export class QueryWorkspaceTypesPubliclyActionRes {
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
  #description: string = "";
  /**
   *
   * @returns {string}
   **/
  get description() {
    return this.#description;
  }
  /**
   *
   * @type {string}
   **/
  set description(value: string) {
    this.#description = String(value);
  }
  setDescription(value: string) {
    this.description = value;
    return this;
  }
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
  #slug: string = "";
  /**
   *
   * @returns {string}
   **/
  get slug() {
    return this.#slug;
  }
  /**
   *
   * @type {string}
   **/
  set slug(value: string) {
    this.#slug = String(value);
  }
  setSlug(value: string) {
    this.slug = value;
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
    const d = data as Partial<QueryWorkspaceTypesPubliclyActionRes>;
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.description !== undefined) {
      this.description = d.description;
    }
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.slug !== undefined) {
      this.slug = d.slug;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      title: this.#title,
      description: this.#description,
      uniqueId: this.#uniqueId,
      slug: this.#slug,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      title: "title",
      description: "description",
      uniqueId: "uniqueId",
      slug: "slug",
    };
  }
  /**
   * Creates an instance of QueryWorkspaceTypesPubliclyActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: QueryWorkspaceTypesPubliclyActionResType) {
    return new QueryWorkspaceTypesPubliclyActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of QueryWorkspaceTypesPubliclyActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<QueryWorkspaceTypesPubliclyActionResType>,
  ) {
    return new QueryWorkspaceTypesPubliclyActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<QueryWorkspaceTypesPubliclyActionResType>,
  ): InstanceType<typeof QueryWorkspaceTypesPubliclyActionRes> {
    return new QueryWorkspaceTypesPubliclyActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof QueryWorkspaceTypesPubliclyActionRes> {
    return new QueryWorkspaceTypesPubliclyActionRes(this.toJSON());
  }
}
export abstract class QueryWorkspaceTypesPubliclyActionResFactory {
  abstract create(data: unknown): QueryWorkspaceTypesPubliclyActionRes;
}
/**
 * The base type definition for queryWorkspaceTypesPubliclyActionRes
 **/
export type QueryWorkspaceTypesPubliclyActionResType = {
  /**
   *
   * @type {string}
   **/
  title: string;
  /**
   *
   * @type {string}
   **/
  description: string;
  /**
   *
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {string}
   **/
  slug: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace QueryWorkspaceTypesPubliclyActionResType {}
