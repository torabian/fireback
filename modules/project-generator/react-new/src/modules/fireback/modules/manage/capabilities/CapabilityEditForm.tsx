import { type EntityFormProps } from "../../../definitions/definitions";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { useContext } from "react";
import { CapabilityDto } from "../../../sdk/abac/CapabilityDto";
import { FormText } from "../../../components/forms/form-text/FormText";
import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { useS } from "../../../hooks/useS";
import { strings } from "./strings/translations";
export const CapabilityForm = ({
  form,
  isEditing,
}: EntityFormProps<CapabilityDto>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) =>
          setFieldValue(CapabilityDto.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={s.capabilities.name}
        hint={s.capabilities.nameHint}
      />
      <FormText
        value={values.description}
        onChange={(value) =>
          setFieldValue(CapabilityDto.Fields.description, value, false)
        }
        errorMessage={errors.description}
        label={s.capabilities.description}
        hint={s.capabilities.descriptionHint}
      />
    </>
  );
};
