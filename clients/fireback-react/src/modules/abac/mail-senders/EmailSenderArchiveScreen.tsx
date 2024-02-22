import { useT } from "@/hooks/useT";
import { EmailSenderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-sender-navigation-tools";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { EmailSenderList } from "./EmailSenderList";

export const EmailSenderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailSenders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailSenderNavigationTools.create(locale));
        }}
      >
        <EmailSenderList />
      </CommonArchiveManager>
    </>
  );
};
