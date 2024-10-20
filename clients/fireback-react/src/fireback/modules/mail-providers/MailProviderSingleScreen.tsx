import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/fireback/components/page-title/PageTitle";
import { useLocale } from "@/fireback/hooks/useLocale";
import { useT } from "@/fireback/hooks/useT";
import { useGetEmailProviderByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetEmailProviderByUniqueId";
import { EmailProviderEntity } from "@/sdk/fireback/modules/workspaces/EmailProviderEntity";

export const EmailProviderSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetEmailProviderByUniqueId({
    query: { uniqueId },
  });
  var d: EmailProviderEntity | undefined = getSingleHook.query.data?.data;

  usePageTitle(d?.type || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(EmailProviderEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: t.mailProvider.type,
              elem: <span>{d?.type}</span>,
            },
            {
              label: t.mailProvider.apiKey,
              elem: <pre dir="ltr">{d?.apiKey}</pre>,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
