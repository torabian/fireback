import { FormEntitySelect3 } from "@/fireback/components/forms/form-select/FormEntitySelect3";
import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/fireback/definitions/definitions";
import { useT } from "@/fireback/hooks/useT";
import { useGetCategories } from "@/sdk/fireback/modules/shop/useGetCategories";
import { useGetProducts } from "@/sdk/fireback/modules/shop/useGetProducts";
import { useContext } from "react";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";

import { useGetBrands } from "@/sdk/fireback/modules/shop/useGetBrands";
import Form from "@rjsf/core";
import validator from "@rjsf/validator-ajv8";
import { useGetTags } from "@/sdk/fireback/modules/shop/useGetTags";
import { FormPriceTag } from "@/fireback/components/forms/form-price-tag/FormPriceTag";
import { FormRichText } from "@/fireback/components/forms/form-richtext/FormRichText";
import { FormUploader } from "@/fireback/components/forms/form-uploader/FormUploader";

function castIErrorToObjectArray(obj: any) {
  const items = {};
  const keys = Object.keys(obj);

  for (let key of keys) {
    items[key] = {
      __errors: [obj[key]],
    };
  }
  return items;
}

function InternalForm({ children, as }) {
  return <div className="rjsf">{children}</div>;
}

const ProductSubmissionFormSelected = ({
  form,
  isEditing,
}: EntityFormProps<ProductSubmissionEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const transformedError = castIErrorToObjectArray(errors);
  const t = useT();

  return (
    <>
      {values.product?.jsonSchema && (
        <Form
          tagName={"div"}
          schema={values.product?.jsonSchema}
          uiSchema={values.product?.uiSchema}
          formData={values.data}
          showErrorList={false}
          _internalFormWrapper={InternalForm}
          liveOmit
          extraErrors={transformedError}
          onChange={(data) => {
            setFieldValue(
              ProductSubmissionEntity.Fields.data,
              data.formData,
              false
            );
          }}
          validator={validator}
        />
      )}
    </>
  );
};

export const ProductSubmissionForm = ({
  form,
  isEditing,
}: EntityFormProps<ProductSubmissionEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <div className="headless-form-entity-manager">
        <FormEntitySelect3
          formEffect={{ form, field: ProductSubmissionEntity.Fields.product$ }}
          useQuery={useGetProducts}
          label={t.productsubmissions.product}
          hint={t.productsubmissions.productHint}
        />

        <FormText
          value={values.name}
          onChange={(value) =>
            setFieldValue(ProductSubmissionEntity.Fields.name, value, false)
          }
          errorMessage={errors.name}
          autoFocus={!isEditing}
          label={t.productsubmissions.name}
          hint={t.productsubmissions.nameHint}
        />
      </div>
      <FormRichText
        value={values.description}
        forceRich
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
      />{" "}
      <div className="headless-form-entity-manager">
        <FormText
          value={values.sku}
          onChange={(value) =>
            setFieldValue(ProductSubmissionEntity.Fields.sku, value, false)
          }
          errorMessage={errors.sku}
          label={t.productsubmissions.sku}
          hint={t.productsubmissions.skuHint}
        />
        <FormEntitySelect3
          formEffect={{ form, field: ProductSubmissionEntity.Fields.brand$ }}
          useQuery={useGetBrands}
          label={t.productsubmissions.brand}
          hint={t.productsubmissions.brandHint}
        />
        <FormPriceTag
          label={t.productsubmissions.price}
          hint={t.productsubmissions.priceHint}
          value={values.price}
          onChange={(value) =>
            setFieldValue(ProductSubmissionEntity.Fields.price$, value, false)
          }
        />
        <FormUploader
          value={values.image}
          onChange={(value) =>
            setFieldValue(ProductSubmissionEntity.Fields.image$, value, false)
          }
        />
        <FormEntitySelect3
          formEffect={{ form, field: ProductSubmissionEntity.Fields.category$ }}
          useQuery={useGetCategories}
          label={t.productsubmissions.category}
          hint={t.productsubmissions.categoryHint}
        />
        <FormEntitySelect3
          formEffect={{ form, field: ProductSubmissionEntity.Fields.tags$ }}
          useQuery={useGetTags}
          multiple
          label={t.productsubmissions.tags}
          hint={t.productsubmissions.tagsHint}
        />
        {values.productId ? (
          <ProductSubmissionFormSelected form={form} isEditing={isEditing} />
        ) : null}
      </div>
    </>
  );
};
