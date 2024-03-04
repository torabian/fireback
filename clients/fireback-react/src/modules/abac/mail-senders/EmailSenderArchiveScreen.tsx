import { useT } from "@/hooks/useT";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { EmailSenderList } from "./EmailSenderList";
import { EmailSenderEntity } from "@/sdk/fireback/modules/workspaces/EmailSenderEntity";

export const EmailSenderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailSenders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailSenderEntity.Navigation.create(locale));
        }}
      >
        <EmailSenderList />
      </CommonArchiveManager>
    </>
  );
};
