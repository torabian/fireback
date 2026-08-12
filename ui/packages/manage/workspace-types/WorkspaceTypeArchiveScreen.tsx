import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";

import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { WorkspaceTypeList } from "./WorkspaceTypeList";
import { WorkspaceTypeNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";

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
