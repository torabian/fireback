import "./LoadingProgress.css";
import { strings as coreStrings } from "../strings/translations";
import { useS } from "../../hooks/useS";

export interface LoadingProgressProps {
  /** Primary status line, e.g. "Loading…" or a column filter's own busy message. */
  message: string;
  /** Rows fetched so far, if known - renders as a live counter under the message. */
  loadedCount?: number;
  /** Total row count, if the server already reported one (e.g. from totalItems). */
  totalCount?: number;
  /** Renders a Cancel button (e.g. to abort the in-flight request) when set. */
  onCancel?: () => void;
  /**
   * Full-viewport overlay (SessionGate's own layout) vs a plain block that fills
   * whatever container it's placed in - DataGridList uses "inline" both for a
   * lone full-page grid and for one boxed inside a fixed-height div, since in
   * both cases the grid's own wrapper (not the viewport) is what should be
   * covered.
   */
  variant?: "inline" | "overlay";
}

// DataGridList's loading/first-fetch indicator - visually matching
// session-gate/SessionGate.tsx's spinner (same proportions, same dark-mode
// handling) without importing it: SessionGate is deliberately self-contained and
// deletable in one piece (own CSS, own translations, "nothing else in the
// codebase references it" - see its own doc comment), so a real ui-core-wide
// component reuses its *look*, not its file, and lives in the normal shared
// components/strings system like everything else in ui-core.
export function LoadingProgress({
  message,
  loadedCount,
  totalCount,
  onCancel,
  variant = "inline",
}: LoadingProgressProps) {
  const cs = useS(coreStrings);

  const hasCount = typeof loadedCount === "number";

  return (
    <div className={`loading-progress loading-progress--${variant}`}>
      <div className="loading-progress__center">
        <div
          className="loading-progress__spinner"
          role="status"
          aria-live="polite"
          aria-label={message}
        />
        <div className="loading-progress__message">{message}</div>
        {hasCount && (
          <div className="loading-progress__count">
            {typeof totalCount === "number"
              ? `${loadedCount} / ${totalCount}`
              : loadedCount}
          </div>
        )}
        {onCancel && (
          <button className="loading-progress__cancel" onClick={onCancel}>
            {cs.common.cancel}
          </button>
        )}
      </div>
    </div>
  );
}

export default LoadingProgress;
