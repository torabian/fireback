import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useEmailSenderGetActionQuery } from "@/modules/fireback/sdk/messaging/EmailSenderGetAction";
import { useEmailSenderCreateAction } from "@/modules/fireback/sdk/messaging/EmailSenderCreateAction";
import { useEmailSenderUpdateAction } from "@/modules/fireback/sdk/messaging/EmailSenderUpdateAction";
import { EmailSenderDto } from "@/modules/fireback/sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
import { EmailSenderEditForm } from "./EmailSenderEditForm";

export const EmailSenderEntityManager = ({
  data,
}: DtoEntity<EmailSenderDto>) => {
  const { router, uniqueId, queryClient, locale, formik } =
    useCommonEntityManager<Partial<EmailSenderDto>>({
      data,
    });
  const t = useT();

  const getSingleHook = useEmailSenderGetActionQuery({
    params: { uniqueId },
  });

  const postHook = useEmailSenderCreateAction({});

  const patchHook = useEmailSenderUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(
          EmailSenderNavigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        EmailSenderNavigation.single(response.data?.uniqueId, locale)
      }
      Form={EmailSenderEditForm}
      onEditTitle={t.fb.editMailSender}
      onCreateTitle={t.fb.newMailSender}
      data={data}
    />
  );
};
