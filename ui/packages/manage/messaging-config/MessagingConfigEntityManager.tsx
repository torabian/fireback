import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useMessagingConfigGetActionQuery } from "@fireback/manage/sdk/messaging/MessagingConfigGetAction";
import { useMessagingConfigUpdateAction } from "@fireback/manage/sdk/messaging/MessagingConfigUpdateAction";
import { MessagingConfigDto } from "@fireback/manage/sdk/messaging/MessagingConfigDto";
import { MessagingConfigNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { MessagingConfigForm } from "./MessagingConfigEditForm";
import { strings } from "./strings/translations";

export const MessagingConfigEntityManager = ({
  data,
}: DtoEntity<MessagingConfigDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<MessagingConfigDto>
  >({
    data,
  });
  const getSingleHook = useMessagingConfigGetActionQuery({});
  const patchHook = useMessagingConfigUpdateAction({});

  return (
    <CommonEntityManager
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      disableOnGetFailed
      forceEdit
      onCancel={() => {
        router.goBackOrDefault(
          MessagingConfigNavigation.single(undefined, locale),
        );
      }}
      onFinishUriResolver={(response, locale) =>
        MessagingConfigNavigation.single(response.data?.uniqueId, locale)
      }
      customClass="w-100"
      Form={MessagingConfigForm}
      onEditTitle={s.messagingConfigs.editMessagingConfig}
      onCreateTitle={s.messagingConfigs.newMessagingConfig}
      data={data}
    />
  );
};
