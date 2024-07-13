import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";

import { useGetPublicJoinKeyByUniqueId } from "../../sdk/modules/workspaces/useGetPublicJoinKeyByUniqueId";
import { usePatchPublicJoinKey } from "../../sdk/modules/workspaces/usePatchPublicJoinKey";
import { usePostPublicJoinKey } from "../../sdk/modules/workspaces/usePostPublicJoinKey";
import { PublicJoinKeyEditForm } from "./PublicJoinKeyEditForm";
import { PublicJoinKeyEntity } from "../../sdk/modules/workspaces/PublicJoinKeyEntity";

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
