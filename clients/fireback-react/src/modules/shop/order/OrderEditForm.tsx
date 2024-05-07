import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/fireback/components/forms/form-select/FormEntitySelect3";
import { useGetPaymentStatuses } from "@/sdk/fireback/modules/shop/useGetPaymentStatuses";
import { useGetOrderStatuses } from "@/sdk/fireback/modules/shop/useGetOrderStatuses";
import { useGetDiscountCodes } from "@/sdk/fireback/modules/shop/useGetDiscountCodes";
export const OrderForm = ({
  form,
  isEditing,
}: EntityFormProps<OrderEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      {/*
          <FormText
            type="?"
            value={values.totalPrice }
            onChange={(value) => setFieldValue(OrderEntity.Fields.totalPrice, value, false)}
            errorMessage={errors.totalPrice }
            label={t.orders.totalPrice }
            hint={t.orders.totalPriceHint}
          />
         */}
      <FormText
        autoFocus={!isEditing}
        value={values.shippingAddress}
        onChange={(value) =>
          setFieldValue(OrderEntity.Fields.shippingAddress, value, false)
        }
        errorMessage={errors.shippingAddress}
        label={t.orders.shippingAddress}
        hint={t.orders.shippingAddressHint}
      />
      <FormEntitySelect3
        formEffect={{ form, field: OrderEntity.Fields.paymentStatus$ }}
        useQuery={useGetPaymentStatuses}
        label={t.orders.paymentStatus}
        hint={t.orders.paymentStatusHint}
      />
      <FormEntitySelect3
        formEffect={{ form, field: OrderEntity.Fields.orderStatus$ }}
        useQuery={useGetOrderStatuses}
        label={t.orders.orderStatus}
        hint={t.orders.orderStatusHint}
      />
      <FormText
        value={values.invoiceNumber}
        onChange={(value) =>
          setFieldValue(OrderEntity.Fields.invoiceNumber, value, false)
        }
        errorMessage={errors.invoiceNumber}
        label={t.orders.invoiceNumber}
        hint={t.orders.invoiceNumberHint}
      />
      <FormEntitySelect3
        formEffect={{ form, field: OrderEntity.Fields.discountCode$ }}
        useQuery={useGetDiscountCodes}
        label={t.orders.discountCode}
        hint={t.orders.discountCodeHint}
      />
      {/*
          <FormText
            type="?"
            value={values.items }
            onChange={(value) => setFieldValue(OrderEntity.Fields.items, value, false)}
            errorMessage={errors.items }
            label={t.orders.items }
            hint={t.orders.itemsHint}
          />
         */}
    </>
  );
};
