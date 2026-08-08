import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  type DtoEntity,
} from "../../fireback-ui/components/entity-manager/CommonEntityManager";
import { GsmProviderForm } from "./GsmProviderEditForm";
import { useGsmProviderGetActionQuery } from "../../sdk/messaging/GsmProviderGetAction";
import { useGsmProviderCreateAction } from "../../sdk/messaging/GsmProviderCreateAction";
import { useGsmProviderUpdateAction } from "../../sdk/messaging/GsmProviderUpdateAction";
import { GsmProviderDto } from "../../sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "../../sdk/navigation/MessagingNavigation";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
export const GsmProviderEntityManager = ({
  data,
}: DtoEntity<GsmProviderDto>) => {
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
        router.goBackOrDefault(GsmProviderNavigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        GsmProviderNavigation.single(response.data?.uniqueId, locale)
      }
      Form={GsmProviderForm}
      onEditTitle={s.gsmProviders.editGsmProvider}
      onCreateTitle={s.gsmProviders.newGsmProvider}
      data={data}
    />
  );
};
