import { useWebPushConfigCreateAction } from "../../sdk/messaging/WebPushConfigCreateAction";
import { useEffect, useState } from "react";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

export function usePushSubscription() {
  const { mutate: submit } = useWebPushConfigCreateAction();
  const s = useS(strings);

  useEffect(() => {
    if (navigator.serviceWorker) {
      navigator.serviceWorker.addEventListener("message", (event) => {
        if (event.data?.type === "PUSH_RECEIVED") {
          console.log("Push message in UI:", event.data.payload);
          // Update state, show toast, etc.
        }
      });
    }
  }, []);

  const [isSubscribing, setIsSubscribing] = useState(false);
  const [isSubscribed, setIsSubscribed] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Detect subscription on mount
  useEffect(() => {
    async function checkSubscription() {
      try {
        const reg = await navigator.serviceWorker.ready;
        const sub = await reg.pushManager.getSubscription();
        setIsSubscribed(!!sub);
      } catch (err) {
        console.error("Failed to check subscription", err);
      }
    }

    checkSubscription();
  }, []);

  const subscribe = async () => {
    setIsSubscribing(true);
    setError(null);
    try {
      const reg = await navigator.serviceWorker.ready;
      const sub = await reg.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey:
          "BAw6oGpr6FoFDj49xOhFbTSOY07zvcqYWyyXeQXUJIFubi5iLQNV0vYsXKLz7J8520o4IjCq8u9tLPBx2NSuu04",
      });

      console.log(25, sub);
      submit({ subscription: sub });
      console.log("Subscribed:", JSON.stringify(sub));
      setIsSubscribed(true);
    } catch (err) {
      setError(s.notifications.failedToSubscribe);
      console.error("Subscription failed:", err);
    } finally {
      setIsSubscribing(false);
    }
  };

  const unsubscribe = async () => {
    setIsSubscribing(true);
    setError(null);
    try {
      const reg = await navigator.serviceWorker.ready;
      const sub = await reg.pushManager.getSubscription();
      if (sub) {
        await sub.unsubscribe();
        setIsSubscribed(false);
      } else {
        setError(s.notifications.noSubscriptionFound);
      }
    } catch (err) {
      setError(s.notifications.failedToUnsubscribe);
      console.error("Unsubscription failed:", err);
    } finally {
      setIsSubscribing(false);
    }
  };

  return {
    isSubscribing,
    isSubscribed,
    error,
    subscribe,
    unsubscribe,
  };
}
