import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { columns } from "./RegionalContentColumns";
import { RegionalContentEntity } from "@/modules/fireback/sdk/modules/abac/RegionalContentEntity";
import { useGetRegionalContents } from "@/modules/fireback/sdk/modules/abac/useGetRegionalContents";
import { useDeleteRegionalContent } from "@/modules/fireback/sdk/modules/abac/useDeleteRegionalContent";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const RegionalContentList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGetRegionalContents}
        uniqueIdHrefHandler={(uniqueId: string) =>
          RegionalContentEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteRegionalContent}
      ></CommonListManager>
    </>
  );
};
