import { useT } from "../../../hooks/useT";
import { useRouter } from "../../../hooks/useRouter";

import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
import { useLocale } from "../../../hooks/useLocale";
import { UserList } from "./UserList";
import { UserEntity } from "../../../sdk/modules/workspaces/UserEntity";

export const UserArchiveScreen = () => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(UserEntity.Navigation.create());
        }}
        pageTitle={t.fbMenu.users}
      >
        <UserList />
      </CommonArchiveManager>
    </>
  );
};
