import { AppConfigContext } from "@fireback/ui-core/hooks/appConfigTools";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "@fireback/ui-core/components/strings/translations";

import { NotFound404 } from "@fireback/ui-core/components/404/NotFound404";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useRtlClass } from "@fireback/ui-core/hooks/useRtlClass";

import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "@fireback/ui-core/components/layouts/Layout";
import { PageTitleProvider } from "@fireback/ui-core/components/page-title/PageTitle";
import { useRemoteMenuResolver } from "@fireback/ui-core/hooks/useRemoteMenuResolver";
import { useManageRoutes } from "@fireback/manage/ManageRoutes";
import { useMobileKitRoutes } from "@fireback/mobile-kit/dashboard/ManageRoutes";
import { SettingsScreen } from "@fireback/selfservice/personal-settings/SettingsScreen";
import { useSelfServiceAuthenticateRoutes } from "@fireback/selfservice/SelfServiceRoutes";
import { AnimatedRouteWrapper } from "@fireback/ui-core/components/swipe-transition/SwipeTransition";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";

export function FirebackEssentialRouterManager({
  children,
  routerId,
}: {
  children?: any;
  routerId?: string;
}) {
  const s = useS(strings);
  useRtlClass();
  const { locale } = useLocale();
  const { config } = useContext(AppConfigContext);
  const sidebarMenu = useRemoteMenuResolver("sidebar");

  const selfServiceAuthenticateRoutes = useSelfServiceAuthenticateRoutes();
  const manageRoutes = useManageRoutes();
  const mobileKitRoutes = useMobileKitRoutes();

  // ~ auto:useRouteDefs

  return (
    <PageTitleProvider affix={s.productName}>
      <Routes>
        <Route
          path="/"
          element={
            <Navigate
              to={(BUILD_VARIABLES.DEFAULT_ROUTE || "/{locale}/signin").replace(
                "{locale}",
                config.interfaceLanguage || locale || "en",
              )}
              replace
            />
          }
        />
        <Route
          path=":locale"
          element={<Layout routerId={routerId} sidebarMenu={sidebarMenu} />}
        >
          <Route
            path="settings"
            element={
              <AnimatedRouteWrapper>
                <SettingsScreen />
              </AnimatedRouteWrapper>
            }
          />

          {selfServiceAuthenticateRoutes}
          {manageRoutes}
          {mobileKitRoutes}

          {children}

          {/* ~ auto:useRouteJsx */}

          <Route path="*" element={<NotFound404 />} />
        </Route>

        <Route path="*" element={<NotFound404 />} />
      </Routes>
    </PageTitleProvider>
  );
}
