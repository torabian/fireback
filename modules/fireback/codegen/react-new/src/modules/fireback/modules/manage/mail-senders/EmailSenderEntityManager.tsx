import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailSenderEntity } from "@/modules/fireback/sdk/modules/abac/EmailSenderEntity";
import { useGetEmailSenderByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetEmailSenderByUniqueId";
import { usePatchEmailSender } from "@/modules/fireback/sdk/modules/abac/usePatchEmailSender";
import { usePostEmailSender } from "@/modules/fireback/sdk/modules/abac/usePostEmailSender";
import { EmailSenderEditForm } from "./EmailSenderEditForm";

export const EmailSenderEntityManager = ({
  data,
}: DtoEntity<EmailSenderEntity>) => {
  const { router, uniqueId, queryClient, locale, formik } =
    useCommonEntityManager<Partial<EmailSenderEntity>>({
      data,
    });
  const t = useT();

  const getSingleHook = useGetEmailSenderByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostEmailSender({
    queryClient,
  });

  const patchHook = usePatchEmailSender({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(
          EmailSenderEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        EmailSenderEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={EmailSenderEditForm}
      onEditTitle={t.fb.editMailSender}
      onCreateTitle={t.fb.newMailSender}
      data={data}
    />
  );
};
