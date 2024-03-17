import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { useGetProducts } from "@/sdk/fireback/modules/shop/useGetProducts";
import { useGetCategories } from "@/sdk/fireback/modules/shop/useGetCategories";
export const ProductSubmissionForm = ({
  form,
  isEditing,
}: EntityFormProps<ProductSubmissionEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormEntitySelect3
        formEffect={{ form, field: ProductSubmissionEntity.Fields.product$ }}
        useQuery={useGetProducts}
        label={t.productsubmissions.product}
        hint={t.productsubmissions.productHint}
      />

      <FormText
        value={values.description}
        onChange={(value) =>
          setFieldValue(
            ProductSubmissionEntity.Fields.description,
            value,
            false
          )
        }
        errorMessage={errors.description}
        label={t.productsubmissions.description}
        hint={t.productsubmissions.descriptionHint}
      />
      <FormText
        value={values.sku}
        onChange={(value) =>
          setFieldValue(ProductSubmissionEntity.Fields.sku, value, false)
        }
        errorMessage={errors.sku}
        label={t.productsubmissions.sku}
        hint={t.productsubmissions.skuHint}
      />
      <FormText
        value={values.brand}
        onChange={(value) =>
          setFieldValue(ProductSubmissionEntity.Fields.brand, value, false)
        }
        errorMessage={errors.brand}
        label={t.productsubmissions.brand}
        hint={t.productsubmissions.brandHint}
      />
      <FormEntitySelect3
        formEffect={{ form, field: ProductSubmissionEntity.Fields.category$ }}
        useQuery={useGetCategories}
        label={t.productsubmissions.category}
        hint={t.productsubmissions.categoryHint}
      />
    </>
  );
};
