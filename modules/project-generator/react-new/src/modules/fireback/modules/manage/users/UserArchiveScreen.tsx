import { useT } from "../../../hooks/useT";
import { useRouter } from "../../../hooks/useRouter";

import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
import { useLocale } from "../../../hooks/useLocale";
import { UserList } from "./UserList";
import { UserNavigation } from "../../../sdk/navigation/AbacNavigation";

export const UserArchiveScreen = () => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(UserNavigation.create());
        }}
        pageTitle={t.fbMenu.users}
      >
        <UserList />
      </CommonArchiveManager>
    </>
  );
};
