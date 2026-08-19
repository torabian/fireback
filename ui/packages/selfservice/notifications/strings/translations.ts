/**
 * Hand-written (not run through the fireback language & translation manager yet -
 * see ui/scripts/translation-manager) - useS falls back to `en` for any locale
 * without its own "$<locale>" key, so this is a valid, if English-only,
 * translations.ts. Run `npm run t` from ui/ once fa/pl copy exists to regenerate
 * this with $fa/$pl included.
 */
export const en = {
  notifications: {
    title: "Notifications",
    empty: "You have no notifications yet.",
    markAsRead: "Mark as read",
    unreadBadge: "unread",
    markedAsRead: "Notification marked as read.",
    loadError: "Could not load your notifications.",
  },
};
export const strings = { ...en };
