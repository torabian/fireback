import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { useEmailSenderGetActionQuery } from "../../../sdk/messaging/EmailSenderGetAction";
import { useEmailSenderCreateAction } from "../../../sdk/messaging/EmailSenderCreateAction";
import { useEmailSenderUpdateAction } from "../../../sdk/messaging/EmailSenderUpdateAction";
import { EmailSenderDto } from "../../../sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "../../../sdk/navigation/MessagingNavigation";
import { EmailSenderEditForm } from "./EmailSenderEditForm";

export const EmailSenderEntityManager = ({
  data,
}: DtoEntity<EmailSenderDto>) => {
  const { router, uniqueId, queryClient, locale, formik } =
    useCommonEntityManager<Partial<EmailSenderDto>>({
      data,
    });
  const s = useS(strings);

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
      onEditTitle={s.editMailSender}
      onCreateTitle={s.newMailSender}
      data={data}
    />
  );
};
