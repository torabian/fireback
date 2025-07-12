import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useGetInvoiceByUniqueId } from "@/modules/fireback/sdk/modules/payment/useGetInvoiceByUniqueId";
import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const InvoiceSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetInvoiceByUniqueId({ query: { uniqueId } });
  var d: InvoiceEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  const payInvoice = (uniqueId: string) => {
    window.open(`http://localhost:4500/payment/invoice/${uniqueId}`, "_blank");
  };
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(InvoiceEntity.Navigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.title,
              label: t.invoices.title,
            },
            {
              elem: d?.amount?.formatted,
              label: t.invoices.amount,
            },
            {
              elem: (
                <>
                  <button
                    className="btn btn-small"
                    onClick={() => payInvoice(uniqueId)}
                  >
                    Pay now
                  </button>
                </>
              ),
              label: "Actions",
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
