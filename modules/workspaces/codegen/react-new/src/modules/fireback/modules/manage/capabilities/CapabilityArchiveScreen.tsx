import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { CapabilityList } from "./CapabilityList";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/workspaces/CapabilityEntity";
export const CapabilityArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.capabilities.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(CapabilityEntity.Navigation.create());
      }}
    >
      <CapabilityList />
    </CommonArchiveManager>
  );
};
