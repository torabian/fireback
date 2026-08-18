import { WebSocketX } from "@fireback/js-remote-ctx/common/WebSocketX";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import { useWebSocketX } from "@fireback/js-remote-ctx/react/useWebSocketX";
/**
 * Action to communicate with the action InternalStatsStream
 */
export type InternalStatsStreamActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export const useInternalStatsStreamAction = (options?: {
  qs?: URLSearchParams;
  overrideUrl?: string;
}) => {
  return useWebSocketX(() =>
    InternalStatsStreamAction.Create(options?.overrideUrl, options?.qs),
  );
};
/**
 * InternalStatsStreamAction
 */
export class InternalStatsStreamAction {
  //
  static URL = "/internal-stats/stream";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(InternalStatsStreamAction.URL, undefined, qs);
  static Method = "REACTIVE";
  static Create = (overrideUrl?: string, qs?: URLSearchParams, options) => {
    const url = overrideUrl ?? InternalStatsStreamAction.NewUrl(qs);
    const Cls = options?.SocketClass
      ? options.SocketClass
      : WebSocketX<unknown, unknown>;
    return new Cls(url, undefined, {
      MessageFactoryClass: undefined,
    });
  };
  static Definition = {
    name: "InternalStatsStream",
    cliName: "stream",
    url: "/internal-stats/stream",
    method: "reactive",
    description:
      "Reactive (websocket) live feed of the same snapshot InternalStatsSnapshot returns, pushed on an interval (see InternalStatsModuleConfig.Interval) until the client disconnects. Each frame is the JSON-encoded snapshot object - identical shape to InternalStatsSnapshot's out fields. Root workspace token required by default - see InternalStatsModuleConfig.Authorize.",
  };
}
