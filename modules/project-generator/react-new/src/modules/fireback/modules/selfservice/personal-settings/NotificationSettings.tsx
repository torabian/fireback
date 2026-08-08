import { ErrorsView } from "../../../components/error-view/ErrorView";
import { PageSection } from "../../../components/page-section/PageSection";
import { useS } from "../../../hooks/useS";
import { strings } from "./strings/translations";
import { usePushSubscription } from "./usePushSubscription";

export function NotificationSettings({}: {}) {
  const { error, isSubscribed, isSubscribing, subscribe, unsubscribe } =
    usePushSubscription();
  const s = useS(strings);

  return (
    <PageSection title={s.notifications.title}>
      <p>{s.notifications.description}</p>
      <ErrorsView error={error} />
      <button
        className="btn"
        disabled={isSubscribing || isSubscribed}
        onClick={() => subscribe()}
      >
        {s.notifications.subscribe}
      </button>
      <button
        disabled={!isSubscribed}
        className="btn"
        onClick={() => unsubscribe()}
      >
        {s.notifications.unsubscribe}
      </button>
    </PageSection>
  );
}
