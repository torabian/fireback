import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./PostTagColumns";
import { PostTagEntity } from "src/sdk/fireback/modules/cms/PostTagEntity";
import { useGetPostTags } from "src/sdk/fireback/modules/cms/useGetPostTags";
import { useDeletePostTag } from "@/sdk/fireback/modules/cms/useDeletePostTag";
export const PostTagList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetPostTags}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PostTagEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePostTag}
      ></CommonListManager>
    </>
  );
};