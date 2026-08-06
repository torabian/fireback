import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { GsmProviderForm } from "./GsmProviderEditForm";
import { GsmProviderEntity } from "@/modules/fireback/sdk/modules/abac/GsmProviderEntity";
import { useGetGsmProviderByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetGsmProviderByUniqueId";
import { usePostGsmProvider } from "@/modules/fireback/sdk/modules/abac/usePostGsmProvider";
import { usePatchGsmProvider } from "@/modules/fireback/sdk/modules/abac/usePatchGsmProvider";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const GsmProviderEntityManager = ({ data }: DtoEntity<GsmProviderEntity>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<GsmProviderEntity>
  >({
    data,
  });
  const getSingleHook = useGetGsmProviderByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostGsmProvider({
    queryClient,
  });
  const patchHook = usePatchGsmProvider({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          GsmProviderEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        GsmProviderEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ GsmProviderForm }
      onEditTitle={s.gsmProviders.editGsmProvider }
      onCreateTitle={s.gsmProviders.newGsmProvider }
      data={data}
    />
  );
};
