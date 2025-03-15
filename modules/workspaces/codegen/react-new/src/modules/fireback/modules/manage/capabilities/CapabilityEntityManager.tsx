import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { CapabilityForm } from "./CapabilityEditForm";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/workspaces/CapabilityEntity";
import { useGetCapabilityByUniqueId } from "@/modules/fireback/sdk/modules/workspaces/useGetCapabilityByUniqueId";
import { usePostCapability } from "@/modules/fireback/sdk/modules/workspaces/usePostCapability";
import { usePatchCapability } from "@/modules/fireback/sdk/modules/workspaces/usePatchCapability";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const CapabilityEntityManager = ({ data }: DtoEntity<CapabilityEntity>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<CapabilityEntity>
  >({
    data,
  });
  const getSingleHook = useGetCapabilityByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostCapability({
    queryClient,
  });
  const patchHook = usePatchCapability({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          CapabilityEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        CapabilityEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ CapabilityForm }
      onEditTitle={s.capabilities.editCapability }
      onCreateTitle={s.capabilities.newCapability }
      data={data}
    />
  );
};
