import { useT } from "@/modules/fireback/hooks/useT";
import { useS } from "@/modules/fireback/hooks/useS";
import { columns } from "./EmailProviderColumns";
import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useGetEmailProviders } from "@/modules/fireback/sdk/modules/abac/useGetEmailProviders";
import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";
import { usePostEmailProviderRemove } from "@/modules/fireback/sdk/modules/abac/usePostEmailProviderRemove";
import { strings } from "./strings/translations";

export const EmailProviderList = () => {
  const t = useT();
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(t, s)}
        queryHook={useGetEmailProviders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailProviderEntity.Navigation.single(uniqueId)
        }
        deleteHook={usePostEmailProviderRemove}
      ></CommonListManager>
    </>
  );
};
