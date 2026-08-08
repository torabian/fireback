import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
import { useT } from "../../../hooks/useT";

import { EmailProviderDto } from "../../../sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "../../../sdk/navigation/MessagingNavigation";
import { EmailProviderList } from "./EmailProviderList";

export const EmailProviderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailProviders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailProviderNavigation.create());
        }}
      >
        <EmailProviderList />
      </CommonArchiveManager>
    </>
  );
};
