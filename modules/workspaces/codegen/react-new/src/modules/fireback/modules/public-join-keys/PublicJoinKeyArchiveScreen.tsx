import { useT } from "@/modules/fireback/hooks/useT";

import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { PublicJoinKeyList } from "./PublicJoinKeyList";

export const PublicJoinKeyArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.publicJoinKey}
        newEntityHandler={({ locale, router }) => {
          router.push(`/${locale}/publicjoinkey/new`);
        }}
      >
        <PublicJoinKeyList />
      </CommonArchiveManager>
    </>
  );
};
