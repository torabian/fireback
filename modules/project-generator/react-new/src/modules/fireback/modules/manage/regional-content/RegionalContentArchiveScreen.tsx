import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { RegionalContentList } from "./RegionalContentList";
import { RegionalContentEntity } from "@/modules/fireback/sdk/modules/abac/RegionalContentEntity";
export const RegionalContentArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.regionalContents.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(RegionalContentEntity.Navigation.create());
      }}
    >
      <RegionalContentList />
    </CommonArchiveManager>
  );
};
