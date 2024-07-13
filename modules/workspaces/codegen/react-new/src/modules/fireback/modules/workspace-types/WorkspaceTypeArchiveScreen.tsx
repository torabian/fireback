import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useT } from "@/modules/fireback/hooks/useT";

import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { WorkspaceTypeList } from "./WorkspaceTypeList";

export const WorkspaceTypeArchiveScreen = () => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(`/${locale}/workspace-type/new`);
        }}
        pageTitle={t.fbMenu.workspaceTypes}
      >
        <WorkspaceTypeList />
      </CommonArchiveManager>
    </>
  );
};
