import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./PageTagColumns";
import { PageTagEntity } from "src/sdk/fireback/modules/cms/PageTagEntity";
import { useGetPageTags } from "src/sdk/fireback/modules/cms/useGetPageTags";
import { useDeletePageTag } from "@/sdk/fireback/modules/cms/useDeletePageTag";
export const PageTagList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetPageTags}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PageTagEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePageTag}
      ></CommonListManager>
    </>
  );
};