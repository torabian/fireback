import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { useT } from "../../../hooks/useT";

import { EmailProviderDto } from "@/modules/fireback/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
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
