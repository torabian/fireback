import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useMessagingConfigGetActionQuery } from "@fireback/manage/sdk/messaging/MessagingConfigGetAction";
import { MessagingConfigNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { strings } from "./strings/translations";

export const MessagingConfigSingleScreen = () => {
  const getSingleHook = useMessagingConfigGetActionQuery({});
  var d = getSingleHook.data?.data?.item;

  const t = useS(strings);

  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(MessagingConfigNavigation.edit(""));
        }}
        noBack
        disableOnGetFailed
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          title={t.messagingConfigs.title}
          description={t.messagingConfigs.description}
          entity={d}
          fields={[
            {
              elem: d?.generalEmailProviderId,
              label: t.messagingConfigs.generalEmailProviderLabel,
            },
            {
              elem: d?.generalGsmProviderId,
              label: t.messagingConfigs.generalGsmProviderLabel,
            },
            {
              elem: d?.inviteToWorkspaceContentId,
              label: t.messagingConfigs.inviteToWorkspaceContentLabel,
            },
            {
              elem: d?.emailOtpContentId,
              label: t.messagingConfigs.emailOtpContentLabel,
            },
            {
              elem: d?.smsOtpContentId,
              label: t.messagingConfigs.smsOtpContentLabel,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
