import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { GsmProviderList } from "./GsmProviderList";
import { GsmProviderEntity } from "@/modules/fireback/sdk/modules/abac/GsmProviderEntity";
export const GsmProviderArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.gsmProviders.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(GsmProviderEntity.Navigation.create());
      }}
    >
      <GsmProviderList />
    </CommonArchiveManager>
  );
};
