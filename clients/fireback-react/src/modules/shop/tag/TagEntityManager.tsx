import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { TagForm } from "./TagEditForm";
import { TagEntity } from "src/sdk/fireback/modules/shop/TagEntity";
import { useGetTagByUniqueId } from "src/sdk/fireback/modules/shop/useGetTagByUniqueId";
import { usePostTag } from "src/sdk/fireback/modules/shop/usePostTag";
import { usePatchTag } from "src/sdk/fireback/modules/shop/usePatchTag";
export const TagEntityManager = ({ data }: DtoEntity<TagEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<TagEntity>
  >({
    data,
  });
  const getSingleHook = useGetTagByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostTag({
    queryClient,
  });
  const patchHook = usePatchTag({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          TagEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        TagEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ TagForm }
      onEditTitle={t.tags.editTag }
      onCreateTitle={t.tags.newTag }
      data={data}
    />
  );
};