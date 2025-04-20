import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/abac/PublicJoinKeyEntity";
import { useGetPublicJoinKeyByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetPublicJoinKeyByUniqueId";
import { usePatchPublicJoinKey } from "@/modules/fireback/sdk/modules/abac/usePatchPublicJoinKey";
import { usePostPublicJoinKey } from "@/modules/fireback/sdk/modules/abac/usePostPublicJoinKey";
import { PublicJoinKeyEditForm } from "./PublicJoinKeyEditForm";

export const PublicJoinKeyEntityManager = ({
  data,
}: DtoEntity<PublicJoinKeyEntity>) => {
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
    Partial<PublicJoinKeyEntity>
  >({
    data,
  });

  const getSingleHook = useGetPublicJoinKeyByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostPublicJoinKey({
    queryClient,
  });

  const patchHook = usePatchPublicJoinKey({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(PublicJoinKeyEntity.Navigation.query());
      }}
      onFinishUriResolver={(response, locale) =>
        PublicJoinKeyEntity.Navigation.single(response.data?.uniqueId)
      }
      Form={PublicJoinKeyEditForm}
      onEditTitle={t.fb.editPublicJoinKey}
      onCreateTitle={t.fb.newPublicJoinKey}
      data={data}
    />
  );
};
