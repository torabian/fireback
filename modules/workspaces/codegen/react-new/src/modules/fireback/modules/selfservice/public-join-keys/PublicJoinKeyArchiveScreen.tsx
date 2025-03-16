import { useT } from "@/modules/fireback/hooks/useT";
import { PublicJoinKeyList } from "./PublicJoinKeyList";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/workspaces/PublicJoinKeyEntity";

export const PublicJoinKeyArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.publicJoinKey}
        newEntityHandler={({ locale, router }) => {
          router.push(PublicJoinKeyEntity.Navigation.create());
        }}
      >
        <PublicJoinKeyList />
      </CommonArchiveManager>
    </>
  );
};
