import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
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
          value={values.series }
          onChange={(value) => setFieldValue(DiscountCodeEntity.Fields.series, value, false)}
          errorMessage={errors.series }
          label={t.discountCodes.series }
          hint={t.discountCodes.seriesHint}
        />
        <FormText
          type="number"
          value={values.limit }
          onChange={(value) => setFieldValue(DiscountCodeEntity.Fields.limit, value, false)}
          errorMessage={errors.limit }
          label={t.discountCodes.limit }
          hint={t.discountCodes.limitHint}
        />
    </>
  );
};