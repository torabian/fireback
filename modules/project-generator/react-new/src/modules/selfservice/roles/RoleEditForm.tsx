import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { RoleDto } from "../../sdk/abac/RoleDto";
import { RolePermissionTree } from "./RolePermissionTree";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

/**
 * Server does not return capabilities list id, because it's used only on post/patch
 * this function casts it regardless to array<string> so form would work.
 */
const normalize = (caps: any, capList: any) => {
  if (caps?.length && !capList?.length) {
    return caps.map((t: any) => t.uniqueId);
  }

  return capList || [];
};

export const RoleEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<RoleDto>>) => {
  const { values, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) => setFieldValue(RoleDto.Fields.name, value, false)}
        errorMessage={errors.name}
        label={s.inviteRoleLabel}
        autoFocus={!isEditing}
        hint={s.inviteRoleHint}
      />

      <RolePermissionTree
        onChange={(value) =>
          setFieldValue(RoleDto.Fields.capabilitiesListId, value, false)
        }
        value={normalize(values.capabilities, values.capabilitiesListId)}
      />
    </>
  );
};
