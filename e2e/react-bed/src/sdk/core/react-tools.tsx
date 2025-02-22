// @ts-nocheck

import { ExecApi, IResponse, RemoteRequestOption, Query } from "./http-tools";
import React, {
  useContext,
  useState,
  useEffect,
  Dispatch,
  SetStateAction,
  useRef,
} from "react";
import { Upload } from "tus-js-client";
import { QueryClient, UseQueryOptions } from "react-query";

/**
 * Removes the workspace id which is default present everywhere
 * @param options
 * @returns
 */
export function noWorkspaceQuery(options: any) {
  return {
    ...options,
    headers: {
      ...options.headers,
      ["workspace-id"]: "",
    },
  };
}

export interface PatchProps {
  queryClient: QueryClient;
  query?: any;
  execFnOverride?: any;
}

export interface DeleteProps {
  queryClient?: QueryClient;
  execFnOverride?: any;
  query?: any;
}

export interface UseRemoteQuery {
  query?: Query;
  queryClient?: QueryClient;
  execFnOverride?: any;
  queryOptions?: UseQueryOptions<any>;
  unauthorized?: boolean;
  UseRemoteQuery?: (options: any) => any;
  optionFn?: (data: RemoteRequestOption) => any;
}

export interface ContextSession {
  token?: string;
}

export interface CapabilityEntity {
  uniqueId: string;
}

export interface CapabilityChild {
  uniqueId: string;
  children: CapabilityChild[];
}

export interface CapabilitiesResult {
  capabilities: CapabilityEntity[];
  nested: CapabilityChild[];
}

export interface UserEntity {
  firstName: string;
  lastName: string;
  photo: string;
  uniqueId: string;
}

export interface Token {
  user: UserEntity | undefined;
  userID?: string | undefined;
  validUntil: string;
  uniqueId: string;
}

export interface Preference {
  itemKey: string;
  value: string;
  valueType: string;
  scope: string;
  user: UserEntity | undefined;
  userID?: string | undefined;
}

interface RoleEntity {
  uniqueId: string;
  name: string;
  parentRoleId?: string | undefined;
  parentRole: RoleEntity | undefined;
  capabilitiesListId: string[];
  capabilities: CapabilityEntity[];
  workspace: WorkspaceEntity | undefined;
  workspaceId?: string | undefined;
}

interface WorkspaceEntity {
  name: string;
  description: string;
  uniqueId: string;
}

interface UserRoleWorkspace {
  workspace: WorkspaceEntity | undefined;
  workspaceId?: string | undefined;
  role: RoleEntity | undefined;
  roleId?: string | undefined;
  user: UserEntity | undefined;
  userId?: string | undefined;
  uniqueId: string;
}

export interface AuthContext {
  workspaceId: string;
  token: string;
  capabilities: string[];
}

export interface IRemoteQueryContext {
  setSession: (session: ContextSession) => void;
  options: RemoteRequestOption;
  session: ContextSession;
  checked: boolean;
  isAuthenticated: boolean;
  selectedUrw?: UserRoleWorkspace;
  signout: () => void;
  selectUrw: (urw: UserRoleWorkspace) => void;
  activeUploads: ActiveUpload[];
  execFn: (options: RemoteRequestOption) => void;
  setActiveUploads: Dispatch<SetStateAction<ActiveUpload[]>>;
}

export interface ActiveUpload {
  uploadId: string;
  bytesSent: number;
  bytesTotal: number;
  filename?: string;
}

export const RemoteQueryContext = React.createContext<IRemoteQueryContext>({
  setSession(session: ContextSession) {},
  options: {},
} as any);

export function useFileUploader() {
  const { session, selectedWorkspace, activeUploads, setActiveUploads } =
    useContext(RemoteQueryContext);
  // const [activeUploads, setActiveUploads] = useState<ActiveUpload[]>([]);

  const upload = (files: File[]): Promise<string>[] => {
    const result = files.map((file) => {
      return new Promise((resolve: (t: string) => void) => {
        const upload = new Upload(file, {
          endpoint: "http://localhost:51230/files/",
          onBeforeRequest(req: any) {
            req.setHeader("authorization", session.token);
            req.setHeader("workspace-id", selectedWorkspace);
          },
          headers: {
            // authorization: authorization,
          },
          metadata: {
            filename: file.name,
            path: "/database/users",
            filetype: file.type,
          },
          onSuccess() {
            const uploadId = upload.url?.match(/([a-z0-9]){10,}/gi);
            resolve(`${uploadId}`);
          },

          onProgress(bytesSent, bytesTotal) {
            const uploadId = upload.url?.match(/([a-z0-9]){10,}/gi)?.toString();
            let updated = false;

            setActiveUploads((items) =>
              items?.map((item) => {
                if (item.uploadId === uploadId) {
                  updated = true;
                  return {
                    uploadId,
                    bytesSent,
                    filename: file.name,
                    bytesTotal,
                  };
                }

                return item;
              })
            );

            if (!updated && uploadId) {
              setActiveUploads((activeUploads) => [
                ...activeUploads,
                { uploadId, bytesSent, bytesTotal, filename: file.name },
              ]);
            }
            console.log(bytesSent, bytesTotal);
          },
        });

        upload.start();
      });
    });

    return result;
  };

  return { upload, activeUploads };
}

export class ReactNativeStorage {
  async setItem(key, value) {}
  async getItem(key) {}
  async removeItem(key) {}
}

/**
 * Kinda module agnostic storage definition,
 * use it to create react native or other platform
 * storage system
 */
export interface CredentialStorage {
  setItem(key, value);
  getItem(key);
  removeItem(key);
}

export class WebStorage implements CredentialStorage {
  async setItem(key, value) {
    return localStorage.setItem(key, value);
  }
  async getItem(key) {
    return localStorage.getItem(key);
  }
  async removeItem(key) {
    return localStorage.removeItem(key);
  }
}

async function saveSession(
  identifier: string,
  session: ContextSession,
  storagex: CredentialStorage
) {
  storagex.setItem("fb_microservice_" + identifier, JSON.stringify(session));
}

function saveWorkspace(
  identifier: string,
  workspaceId: UserRoleWorkspace,
  storagex: CredentialStorage
) {
  storagex.setItem(
    "fb_selected_workspace_" + identifier,
    JSON.stringify(workspaceId)
  );
}

async function getSession(identifier: string, storagex: CredentialStorage) {
  let data = null;
  try {
    data = JSON.parse(await storagex.getItem("fb_microservice_" + identifier));
  } catch (err) {}
  return data;
}

async function getWorkspace(
  identifier: string,
  storagex: CredentialStorage
): UserRoleWorkspace | undefined {
  let data = null;
  try {
    data = JSON.parse(
      await storagex.getItem("fb_selected_workspace_" + identifier)
    );
  } catch (err) {}
  return data;
}

interface IRemoteQueryProvider {
  /**
   * Rest of the application code, which will have access to sdk
   * Make sure the react query provider is outside, and pass the queryclient
   * via queryClient prop
   */
  children?: React.ReactNode;

  /**
   * Address of the web server is running. You can change the value based on development or production
   * environment. Make sure you'll have trailing slash, for example remote="http://localhost:4500/"
   * Fireback won't add slash if you do not provide.
   */
  remote?: string;

  /**
   * Force the accept-language for each request. Fireback might have some javascript code to detect
   * the accept language at best it can, but if user has chosen a language then you can
   * put it here to avoid auto detect or going empty.
   */
  preferredAcceptLanguage?: string;

  /**
   * unique identifier of the sdk upon saving the caches to the localstorage.
   * this is required, leave your app name. It's to make sure different apps on
   * development environment would never collide.
   */
  identifier: string;

  /**
   * SDK can keep track of user role workspace, so role-id, workspace-id will be sent
   * with headers automatically. After authentication, make sure you keep that information,
   * and if you give the option for users to select their current workspace and role,
   * update the selectedUrw object as well.
   */
  selectedUrw?: UserRoleWorkspace;

  /**
   * Fireback does not keep the token or provide authentication function.
   * After authenticating the client using hooks such as usePostPassportsSigninClassic,
   * make sure you store the token in localstorage, and provide it using this prop
   *
   * This is due to fact, multiple sdks can be present in a single app, and you need to use the
   * same token often will all of them.
   *
   * token might be not required, if the app is fully public, and if there will be a specific function to
   * authentication for sdk, just used as forced.
   */
  token?: string;

  /**
   * Same object that you provide to QueryClientProvider, will store caches and other
   * necessary things with react-query
   */
  queryClient?: QueryClient;

  /**
   * Allows developer to override the http call function. Basically all http requests
   * going through a single function, and you can override it, for cases such a mock server:
   *
   * defaultExecFn={() => {
   *  return (options: any) => mockExecFn(options, mockServer.current);
   * }}
   */
  defaultExecFn?: any;

  /**
   * For applications using socket (Fireback reactive), it would subscribe the the socket server.
   */
  socket?: boolean;

  /**
   * CredentialStorage is class which can provide a way to store the data in localstorage.
   * By default, it's a WebStorage, for react native, you can implement the interface,
   * and for example use async storage library instead.
   */
  credentialStorage?: CredentialStorage;

  /**
   * Prefixes all of the api addresses with this string. Note that it will be added after remote,
   * and before the custom API call: remote + prefix + /my/another/function
   */
  prefix?: string;
}

export function RemoteQueryProvider({
  children,
  remote,
  selectedUrw,
  identifier,
  token,
  preferredAcceptLanguage,
  queryClient,
  defaultExecFn,
  socket,
  credentialStorage,
  prefix,
}: IRemoteQueryProvider) {
  const [checked, setChecked] = useState(false);
  const [session, setSession$] = useState<ContextSession>();
  const [selectedWorkspaceInternal, selectWorkspace$] =
    useState<UserRoleWorkspace>();

  const storage = useRef(
    credentialStorage ? credentialStorage : new WebStorage()
  );

  const beginPreCatch = async () => {
    const workspace = await getWorkspace(identifier, storage.current);
    const session = await getSession(identifier, storage.current);

    selectWorkspace$(workspace);
    setSession$(session);
    setChecked(true);
  };

  useEffect(() => {
    beginPreCatch();
  }, []);

  const [activeUploads, setActiveUploads] = useState<ActiveUpload[]>([]);

  const [execFn, setExecFn] = useState<ExecApi>(defaultExecFn);

  const isAuthenticated = !!session;

  const selectUrw = (urw: UserRoleWorkspace) => {
    saveWorkspace(identifier, urw, storage.current);
    selectWorkspace$(urw);
  };

  const setSession = (session: ContextSession) => {
    setSession$(() => {
      saveSession(identifier, session, storage.current);
      return session;
    });
  };

  const options = {
    headers: {
      authorization: token || session?.token,
    },
    prefix: remote + (prefix || ""),
  };

  if (selectedWorkspaceInternal) {
    options.headers["workspace-id"] = selectedWorkspaceInternal.workspaceId;
    options.headers["role-id"] = selectedWorkspaceInternal.roleId;
  } else if (selectedUrw) {
    options.headers["workspace-id"] = selectedUrw.workspaceId;
    options.headers["role-id"] = selectedUrw.roleId;
  } else if (session?.userWorkspaces && session.userWorkspaces.length > 0) {
    const sess2 = session.userWorkspaces[0];
    options.headers["workspace-id"] = sess2.workspaceId;
    options.headers["role-id"] = sess2.roleId;
  }

  if (preferredAcceptLanguage) {
    options.headers["accept-language"] = preferredAcceptLanguage;
  }

  useEffect(() => {
    if (token) {
      setSession$({
        ...(session || {}),
        token,
      });
    }
  }, [token]);

  const signout = () => {
    setSession$(null);
    storage.current?.removeItem("fb_microservice_" + identifier);
    selectUrw(undefined);
  };

  const discardActiveUploads = () => {
    setActiveUploads([]);
  };

  const { socketState } = useSocket(
    remote,
    options.headers?.authorization,
    (options.headers as any)["workspace-id"],
    queryClient
  );

  return (
    <RemoteQueryContext.Provider
      value={{
        options,
        signout,
        setSession,
        socketState,
        checked,
        selectedUrw: selectedWorkspaceInternal,
        selectUrw,
        session,
        preferredAcceptLanguage,
        activeUploads,
        setActiveUploads,
        execFn,
        setExecFn,
        discardActiveUploads,
        isAuthenticated,
      }}
    >
      {children}
    </RemoteQueryContext.Provider>
  );
}

export interface PossibleStoreData<T> {
  data: IResponse<T>;
  jsonQuery: string;
}

export function useSocket(remote, token, workspaceId, queryClient) {
  const [socketState, setSocketState] = useState({ state: "unknown" });

  useEffect(() => {
    if (!remote || process.env.REACT_APP_INACCURATE_MOCK_MODE == "true") {
      return;
    }
    const wsRemote = remote.replace("https", "wss").replace("http", "ws");
    let conn: WebSocket;
    try {
      conn = new WebSocket(
        `${wsRemote}ws?token=${token}&workspaceId=${workspaceId}`
      );
      conn.onerror = function (evt) {
        console.log("Closed", evt);
        setSocketState({ state: "error" });
      };
      conn.onclose = function (evt) {
        setSocketState({ state: "closed" });
      };
      conn.onmessage = function (evt: any) {
        try {
          const msg = JSON.parse(evt.data);

          if (msg?.data.entityKey) {
            queryClient.invalidateQueries(msg?.data.entityKey);
          }
        } catch (e: any) {
          console.log(evt);
        }
      };
      conn.onopen = function (evt) {
        setSocketState({ state: "connected" });
      };
    } catch (err) {}

    return () => {
      if (conn?.readyState === 1) {
        conn.close();
      }
    };
  }, [token, workspaceId]);

  return { socketState };
}

export function queryBeforeSend(query: any) {
  if (!query) {
    return {};
  }

  const newQuery = {};

  if (query.startIndex) {
    newQuery.startIndex = query.startIndex;
  }
  if (query.itemsPerPage) {
    newQuery.itemsPerPage = query.itemsPerPage;
  }
  if (query.query) {
    newQuery.query = query.query;
  }
  if (query.deep) {
    newQuery.deep = query.deep;
  }
  if (query.jsonQuery) {
    newQuery.jsonQuery = JSON.stringify(query.jsonQuery);
  }
  if (query.withPreloads) {
    newQuery.withPreloads = query.withPreloads;
  }
  if (query.uniqueId) {
    newQuery.uniqueId = query.uniqueId;
  }
  if (query.sort) {
    newQuery.sort = query.sort;
  }

  return newQuery;
}
