import { useRouter } from "../../../../fireback-ui/hooks/useRouter";
import { useLocale } from "../../../../fireback-ui/hooks/useLocale";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { WorkspaceTypeList } from "./WorkspaceTypeList";
import { WorkspaceTypeNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceTypeArchiveScreen = () => {
  const s = useS(strings);
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(WorkspaceTypeNavigation.create());
        }}
        pageTitle={s.menuTitle}
      >
        <WorkspaceTypeList />
      </CommonArchiveManager>
    </>
  );
};
