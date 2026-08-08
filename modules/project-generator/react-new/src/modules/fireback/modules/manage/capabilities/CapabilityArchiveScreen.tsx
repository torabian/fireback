import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { CapabilityList } from "./CapabilityList";
import { CapabilityDto } from "../../../sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "../../../sdk/navigation/AbacNavigation";
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
