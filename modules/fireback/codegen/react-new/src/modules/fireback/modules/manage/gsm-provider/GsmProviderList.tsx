import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { columns } from "./GsmProviderColumns";
import { GsmProviderEntity } from "@/modules/fireback/sdk/modules/abac/GsmProviderEntity";
import { useGetGsmProviders } from "@/modules/fireback/sdk/modules/abac/useGetGsmProviders";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { usePostGsmProviderRemove } from "@/modules/fireback/sdk/modules/abac/usePostGsmProviderRemove";
export const GsmProviderList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGetGsmProviders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          GsmProviderEntity.Navigation.single(uniqueId)
        }
        deleteHook={usePostGsmProviderRemove}
      ></CommonListManager>
    </>
  );
};
