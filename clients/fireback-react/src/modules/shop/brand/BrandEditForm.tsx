import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { BrandEntity } from "src/sdk/fireback/modules/shop/BrandEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
export const BrandForm = ({
  form,
  isEditing,
}: EntityFormProps<BrandEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) =>
          setFieldValue(BrandEntity.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.brands.name}
        hint={t.brands.nameHint}
        autoFocus={!isEditing}
      />
    </>
  );
};
