import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { usePublicJoinKeyGetActionQuery } from "@fireback/selfservice/sdk/abac/PublicJoinKeyGetAction";
import { usePublicJoinKeyCreateAction } from "@fireback/selfservice/sdk/abac/PublicJoinKeyCreateAction";
import { usePublicJoinKeyUpdateAction } from "@fireback/selfservice/sdk/abac/PublicJoinKeyUpdateAction";
import { PublicJoinKeyDto } from "@fireback/selfservice/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { PublicJoinKeyEditForm } from "./PublicJoinKeyEditForm";

export const PublicJoinKeyEntityManager = ({
  data,
}: DtoEntity<PublicJoinKeyDto>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<PublicJoinKeyDto>
  >({
    data,
  });
  const s = useS(strings);

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
      onEditTitle={s.editPublicJoinKey}
      onCreateTitle={s.newPublicJoinKey}
      data={data}
    />
  );
};
