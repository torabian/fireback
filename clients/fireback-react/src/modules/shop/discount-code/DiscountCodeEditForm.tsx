import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { FormDate } from "@/components/forms/form-date/FormDate";
import { useGetCategories } from "@/sdk/fireback/modules/shop/useGetCategories";
import { useGetProductSubmissions } from "@/sdk/fireback/modules/shop/useGetProductSubmissions";
export const DiscountCodeForm = ({
  form,
  isEditing,
}: EntityFormProps<DiscountCodeEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.series}
        onChange={(value) =>
          setFieldValue(DiscountCodeEntity.Fields.series, value, false)
        }
        errorMessage={errors.series}
        label={t.discountCodes.series}
        hint={t.discountCodes.seriesHint}
      />
      <FormText
        type="number"
        value={values.limit}
        onChange={(value) =>
          setFieldValue(DiscountCodeEntity.Fields.limit, value, false)
        }
        errorMessage={errors.limit}
        label={t.discountCodes.limit}
        hint={t.discountCodes.limitHint}
      />

      <FormDate
        value={values.validStart}
        onChange={(value) => setFieldValue("validStart", value, false)}
        label={t.discountCodes.validFrom}
        hint={t.discountCodes.validFromHint}
      />

      <FormDate
        value={values.validEnd}
        onChange={(value) => setFieldValue("validEnd", value, false)}
        label={t.discountCodes.validUntil}
        hint={t.discountCodes.validUntilHint}
      />

      {/*
          <FormText
            type="?"
            value={values.validUntil }
            onChange={(value) => setFieldValue(DiscountCodeEntity.Fields.validUntil, value, false)}
            errorMessage={errors.validUntil }
            label={t.discountCodes.validUntil }
            hint={t.discountCodes.validUntilHint}
          />
         */}
      <FormEntitySelect3
        multiple
        formEffect={{ form, field: DiscountCodeEntity.Fields.appliedProducts$ }}
        useQuery={useGetProductSubmissions}
        label={t.discountCodes.appliedProducts}
        hint={t.discountCodes.appliedProductsHint}
      />
      <FormEntitySelect3
        multiple
        formEffect={{
          form,
          field: DiscountCodeEntity.Fields.excludedProducts$,
        }}
        useQuery={useGetProductSubmissions}
        label={t.discountCodes.excludedProducts}
        hint={t.discountCodes.excludedProductsHint}
      />
      <FormEntitySelect3
        multiple
        formEffect={{
          form,
          field: DiscountCodeEntity.Fields.appliedCategories$,
        }}
        useQuery={useGetCategories}
        label={t.discountCodes.appliedCategories}
        hint={t.discountCodes.appliedCategoriesHint}
      />
      <FormEntitySelect3
        multiple
        formEffect={{
          form,
          field: DiscountCodeEntity.Fields.excludedCategories$,
        }}
        useQuery={useGetCategories}
        label={t.discountCodes.excludedCategories}
        hint={t.discountCodes.excludedCategoriesHint}
      />
    </>
  );
};
