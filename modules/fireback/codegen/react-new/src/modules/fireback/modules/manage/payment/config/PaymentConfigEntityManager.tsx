import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { PaymentConfigEntity } from "@/modules/fireback/sdk/modules/payment/PaymentConfigEntity";
import { useGetPaymentConfigDistinct } from "@/modules/fireback/sdk/modules/payment/useGetPaymentConfigDistinct";
import { usePatchPaymentConfigDistinct } from "@/modules/fireback/sdk/modules/payment/usePatchPaymentConfigDistinct";
import { PaymentConfigForm } from "./PaymentConfigEditForm";
import { strings } from "./strings/translations";

export const PaymentConfigEntityManager = ({
  data,
}: DtoEntity<PaymentConfigEntity>) => {
  const s = useS(strings);
  const { router, queryClient, locale } = useCommonEntityManager<
    Partial<PaymentConfigEntity>
  >({
    data,
  });

  const uniqueId = "workspace";

  const getSingleHook = useGetPaymentConfigDistinct({
    query: { uniqueId },
  });

  const patchHook = usePatchPaymentConfigDistinct({
    queryClient,
  });
  return (
    <CommonEntityManager
      patchHook={patchHook}
      forceEdit
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          PaymentConfigEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PaymentConfigEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PaymentConfigForm}
      onEditTitle={s.paymentConfigs.editPaymentConfig}
      onCreateTitle={s.paymentConfigs.newPaymentConfig}
      data={data}
    />
  );
};
