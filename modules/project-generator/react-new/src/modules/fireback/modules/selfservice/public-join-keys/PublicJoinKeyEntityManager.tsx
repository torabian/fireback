import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { usePublicJoinKeyGetActionQuery } from "@/modules/fireback/sdk/abac/PublicJoinKeyGetAction";
import { usePublicJoinKeyCreateAction } from "@/modules/fireback/sdk/abac/PublicJoinKeyCreateAction";
import { usePublicJoinKeyUpdateAction } from "@/modules/fireback/sdk/abac/PublicJoinKeyUpdateAction";
import { PublicJoinKeyDto } from "@/modules/fireback/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { PublicJoinKeyEditForm } from "./PublicJoinKeyEditForm";

export const PublicJoinKeyEntityManager = ({
  data,
}: DtoEntity<PublicJoinKeyDto>) => {
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
    Partial<PublicJoinKeyDto>
  >({
    data,
  });

  const getSingleHook = usePublicJoinKeyGetActionQuery({
    params: { uniqueId },
  });

  const postHook = usePublicJoinKeyCreateAction({});

  const patchHook = usePublicJoinKeyUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(PublicJoinKeyNavigation.query());
      }}
      onFinishUriResolver={(response, locale) =>
        PublicJoinKeyNavigation.single(response.data?.uniqueId)
      }
      Form={PublicJoinKeyEditForm}
      onEditTitle={t.fb.editPublicJoinKey}
      onCreateTitle={t.fb.newPublicJoinKey}
      data={data}
    />
  );
};
