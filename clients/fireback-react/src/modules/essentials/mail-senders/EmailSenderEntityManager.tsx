import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useGetEmailSenderByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetEmailSenderByUniqueId";
import { usePatchEmailSender } from "src/sdk/fireback/modules/workspaces/usePatchEmailSender";
import { usePostEmailSender } from "src/sdk/fireback/modules/workspaces/usePostEmailSender";

import { useT } from "@/hooks/useT";
import { EmailSenderEditForm } from "./EmailSenderEditForm";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { EmailSenderEntity } from "@/sdk/fireback/modules/workspaces/EmailSenderEntity";

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
