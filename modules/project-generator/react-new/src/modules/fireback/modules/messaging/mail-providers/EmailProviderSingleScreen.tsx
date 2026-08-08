import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/modules/fireback/hooks/authContext";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { useT } from "@/modules/fireback/hooks/useT";
import { useS } from "@/modules/fireback/hooks/useS";
import { EmailProviderDto } from "@/modules/fireback/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
import {
  SendEmailActionReq,
  useSendEmailAction,
} from "@/modules/fireback/sdk/messaging/SendEmailAction";
import { useEmailProviderGetActionQuery } from "@/modules/fireback/sdk/messaging/EmailProviderGetAction";
import { strings } from "./strings/translations";

export const EmailProviderSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const { mutateAsync } = useSendEmailAction();

  const getSingleHook = useGetEmailProviderByUniqueId({
    query: { uniqueId },
  });
  var d: EmailProviderDto | undefined = getSingleHook.query.data?.data;

  usePageTitle(d?.type || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(EmailProviderNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: s.emailProviders.title,
              elem: d?.title,
            },
            {
              label: t.mailProvider.type,
              elem: d?.type,
            },
          ]}
        />

        <button
          className="btn mt-5 btn-success btn-sm"
          onClick={() => {
            mutateAsync(
              SendEmailActionReq.from({
                providerId: uniqueId,
                body: "±",
                toAddress: "asdad",
              }),
            )
              .then((res) => {
                console.log(res);
                alert(s.emailProviders.mailSent);
              })
              .catch((err) => {
                alert(`${err}`);
              });
          }}
        >
          {s.emailProviders.sendTestEmail}
        </button>
      </CommonSingleManager>
    </>
  );
};
