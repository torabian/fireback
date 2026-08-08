import { useContext } from "react";
import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { useS } from "../../fireback-ui/hooks/useS";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { CapabilityDto } from "../../sdk/abac/CapabilityDto";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
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
