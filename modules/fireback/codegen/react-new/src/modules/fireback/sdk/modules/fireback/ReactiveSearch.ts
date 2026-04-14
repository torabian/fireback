import { WebSocketX } from "../../sdk/common/WebSocketX";
import { buildUrl } from "../../sdk/common/buildUrl";
import { useWebSocketX } from "../../sdk/react/useWebSocketX";
/**
 * Action to communicate with the action ReactiveSearch
 */
export type ReactiveSearchActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export const useReactiveSearchAction = (options?: {
  qs?: URLSearchParams;
  overrideUrl?: string;
}) => {
  return useWebSocketX(() =>
    ReactiveSearchAction.Create(options?.overrideUrl, options?.qs),
  );
};
/**
 * ReactiveSearchAction
 */
export class ReactiveSearchAction {
  //
  static URL = "/reactive-search";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ReactiveSearchAction.URL, undefined, qs);
  static Method = "reactive";
  static Create = (overrideUrl?: string, qs?: URLSearchParams) => {
    const url = overrideUrl ?? ReactiveSearchAction.NewUrl(qs);
    return new WebSocketX<unknown, unknown>(url, undefined, {
      MessageFactoryClass: undefined,
    });
  };
  static Definition = {
    name: "ReactiveSearch",
    url: "/reactive-search",
    method: "reactive",
    description:
      "Reactive search is a general purpose search mechanism for different modules, and could be used in mobile apps or front-end to quickly search for a entity.",
  };
}
