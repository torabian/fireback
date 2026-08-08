import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { GsmProviderList } from "./GsmProviderList";
import { GsmProviderDto } from "@/modules/fireback/sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
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
