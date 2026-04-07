import { WebSocketX } from "../../sdk/common/WebSocketX";
import { buildUrl } from "../../sdk/common/buildUrl";
import { useWebSocketX } from "../../sdk/react/useWebSocketX";
/**
 * Action to communicate with the action EventBusSubscription
 */
export type EventBusSubscriptionActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export const useEventBusSubscriptionAction = (options?: {
  qs?: URLSearchParams;
  overrideUrl?: string;
}) => {
  return useWebSocketX(() =>
    EventBusSubscriptionAction.Create(options?.overrideUrl, options?.qs),
  );
};
/**
 * EventBusSubscriptionAction
 */
export class EventBusSubscriptionAction {
  //
  static URL = "/ws";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(EventBusSubscriptionAction.URL, undefined, qs);
  static Method = "reactive";
  static Create = (overrideUrl?: string, qs?: URLSearchParams) => {
    const url = overrideUrl ?? EventBusSubscriptionAction.NewUrl(qs);
    return new WebSocketX<unknown, unknown>(url, undefined, {
      MessageFactoryClass: undefined,
    });
  };
  static Definition = {
    name: "EventBusSubscription",
    url: "/ws",
    method: "reactive",
    description:
      "Connects a client to all events related to their user profile, or workspace they are in",
  };
}
