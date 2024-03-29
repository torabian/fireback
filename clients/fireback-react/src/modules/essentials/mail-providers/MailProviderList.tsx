import { useT } from "@/hooks/useT";
import { useDeleteEmailProvider } from "src/sdk/fireback/modules/workspaces/useDeleteEmailProvider";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useGetEmailProviders } from "src/sdk/fireback/modules/workspaces/useGetEmailProviders";
import { columns } from "./MailProviderColumns";
import { EmailProviderEntity } from "@/sdk/fireback/modules/workspaces/EmailProviderEntity";

export const EmailProviderList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetEmailProviders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailProviderEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteEmailProvider}
      ></CommonListManager>
    </>
  );
};
