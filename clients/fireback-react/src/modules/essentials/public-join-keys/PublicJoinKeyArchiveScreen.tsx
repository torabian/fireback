import { useT } from "@/hooks/useT";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
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
