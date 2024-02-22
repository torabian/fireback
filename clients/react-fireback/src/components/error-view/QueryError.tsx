import { useT } from "@/hooks/useT";
import { enTranslations } from "@/translations/en";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { UseMutationResult, UseQueryResult } from "react-query";

export function getQueryErrorString(
  t: typeof enTranslations,
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
      unknownStr =
        "There is an unknown error while getting information, please contact your software provider if issue persists.";
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
  const { options } = useContext(RemoteQueryContext);

  if (!query) {
    return null;
  }
  return (
    <>
      {query.isError && (
        <div className="basic-error-box fadein">
          {getQueryErrorString(t, query, { remote: options.prefix }) || ""}
        </div>
      )}
      {/* Now this is to debate, if there is an error, and no data, then hide it. */}
      {!query.isError || (query as any).isPreviousData ? children : null}
    </>
  );
}
