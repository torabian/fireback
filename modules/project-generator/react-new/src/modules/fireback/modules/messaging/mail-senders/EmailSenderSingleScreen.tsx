import { CommonSingleManager } from "../../../components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../../hooks/authContext";
import { useLocale } from "../../../hooks/useLocale";
import { useRouter } from "../../../hooks/useRouter";
import { useT } from "../../../hooks/useT";
import { EmailSenderDto } from "../../../sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "../../../sdk/navigation/MessagingNavigation";
import { useEmailSenderGetActionQuery } from "../../../sdk/messaging/EmailSenderGetAction";
import { useState } from "react";

export const EmailSenderSingleScreen = () => {
  const router = useRouter();
  const t = useT();
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
