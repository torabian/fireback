import { Context as C, JsonQuery } from "@/definitions/definitions";
import { enTranslations } from "@/translations/en";
import { IResponseList, RemoteRequestOption } from "src/sdk/fireback";
import { withJsonQuery } from "./withJsonQuery";
const { matchPattern } = require("url-matcher");

export type Context = C;

export function uriMatch(value: string) {
  return function (target: any, propertyKey: string) {
    if (!target["url"]) {
      target["url"] = {};
    }
    target.url[propertyKey as any] = value;
  };
}

export function method(value: string) {
  return function (target: any, propertyKey: string) {
    if (!target["method"]) {
      target["method"] = {};
    }
    target.method[propertyKey as any] = value;
  };
}

export const mockExecFn = (
  options: RemoteRequestOption,
  mockServerInstances: any[],
  t: typeof enTranslations
) => {
  return function (method: string, url: string, body: any) {
    const searchParams = new URLSearchParams(url);
    const itp = searchParams.get("itemsPerPage");
    const si = searchParams.get("startIndex");
    const itemsPerPage = itp === null ? 10 : +itp;
    const startIndex = si === null ? 0 : +si;

    /**
     * We scan the methods inside the mock instance.
     * They are decorated with method and url. If there is a match, we call that function
     * and pass the context to it, which has some details
     */
    for (let mockServerInstance of mockServerInstances) {
      let protoOfTest = Object.getPrototypeOf(mockServerInstance);
      for (let item of Object.getOwnPropertyNames(protoOfTest)) {
        if (
          typeof mockServerInstance[item as any] == "function" &&
          item !== "constructor" &&
          mockServerInstance.url[item]
        ) {
          console.log(1, mockServerInstance.url[item], url);
          const matchData = matchPattern(mockServerInstance.url[item], url);
          if (
            matchData &&
            method?.toLocaleLowerCase() ===
              mockServerInstance.method[item]?.toLocaleLowerCase()
          ) {
            const lang = (options as any).headers["accept-language"] || "en";
            const context = {
              url,
              token: options.headers?.authorization,
              acceptLanguage: lang,
              workspaceId:
                (options.headers && options.headers["workspace-id"]) || "",
              body,
              method,
              startIndex,
              itemsPerPage,
              paramValues: matchData.paramValues,
            };

            return mockServerInstance[item](context);
          }
        }
      }
    }

    return Promise.reject({
      error: {
        message:
          t.featureNotAvailableOnMock + "(" + method + " on " + url + ")",
      },
    });
  };
};

export type DeepPartial<T> = T extends object
  ? {
      [P in keyof T]?: DeepPartial<T[P]>;
    }
  : T;

export const emptyList = { data: { items: [] } } as any;

export async function getJsonRaw(entity: string, ctx: Context) {
  return fetch(
    process.env.REACT_APP_PUBLIC_URL +
      `md/${ctx.acceptLanguage || "en"}/${entity}.json`
  ).then((t) => t.json());
}

function paginate(items: Array<any>, ctx: Context): Array<any> {
  return items.filter((item: any, index: number) => {
    if (index < (ctx.startIndex || 0)) {
      return false;
    }
    if (index - (ctx.startIndex || 0) > ctx.itemsPerPage - 1) {
      return false;
    }

    return true;
  });
}

function applyQueryDSL(items: Array<any>, ctx: Context): Array<any> {
  return paginate(withJsonQuery(items, ctx), ctx);
}

export async function getJsonList(entity: string, ctx: Context) {
  return fetch(
    process.env.REACT_APP_PUBLIC_URL +
      `md/${ctx.acceptLanguage || "en"}/${entity}.json`
  ).then((t) => t.json());
}
export async function getJson(entity: string, ctx: Context) {
  return getJsonList(entity, ctx).then((resp) => {
    resp.data.items = applyQueryDSL(resp.data.items || [], ctx);

    return resp;
  });
}

export async function getItemUid(entity: string, ctx: Context) {
  const uniqueId: string = ctx.paramValues[0];
  return getJsonList(entity, ctx)
    .then((resp) => resp.data.items.find((t: any) => t.uniqueId === uniqueId))
    .then((item) => ({ data: item }));
}
