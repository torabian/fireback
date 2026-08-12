import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { PublicJoinKeyList } from "./PublicJoinKeyList";
import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { PublicJoinKeyDto } from "@fireback/selfservice/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";

export const PublicJoinKeyArchiveScreen = () => {
  const s = useS(strings);

  return (
    <>
      <CommonArchiveManager
        pageTitle={s.menuTitle}
        newEntityHandler={({ locale, router }) => {
          router.push(PublicJoinKeyNavigation.create());
        }}
      >
        <PublicJoinKeyList />
      </CommonArchiveManager>
    </>
  );
};
