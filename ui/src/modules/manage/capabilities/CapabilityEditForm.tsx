import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { useS } from "../../fireback-ui/hooks/useS";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { CapabilityDto } from "../../sdk/abac/CapabilityDto";
import { strings } from "./strings/translations";

export const CapabilityForm = ({
  form,
  isEditing,
}: EntityFormProps<CapabilityDto>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
      <FormText
        value={values.uniqueId}
        onChange={(value) =>
          setFieldValue(CapabilityDto.Fields.uniqueId, value, false)
        }
        // The id is the capability's real, meaningful key (e.g. "students.query") - it
        // can only be chosen once, at creation. Once a role/workspace-type/etc. has
        // actually granted or referenced it, silently renaming it out from under them
        // would break that reference, so it's locked (not just hidden) once editing.
        disabled={isEditing}
        errorMessage={errors.uniqueId}
        label={s.capabilities.uniqueId}
        hint={isEditing ? s.capabilities.uniqueIdLockedHint : s.capabilities.uniqueIdHint}
      />
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
