import { useT } from "@/modules/fireback/hooks/useT";
import { useRouter } from "@/modules/fireback/hooks/useRouter";

import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { UserList } from "./UserList";
import { UserEntity } from "../../sdk/modules/workspaces/UserEntity";

export const UserArchiveScreen = () => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(UserEntity.Navigation.create(locale));
        }}
        pageTitle={t.fbMenu.users}
      >
        <UserList />
      </CommonArchiveManager>
    </>
  );
};
