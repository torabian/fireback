import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/modules/fireback/components/page-title/PageTitle";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useT } from "@/modules/fireback/hooks/useT";
import { useGetEmailProviderByUniqueId } from "../../sdk/modules/workspaces/useGetEmailProviderByUniqueId";
import { EmailProviderEntity } from "../../sdk/modules/workspaces/EmailProviderEntity";

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
