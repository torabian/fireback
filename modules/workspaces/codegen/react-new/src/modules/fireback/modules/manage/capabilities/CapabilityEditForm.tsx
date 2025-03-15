import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { useContext } from "react";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/workspaces/CapabilityEntity";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const CapabilityForm = ({
  form,
  isEditing,
}: EntityFormProps<CapabilityEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
        <FormText
          value={values.name }
          onChange={(value) => setFieldValue(CapabilityEntity.Fields.name, value, false)}
          errorMessage={errors.name }
          label={s.capabilities.name }
          hint={s.capabilities.nameHint}
        />
        <FormText
          value={values.description }
          onChange={(value) => setFieldValue(CapabilityEntity.Fields.description, value, false)}
          errorMessage={errors.description }
          label={s.capabilities.description }
          hint={s.capabilities.descriptionHint}
        />
    </>
  );
};