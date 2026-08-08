import { CommonArchiveManager } from "../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { RegionalContentNavigation } from "../../sdk/navigation/AbacNavigation";
import { RegionalContentList } from "./RegionalContentList";
import { strings } from "./strings/translations";

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
