import { useT } from "@/hooks/useT";
import { EmailProviderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-provider-navigation-tools";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { EmailProviderList } from "./MailProviderList";

export const EmailProviderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailProviders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailProviderNavigationTools.create(locale));
        }}
      >
        <EmailProviderList />
      </CommonArchiveManager>
    </>
  );
};
