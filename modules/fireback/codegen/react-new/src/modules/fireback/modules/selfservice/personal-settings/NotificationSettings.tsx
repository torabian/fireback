import { ErrorsView } from "@/modules/fireback/components/error-view/ErrorView";
import { PageSection } from "../../../components/page-section/PageSection";
import { usePushSubscription } from "./usePushSubscription";

export function NotificationSettings({}: {}) {
  const { error, isSubscribed, isSubscribing, subscribe, unsubscribe } =
    usePushSubscription();

  return (
    <PageSection title={"Notification settings"}>
      <p>Here you can manage your notifications</p>
      <ErrorsView error={error} />
      <button
        className="btn"
        disabled={isSubscribing || isSubscribed}
        onClick={() => subscribe()}
      >
        Subscribe
      </button>
      <button
        disabled={!isSubscribed}
        className="btn"
        onClick={() => unsubscribe()}
      >
        Unsubscribe
      </button>
    </PageSection>
  );
}
