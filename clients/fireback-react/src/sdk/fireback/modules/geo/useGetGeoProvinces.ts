import { useContext } from "react";
import { useQuery } from "react-query";
import { 
  RemoteQueryContext,
  UseRemoteQuery,
  queryBeforeSend,
} from "../../core/react-tools";
import { execApiFn, IResponseList } from "../../core/http-tools";
import {
    GeoProvinceEntity,
} from "../geo/GeoProvinceEntity"
export function useGetGeoProvinces({
  queryOptions,
  query,
  queryClient,
  execFnOverride,
  unauthorized,
  optionFn
}: UseRemoteQuery) {
  const { options, execFn } = useContext(RemoteQueryContext);
  const computedOptions = optionFn ? optionFn(options) : options;
  // Calculare the function which will do the remote calls.
  // We consider to use global override, this specific override, or default which
  // comes with the sdk.
  const rpcFn = execFnOverride
    ? execFnOverride(computedOptions)
    : execFn
    ? execFn(computedOptions)
    : execApiFn(computedOptions);
  // Url of the remote affix.
  const url = "/geoProvinces".substr(1);
  let computedUrl = `${url}?${new URLSearchParams(
    queryBeforeSend(query)
  ).toString()}`;
  // Attach the details of the request to the fn
  const fn = () => rpcFn("GET", computedUrl);
  const auth = computedOptions?.headers?.authorization
  const hasKey = auth != "undefined" && auth != undefined && auth !=null && auth != "null" && !!auth
  const query$ = useQuery<any, any, IResponseList<GeoProvinceEntity>, any>(["*geo.GeoProvinceEntity", computedOptions, query], fn, {
    cacheTime: 1000,
    retry: false,
    keepPreviousData: true,
    enabled: hasKey || unauthorized || false,
    ...(queryOptions || {})
  } as any);
  const items: Array<GeoProvinceEntity> = query$.data?.data?.items || [];
  return { query: query$, items};
}
useGetGeoProvinces.UKEY = "*geo.GeoProvinceEntity"