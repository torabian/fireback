import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { ShoppingCartEntity } from "src/sdk/fireback/modules/shop/ShoppingCartEntity";
import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/fireback/components/forms/form-select/FormEntitySelect3";
import { Repeater } from "@/fireback/components/repeater/Repeater";
import { useGetProductSubmissions } from "@/sdk/fireback/modules/shop/useGetProductSubmissions";
import { FormikProps } from "formik";

export const ShoppingCartItem = ({
  form,
  isEditing,
  index,
  disabled,
}: {
  isEditing?: boolean;
  form: FormikProps<Partial<ShoppingCartEntity>>;
  index: number;
  disabled?: boolean;
}) => {
  const { values, setFieldValue } = form;
  const t = useT();
  if (!values.items) {
    return null;
  }

  return (
    <>
      <FormEntitySelect3
        formEffect={{ form, field: `items[${index}].product` }}
        useQuery={useGetProductSubmissions}
        disabled={disabled}
        label={t.shoppingCarts.product}
        hint={t.shoppingCarts.productHint}
      />
      <FormText
        type="number"
        value={values.items[index].quantity}
        onChange={(value) =>
          setFieldValue(`items[${index}].quantity`, value, false)
        }
        label={t.shoppingCarts.quantity}
        hint={t.shoppingCarts.quantityHint}
      />
    </>
  );
};

export const ShoppingCartForm = ({
  form,
  isEditing,
}: EntityFormProps<ShoppingCartEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <Repeater
        value={values.items}
        label={"New Item"}
        form={form}
        Component={ShoppingCartItem}
        onChange={(value) => {
          setFieldValue(ShoppingCartEntity.Fields.items$, value, false);
        }}
      />

      {/*
          <FormText
            type="?"
            value={values.items }
            onChange={(value) => setFieldValue(ShoppingCartEntity.Fields.items, value, false)}
            errorMessage={errors.items }
            label={t.shoppingCarts.items }
            hint={t.shoppingCarts.itemsHint}
          />
         */}
    </>
  );
};
