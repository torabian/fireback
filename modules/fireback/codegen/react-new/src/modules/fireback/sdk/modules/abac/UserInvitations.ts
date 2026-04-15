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
 * Action to communicate with the action UserInvitations
 */
export type UserInvitationsActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type UserInvitationsActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<UserInvitationsActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  UserInvitationsActionOptions &
  Partial<{
    creatorFn: (item: unknown) => UserInvitationsActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useUserInvitationsActionQuery = (
  options: UserInvitationsActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return UserInvitationsAction.Fetch(
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
    queryKey: [UserInvitationsAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type UserInvitationsActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserInvitationsActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserInvitationsActionRes;
  }>;
export const useUserInvitationsAction = (
  options?: UserInvitationsActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return UserInvitationsAction.Fetch(
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
 * UserInvitationsAction
 */
export class UserInvitationsAction {
  //
  static URL = "/users/invitations";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(UserInvitationsAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserInvitationsActionRes>, unknown, unknown>(
      overrideUrl ?? UserInvitationsAction.NewUrl(qs),
      {
        method: UserInvitationsAction.Method,
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
      creatorFn?: ((item: unknown) => UserInvitationsActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new UserInvitationsActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new UserInvitationsActionRes(item));
    const res = await UserInvitationsAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<UserInvitationsActionRes>();
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
    name: "UserInvitations",
    url: "/users/invitations",
    method: "get",
    description:
      "Shows the invitations for an specific user, if the invited member already has a account. It's based on the passports, so if the passport is authenticated we will show them.",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "userId",
          description: "UserUniqueId",
          type: "string",
        },
        {
          name: "uniqueId",
          description: "Invitation unique id",
          type: "string",
        },
        {
          name: "value",
          description: "The value of the passport (email/phone)",
          type: "string",
        },
        {
          name: "roleName",
          description: "Name of the role that user will get",
          type: "string",
        },
        {
          name: "workspaceName",
          description: "Name of the workspace which user is invited to.",
          type: "string",
        },
        {
          name: "type",
          description: "The method of the invitation, such as email.",
          type: "string",
        },
        {
          name: "coverLetter",
          description:
            "The content that user will receive to understand the reason of the letter.",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for userInvitationsActionRes
 **/
export class UserInvitationsActionRes {
  /**
   * UserUniqueId
   * @type {string}
   **/
  #userId: string = "";
  /**
   * UserUniqueId
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * UserUniqueId
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
   * Invitation unique id
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * Invitation unique id
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * Invitation unique id
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
   * The value of the passport (email/phone)
   * @type {string}
   **/
  #value: string = "";
  /**
   * The value of the passport (email/phone)
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   * The value of the passport (email/phone)
   * @type {string}
   **/
  set value(value: string) {
    this.#value = String(value);
  }
  setValue(value: string) {
    this.value = value;
    return this;
  }
  /**
   * Name of the role that user will get
   * @type {string}
   **/
  #roleName: string = "";
  /**
   * Name of the role that user will get
   * @returns {string}
   **/
  get roleName() {
    return this.#roleName;
  }
  /**
   * Name of the role that user will get
   * @type {string}
   **/
  set roleName(value: string) {
    this.#roleName = String(value);
  }
  setRoleName(value: string) {
    this.roleName = value;
    return this;
  }
  /**
   * Name of the workspace which user is invited to.
   * @type {string}
   **/
  #workspaceName: string = "";
  /**
   * Name of the workspace which user is invited to.
   * @returns {string}
   **/
  get workspaceName() {
    return this.#workspaceName;
  }
  /**
   * Name of the workspace which user is invited to.
   * @type {string}
   **/
  set workspaceName(value: string) {
    this.#workspaceName = String(value);
  }
  setWorkspaceName(value: string) {
    this.workspaceName = value;
    return this;
  }
  /**
   * The method of the invitation, such as email.
   * @type {string}
   **/
  #type: string = "";
  /**
   * The method of the invitation, such as email.
   * @returns {string}
   **/
  get type() {
    return this.#type;
  }
  /**
   * The method of the invitation, such as email.
   * @type {string}
   **/
  set type(value: string) {
    this.#type = String(value);
  }
  setType(value: string) {
    this.type = value;
    return this;
  }
  /**
   * The content that user will receive to understand the reason of the letter.
   * @type {string}
   **/
  #coverLetter: string = "";
  /**
   * The content that user will receive to understand the reason of the letter.
   * @returns {string}
   **/
  get coverLetter() {
    return this.#coverLetter;
  }
  /**
   * The content that user will receive to understand the reason of the letter.
   * @type {string}
   **/
  set coverLetter(value: string) {
    this.#coverLetter = String(value);
  }
  setCoverLetter(value: string) {
    this.coverLetter = value;
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
    const d = data as Partial<UserInvitationsActionRes>;
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.roleName !== undefined) {
      this.roleName = d.roleName;
    }
    if (d.workspaceName !== undefined) {
      this.workspaceName = d.workspaceName;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.coverLetter !== undefined) {
      this.coverLetter = d.coverLetter;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      userId: this.#userId,
      uniqueId: this.#uniqueId,
      value: this.#value,
      roleName: this.#roleName,
      workspaceName: this.#workspaceName,
      type: this.#type,
      coverLetter: this.#coverLetter,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      userId: "userId",
      uniqueId: "uniqueId",
      value: "value",
      roleName: "roleName",
      workspaceName: "workspaceName",
      type: "type",
      coverLetter: "coverLetter",
    };
  }
  /**
   * Creates an instance of UserInvitationsActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserInvitationsActionResType) {
    return new UserInvitationsActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of UserInvitationsActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserInvitationsActionResType>) {
    return new UserInvitationsActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserInvitationsActionResType>,
  ): InstanceType<typeof UserInvitationsActionRes> {
    return new UserInvitationsActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserInvitationsActionRes> {
    return new UserInvitationsActionRes(this.toJSON());
  }
}
export abstract class UserInvitationsActionResFactory {
  abstract create(data: unknown): UserInvitationsActionRes;
}
/**
 * The base type definition for userInvitationsActionRes
 **/
export type UserInvitationsActionResType = {
  /**
   * UserUniqueId
   * @type {string}
   **/
  userId: string;
  /**
   * Invitation unique id
   * @type {string}
   **/
  uniqueId: string;
  /**
   * The value of the passport (email/phone)
   * @type {string}
   **/
  value: string;
  /**
   * Name of the role that user will get
   * @type {string}
   **/
  roleName: string;
  /**
   * Name of the workspace which user is invited to.
   * @type {string}
   **/
  workspaceName: string;
  /**
   * The method of the invitation, such as email.
   * @type {string}
   **/
  type: string;
  /**
   * The content that user will receive to understand the reason of the letter.
   * @type {string}
   **/
  coverLetter: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserInvitationsActionResType {}
