import { useSyncExternalStore } from "react";
import { useRouter } from "./useRouter";
import { getLocale, setLocale, subscribeToLocale } from "./localeStore";

/**
 * The app's current locale, region and text direction.
 *
 * Locale used to be read out of the URL (a leading "/en/"/"/fa/" segment
 * every route was wrapped in - see EssentialRouter.tsx/App.tsx's own history
 * for why that's gone). It now comes from localeStore instead: a forced
 * build locale, else whatever the user last chose (persisted, and shared
 * live with every other useLocale()/usePureLocale() instance via
 * useSyncExternalStore - no Provider needed), else the build's default. See
 * personal-settings/InterfaceSettings.tsx for the switcher that calls
 * setLocale().
 *
 * asPath is still sourced from the router (unrelated to locale resolution
 * itself) purely for the handful of existing consumers that destructure it
 * off this hook already (ActiveLink, MenuParticle, useRtlClass).
 */
export function useLocale() {
  const router = useRouter();
  const locale = useSyncExternalStore(subscribeToLocale, getLocale, getLocale);

  let region = "us";
  let dir = "ltr";

  // Matches the pre-existing behavior exactly (only "fa" was ever special-cased
  // here) - not the place to also start covering "ar", a separate concern from
  // this hook's actual job in this change: where the locale value comes from.
  if (locale === "fa") {
    region = "ir";
    dir = "rtl";
  }

  return { locale, asPath: router.asPath, region, dir, setLocale };
}
