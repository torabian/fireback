import { useRouter } from "../../../hooks/useRouter";
import { useLocale } from "../../../hooks/useLocale";
import { useT } from "../../../hooks/useT";

import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
import { WorkspaceTypeList } from "./WorkspaceTypeList";
import { WorkspaceTypeEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceTypeEntity";

export const WorkspaceTypeArchiveScreen = () => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(WorkspaceTypeEntity.Navigation.create());
        }}
        pageTitle={t.fbMenu.workspaceTypes}
      >
        <WorkspaceTypeList />
      </CommonArchiveManager>
    </>
  );
};
