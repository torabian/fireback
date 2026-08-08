import { CommonListManager } from "../../../../fireback-ui/components/entity-manager/CommonListManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { usePassportAwareDeleteAction } from "../../../sdk/abac/PassportAwareDeleteAction";
import { usePassportMethodBrowseActionQuery } from "../../../sdk/abac/PassportMethodBrowseAction";
import { PassportMethodNavigation } from "../../../sdk/navigation/AbacNavigation";
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
