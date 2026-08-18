import { useEffect, useRef, type RefObject } from "react";

// Walks up from `from` looking for the nearest ancestor that's actually a
// scroll container (overflow-y auto/scroll), for use as IntersectionObserver's
// `root`. Needed because "the page" isn't necessarily `<body>`/the viewport -
// in this app's own admin shell (enterprise-shell/ApplicationPanels), the real
// scrollable region is an inner ".content-container" div with a fixed height
// and overflow:auto; <body>/<html> never scroll at all. Verified empirically:
// without this, a DataGridList given no containerRef (the "just put it on the
// page" case) never received a single IntersectionObserver callback, because
// `root: null` observes the *viewport*, and the viewport itself never scrolls
// in this shell - only content-container does. Falls back to null (the
// viewport) when no scrollable ancestor is found, which is correct for a
// genuinely plain page.
function findScrollableAncestor(from: HTMLElement | null): HTMLElement | null {
  let el = from?.parentElement ?? null;
  while (el && el !== document.body) {
    const style = getComputedStyle(el);
    if (style.overflowY === "auto" || style.overflowY === "scroll") {
      return el;
    }
    el = el.parentElement;
  }
  return null;
}

// Triggers onLoadMore when a sentinel element near the end of the list scrolls
// into view - via IntersectionObserver, not scrollTop/scrollHeight math, so the
// exact same code path works whether DataGridList sits inside a caller-given
// scroll container (containerRef) or is just placed in the page - in the
// latter case the *real* scrolling ancestor is auto-detected (see
// findScrollableAncestor above) rather than assumed to be the browser
// viewport, which a real app shell usually isn't.
export function useInfiniteScroll({
  containerRef,
  onLoadMore,
  enabled,
  rootMargin = "300px",
}: {
  /** Scroll container to observe within. Auto-detected when omitted. */
  containerRef?: RefObject<HTMLElement | null>;
  onLoadMore: () => void;
  /** false while there's nothing more to load, or a fetch is already in flight. */
  enabled: boolean;
  /** How far before the sentinel actually enters view to fire early, like a prefetch margin. */
  rootMargin?: string;
}) {
  const sentinelRef = useRef<HTMLDivElement | null>(null);
  const onLoadMoreRef = useRef(onLoadMore);
  onLoadMoreRef.current = onLoadMore;

  useEffect(() => {
    const sentinel = sentinelRef.current;
    if (!sentinel || !enabled) {
      return;
    }

    const root = containerRef?.current ?? findScrollableAncestor(sentinel);

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries.some((entry) => entry.isIntersecting)) {
          onLoadMoreRef.current();
        }
      },
      { root, rootMargin },
    );

    observer.observe(sentinel);
    return () => observer.disconnect();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [enabled, containerRef?.current, rootMargin]);

  return { sentinelRef };
}
