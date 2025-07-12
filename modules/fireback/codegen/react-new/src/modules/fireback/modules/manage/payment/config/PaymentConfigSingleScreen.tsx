import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { PaymentConfigEntity } from "@/modules/fireback/sdk/modules/payment/PaymentConfigEntity";
import { useGetPaymentConfigDistinct } from "@/modules/fireback/sdk/modules/payment/useGetPaymentConfigDistinct";
import { strings } from "./strings/translations";

export const PaymentConfigSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetPaymentConfigDistinct({ query: { uniqueId } });
  var d: PaymentConfigEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push("../config/edit");
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.stripeSecretKey,
              label: t.paymentConfigs.stripeSecretKey,
            },
            {
              elem: d?.stripeCallbackUrl,
              label: t.paymentConfigs.stripeCallbackUrl,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
