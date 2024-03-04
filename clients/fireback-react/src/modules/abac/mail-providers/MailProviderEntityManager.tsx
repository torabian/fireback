import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useGetEmailProviderByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetEmailProviderByUniqueId";
import { usePatchEmailProvider } from "src/sdk/fireback/modules/workspaces/usePatchEmailProvider";
import { usePostEmailProvider } from "src/sdk/fireback/modules/workspaces/usePostEmailProvider";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { EmailProviderEditForm } from "./MailProviderEditForm";
import { EmailProviderEntity } from "@/sdk/fireback/modules/workspaces/EmailProviderEntity";

export const EmailProviderEntityManager = ({
  data,
}: DtoEntity<EmailProviderEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<EmailProviderEntity>
  >({
    data,
  });

  const getSingleHook = useGetEmailProviderByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostEmailProvider({
    queryClient,
  });

  const patchHook = usePatchEmailProvider({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(
          EmailProviderEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        EmailProviderEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={EmailProviderEditForm}
      onEditTitle={t.fb.editMailProvider}
      onCreateTitle={t.fb.newMailProvider}
      data={data}
    />
  );
};
