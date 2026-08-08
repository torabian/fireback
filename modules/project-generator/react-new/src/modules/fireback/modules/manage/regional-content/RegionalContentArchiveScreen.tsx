import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { RegionalContentList } from "./RegionalContentList";
import { RegionalContentDto } from "@/modules/fireback/sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
export const RegionalContentArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.regionalContents.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(RegionalContentNavigation.create());
      }}
    >
      <RegionalContentList />
    </CommonArchiveManager>
  );
};
