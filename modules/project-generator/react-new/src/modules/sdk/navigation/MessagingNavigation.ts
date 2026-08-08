import { createEntityNavigation } from "./createEntityNavigation";

export const EmailProviderNavigation = createEntityNavigation(
  "email-provider",
  "email-providers"
);
export const EmailSenderNavigation = createEntityNavigation(
  "email-sender",
  "email-senders"
);
export const GsmProviderNavigation = createEntityNavigation(
  "gsm-provider",
  "gsm-providers"
);
export const WebPushConfigNavigation = createEntityNavigation(
  "web-push-config",
  "web-push-configs"
);
