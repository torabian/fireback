import { usePageTitle } from "@/modules/fireback-ui/components/page-title/PageTitle";
import { CommonSingleManager } from "../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { useLocale } from "../../fireback-ui/hooks/useLocale";
import { useRouter } from "../../fireback-ui/hooks/useRouter";
import { useS } from "../../fireback-ui/hooks/useS";
import { EmailProviderDto } from "../../sdk/messaging/EmailProviderDto";
import { useEmailProviderGetActionQuery } from "../../sdk/messaging/EmailProviderGetAction";
import {
  SendEmailActionReq,
  useSendEmailAction,
} from "../../sdk/messaging/SendEmailAction";
import { EmailProviderNavigation } from "../../sdk/navigation/MessagingNavigation";
import { strings } from "./strings/translations";

export const EmailProviderSingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const { mutateAsync } = useSendEmailAction();

  const getSingleHook = useEmailProviderGetActionQuery({
    params: { uniqueId },
  });
  var d: EmailProviderDto | undefined = getSingleHook.data?.data?.item;

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
              label: s.type,
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
