import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { useEmailProviderGetActionQuery } from "@fireback/messaging/sdk/messaging/EmailProviderGetAction";
import { useEmailProviderCreateAction } from "@fireback/messaging/sdk/messaging/EmailProviderCreateAction";
import { useEmailProviderUpdateAction } from "@fireback/messaging/sdk/messaging/EmailProviderUpdateAction";
import { EmailProviderDto } from "@fireback/messaging/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { EmailProviderEditForm } from "./EmailProviderEditForm";

export const EmailProviderEntityManager = ({
  data,
}: DtoEntity<EmailProviderDto>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<EmailProviderDto>
  >({
    data,
  });
  const s = useS(strings);

  const getSingleHook = useEmailProviderGetActionQuery({
    params: { uniqueId },
  });

  const postHook = useEmailProviderCreateAction({});

  const patchHook = useEmailProviderUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(
          EmailProviderNavigation.query(undefined, locale),
        );
      }}
      onFinishUriResolver={(response, locale) =>
        EmailProviderNavigation.single(response.data?.uniqueId, locale)
      }
      Form={EmailProviderEditForm}
      onEditTitle={s.editMailProvider}
      onCreateTitle={s.newMailProvider}
      data={data}
    />
  );
};
