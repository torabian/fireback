import { useState } from "react";
import {
  PageTitle,
  usePageTitle,
} from "@fireback/ui-core/components/page-title/PageTitle";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { useS } from "@fireback/ui-core/hooks/useS";
import { httpErrorHanlder } from "@fireback/ui-core/hooks/api";
import { Toast } from "@fireback/ui-core/hooks/toast";
import { strings as coreStrings } from "@fireback/ui-core/components/strings/translations";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { FormSelectMultiple } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { useUsersQuerySource } from "@fireback/ui-core/hooks/useUsersQuerySource";
import { type UserDto } from "@fireback/manage/sdk/abac/UserDto";
import {
  SendNotificationActionReq,
  useSendNotificationAction,
} from "@fireback/manage/sdk/abac/SendNotificationAction";
import { strings } from "./strings/translations";

function userLabel(user: UserDto): string {
  const name = [user.firstName, user.lastName].filter(Boolean).join(" ");
  return name ? `${name} (${user.uniqueId})` : user.uniqueId || "";
}

// Root-only admin screen for SendNotificationAction (see modules/abac/
// NotificationActions.go) - the backend + generated SDK for this action have
// existed since the notification module was built, but no screen ever called it:
// the only notification UI that got built was self-service's own "my
// notifications" list (NotificationsSection.tsx). This is the missing piece -
// picks one or more recipients (FormSelectMultiple over useUsersQuerySource,
// same async-search pattern WorkspaceConfigEditForm's dropdowns use) plus a
// title/body, and fires them all in one SendNotificationAction call.
export const SendNotification = () => {
  const s = useS(strings);
  const cs = useS(coreStrings);
  usePageTitle(s.notifications.title);

  const [recipients, setRecipients] = useState<UserDto[]>([]);
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");

  const sendMutation = useSendNotificationAction({});

  const canSubmit =
    recipients.length > 0 && title.trim().length > 0 && body.trim().length > 0;

  const onSubmit = () => {
    if (!canSubmit) {
      return;
    }
    sendMutation
      .mutateAsync(
        new SendNotificationActionReq({
          userIds: recipients.map((u) => u.uniqueId!),
          title: title.trim(),
          body: body.trim(),
        }),
      )
      .then((res: any) => {
        const sentCount = res?.data?.item?.sentCount ?? recipients.length;
        const skipped = res?.data?.item?.skippedUserIds ?? [];
        Toast(s.notifications.sent.replace("{count}", String(sentCount)), {
          type: "success",
        });
        if (skipped.length > 0) {
          Toast(
            s.notifications.someSkipped.replace(
              "{count}",
              String(skipped.length),
            ),
            { type: "warning" },
          );
        }
        setRecipients([]);
        setTitle("");
        setBody("");
      })
      .catch((err) => httpErrorHanlder(err, cs));
  };

  return (
    <div>
      <PageTitle description={s.notifications.description} />

      <PageSection title={s.notifications.composeSection}>
        <FormSelectMultiple<UserDto, string>
          value={recipients}
          onChange={setRecipients}
          querySource={useUsersQuerySource}
          keyExtractor={(u) => u.uniqueId!}
          fnLabelFormat={userLabel}
          label={s.notifications.recipients}
          hint={s.notifications.recipientsHint}
        />

        <FormText
          value={title}
          onChange={setTitle}
          label={s.notifications.notificationTitle}
        />

        <div className="mb-3">
          <label className="form-label">{s.notifications.notificationBody}</label>
          <textarea
            className="form-control"
            rows={5}
            value={body}
            onChange={(e) => setBody(e.target.value)}
          />
        </div>

        <button
          className="btn btn-primary"
          disabled={!canSubmit || sendMutation.isPending}
          onClick={onSubmit}
        >
          {sendMutation.isPending
            ? s.notifications.sending
            : s.notifications.send}
        </button>
      </PageSection>
    </div>
  );
};
