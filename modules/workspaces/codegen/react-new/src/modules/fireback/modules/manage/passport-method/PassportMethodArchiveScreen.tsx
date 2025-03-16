import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { PassportMethodList } from "./PassportMethodList";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/workspaces/PassportMethodEntity";
export const PassportMethodArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.passportMethods.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PassportMethodEntity.Navigation.create());
      }}
    >
      <PassportMethodList />
    </CommonArchiveManager>
  );
};
