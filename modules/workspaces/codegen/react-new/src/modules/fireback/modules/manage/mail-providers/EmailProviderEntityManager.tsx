import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";
import { useGetEmailProviderByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetEmailProviderByUniqueId";
import { usePatchEmailProvider } from "@/modules/fireback/sdk/modules/abac/usePatchEmailProvider";
import { usePostEmailProvider } from "@/modules/fireback/sdk/modules/abac/usePostEmailProvider";
import { EmailProviderEditForm } from "./EmailProviderEditForm";

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
