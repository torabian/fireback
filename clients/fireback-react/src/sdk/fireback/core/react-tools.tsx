// @ts-nocheck

import { ExecApi, IResponse, RemoteRequestOption } from "./http-tools";
import React, {
  useContext,
  useState,
  useEffect,
  Dispatch,
  SetStateAction,
} from "react";
import { Upload } from "tus-js-client";
import { QueryClient } from "react-query";

export interface Query {
  userId?: string | null;
  uniqueId?: string | null;
}

/**
 * Removes the workspace id which is default present everywhere
 * @param options
 * @returns
 */
export function noWorkspaceQuery(options) {
  return {
    ...options,
    headers: {
      ...options.headers,
      ["workspace-id"]: "",
    },
  };
}

export interface PatchProps {
  queryClient: QueryClient, query?: any, execFnOverride?: any
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
  optionFn?: (data: RemoteRequestOption) => any,
}

export interface ContextSession {
  token?: string;
}

export interface CapabilityEntity {
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:2500;autoIncrement:false") */
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
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;") */
  uniqueId: string;
}

export interface Token {
  /** @tag(gorm:"foreignKey:UserID;references:UniqueId") */
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
  /** @tag(gorm:"foreignKey:UserID;references:UniqueId") */
  user: UserEntity | undefined;
  userID?: string | undefined;
}

interface RoleEntity {
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;") */
  uniqueId: string;
  /** @tag(polyglot:"name") */
  name: string;
  parentRoleId?: string | undefined;
  parentRole: RoleEntity | undefined;
  /** @tag(gorm:"-" sql:"-") */
  capabilitiesListId: string[];
  /** @tag(gorm:"many2many:role_capability;foreignKey:UniqueId;references:UniqueId" json:"capabilities") */
  capabilities: CapabilityEntity[];
  /** @tag(gorm:"foreignKey:WorkspaceId;references:UniqueId" json:"-") */
  workspace: WorkspaceEntity | undefined;
  /** @tag(json:"workspaceId" gorm:"size:100;") */
  workspaceId?: string | undefined;
}

interface WorkspaceEntity {
  name: string;
  description: string;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;") */
  uniqueId: string;
}

interface PendingWorkspaceInvite {
  id: number;
  email: string;
}

interface AcceptInviteDto {
  inviteUniqueId: string;
}

interface AssignRoleDto {
  /** @tag(validate:"required") */
  roleId: string;
  /** @tag(validate:"required") */
  userId: string;
}

interface WorkspaceInviteEntity {
  coverLetter: string;
  targetUserLocale: string;
  email: string;
  phoneNumber: string;
  /** @tag(gorm:"foreignKey:WorkspaceId;references:UniqueId" json:"-") */
  workspace: WorkspaceEntity | undefined;
  /** @tag(gorm:"size:100;") */
  workspaceId?: string | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;") */
  uniqueId: string;
  /** @tag(gorm:"foreignKey:RoleId;references:UniqueId") */
  role: RoleEntity | undefined;
  /** @tag(gorm:"size:100;") */
  roleId?: string | undefined;
  /** @tag(validate:"required") */
  firstName: string;
  /** @tag(validate:"required") */
  lastName: string;
}

/** Used when creating a new workspace */
interface WorkspaceInviteDto {
  /** @tag(validate:"required") */
  email: string;
  /** @tag(validate:"required") */
  roleId: string;
}

interface UserRoleWorkspace {
  /** @tag(gorm:"foreignKey:WorkspaceId;references:UniqueId") */
  workspace: WorkspaceEntity | undefined;
  /** @tag(gorm:"size:100;") */
  workspaceId?: string | undefined;
  /** @tag(gorm:"foreignKey:RoleId;references:UniqueId") */
  role: RoleEntity | undefined;
  /** @tag(gorm:"size:100;") */
  roleId?: string | undefined;
  /** @tag(gorm:"foreignKey:UserId;references:UniqueId") */
  user: UserEntity | undefined;
  /** @tag(gorm:"size:100;") */
  userId?: string | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;") */
  uniqueId: string;
}

/** Use this when sending information about an specific workspace */
interface WorkspaceDto {
  relations: UserRoleWorkspace[];
}

/** Dump every workspce settings you might think here */
interface WorkspaceConfigEntity {
  /** @tag(gorm:"foreignKey:WorkspaceId;references:UniqueId") */
  workspace: WorkspaceEntity | undefined;
  /** @tag(gorm:"size:100;") */
  workspaceId?: string | undefined;
  zoomClientId?: string | undefined;
  zoomClientSecret?: string | undefined;
  allowPublicToJoinTheWorkspace?: boolean | undefined;
}

interface UserProfileEntity {
  firstName?: string | undefined;
  lastName?: string | undefined;
}

interface ExchangeKeyInformationDto {
  key: string;
}

interface UserAccessLevel {
  /** @tag(json:"capabilities") */
  capabilities: string[];
  /** @tag(json:"workspaces") */
  workspaces: string[];
  /** @tag(json:"sql") */
  SQL: string;
}

interface AuthResult {
  /** @tag(json:"workspaceId") */
  workspaceId: string;
  /** @tag(json:"internalSql") */
  internalSql: string;
  /** @tag(json:"userId") */
  userId: string;
  /** @tag(json:"user") */
  user: UserEntity | undefined;
  /** @tag(json:"accessLevel") */
  accessLevel?: UserAccessLevel | undefined;
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
});

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

function saveSession(identifier: string, session: ContextSession) {
  localStorage.setItem(
    "fb_microservice_" + identifier,
    JSON.stringify(session)
  );
}

function saveWorkspace(identifier: string, workspaceId: UserRoleWorkspace) {
  localStorage.setItem(
    "fb_selected_workspace_" + identifier,
    JSON.stringify(workspaceId)
  );
}

function getSession(identifier: string) {
  let data = null;
  try {
    data = JSON.parse(localStorage.getItem("fb_microservice_" + identifier));
  } catch (err) {}
  return data;
}

function getWorkspace(identifier: string): UserRoleWorkspace | undefined {
  let data = null;
  try {
    data = JSON.parse(
      localStorage.getItem("fb_selected_workspace_" + identifier)
    );
  } catch (err) {}
  return data;
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
}: {
  children: React.ReactNode;
  remote?: string;
  preferredAcceptLanguage?: string;
  identifier: string;
  selectedUrw?: UserRoleWorkspace;
  token?: string;
  queryClient?: QueryClient;
  defaultExecFn?: any;
  socket?: boolean;
}) {
  const [session, setSession$] = useState<ContextSession>(
    getSession(identifier)
  );
  const [selectedWorkspaceInternal, selectWorkspace$] =
    useState<UserRoleWorkspace>(getWorkspace(identifier));

  const [activeUploads, setActiveUploads] = useState<ActiveUpload[]>([]);

  const [execFn, setExecFn] = useState<ExecApi>(defaultExecFn);

  const isAuthenticated = !!session;

  const selectUrw = (urw: UserRoleWorkspace) => {
    saveWorkspace(identifier, urw);
    selectWorkspace$(urw);
  };

  const setSession = (session: ContextSession) => {
    setSession$(() => {
      saveSession(identifier, session);
      return session;
    });
  };

  const options = {
    headers: {
      authorization: session?.token,
    },
    prefix: remote,
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
    localStorage.removeItem("fb_microservice_" + identifier);
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
