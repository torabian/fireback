import { AppConfigContext } from "@/fireback/hooks/appConfigTools";
import { useT } from "@/fireback/hooks/useT";
import { CommonProfileEntityManager } from "@/fireback/modules/common-profile/CommonProfileEntityManager";
import { SettingsScreen } from "@/fireback/modules/desktop-app-settings/SettingsScreen";

import { NotFound404 } from "@/fireback/components/404/NotFound404";
import { useLocale } from "@/fireback/hooks/useLocale";
import { useRtlClass } from "@/fireback/hooks/useRtlClass";
import {
  useAbacAuthenticatedRoutes,
  useAbacModulePublicRoutes,
} from "@/fireback/modules/AbacModuleRoutes";
import { useDriveRoutes } from "@/fireback/modules/drive/DriveRoutes";

import { useRemoteMenuResolver } from "@/fireback/hooks/useRemoteMenuResolver";
import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "@/fireback/components/layouts/Layout";
import { AboutScreen } from "../../../apps/fireback/AboutScreen";
import { PageTitleProvider } from "@/fireback/components/page-title/PageTitle";

export function FirebackEssentialRouterManager({
  children,
}: {
  children?: any;
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

        <Route path=":locale" element={<Layout sidebarMenu={sidebarMenu} />}>
          <Route
            path={"profile"}
            element={<CommonProfileEntityManager />}
          ></Route>
          <Route path="about" element={<AboutScreen />}></Route>

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
