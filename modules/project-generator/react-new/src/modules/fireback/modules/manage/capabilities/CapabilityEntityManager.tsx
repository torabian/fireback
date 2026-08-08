import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { useCapabilityCreateAction } from "@/modules/fireback/sdk/abac/CapabilityCreateAction";
import { useCapabilityGetActionQuery } from "@/modules/fireback/sdk/abac/CapabilityGetAction";
import { useCapabilityUpdateAction } from "@/modules/fireback/sdk/abac/CapabilityUpdateAction";
import { CapabilityDto } from "@/modules/fireback/sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { CapabilityForm } from "./CapabilityEditForm";
import { strings } from "./strings/translations";
export const CapabilityEntityManager = ({
  data,
}: DtoEntity<CapabilityDto>) => {
  const s = useS(strings);
  const { router, uniqueId, locale } = useCommonEntityManager<
    Partial<CapabilityDto>
  >({
    data,
  });
  const getSingleHook = useCapabilityGetActionQuery({
    params: { uniqueId },
  });
  const postHook = useCapabilityCreateAction({});
  const patchHook = useCapabilityUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          CapabilityNavigation.query(undefined, locale),
        );
      }}
      onFinishUriResolver={(response, locale) =>
        CapabilityNavigation.single(response.data?.uniqueId, locale)
      }
      Form={CapabilityForm}
      onEditTitle={s.capabilities.editCapability}
      onCreateTitle={s.capabilities.newCapability}
      data={data}
    />
  );
};
