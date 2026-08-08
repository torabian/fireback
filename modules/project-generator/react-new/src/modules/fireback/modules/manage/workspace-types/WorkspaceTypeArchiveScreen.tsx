import { useRouter } from "../../../../fireback-ui/hooks/useRouter";
import { useLocale } from "../../../../fireback-ui/hooks/useLocale";
import { useT } from "../../../../fireback-ui/hooks/useT";

import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
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
