import { useRouter } from "../../../hooks/useRouter";
import { useLocale } from "../../../hooks/useLocale";
import { useT } from "../../../hooks/useT";

import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
import { WorkspaceTypeList } from "./WorkspaceTypeList";
import { WorkspaceTypeNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceTypeArchiveScreen = () => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(WorkspaceTypeNavigation.create());
        }}
        pageTitle={t.fbMenu.workspaceTypes}
      >
        <WorkspaceTypeList />
      </CommonArchiveManager>
    </>
  );
};
