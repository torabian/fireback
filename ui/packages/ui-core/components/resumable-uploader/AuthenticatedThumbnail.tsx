import { useEffect, useState } from "react";
import type { CSSProperties } from "react";
import type { HeaderProvider } from "./types";

async function resolveHeaders(
  headers?: HeaderProvider,
): Promise<Record<string, string>> {
  if (!headers) return {};
  if (typeof headers === "function") {
    return (await headers()) || {};
  }
  return headers;
}

/**
 * An <img> whose source is fetched via `fetch()` with explicit headers,
 * instead of a plain `<img src>` (which the browser sends with no way to
 * attach a custom header at all). Needed for any thumbnail/download URL
 * that requires an Authorization header - e.g. an owned upload from the
 * storage module, which the README documents as its own caveat ("Caveat:
 * embedding an owned file in <img src=...>", storage/README.md §6):
 * `<img>` resolves anonymously, and an owned file 403s an anonymous
 * request. This is the fetch-then-createObjectURL workaround that README
 * suggests, packaged so every caller of getThumbnailUrl doesn't have to
 * reimplement it themselves.
 *
 * Deliberately takes `headers` as a plain prop rather than reading
 * `useUploaderConfig()` itself, so this stays usable outside a
 * <UploaderConfigProvider> too (e.g. a read-only detail view that only
 * ever needs the thumbnail, never the upload widget).
 */
export function AuthenticatedThumbnail({
  src,
  alt,
  className,
  style,
  headers,
  onError,
}: {
  src: string;
  alt?: string;
  className?: string;
  style?: CSSProperties;
  headers?: HeaderProvider;
  onError?: () => void;
}) {
  const [objectUrl, setObjectUrl] = useState<string | null>(null);

  useEffect(() => {
    let revokeUrl: string | null = null;
    let cancelled = false;

    setObjectUrl(null);

    (async () => {
      try {
        const resolvedHeaders = await resolveHeaders(headers);
        const resp = await fetch(src, { headers: resolvedHeaders });
        if (!resp.ok) {
          throw new Error(`HTTP ${resp.status}`);
        }
        const blob = await resp.blob();
        if (cancelled) return;
        revokeUrl = URL.createObjectURL(blob);
        setObjectUrl(revokeUrl);
      } catch {
        if (!cancelled) onError?.();
      }
    })();

    return () => {
      cancelled = true;
      if (revokeUrl) URL.revokeObjectURL(revokeUrl);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps -- headers is
    // read at fetch time from the render that triggered this effect; it
    // deliberately isn't a dep, or an inline `headers={() => ({...})}`
    // (a fresh function identity every render) would refetch on every
    // parent re-render instead of only when the image itself changes.
  }, [src]);

  if (!objectUrl) {
    return null;
  }

  return <img src={objectUrl} alt={alt} className={className} style={style} />;
}
