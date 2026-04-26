import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/modules/fireback/hooks/authContext";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";
import {
  SendEmailActionReq,
  useSendEmailAction,
} from "@/modules/fireback/sdk/modules/abac/SendEmail";
import { useGetEmailProviderByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetEmailProviderByUniqueId";

export const EmailProviderSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const { mutateAsync } = useSendEmailAction();

  const getSingleHook = useGetEmailProviderByUniqueId({
    query: { uniqueId },
  });
  var d: EmailProviderEntity | undefined = getSingleHook.query.data?.data;

  usePageTitle(d?.type || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(EmailProviderEntity.Navigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: "Title",
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
                alert("Mail has been sent");
              })
              .catch((err) => {
                alert(`${err}`);
              });
          }}
        >
          Send a test email
        </button>
      </CommonSingleManager>
    </>
  );
};
