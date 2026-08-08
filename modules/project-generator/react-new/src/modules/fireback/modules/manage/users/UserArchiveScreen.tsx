import { useT } from "../../../../fireback-ui/hooks/useT";
import { useRouter } from "../../../../fireback-ui/hooks/useRouter";

import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { useLocale } from "../../../../fireback-ui/hooks/useLocale";
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
