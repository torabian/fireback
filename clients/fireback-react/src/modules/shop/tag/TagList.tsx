import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./TagColumns";
import { TagEntity } from "src/sdk/fireback/modules/shop/TagEntity";
import { useGetTags } from "src/sdk/fireback/modules/shop/useGetTags";
import { useDeleteTag } from "@/sdk/fireback/modules/shop/useDeleteTag";
export const TagList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetTags}
        uniqueIdHrefHandler={(uniqueId: string) =>
          TagEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteTag}
      ></CommonListManager>
    </>
  );
};