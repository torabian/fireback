import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { useCapabilityCreateAction } from "../../../sdk/abac/CapabilityCreateAction";
import { useCapabilityGetActionQuery } from "../../../sdk/abac/CapabilityGetAction";
import { useCapabilityUpdateAction } from "../../../sdk/abac/CapabilityUpdateAction";
import { CapabilityDto } from "../../../sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "../../../sdk/navigation/AbacNavigation";
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
