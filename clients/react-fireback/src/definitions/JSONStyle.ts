export interface IResponseData<T> {
  kind?: string;
  fields?: string;
  etag?: string;
  id?: string;
  lang?: string;
  updated?: string;
  deleted?: boolean;
  currentItemCount?: number;
  itemsPerPage?: number;
  startIndex?: number;
  totalItems?: number;
  pageIndex?: number;
  totalPages?: number;
  items?: Array<T>;
}
export interface IResponseErrorItem {
  domain?: string;
  reason?: string;
  message?: string;
  location?: string;
  locationType?: string;
  extendedHelp?: string;
  sendReport?: string;
}
export interface IResponseError {
  code?: Number;
  message: string;
  errors?: Array<IResponseErrorItem>;
}

export interface IResponse<T> {
  apiVersion?: string;
  context?: string;
  id?: string;
  params?: {
    id?: string;
  };
  data?: T;
  error?: IResponseError;
}

export interface IResponseList<T> {
  apiVersion?: string;
  context?: string;
  id?: string;
  params?: {
    id?: string;
  };
  data?: IResponseData<T>;
  error?: IResponseError;
}

export type IDeleteResponse = IResponse<{ rowsAffected: number }>;

export interface RemoteRequestOption {
  headers?: {
    [key: string]: string;
  };
  prefix?: string;
}

export type ExecApi = (
  method: "post" | "get" | "put" | "delete" | "patch",
  affix: string,
  body?: any,
  headers?: any
) => Promise<any>;

export const execApiFn =
  (options: RemoteRequestOption) =>
  (
    method: "post" | "get" | "put" | "delete" | "patch",
    affix: string,
    body?: any
  ) => {
    return fetch(
      [options.prefix, affix.startsWith("/") ? affix.substring(1) : affix]
        .filter(Boolean)
        .join("/"),
      {
        method,
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          ...(options.headers || {}),
        },
        body: JSON.stringify(body),
      }
    ).then((response) => {
      const contentType = response.headers.get("content-type");
      if (contentType && contentType.indexOf("application/json") !== -1) {
        return response.json().then((data) => {
          if (response.ok) {
            return data;
          } else {
            throw data;
          }
        });
      } else {
        throw response;
      }
    });
  };
