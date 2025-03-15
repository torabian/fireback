import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { useContext } from "react";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/workspaces/PassportMethodEntity";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const PassportMethodForm = ({
  form,
  isEditing,
}: EntityFormProps<PassportMethodEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
        {/*
          <FormText
            type="?"
            value={values.type }
            onChange={(value) => setFieldValue(PassportMethodEntity.Fields.type, value, false)}
            errorMessage={errors.type }
            label={s.passportMethods.type }
            hint={s.passportMethods.typeHint}
          />
         */}
        {/*
          <FormText
            type="?"
            value={values.region }
            onChange={(value) => setFieldValue(PassportMethodEntity.Fields.region, value, false)}
            errorMessage={errors.region }
            label={s.passportMethods.region }
            hint={s.passportMethods.regionHint}
          />
         */}
    </>
  );
};