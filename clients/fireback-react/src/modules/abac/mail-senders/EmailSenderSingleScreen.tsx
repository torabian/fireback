import { useRouter } from "@/Router";
import { useActions, useEditAction } from "@/components/action-menu/ActionMenu";
import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { QueryErrorView } from "@/components/error-view/QueryError";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { PageSection } from "@/components/page-section/PageSection";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import { EmailSenderEntity } from "src/sdk/fireback";
import { EmailSenderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-sender-navigation-tools";
import { useGetEmailSenderByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetEmailSenderByUniqueId";
import { useState } from "react";
import { useQueryClient } from "react-query";

export const EmailSenderSingleScreen = () => {
  const router = useRouter();
  const queryClient = useQueryClient();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();
  const [value, setValue] = useState<string[]>([]);

  const getSingleHook = useGetEmailSenderByUniqueId({
    query: { uniqueId },
  });
  var d: EmailSenderEntity | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.fromName || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(EmailSenderNavigationTools.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: t.mailProvider.fromName,
              elem: d?.fromName,
            },
            {
              label: t.mailProvider.fromEmailAddress,
              elem: d?.fromEmailAddress,
            },
            {
              label: t.mailProvider.nickName,
              elem: d?.nickName,
            },
            {
              label: t.mailProvider.replyTo,
              elem: d?.replyTo,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
