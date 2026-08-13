import { useEffect, useRef, useState, type ReactNode } from "react";

import "./SessionGate.css";
import { SESSION_GATE_RTL_LOCALES, sessionGateStrings } from "./translations";
import { BUILD_VARIABLES } from "../../hooks/build-variables";

/**
 * Resolves a UI locale for this component only, without depending on
 * anything from the app's own routing/context (this gate is meant to sit
 * above all of that - it may render before a Router or any provider is
 * mounted). Looks at the URL first (path or hash segment), then the
 * browser's language, then falls back to English.
 */
function detectLocale(): string {
  const supported = Object.keys(sessionGateStrings);
  const fromUrl = `${window.location.pathname}${window.location.hash}`.match(
    /\/(en|pl|ru|es|fa)(?:\/|$)/,
  )?.[1];
  if (fromUrl) return fromUrl;

  const fromNavigator = (navigator.language || BUILD_VARIABLES.DEFAULT_LOCALE)
    .slice(0, 2)
    .toLowerCase();
  return supported.includes(fromNavigator)
    ? fromNavigator
    : BUILD_VARIABLES.DEFAULT_LOCALE;
}

export interface SessionGateProps {
  /**
   * Called once on mount, and again after every failed attempt, to check
   * whether the session is ready to use. Must be a *function that returns a
   * new promise each call* - a single already-settled Promise object can't
   * be retried, so passing one in directly isn't supported.
   *
   * Resolve it to let `children` render. Reject it (any reason) to keep the
   * gate up, bump the on-screen attempt counter, and try again after
   * `retryDelayMs`. Retries continue indefinitely - there's no attempt cap,
   * matching "keep checking and tell the user how many times it's tried".
   */
  checkSession: () => Promise<unknown>;
  children: ReactNode;
  /** Delay between a failed attempt and the next one, in ms. Default 1500. */
  retryDelayMs?: number;
}

/**
 * Gates rendering of `children` behind a session check. Meant to wrap an
 * entire app (see apps/*\/App.tsx) so nothing underneath it - no routes, no
 * providers that assume a session already exists - mounts before that first
 * check settles.
 *
 * Self-contained on purpose: own CSS, own translations (en/pl/ru/es/fa),
 * no shared state, no dependency on any other app provider. To remove this
 * mechanism entirely: delete this folder, then in App.tsx drop the
 * `<SessionGate checkSession={...}>` wrapper around the rest of the tree.
 * Nothing else in the codebase references it.
 */
export function SessionGate({
  checkSession,
  children,
  retryDelayMs = 1500,
}: SessionGateProps) {
  const [ready, setReady] = useState(false);
  const [attempt, setAttempt] = useState(0);

  useEffect(() => {
    let alive = true;
    let timer: ReturnType<typeof setTimeout>;

    const attemptCheck = (n: number) => {
      checkSession().then(
        () => {
          if (alive) setReady(true);
        },
        () => {
          if (!alive) return;
          setAttempt(n);
          timer = setTimeout(() => attemptCheck(n + 1), retryDelayMs);
        },
      );
    };

    attemptCheck(1);

    return () => {
      alive = false;
      clearTimeout(timer);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [retryDelayMs]);

  if (ready) {
    return <>{children}</>;
  }

  const locale = detectLocale();
  const s = sessionGateStrings[locale] ?? sessionGateStrings.en;
  const dir = SESSION_GATE_RTL_LOCALES.has(locale) ? "rtl" : "ltr";

  return (
    <div className="session-gate" dir={dir}>
      <div className="session-gate__center">
        <div
          className="session-gate__spinner"
          role="status"
          aria-live="polite"
          aria-label={s.checking}
        />
        <div className="session-gate__message">{s.checking}</div>
        {attempt > 0 && (
          <div className="session-gate__attempt">
            {s.attempt.replace("{n}", String(attempt))}
          </div>
        )}
      </div>
    </div>
  );
}

/**
 * Default `checkSession` - resolves immediately, so wiring `SessionGate` in
 * doesn't change app boot behavior until it's swapped for a real check.
 * Replace with a call to your actual session/whoami endpoint, e.g.:
 *
 *   () => fetchx.get("/session").then(res => { if (!res.ok) throw res; })
 */
export function noopCheckSession(): Promise<void> {
  return Promise.resolve();
}

export default SessionGate;
