import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { CapabilityList } from "./CapabilityList";
import { CapabilityDto } from "@fireback/manage/sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
export const CapabilityArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.capabilities.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(CapabilityNavigation.create());
      }}
    >
      <CapabilityList />
    </CommonArchiveManager>
  );
};
