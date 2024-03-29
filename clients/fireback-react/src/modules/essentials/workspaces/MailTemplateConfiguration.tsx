import { ModalContext } from "@/components/modal/Modal";
import { useT } from "@/hooks/useT";
import { FormikProps } from "formik";
import { useContext, useRef } from "react";
import {
  MailTemplateEditor,
  MailTemplateEntityManager,
} from "./MailTemplateEntityManager";
import { FormEmailSenderPicker } from "./FormEmailSenderPicker";
import { enTranslations } from "@/translations/en";
import { NotificationConfigEntity } from "@/sdk/fireback/modules/workspaces/NotificationConfigEntity";
import { EmailSenderEntity } from "@/sdk/fireback/modules/workspaces/EmailSenderEntity";

type Templates =
  | "inviteToWorkspaceSender"
  | "confirmEmailSender"
  | "forgetPasswordSender";

interface WorkspaceEmailConfigItem {
  title: string;
  key: string;
}

const mailTypes = (t: typeof enTranslations): WorkspaceEmailConfigItem[] => [
  {
    title: t.wokspaces.inviteToWorkspace,
    key: "inviteToWorkspaceSender",
  },
  {
    title: t.wokspaces.confirmEmailSender,
    key: "confirmEmailSender",
  },
  {
    key: "forgetPasswordSender",
    title: t.wokspaces.forgetPasswordSender,
  },
];

export function MailTemplateConfiguration({
  form,
}: {
  form: FormikProps<Partial<NotificationConfigEntity>>;
}) {
  const t = useT();
  const useModal = useContext(ModalContext);
  const templateFormik = useRef<FormikProps<
    Partial<MailTemplateEditor>
  > | null>();

  const customizeTemplate = (
    item: Partial<NotificationConfigEntity>,
    key: Templates
  ) => {
    let body = "",
      title = "",
      defaultBody = "",
      defaultTitle = "";

    if (key === "confirmEmailSender") {
      body = item.confirmEmailContent || "";
      title = item.confirmEmailTitle || "";
      defaultBody = item.confirmEmailContentDefault || "";
      defaultTitle = item.confirmEmailTitleDefault || "";
    }

    if (key === "forgetPasswordSender") {
      body = item.forgetPasswordContent || "";
      title = item.forgetPasswordTitle || "";
      defaultBody = item.forgetPasswordContentDefault || "";
      defaultTitle = item.forgetPasswordTitleDefault || "";
    }

    if (key === "inviteToWorkspaceSender") {
      body = item.inviteToWorkspaceContent || "";
      title = item.inviteToWorkspaceTitle || "";
      defaultBody = item.inviteToWorkspaceContentDefault || "";
      defaultTitle = item.inviteToWorkspaceTitleDefault || "";
    }

    useModal.openModal({
      title: t.wokspaces.notification.dialogTitle,
      component: () => (
        <MailTemplateEntityManager
          setInnerRef={(r) => (templateFormik.current = r)}
          data={{ body, title, defaultBody, defaultTitle }}
        />
        // <UnitEntityManager
        //   enabledFields={{ title: true, content: true }}
        //   data={{ courseId: d?.uniqueId }}
        //   showSubmit={false}
        //   setInnerRef={(r) => (formik.current = r)}
        //   onSuccess={() => {
        //     queryUnits.refetch();
        //   }}
        // />
      ),
      onSubmit: async () => {
        const body = templateFormik.current?.values.body;
        const title = templateFormik.current?.values.title;

        if (key === "confirmEmailSender") {
          form.setFieldValue(
            NotificationConfigEntity.Fields.confirmEmailContent,
            body
          );
          form.setFieldValue(
            NotificationConfigEntity.Fields.confirmEmailTitle,
            title
          );
        }

        if (key === "forgetPasswordSender") {
          form.setFieldValue(
            NotificationConfigEntity.Fields.forgetPasswordContent,
            body
          );
          form.setFieldValue(
            NotificationConfigEntity.Fields.forgetPasswordTitle,
            title
          );
        }

        if (key === "inviteToWorkspaceSender") {
          form.setFieldValue(
            NotificationConfigEntity.Fields.inviteToWorkspaceContent,
            body
          );
          form.setFieldValue(
            NotificationConfigEntity.Fields.inviteToWorkspaceTitle,
            title
          );
        }

        return true;
      },
    });
  };

  return (
    <table className="table">
      <thead>
        <tr>
          <th>{t.wokspaces.type}</th>
          <th>{t.wokspaces.sender}</th>
          <th>{t.wokspaces.customizedTemplate}</th>
        </tr>
      </thead>
      <tbody>
        {mailTypes(t).map((item) => (
          <tr key={item.key}>
            <td>{item.title}</td>
            <td width={400}>
              <FormEmailSenderPicker
                value={(form.values as any)[item.key]}
                onChange={(entity: EmailSenderEntity) => {
                  form.setValues({
                    ...form.values,
                    [item.key]: entity,
                    [`${item.key}Id`]: entity.uniqueId,
                  });
                }}
              />
            </td>
            <td>
              <button
                onClick={() =>
                  customizeTemplate(form.values, item.key as Templates)
                }
                className="btn btn-secondary"
              >
                {t.wokspaces.customizedTemplate}
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
