import { FormEntitySelect2 } from "@/components/forms/form-select/FormEntitySelect2";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";
import {
  WorkspaceTypeEntity,
  RoleEntity,
  WorkspaceEntity,
} from "src/sdk/fireback";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { RoleActions } from "src/sdk/fireback/modules/workspaces/role-actions";
import { WorkspaceActions } from "src/sdk/fireback/modules/workspaces/workspace-actions";
import { WorkspaceTypeEntityFields } from "src/sdk/fireback/modules/workspaces/workspace-type-fields";
import { FormikProps } from "formik";
import { useContext } from "react";

export const WorkspaceTypeEditForm = ({
  form,
}: {
  form: FormikProps<Partial<WorkspaceTypeEntity>>;
}) => {
  const { values, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <FormText
        value={values.title}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeEntityFields.title, value, false)
        }
        errorMessage={form.errors.title}
        label={t.wokspaces.workspaceTypeTitle}
        hint={t.wokspaces.workspaceTypeTitleHint}
      />
      <FormText
        value={values.slug}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeEntityFields.slug, value, false)
        }
        errorMessage={form.errors.slug}
        label={t.wokspaces.workspaceTypeSlug}
        hint={t.wokspaces.workspaceTypeSlugHint}
      />
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
        errorMessage={form.errors.roleId}
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
