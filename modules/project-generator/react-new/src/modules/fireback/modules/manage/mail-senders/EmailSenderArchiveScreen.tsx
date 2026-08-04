import { useT } from "@/modules/fireback/hooks/useT";
import { EmailSenderList } from "./EmailSenderList";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { EmailSenderEntity } from "@/modules/fireback/sdk/modules/abac/EmailSenderEntity";

export const EmailSenderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailSenders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailSenderEntity.Navigation.create());
        }}
      >
        <EmailSenderList />
      </CommonArchiveManager>
    </>
  );
};
