import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";

import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { UserList } from "./UserList";
import { UserNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";

export const UserArchiveScreen = () => {
  const s = useS(strings);
  const router = useRouter();
  const { locale } = useLocale();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={() => {
          router.push(UserNavigation.create());
        }}
        pageTitle={s.menuTitle}
      >
        <UserList />
      </CommonArchiveManager>
    </>
  );
};
