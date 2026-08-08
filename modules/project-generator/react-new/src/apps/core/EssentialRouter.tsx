import { AppConfigContext } from "../../modules/fireback-ui/hooks/appConfigTools";
import { useS } from "../../modules/fireback-ui/hooks/useS";
import { strings } from "../../modules/fireback-ui/components/strings/translations";

import { NotFound404 } from "../../modules/fireback-ui/components/404/NotFound404";
import { useLocale } from "../../modules/fireback-ui/hooks/useLocale";
import { useRtlClass } from "../../modules/fireback-ui/hooks/useRtlClass";

import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "../../modules/fireback-ui/components/layouts/Layout";
import { PageTitleProvider } from "../../modules/fireback-ui/components/page-title/PageTitle";
import { useRemoteMenuResolver } from "../../modules/fireback-ui/hooks/useRemoteMenuResolver";
import { useManageRoutes } from "../../modules/manage/ManageRoutes";
import { useMobileKitRoutes } from "../../modules/mobile-kit/dashboard/ManageRoutes";
import { SettingsScreen } from "../../modules/selfservice/personal-settings/SettingsScreen";
import { useSelfServiceAuthenticateRoutes } from "../../modules/selfservice/SelfServiceRoutes";
import { AnimatedRouteWrapper } from "./SwipeTransition";
import { BUILD_VARIABLES } from "../../modules/fireback-ui/hooks/build-variables";

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
