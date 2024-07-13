import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/modules/fireback/components/page-title/PageTitle";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useT } from "@/modules/fireback/hooks/useT";
import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { EmailSenderEntity } from "../../sdk/modules/workspaces/EmailSenderEntity";
import { useState } from "react";
import { useGetEmailSenderByUniqueId } from "../../sdk/modules/workspaces/useGetEmailSenderByUniqueId";

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
