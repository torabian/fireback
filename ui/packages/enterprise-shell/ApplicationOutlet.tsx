import { QueryClient } from "@tanstack/react-query";
import { ToastContainer } from "react-toastify";
import { ActionMenuProvider } from "@fireback/ui-core/components/action-menu/ActionMenu";
import {
  ModalManager,
  ModalProvider,
} from "@fireback/ui-core/components/modal/Modal";
import { OverlayProvider } from "@fireback/ui-core/components/overlay/OverlayProvider";
import { ReactiveSearchProvider } from "@fireback/ui-core/components/reactive-search/ReactiveSearchContext";
import { AppConfigProvider } from "@fireback/ui-core/hooks/appConfigTools";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";

/**
 * Shows routes of the application, can be independently used,
 * needs to be wrapped in a router
 * @param param0
 * @returns
 */
export const ApplicationOutlet = ({
  routerId,
  ApplicationRoutes,
  queryClient,
}: {
  routerId: string;
  ApplicationRoutes: any;
  queryClient: QueryClient;
}) => {
  return (
    <AppConfigProvider
      initialConfig={{
        remote: BUILD_VARIABLES.REMOTE_SERVICE,
      }}
    >
      <ReactiveSearchProvider>
        <ActionMenuProvider>
          <ModalProvider>
            <OverlayProvider>
              <ApplicationRoutes routerId={routerId} />
              <ModalManager />
            </OverlayProvider>
          </ModalProvider>
          <ToastContainer />
        </ActionMenuProvider>
      </ReactiveSearchProvider>
    </AppConfigProvider>
  );
};
