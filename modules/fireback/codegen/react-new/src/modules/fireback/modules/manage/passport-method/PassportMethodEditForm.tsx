import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useS } from "@/modules/fireback/hooks/useS";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/abac/PassportMethodEntity";
import { useContext } from "react";
import { strings } from "./strings/translations";
import { createQuerySource } from "@/modules/fireback/hooks/useAsQuery";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";

export const PassportMethodForm = ({
  form,
  isEditing,
}: EntityFormProps<PassportMethodEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);

  const source = createQuerySource([
    { name: "Google", uniqueId: "google" },
    { name: "Facebook", uniqueId: "facebook" },
    { name: "Email", uniqueId: "email" },
    { name: "Phone", uniqueId: "phone" },
  ]);

  return (
    <>
      <FormSelect
        querySource={source}
        formEffect={{
          form,
          field: PassportMethodEntity.Fields.type,
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
          setFieldValue(PassportMethodEntity.Fields.region, value, false)
        }
        errorMessage={errors.region}
        label={s.passportMethods.region}
        hint={s.passportMethods.regionHint}
      />
      {values.type === "google" || (values.type as any) === "facebook" ? (
        <FormText
          value={values.clientKey}
          onChange={(value) =>
            setFieldValue(PassportMethodEntity.Fields.clientKey, value, false)
          }
          errorMessage={errors.clientKey}
          label={s.passportMethods.clientKey}
          hint={s.passportMethods.clientKeyHint}
        />
      ) : null}
    </>
  );
};
