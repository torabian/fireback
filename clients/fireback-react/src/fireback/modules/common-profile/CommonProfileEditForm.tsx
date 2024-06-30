import { FormText } from "@/fireback/components/forms/form-text/FormText";
import { CommonProfileEntity } from "@/sdk/fireback/modules/commonprofile/CommonProfileEntity";

import { FormikProps } from "formik";

export const CommonProfileEditForm = ({
  form,
}: {
  form: FormikProps<Partial<CommonProfileEntity>>;
}) => {
  const { values, setFieldValue, errors } = form;

  return (
    <>
      <FormText
        value={values.firstName}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.firstName, value, false)
        }
        errorMessage={errors.firstName}
        label="First name"
        hint="Public name of yours firstname"
      />

      <FormText
        value={values.lastName}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.lastName, value, false)
        }
        errorMessage={errors.lastName}
        label="Last name"
        hint="Public last name"
      />
      <FormText
        value={values.company}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.company, value, false)
        }
        errorMessage={errors.company}
        label="Company"
        hint="Company name"
      />
      <FormText
        value={values.phoneNumber}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.phoneNumber, value, false)
        }
        errorMessage={errors.phoneNumber}
        label="Phone number"
        hint="Enter your phone number"
      />
      <FormText
        value={values.email}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.email, value, false)
        }
        errorMessage={errors.email}
        label="Your public email address"
        hint="Public email address"
      />
      <FormText
        value={values.street}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.street, value, false)
        }
        errorMessage={errors.street}
        label="Your street address"
        hint="Street address"
      />
      <FormText
        value={values.houseNumber}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.houseNumber, value, false)
        }
        errorMessage={errors.houseNumber}
        label="House number"
        hint="House number"
      />
      <FormText
        value={values.zipCode}
        onChange={(value) =>
          setFieldValue(CommonProfileEntity.Fields.zipCode, value, false)
        }
        errorMessage={errors.zipCode}
        label="Zip code"
        hint="Zip code"
      />
    </>
  );
};
