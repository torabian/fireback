import {
  CommonEntityManager,
  type DtoEntity,
} from "../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { useEmailProviderGetActionQuery } from "../../sdk/messaging/EmailProviderGetAction";
import { useEmailProviderCreateAction } from "../../sdk/messaging/EmailProviderCreateAction";
import { useEmailProviderUpdateAction } from "../../sdk/messaging/EmailProviderUpdateAction";
import { EmailProviderDto } from "../../sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "../../sdk/navigation/MessagingNavigation";
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
