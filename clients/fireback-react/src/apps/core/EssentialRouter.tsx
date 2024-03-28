import { AppConfigContext } from "@/hooks/appConfigTools";
import { useT } from "@/hooks/useT";
import { CommonProfileEntityManager } from "@/modules/abac/common-profile/CommonProfileEntityManager";
import { SettingsScreen } from "@/modules/desktop-app-settings/SettingsScreen";

import { NotFound404 } from "@/components/404/NotFound404";
import { useLocale } from "@/hooks/useLocale";
import { useRtlClass } from "@/hooks/useRtlClass";
import {
  useAbacAuthenticatedRoutes,
  useAbacModulePublicRoutes,
} from "@/modules/abac/AbacModuleRoutes";
import { useDriveRoutes } from "@/modules/drive/DriveRoutes";

import { useRemoteMenuResolver } from "@/hooks/useRemoteMenuResolver";
import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "../../components/layouts/Layout";
import { AboutScreen } from "../fireback/AboutScreen";
import { PageTitleProvider } from "@/components/page-title/PageTitle";

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
