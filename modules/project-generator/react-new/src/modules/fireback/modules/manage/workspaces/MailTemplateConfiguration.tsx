import { ModalContext } from "../../../../fireback-ui/components/modal/Modal";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { type FormikProps } from "formik";
import { useContext, useRef } from "react";
import {
  type MailTemplateEditor,
  MailTemplateEntityManager,
} from "./MailTemplateEntityManager";
import { FormEmailSenderPicker } from "./FormEmailSenderPicker";
import { strings } from "./strings/translations";
import { NotificationConfigDto } from "../../../sdk/abac/NotificationConfigDto";
import { EmailSenderDto } from "../../../sdk/messaging/EmailSenderDto";

type Templates =
  | "inviteToWorkspaceSender"
  | "confirmEmailSender"
  | "forgetPasswordSender";

interface WorkspaceEmailConfigItem {
  title: string;
  key: string;
}

const mailTypes = (s: typeof strings): WorkspaceEmailConfigItem[] => [
  {
    title: s.inviteToWorkspace,
    key: "inviteToWorkspaceSender",
  },
  {
    title: s.confirmEmailSender,
    key: "confirmEmailSender",
  },
  {
    key: "forgetPasswordSender",
    title: s.forgetPasswordSender,
  },
];

export function MailTemplateConfiguration({
  form,
}: {
  form: FormikProps<Partial<NotificationConfigDto>>;
}) {
  const s = useS(strings);
  const useModal = useContext(ModalContext);
  const templateFormik = useRef<FormikProps<
    Partial<MailTemplateEditor>
  > | null>();

  const customizeTemplate = (
    item: Partial<NotificationConfigDto>,
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
      title: s.notificationDialogTitle,
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
            NotificationConfigDto.Fields.confirmEmailContent,
            body
          );
          form.setFieldValue(
            NotificationConfigDto.Fields.confirmEmailTitle,
            title
          );
        }

        if (key === "forgetPasswordSender") {
          form.setFieldValue(
            NotificationConfigDto.Fields.forgetPasswordContent,
            body
          );
          form.setFieldValue(
            NotificationConfigDto.Fields.forgetPasswordTitle,
            title
          );
        }

        if (key === "inviteToWorkspaceSender") {
          form.setFieldValue(
            NotificationConfigDto.Fields.inviteToWorkspaceContent,
            body
          );
          form.setFieldValue(
            NotificationConfigDto.Fields.inviteToWorkspaceTitle,
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
          <th>{s.type}</th>
          <th>{s.sender}</th>
          <th>{s.customizedTemplate}</th>
        </tr>
      </thead>
      <tbody>
        {mailTypes(s).map((item) => (
          <tr key={item.key}>
            <td>{item.title}</td>
            <td width={400}>
              <FormEmailSenderPicker
                value={(form.values as any)[item.key]}
                onChange={(entity: EmailSenderDto) => {
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
                {s.customizedTemplate}
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
