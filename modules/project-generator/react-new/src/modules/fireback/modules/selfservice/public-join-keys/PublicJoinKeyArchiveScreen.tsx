import { useT } from "../../../../fireback-ui/hooks/useT";
import { PublicJoinKeyList } from "./PublicJoinKeyList";
import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { PublicJoinKeyDto } from "../../../sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "../../../sdk/navigation/AbacNavigation";

export const PublicJoinKeyArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.publicJoinKey}
        newEntityHandler={({ locale, router }) => {
          router.push(PublicJoinKeyNavigation.create());
        }}
      >
        <PublicJoinKeyList />
      </CommonArchiveManager>
    </>
  );
};
