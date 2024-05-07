import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { PostTagForm } from "./PostTagEditForm";
import { PostTagEntity } from "src/sdk/fireback/modules/cms/PostTagEntity";
import { useGetPostTagByUniqueId } from "src/sdk/fireback/modules/cms/useGetPostTagByUniqueId";
import { usePostPostTag } from "src/sdk/fireback/modules/cms/usePostPostTag";
import { usePatchPostTag } from "src/sdk/fireback/modules/cms/usePatchPostTag";
export const PostTagEntityManager = ({ data }: DtoEntity<PostTagEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<PostTagEntity>
  >({
    data,
  });
  const getSingleHook = useGetPostTagByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostPostTag({
    queryClient,
  });
  const patchHook = usePatchPostTag({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          PostTagEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PostTagEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PostTagForm}
      onEditTitle={t.posttags.editpostTag}
      onCreateTitle={t.posttags.newpostTag}
      data={data}
    />
  );
};
