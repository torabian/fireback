import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useS } from "@/modules/fireback/hooks/useS";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
import { useContext } from "react";
import { strings } from "./strings/translations";
import { FormCurrency } from "@/modules/fireback/components/forms/form-currency/FormCurrency";

export const InvoiceForm = ({
  form,
  isEditing,
}: EntityFormProps<InvoiceEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
      <FormText
        value={values.title}
        onChange={(value) =>
          setFieldValue(InvoiceEntity.Fields.title, value, false)
        }
        errorMessage={errors.title}
        label={s.invoices.title}
        hint={s.invoices.titleHint}
      />

      <FormCurrency
        value={values.amount}
        onChange={(value) =>
          setFieldValue(InvoiceEntity.Fields.amount, value, false)
        }
        label={s.invoices.amount}
        hint={s.invoices.amountHint}
      />

      <FormText
        value={values.finalStatus}
        onChange={(value) =>
          setFieldValue(InvoiceEntity.Fields.finalStatus, value, false)
        }
        errorMessage={errors.finalStatus}
        label={s.invoices.finalStatus}
        hint={s.invoices.finalStatusHint}
      />
    </>
  );
};
