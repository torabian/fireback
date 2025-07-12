import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { InvoiceForm } from "./InvoiceEditForm";
import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
import { useGetInvoiceByUniqueId } from "@/modules/fireback/sdk/modules/payment/useGetInvoiceByUniqueId";
import { usePostInvoice } from "@/modules/fireback/sdk/modules/payment/usePostInvoice";
import { usePatchInvoice } from "@/modules/fireback/sdk/modules/payment/usePatchInvoice";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const InvoiceEntityManager = ({ data }: DtoEntity<InvoiceEntity>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<InvoiceEntity>
  >({
    data,
  });
  const getSingleHook = useGetInvoiceByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostInvoice({
    queryClient,
  });
  const patchHook = usePatchInvoice({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          InvoiceEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        InvoiceEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ InvoiceForm }
      onEditTitle={s.invoices.editInvoice }
      onCreateTitle={s.invoices.newInvoice }
      data={data}
    />
  );
};
