import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { PublicJoinKeyList } from "./PublicJoinKeyList";
import { CommonArchiveManager } from "../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { PublicJoinKeyDto } from "../../sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "../../sdk/navigation/AbacNavigation";

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
