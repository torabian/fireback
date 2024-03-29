import { useRouter } from "@/Router";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
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
