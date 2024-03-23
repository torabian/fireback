import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./PostColumns";
import { PostEntity } from "src/sdk/fireback/modules/cms/PostEntity";
import { useGetPosts } from "src/sdk/fireback/modules/cms/useGetPosts";
import { useDeletePost } from "@/sdk/fireback/modules/cms/useDeletePost";
export const PostList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t) as any}
        queryHook={useGetPosts}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PostEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePost}
      ></CommonListManager>
    </>
  );
};
