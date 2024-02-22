import { FormText } from "@/components/forms/form-text/FormText";
import { RoleEntity } from "src/sdk/fireback";
import { RoleEntityFields } from "src/sdk/fireback/modules/workspaces/role-fields";
import { FormikProps } from "formik";
import { RolePermissionTree } from "./RolePermissionTree";
import { useT } from "@/hooks/useT";

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
}: {
  form: FormikProps<Partial<RoleEntity>>;
  isEditing?: boolean;
}) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.name}
        onChange={(value) => setFieldValue(RoleEntityFields.name, value, false)}
        errorMessage={errors.name}
        label={t.wokspaces.invite.role}
        autoFocus={!isEditing}
        hint={t.wokspaces.invite.roleHint}
      />

      <RolePermissionTree
        onChange={(value) =>
          setFieldValue(RoleEntityFields.capabilitiesListId, value, false)
        }
        value={normalize(values.capabilities, values.capabilitiesListId)}
      />
    </>
  );
};
