import { FirebackEssentialRouterManager } from "../../modules/fireback/apps/core/EssentialRouter";

// ~ auto:useRouteImport

export function ApplicationRoutes({ routerId }: { routerId?: string }) {
  // ~ auto:useRouteDefs

  return (
    <FirebackEssentialRouterManager routerId={routerId}>
      {/* ~ auto:useRouteJsx */}
    </FirebackEssentialRouterManager>
  );
}
