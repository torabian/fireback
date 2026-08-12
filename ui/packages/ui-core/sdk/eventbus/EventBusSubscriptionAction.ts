import { WebSocketX } from "@fireback/js-remote-ctx/common/WebSocketX";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import { useWebSocketX } from "@fireback/js-remote-ctx/react/useWebSocketX";
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
  static Method = "REACTIVE";
  static Create = (overrideUrl?: string, qs?: URLSearchParams, options) => {
    const url = overrideUrl ?? EventBusSubscriptionAction.NewUrl(qs);
    const Cls = options?.SocketClass
      ? options.SocketClass
      : WebSocketX<unknown, unknown>;
    return new Cls(url, undefined, {
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
