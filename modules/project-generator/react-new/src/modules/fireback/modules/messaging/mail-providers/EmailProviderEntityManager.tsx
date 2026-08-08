import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useEmailProviderGetActionQuery } from "@/modules/fireback/sdk/messaging/EmailProviderGetAction";
import { useEmailProviderCreateAction } from "@/modules/fireback/sdk/messaging/EmailProviderCreateAction";
import { useEmailProviderUpdateAction } from "@/modules/fireback/sdk/messaging/EmailProviderUpdateAction";
import { EmailProviderDto } from "@/modules/fireback/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
import { EmailProviderEditForm } from "./EmailProviderEditForm";

export const EmailProviderEntityManager = ({
  data,
}: DtoEntity<EmailProviderDto>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<EmailProviderDto>
  >({
    data,
  });

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
      onEditTitle={t.fb.editMailProvider}
      onCreateTitle={t.fb.newMailProvider}
      data={data}
    />
  );
};
