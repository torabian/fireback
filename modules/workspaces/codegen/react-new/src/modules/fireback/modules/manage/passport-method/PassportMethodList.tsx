import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { columns } from "./PassportMethodColumns";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/workspaces/PassportMethodEntity";
import { useGetPassportMethods } from "@/modules/fireback/sdk/modules/workspaces/useGetPassportMethods";
import { useDeletePassportMethod } from "@/modules/fireback/sdk/modules/workspaces/useDeletePassportMethod";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const PassportMethodList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGetPassportMethods}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PassportMethodEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePassportMethod}
      ></CommonListManager>
    </>
  );
};
