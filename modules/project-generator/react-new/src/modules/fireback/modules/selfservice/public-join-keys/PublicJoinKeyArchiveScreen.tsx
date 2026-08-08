import { useT } from "@/modules/fireback/hooks/useT";
import { PublicJoinKeyList } from "./PublicJoinKeyList";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { PublicJoinKeyDto } from "@/modules/fireback/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";

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
