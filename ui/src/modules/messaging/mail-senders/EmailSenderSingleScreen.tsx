import { CommonSingleManager } from "../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../fireback-ui/hooks/authContext";
import { useLocale } from "../../fireback-ui/hooks/useLocale";
import { useRouter } from "../../fireback-ui/hooks/useRouter";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { EmailSenderDto } from "../../sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "../../sdk/navigation/MessagingNavigation";
import { useEmailSenderGetActionQuery } from "../../sdk/messaging/EmailSenderGetAction";
import { useState } from "react";

export const EmailSenderSingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();
  const [value, setValue] = useState<string[]>([]);

  const getSingleHook = useGetEmailSenderByUniqueId({
    query: { uniqueId },
  });
  var d: EmailSenderDto | undefined = getSingleHook.query.data?.data;
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
