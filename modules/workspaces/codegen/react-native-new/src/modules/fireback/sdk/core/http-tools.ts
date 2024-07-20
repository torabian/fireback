// @ts-nocheck

export interface EmptyRequest {}
export interface OkayResponse {}
export interface BulkRecordRequest<T> {
  records: Array<T>;
}
export interface DeleteRequest {
  uniqueId: string | string[];
  mode?: "immediate" | "background";
  query?: string;
}

export interface DeleteResponse {
  rowsAffected: number;
}

export interface RemoteRequestOption {
  headers?: {
    [key: string]: string;
  };
  prefix?: string;
}

/**
 * Converts all errors, network, api into an object that can
 * be passed to setErrors of formik ref.
 */
export function mutationErrorsToFormik(errors: any): {
  [key: string]: string;
} {
  const err: { [key: string]: string } = {};

  if (errors.error && Array.isArray(errors.error?.errors)) {
    for (const field of errors.error?.errors) {
      err[field.location] = field.messageTranslated || field.message;
    }
  }

  // This is when a network failure happens
  if (errors.status && errors.ok === false) {
    return {
      form: `${errors.status}`,
    };
  }

  if (errors?.error?.code) {
    err.form = errors?.error?.code;
  }

  if (errors?.error?.message) {
    err.form = errors?.error?.message;
  }

  if (errors?.error?.messageTranslated) {
    err.form = errors?.error?.messageTranslated;
  }

  if (errors.message) {
    return {
      form: `${errors.message}`,
    };
  }

  return err;
}

export type ExecApi = (
  method: "post" | "get" | "put" | "delete" | "patch",
  affix: string,
  body?: any
) => Promise<any>;

export interface IResponseData<T> {
  kind?: string;
  fields?: string;
  etag?: string;
  id?: string;
  lang?: string;
  updated?: string;
  deleted?: boolean;
  currentItemCount?: Number;
  itemsPerPage?: Number;
  startIndex?: Number;
  totalItems?: Number;
  pageIndex?: Number;
  totalPages?: Number;
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

export type IDeleteResponse = IResponse<{ rowsAffected: number }>;

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

export const execApiFn =
  (options: RemoteRequestOption) =>
  (
    method: "post" | "get" | "put" | "delete" | "patch",
    affix: string,
    body?: any
  ) => {
    const actualUrl = options.prefix + affix;
    return fetch(actualUrl, {
      method,
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        ...(options.headers || {}),
      },
      body: JSON.stringify(body),
    }).then((response) => {
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

export interface Query {
  withPreloads?: string;
  itemsPerPage?: number;
  deep?: boolean;
  startIndex?: number;
  query?: string;
  jsonQuery?: any;
  uniqueId?: string;
}
