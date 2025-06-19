import { AppConfigContext } from "../../hooks/appConfigTools";
import { useT } from "../../hooks/useT";

import { NotFound404 } from "../../components/404/NotFound404";
import { useLocale } from "../../hooks/useLocale";
import { useRtlClass } from "../../hooks/useRtlClass";
import { useDriveRoutes } from "../../modules/manage/drive/DriveRoutes";

import { useRemoteMenuResolver } from "../../hooks/useRemoteMenuResolver";
import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "../../components/layouts/Layout";
import { PageTitleProvider } from "../../components/page-title/PageTitle";
import { SettingsScreen } from "../../modules/selfservice/personal-settings/SettingsScreen";
import {
  useSelfServiceAuthenticateRoutes,
  useSelfServicePublicRoutes,
} from "../../modules/selfservice/SelfServiceRoutes";
import { useManageRoutes } from "../../modules/manage/ManageRoutes";
import { useMobileKitRoutes } from "../../modules/mobile-kit/dashboard/ManageRoutes";

export function FirebackEssentialRouterManager({
  children,
  routerId,
}: {
  children?: any;
  routerId?: string;
}) {
  const t = useT();
  useRtlClass();
  const { locale } = useLocale();
  const { config } = useContext(AppConfigContext);
  const sidebarMenu = useRemoteMenuResolver("sidebar");

  const selfServiceAuthenticateRoutes = useSelfServiceAuthenticateRoutes();
  const manageRoutes = useManageRoutes();
  const mobileKitRoutes = useMobileKitRoutes();

  // ~ auto:useRouteDefs

  return (
    <PageTitleProvider affix={t.productName}>
      <Routes>
        <Route
          path="/"
          element={
            <Navigate
              to={(
                process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
              ).replace("{locale}", config.interfaceLanguage || locale || "en")}
              replace
            />
          }
        />
        <Route
          path=":locale"
          element={<Layout routerId={routerId} sidebarMenu={sidebarMenu} />}
        >
          <Route path={"settings"} element={<SettingsScreen />}></Route>

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
