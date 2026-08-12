import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { EmailSenderDto } from "@fireback/messaging/sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { useEmailSenderGetActionQuery } from "@fireback/messaging/sdk/messaging/EmailSenderGetAction";
import { useState } from "react";
import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";

export const EmailSenderSingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();
  const [value, setValue] = useState<string[]>([]);

  const getSingleHook = useEmailSenderGetActionQuery({
    params: { uniqueId },
  });
  var d: EmailSenderDto | undefined = getSingleHook.data?.data?.item;
  usePageTitle(d?.fromName || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(EmailSenderNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: s.fromName,
              elem: d?.fromName,
            },
            {
              label: s.fromEmailAddress,
              elem: d?.fromEmailAddress,
            },
            {
              label: s.nickName,
              elem: d?.nickName,
            },
            {
              label: s.replyTo,
              elem: d?.replyTo,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
