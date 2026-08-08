import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { RegionalContentList } from "./RegionalContentList";
import { RegionalContentDto } from "../../sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "../../sdk/navigation/AbacNavigation";
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
