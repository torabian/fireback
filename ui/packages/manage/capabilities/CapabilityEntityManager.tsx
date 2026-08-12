import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useCapabilityCreateAction } from "@fireback/manage/sdk/abac/CapabilityCreateAction";
import { useCapabilityGetActionQuery } from "@fireback/manage/sdk/abac/CapabilityGetAction";
import { useCapabilityUpdateAction } from "@fireback/manage/sdk/abac/CapabilityUpdateAction";
import { CapabilityDto } from "@fireback/manage/sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { CapabilityForm } from "./CapabilityEditForm";
import { strings } from "./strings/translations";
export const CapabilityEntityManager = ({ data }: DtoEntity<CapabilityDto>) => {
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
        router.goBackOrDefault(CapabilityNavigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        CapabilityNavigation.single(response.data?.item?.uniqueId, locale)
      }
      Form={CapabilityForm}
      onEditTitle={s.capabilities.editCapability}
      onCreateTitle={s.capabilities.newCapability}
      data={data}
    />
  );
};
