import { FormText } from "../../../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../definitions/definitions";
import { RoleDto } from "../../../sdk/abac/RoleDto";
import { RolePermissionTree } from "./RolePermissionTree";
import { useT } from "../../../../fireback-ui/hooks/useT";

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
  const t = useT();
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) =>
          setFieldValue(RoleDto.Fields.name, value, false)
        }
        errorMessage={errors.name}
        label={t.wokspaces.invite.role}
        autoFocus={!isEditing}
        hint={t.wokspaces.invite.roleHint}
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
