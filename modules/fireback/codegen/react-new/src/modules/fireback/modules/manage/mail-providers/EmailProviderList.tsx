import { useT } from "@/modules/fireback/hooks/useT";
import { columns } from "./EmailProviderColumns";
import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useGetEmailProviders } from "@/modules/fireback/sdk/modules/abac/useGetEmailProviders";
import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";
import { useDeleteEmailProvider } from "@/modules/fireback/sdk/modules/abac/useDeleteEmailProvider";

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
