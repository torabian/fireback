import { FormEntitySelect2 } from "@/components/forms/form-select/FormEntitySelect2";
import { useT } from "@/hooks/useT";
import {
  PublicJoinKeyEntity,
  RoleEntity,
  WorkspaceEntity,
} from "src/sdk/fireback";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { RoleActions } from "src/sdk/fireback/modules/workspaces/role-actions";
import { WorkspaceActions } from "src/sdk/fireback/modules/workspaces/workspace-actions";
import { FormikProps } from "formik";
import { useContext } from "react";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";

export const PublicJoinKeyEditForm = ({
  form,
}: {
  form: FormikProps<Partial<PublicJoinKeyEntity>>;
}) => {
  const { values, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <FormEntitySelect2
        label={t.wokspaces.invite.role}
        hint={t.wokspaces.invite.roleHint}
        fnLoadOptions={async (keyword) => {
          return (
            (
              await RoleActions.fn(options)
                .query(`name %"${keyword}"%`)
                .getRoles()
            ).data?.items || []
          );
        }}
        value={values.role}
        onChange={(entity) => {
          setValues({
            ...values,
            role: entity,
            roleId: entity.uniqueId,
          });
        }}
        labelFn={(t: RoleEntity) => [t?.name].join(" ")}
      />

      <FormEntitySelect2
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
      />
      {/* <FormSelect
        value={values.type}
        onChange={(value) =>
          setFieldValue(EmailProviderEntityFields.type, value, false)
        }
        options={[{ label: "Sendgrid", value: "sendgrid" }]}
        errorMessage={errors.type}
        label="Provider type"
        hint="Select the mail provider from list. Under the list you can find all providers we support."
      /> */}
      {/* <FormText
        value={values.senderAddress}
        onChange={(value) =>
          setFieldValue(EmailProviderEntityFields.senderAddress, value, false)
        }
        errorMessage={errors.senderAddress}
        label="MailProvider.senderAddress"
        hint="Mail provider.senderAddress"
      />
      <FormText
        value={values.senderName}
        onChange={(value) =>
          setFieldValue(EmailProviderEntityFields.senderName, value, false)
        }
        errorMessage={errors.senderName}
        label="MailProvider.senderName"
        hint="Mail provider.senderName"
      /> */}
      {/* <FormText
        value={values.apiKey}
        onChange={(value) =>
          setFieldValue(EmailProviderEntityFields.apiKey, value, false)
        }
        errorMessage={errors.apiKey}
        label="API Key"
        hint="The API key related to the mail service provider, if applicable"
      /> */}
    </>
  );
};
