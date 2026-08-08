import { useT } from "@/modules/fireback/hooks/useT";
import { EmailSenderList } from "./EmailSenderList";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { EmailSenderDto } from "@/modules/fireback/sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";

export const EmailSenderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailSenders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailSenderNavigation.create());
        }}
      >
        <EmailSenderList />
      </CommonArchiveManager>
    </>
  );
};
