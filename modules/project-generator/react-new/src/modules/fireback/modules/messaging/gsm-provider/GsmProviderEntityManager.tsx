import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { GsmProviderForm } from "./GsmProviderEditForm";
import { useGsmProviderGetActionQuery } from "@/modules/fireback/sdk/messaging/GsmProviderGetAction";
import { useGsmProviderCreateAction } from "@/modules/fireback/sdk/messaging/GsmProviderCreateAction";
import { useGsmProviderUpdateAction } from "@/modules/fireback/sdk/messaging/GsmProviderUpdateAction";
import { GsmProviderDto } from "@/modules/fireback/sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const GsmProviderEntityManager = ({ data }: DtoEntity<GsmProviderDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<GsmProviderDto>
  >({
    data,
  });
  const getSingleHook = useGsmProviderGetActionQuery({
    params: { uniqueId },
  });
  const postHook = useGsmProviderCreateAction({});
  const patchHook = useGsmProviderUpdateAction({ params: { uniqueId } });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          GsmProviderNavigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        GsmProviderNavigation.single(response.data?.uniqueId, locale)
      }
      Form={ GsmProviderForm }
      onEditTitle={s.gsmProviders.editGsmProvider }
      onCreateTitle={s.gsmProviders.newGsmProvider }
      data={data}
    />
  );
};
