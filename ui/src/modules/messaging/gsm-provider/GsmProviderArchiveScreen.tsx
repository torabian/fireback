import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { GsmProviderList } from "./GsmProviderList";
import { GsmProviderDto } from "../../sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "../../sdk/navigation/MessagingNavigation";
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
