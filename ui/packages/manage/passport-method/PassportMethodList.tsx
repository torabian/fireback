import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { usePassportMethodAwareDeleteAction } from "@fireback/manage/sdk/abac/PassportMethodAwareDeleteAction";
import { usePassportMethodBrowseActionQuery } from "@fireback/manage/sdk/abac/PassportMethodBrowseAction";
import { PassportMethodNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./PassportMethodColumns";
import { strings } from "./strings/translations";
import { createUdfBrowseQueryHook } from "@fireback/ui-core/hooks/useUdfBrowseQuery";
export const PassportMethodList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={createUdfBrowseQueryHook(usePassportMethodBrowseActionQuery)}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PassportMethodNavigation.single(uniqueId)
        }
        deleteHook={usePassportMethodAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
