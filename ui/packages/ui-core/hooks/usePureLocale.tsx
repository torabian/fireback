import { useSyncExternalStore } from "react";
import { getLocale, subscribeToLocale, setLocale } from "./localeStore";

/**
 * The Router-independent twin of useLocale() - for the handful of call sites
 * that resolve locale *before* any <Router> (or even a session) is known to
 * exist yet, e.g. WithSelfServiceRoutes.tsx deciding which route table to
 * mount in the first place.
 *
 * Historically this parsed a locale straight out of window.location.hash
 * (via its own useWindowHash listener), since that was the only locale
 * signal available pre-Router. Now that locale comes from localeStore
 * instead of the URL at all, there's no URL to watch anymore - this is just
 * useLocale() minus the router-derived asPath/region-of-router-context
 * requirement, so both hooks now genuinely agree on the same value at every
 * point in the app, not just after the URL and this hook happen to sync up.
 */
export function usePureLocale() {
  const locale = useSyncExternalStore(subscribeToLocale, getLocale, getLocale);

  let region = "us";
  let dir = "ltr";

  if (locale === "fa") {
    region = "ir";
    dir = "rtl";
  }

  return { locale, region, dir, setLocale };
}
