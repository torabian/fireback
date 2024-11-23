import { AppConfigContext } from "@/modules/fireback/hooks/appConfigTools";
import { useT } from "@/modules/fireback/hooks/useT";
import { CommonProfileEntityManager } from "@/modules/fireback/modules/common-profile/CommonProfileEntityManager";
import { SettingsScreen } from "@/modules/fireback/modules/desktop-app-settings/SettingsScreen";

import { NotFound404 } from "@/modules/fireback/components/404/NotFound404";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useRtlClass } from "@/modules/fireback/hooks/useRtlClass";
import {
  useAbacAuthenticatedRoutes,
  useAbacModulePublicRoutes,
} from "@/modules/fireback/modules/AbacModuleRoutes";
import { useDriveRoutes } from "@/modules/fireback/modules/drive/DriveRoutes";

import { useRemoteMenuResolver } from "@/modules/fireback/hooks/useRemoteMenuResolver";
import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "@/modules/fireback/components/layouts/Layout";
import { PageTitleProvider } from "@/modules/fireback/components/page-title/PageTitle";

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

  const abacModulePublicRoutes = useAbacModulePublicRoutes();
  const abacAuthenticatedRoutes = useAbacAuthenticatedRoutes();
  const driveRoutes = useDriveRoutes();
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
        <Route path=":locale">{abacModulePublicRoutes}</Route>

        <Route
          path=":locale"
          element={<Layout routerId={routerId} sidebarMenu={sidebarMenu} />}
        >
          <Route
            path={"profile"}
            element={<CommonProfileEntityManager />}
          ></Route>

          <Route path={"settings"} element={<SettingsScreen />}></Route>

          {driveRoutes}
          {abacAuthenticatedRoutes}

          {children}

          {/* ~ auto:useRouteJsx */}

          <Route path="*" element={<NotFound404 />} />
        </Route>
        <Route path="*" element={<NotFound404 />} />
      </Routes>
    </PageTitleProvider>
  );
}
