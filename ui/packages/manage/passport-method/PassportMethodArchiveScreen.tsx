import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { PassportMethodList } from "./PassportMethodList";
import { PassportMethodNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
export const PassportMethodArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.passportMethods.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PassportMethodNavigation.create());
      }}
    >
      <PassportMethodList />
    </CommonArchiveManager>
  );
};
