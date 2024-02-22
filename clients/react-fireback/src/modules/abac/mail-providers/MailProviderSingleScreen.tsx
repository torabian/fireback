import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import { EmailProviderEntity } from "src/sdk/fireback";
import { EmailProviderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-provider-navigation-tools";
import { useGetEmailProviderByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetEmailProviderByUniqueId";

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
          router.push(EmailProviderNavigationTools.edit(uniqueId, locale));
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
