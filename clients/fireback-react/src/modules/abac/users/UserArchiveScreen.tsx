import { useT } from "@/hooks/useT";
import { useRouter } from "@/Router";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { useLocale } from "@/hooks/useLocale";
import { UserNavigationTools } from "src/sdk/fireback/modules/workspaces/user-navigation-tools";
import { UserList } from "./UserList";

export const UserArchiveScreen = () => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(UserNavigationTools.create(locale));
        }}
        pageTitle={t.fbMenu.users}
      >
        <UserList />
      </CommonArchiveManager>
    </>
  );
};
