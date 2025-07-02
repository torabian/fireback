import { ToastContainer } from "react-toastify";
import { ActionMenuProvider } from "../../components/action-menu/ActionMenu";
import { ModalManager, ModalProvider } from "../../components/modal/Modal";
import { OverlayProvider } from "../../components/overlay/OverlayProvider";
import { ReactiveSearchProvider } from "../../components/reactive-search/ReactiveSearchContext";
import { AppConfigProvider } from "../../hooks/appConfigTools";
import { QueryClient, QueryClientProvider } from "react-query";

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
        remote: process.env.REACT_APP_REMOTE_SERVICE,
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
