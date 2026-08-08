import { useT } from "../../hooks/useT";
import { useS } from "../../hooks/useS";
import { enTranslations } from "../../translations/en";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useContext } from "react";
import type { UseMutationResult, UseQueryResult } from "@tanstack/react-query";
import { FormButton } from "../forms/form-button/FormButton";
import { strings } from "../strings/translations";

export function getQueryErrorString(
  t: typeof enTranslations,
  s: typeof strings,
  query: UseQueryResult<any, any> | UseMutationResult<any, any>,
  params: any = {}
): string | null {
  if (query.isError) {
    if (query.error?.status === 404) {
      return t.notfound + "(" + params.remote + ")";
    }
    if (query.error.message === "Failed to fetch") {
      return t.networkError + "(" + params.remote + ")";
    }

    if (query.error?.error?.messageTranslated) {
      return query.error?.error?.messageTranslated;
    }
    if (query.error?.error?.message) {
      return query.error?.error?.message;
    }

    let unknownStr = query.error?.toString();

    if ((unknownStr + "").includes("object Object")) {
      unknownStr = s.components.unknownError;
    }

    return unknownStr;
  }

  return null;
}

export function QueryErrorView({
  query,
  children,
}: {
  query: UseQueryResult<any, any> | UseMutationResult<any, any> | any;
  children?: React.ReactNode;
}) {
  const t = useT();
  const s = useS(strings);
  const { options, setOverrideRemoteUrl, overrideRemoteUrl } =
    useContext(RemoteQueryContext);

  let showAutoAdjustTheUrl = false;
  let port = "80";

  try {
    if (options?.prefix) {
      const url = new URL(options?.prefix);
      port = url.port || (url.protocol === "https:" ? "443" : "80");
      showAutoAdjustTheUrl =
        (location.host.includes("192.168") ||
          location.host.includes("127.0")) &&
        query.error?.message?.includes("Failed to fetch");
    }
  } catch (err) {}

  const autoAdjust = () => {
    setOverrideRemoteUrl("http://" + location.hostname + ":" + port + "/");
  };

  if (!query) {
    return null;
  }
  return (
    <>
      {query.isError && (
        <div className="basic-error-box fadein">
          {getQueryErrorString(t, s, query, { remote: options.prefix }) || ""}
          {showAutoAdjustTheUrl && (
            <button className="btn btn-sm btn-secondary" onClick={autoAdjust}>
              {s.components.autoReroute}
            </button>
          )}
          {overrideRemoteUrl && (
            <button
              className="btn btn-sm btn-secondary"
              onClick={() => setOverrideRemoteUrl(undefined)}
            >
              {s.components.reset}
            </button>
          )}
          <ul>
            {(query.error?.error?.errors || []).map((item) => {
              return (
                <li key={item.location}>
                  {item.messageTranslated || item.message} ({item.location})
                </li>
              );
            })}
          </ul>
          {query.refetch && (
            <FormButton onClick={query.refetch}>{s.components.retry}</FormButton>
          )}
        </div>
      )}
      {/* Now this is to debate, if there is an error, and no data, then hide it. */}
      {!query.isError || (query as any).isPreviousData ? children : null}
    </>
  );
}
