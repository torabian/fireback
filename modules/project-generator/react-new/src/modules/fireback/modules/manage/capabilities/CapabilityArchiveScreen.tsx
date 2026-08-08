import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { CapabilityList } from "./CapabilityList";
import { CapabilityDto } from "@/modules/fireback/sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
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
