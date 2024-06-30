import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { PostForm } from "./PostEditForm";
import { PostEntity } from "src/sdk/fireback/modules/cms/PostEntity";
import { useGetPostByUniqueId } from "src/sdk/fireback/modules/cms/useGetPostByUniqueId";
import { usePostPost } from "src/sdk/fireback/modules/cms/usePostPost";
import { usePatchPost } from "src/sdk/fireback/modules/cms/usePatchPost";
export const PostEntityManager = ({ data }: DtoEntity<PostEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<PostEntity>
  >({
    data,
  });
  const getSingleHook = useGetPostByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostPost({
    queryClient,
  });
  const patchHook = usePatchPost({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(PostEntity.Navigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        PostEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PostForm}
      onEditTitle={t.posts.editpost}
      onCreateTitle={t.posts.newpost}
      data={data}
    />
  );
};
