import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { usePassportAwareDeleteAction } from "@/modules/fireback/sdk/abac/PassportAwareDeleteAction";
import { usePassportMethodBrowseActionQuery } from "@/modules/fireback/sdk/abac/PassportMethodBrowseAction";
import { PassportMethodNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { columns } from "./PassportMethodColumns";
import { strings } from "./strings/translations";
export const PassportMethodList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={usePassportMethodBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PassportMethodNavigation.single(uniqueId)
        }
        deleteHook={usePassportAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
