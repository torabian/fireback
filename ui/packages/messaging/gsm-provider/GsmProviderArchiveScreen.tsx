import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { GsmProviderList } from "./GsmProviderList";
import { GsmProviderDto } from "@fireback/messaging/sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
export const GsmProviderArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.gsmProviders.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(GsmProviderNavigation.create());
      }}
    >
      <GsmProviderList />
    </CommonArchiveManager>
  );
};
