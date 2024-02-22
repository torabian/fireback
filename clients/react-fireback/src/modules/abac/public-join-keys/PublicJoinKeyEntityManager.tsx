import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { PublicJoinKeyEntity } from "src/sdk/fireback";

import { useGetPublicJoinKeyByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetPublicJoinKeyByUniqueId";
import { usePatchPublicJoinKey } from "@/sdk/fireback/modules/workspaces/usePatchPublicJoinKey";
import { usePostPublicJoinKey } from "@/sdk/fireback/modules/workspaces/usePostPublicJoinKey";
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
        router.goBackOrDefault(`/${locale}/publicjoinkeys`);
      }}
      onFinishUriResolver={(response, locale) =>
        `/${locale}/publicjoinkey/${response.data?.uniqueId}`
      }
      Form={PublicJoinKeyEditForm}
      onEditTitle={t.fb.editPublicJoinKey}
      onCreateTitle={t.fb.newPublicJoinKey}
      data={data}
    />
  );
};
