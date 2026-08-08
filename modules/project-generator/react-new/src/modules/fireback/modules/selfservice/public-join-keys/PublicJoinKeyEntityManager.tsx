import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { usePublicJoinKeyGetActionQuery } from "../../../sdk/abac/PublicJoinKeyGetAction";
import { usePublicJoinKeyCreateAction } from "../../../sdk/abac/PublicJoinKeyCreateAction";
import { usePublicJoinKeyUpdateAction } from "../../../sdk/abac/PublicJoinKeyUpdateAction";
import { PublicJoinKeyDto } from "../../../sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "../../../sdk/navigation/AbacNavigation";
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
