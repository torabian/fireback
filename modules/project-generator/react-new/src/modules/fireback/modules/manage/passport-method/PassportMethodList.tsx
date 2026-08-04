import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/abac/PassportMethodEntity";
import { useGetPassportMethods } from "@/modules/fireback/sdk/modules/abac/useGetPassportMethods";
import { usePostPassportMethodRemove } from "@/modules/fireback/sdk/modules/abac/usePostPassportMethodRemove";
import { columns } from "./PassportMethodColumns";
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
        deleteHook={usePostPassportMethodRemove}
      ></CommonListManager>
    </>
  );
};
