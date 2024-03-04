import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";

import { PublicJoinKeyEntity } from "@/sdk/fireback/modules/workspaces/PublicJoinKeyEntity";
import { RoleEntity } from "@/sdk/fireback/modules/workspaces/RoleEntity";
import { useGetRoles } from "@/sdk/fireback/modules/workspaces/useGetRoles";
import { FormikProps } from "formik";
import { useContext } from "react";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";

export const PublicJoinKeyEditForm = ({
  form,
  isEditing,
}: {
  form: FormikProps<Partial<PublicJoinKeyEntity>>;
  isEditing?: boolean;
}) => {
  const { values, setValues, setFieldValue, errors } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      {/* <FormEntitySelect2
        label={t.wokspaces.joinKeyWorkspace}
        hint={t.wokspaces.joinKeyWorkspaceHint}
        fnLoadOptions={async (keyword) => {
          return (
            (
              await WorkspaceActions.fn(options)
                .query(`name %"${keyword}"%`)
                .getWorkspaces()
            ).data?.items || []
          );
        }}
        value={values.workspace}
        onChange={(entity) => {
          setValues({
            ...values,
            workspace: entity,
            workspaceId: entity.uniqueId,
          });
        }}
        labelFn={(t: WorkspaceEntity) => [t?.name].join(" ")}
      /> */}
      <FormEntitySelect3
        formEffect={{ field: "role", form }}
        useQuery={useGetRoles}
        label="Role"
        hint="Select the role with public join key"
      />
      {/* <FormText
        value={values.uniqueId}
        onChange={(value) => setFieldValue("uniqueId", value, false)}
        errorMessage={errors.uniqueId}
        label="Unique Id"
        autoFocus={!isEditing}
        hint="Unique id the public join key"
      /> */}
      {/*<FormText
        value={values.senderName}
        onChange={(value) =>
          setFieldValue(EmailProviderEntity.Fields.senderName, value, false)
        }
        errorMessage={errors.senderName}
        label="MailProvider.senderName"
        hint="Mail provider.senderName"
      /> */}
      {/* <FormText
        value={values.apiKey}
        onChange={(value) =>
          setFieldValue(EmailProviderEntity.Fields.apiKey, value, false)
        }
        errorMessage={errors.apiKey}
        label="API Key"
        hint="The API key related to the mail service provider, if applicable"
      /> */}
    </>
  );
};
