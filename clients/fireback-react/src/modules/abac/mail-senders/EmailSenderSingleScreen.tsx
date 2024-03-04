import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import { useRouter } from "@/Router";
import { EmailSenderEntity } from "@/sdk/fireback/modules/workspaces/EmailSenderEntity";
import { useState } from "react";
import { useGetEmailSenderByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetEmailSenderByUniqueId";

export const EmailSenderSingleScreen = () => {
  const router = useRouter();
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
          router.push(EmailSenderEntity.Navigation.edit(uniqueId, locale));
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
