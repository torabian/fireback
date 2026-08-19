import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { useS } from "@fireback/ui-core/hooks/useS";
import { httpErrorHanlder } from "@fireback/ui-core/hooks/api";
import { Toast } from "@fireback/ui-core/hooks/toast";
import { strings as coreStrings } from "@fireback/ui-core/components/strings/translations";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import {
  MarkNotificationReadActionReq,
  useMarkNotificationReadAction,
} from "@fireback/selfservice/sdk/abac/MarkNotificationReadAction";
import {
  type MyNotificationsActionRes,
  useMyNotificationsActionQuery,
} from "@fireback/selfservice/sdk/abac/MyNotificationsAction";
import { strings } from "./strings/translations";

// Poll interval for the notification list - a plain GET+react-query refetch (like
// every other list in the app) rather than wiring up NotificationStream/a websocket,
// so a notification root just sent shows up here within a few seconds without a
// manual refresh.
const REFRESH_INTERVAL_MS = 10000;

export const NotificationsSection = () => {
  const s = useS(strings);
  const cs = useS(coreStrings);
  const { locale } = useLocale();

  const query = useMyNotificationsActionQuery({
    refetchInterval: REFRESH_INTERVAL_MS,
  });
  const items =
    ((query.data as any)?.data?.items as MyNotificationsActionRes[]) || [];
  const unreadCount = items.filter((item) => !item.isRead).length;

  const markReadMutation = useMarkNotificationReadAction({});

  const markAsRead = (uniqueId: string) => {
    markReadMutation
      .mutateAsync(new MarkNotificationReadActionReq({ uniqueId }))
      .then(() => {
        Toast(s.notifications.markedAsRead, { type: "success" });
        query.refetch();
      })
      .catch((err) => httpErrorHanlder(err, cs));
  };

  return (
    <PageSection
      title={
        unreadCount > 0
          ? `${s.notifications.title} (${unreadCount} ${s.notifications.unreadBadge})`
          : s.notifications.title
      }
    >
      {query.isError && (
        <div className="alert alert-danger">{s.notifications.loadError}</div>
      )}

      {!query.isError && items.length === 0 && (
        <p className="text-muted">{s.notifications.empty}</p>
      )}

      {items.map((item) => (
        <div
          key={item.uniqueId}
          className="row mt-2 pb-2"
          style={{ borderBottom: "1px solid var(--bs-border-color, #dee2e6)" }}
        >
          <div className="col">
            <div className="d-flex align-items-center gap-2">
              {!item.isRead && (
                <span
                  aria-hidden="true"
                  style={{
                    display: "inline-block",
                    width: 8,
                    height: 8,
                    borderRadius: "50%",
                    background: "var(--bs-primary, #0d6efd)",
                    flex: "none",
                  }}
                />
              )}
              <strong>{item.title}</strong>
            </div>
            <div>{item.body}</div>
            <small className="text-muted">
              {new Date(item.createdAt).toLocaleString(locale)}
            </small>
          </div>
          {!item.isRead && (
            <div className="col-auto align-self-center">
              <button
                className="btn btn-sm btn-outline-primary"
                onClick={() => markAsRead(item.uniqueId)}
              >
                {s.notifications.markAsRead}
              </button>
            </div>
          )}
        </div>
      ))}
    </PageSection>
  );
};
