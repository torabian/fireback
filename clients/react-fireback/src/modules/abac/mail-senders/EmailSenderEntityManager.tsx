import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useGetEmailSenderByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetEmailSenderByUniqueId";
import { usePatchEmailSender } from "src/sdk/fireback/modules/workspaces/usePatchEmailSender";
import { usePostEmailSender } from "src/sdk/fireback/modules/workspaces/usePostEmailSender";

import { useT } from "@/hooks/useT";
import { EmailSenderEntity } from "src/sdk/fireback";
import { EmailSenderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-sender-navigation-tools";
import { EmailSenderEditForm } from "./EmailSenderEditForm";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";

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
          EmailSenderNavigationTools.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        EmailSenderNavigationTools.single(response.data?.uniqueId, locale)
      }
      Form={EmailSenderEditForm}
      onEditTitle={t.fb.editMailSender}
      onCreateTitle={t.fb.newMailSender}
      data={data}
    />
  );
};
