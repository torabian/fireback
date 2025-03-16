import { FormText } from "../../../components/forms/form-text/FormText";
import { EntityFormProps } from "../../../definitions/definitions";
import { useT } from "../../../hooks/useT";
import { WorkspaceTypeEntity } from "../../../sdk/modules/workspaces/WorkspaceTypeEntity";

import { useContext } from "react";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { useGetRoles } from "../../../sdk/modules/workspaces/useGetRoles";

export const WorkspaceTypeEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<WorkspaceTypeEntity>>) => {
  const { values, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const t = useT();

  return (
    <>
      <FormText
        value={values.title}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeEntity.Fields.title, value, false)
        }
        errorMessage={form.errors.title}
        label={t.wokspaces.workspaceTypeTitle}
        autoFocus={!isEditing}
        hint={t.wokspaces.workspaceTypeTitleHint}
      />
      <FormText
        value={values.slug}
        onChange={(value) =>
          form.setFieldValue(WorkspaceTypeEntity.Fields.slug, value, false)
        }
        errorMessage={form.errors.slug}
        label={t.wokspaces.workspaceTypeSlug}
        hint={t.wokspaces.workspaceTypeSlugHint}
      />
      <FormSelect
        label={t.wokspaces.invite.role}
        hint={t.wokspaces.invite.roleHint}
        fnLabelFormat={(role) => role.name}
        querySource={useGetRoles}
        formEffect={{ form, field: WorkspaceTypeEntity.Fields.role$ }}
        errorMessage={form.errors.roleId}
      />

      {/* <FormSelect
        value={values.type}
        onChange={(value) =>
          setFieldValue(EmailProviderEntity.Fields.type, value, false)
        }
        options={[{ label: "Sendgrid", value: "sendgrid" }]}
        errorMessage={errors.type}
        label="Provider type"
        hint="Select the mail provider from list. Under the list you can find all providers we support."
      /> */}
      {/* <FormText
        value={values.senderAddress}
        onChange={(value) =>
          setFieldValue(EmailProviderEntity.Fields.senderAddress, value, false)
        }
        errorMessage={errors.senderAddress}
        label="MailProvider.senderAddress"
        hint="Mail provider.senderAddress"
      />
      <FormText
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
