import { FormText } from "../../../components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../definitions/definitions";
import { useS } from "../../../hooks/useS";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { PassportMethodDto } from "../../../sdk/abac/PassportMethodDto";
import { useContext } from "react";
import { strings } from "./strings/translations";
import { createQuerySource } from "../../../hooks/useAsQuery";
import { FormSelect } from "../../../components/forms/form-select/FormSelect";

export const PassportMethodForm = ({
  form,
  isEditing,
}: EntityFormProps<PassportMethodDto>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);

  const source = createQuerySource([
    { name: s.passportMethods.methodGoogle, uniqueId: "google" },
    { name: s.passportMethods.methodFacebook, uniqueId: "facebook" },
    { name: s.passportMethods.methodEmail, uniqueId: "email" },
    { name: s.passportMethods.methodPhone, uniqueId: "phone" },
  ]);

  return (
    <>
      <FormSelect
        querySource={source}
        formEffect={{
          form,
          field: PassportMethodDto.Fields.type,
          beforeSet(item) {
            return item.uniqueId;
          },
        }}
        keyExtractor={(v) => v.uniqueId}
        fnLabelFormat={(v) => v.name}
        errorMessage={errors.type}
        label={s.passportMethods.type}
        hint={s.passportMethods.typeHint}
      />

      <FormText
        value={values.region}
        onChange={(value) =>
          setFieldValue(PassportMethodDto.Fields.region, value, false)
        }
        errorMessage={errors.region}
        label={s.passportMethods.region}
        hint={s.passportMethods.regionHint}
      />
      {values.type === "google" || (values.type as any) === "facebook" ? (
        <FormText
          value={values.clientKey}
          onChange={(value) =>
            setFieldValue(PassportMethodDto.Fields.clientKey, value, false)
          }
          errorMessage={errors.clientKey}
          label={s.passportMethods.clientKey}
          hint={s.passportMethods.clientKeyHint}
        />
      ) : null}
    </>
  );
};
