import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import { UserSessionDto } from "./UserSessionDto";
import { buildUrl } from "../../sdk/common/buildUrl";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action OsLoginAuthenticate
 */
export type OsLoginAuthenticateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type OsLoginAuthenticateActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, UserSessionDto, unknown[]>,
  "queryKey"
> &
  OsLoginAuthenticateActionOptions &
  Partial<{
    creatorFn: (item: unknown) => UserSessionDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useOsLoginAuthenticateActionQuery = (
  options: OsLoginAuthenticateActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return OsLoginAuthenticateAction.Fetch(
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
    queryKey: [OsLoginAuthenticateAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type OsLoginAuthenticateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  OsLoginAuthenticateActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserSessionDto;
  }>;
export const useOsLoginAuthenticateAction = (
  options?: OsLoginAuthenticateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return OsLoginAuthenticateAction.Fetch(
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
 * OsLoginAuthenticateAction
 */
export class OsLoginAuthenticateAction {
  //
  static URL = "/passports/os/login";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(OsLoginAuthenticateAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<UserSessionDto, unknown, unknown>(
      overrideUrl ?? OsLoginAuthenticateAction.NewUrl(qs),
      {
        method: OsLoginAuthenticateAction.Method,
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
      creatorFn?: ((item: unknown) => UserSessionDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new UserSessionDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new UserSessionDto(item));
    const res = await OsLoginAuthenticateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "OsLoginAuthenticate",
    cliName: "oslogin",
    url: "/passports/os/login",
    method: "get",
    description:
      "Logins into the system using operating system (current) user, and store the information for them. Useful for desktop applications.",
    out: {
      dto: "UserSessionDto",
    },
  };
}
