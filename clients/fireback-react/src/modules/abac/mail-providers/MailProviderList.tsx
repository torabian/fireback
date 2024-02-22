import { useT } from "@/hooks/useT";
import { EmailProviderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-provider-navigation-tools";
import { useDeleteEmailProvider } from "src/sdk/fireback/modules/workspaces/useDeleteEmailProvider";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useGetEmailProviders } from "src/sdk/fireback/modules/workspaces/useGetEmailProviders";
import { columns } from "./MailProviderColumns";

export const EmailProviderList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetEmailProviders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailProviderNavigationTools.single(uniqueId)
        }
        deleteHook={useDeleteEmailProvider}
      ></CommonListManager>
    </>
  );
};
