import { useDebouncedEffect } from "../../hooks/useDebouncedEffect";
import { useKeyPress } from "../../hooks/useKeyPress";
import { useLocale } from "../../hooks/useLocale";
import { useT } from "../../hooks/useT";
import { useRouter } from "../../hooks/useRouter";
import { useContext, useEffect, useRef, useState } from "react";
import { ReactiveSearchContext } from "./ReactiveSearchContext";
import { detectDeviceType } from "../../hooks/deviceInformation";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useWebSocketX } from "../../sdk/sdk/react/useWebSocketX";
import { ReactiveSearchAction } from "../../sdk/reactivesearch/ReactiveSearchAction";
import type { ReactiveSearchResultDtoType } from "../../sdk/reactivesearch/ReactiveSearchResultDto";

export function ReactiveSearch() {
  const t = useT();
  const { withDebounce } = useDebouncedEffect();
  const { setResult, setPhrase, phrase, result, reset } = useContext(
    ReactiveSearchContext
  );
  const { options } = useContext(RemoteQueryContext);
  // Holds whatever query params the next connection should carry - read
  // synchronously by the factory below at restart() time, so a phrase typed
  // right before calling restart() is never one render behind (a plain
  // useState here would be, since restart() runs in the same tick as the
  // setState call that would update it).
  const qsRef = useRef<URLSearchParams>();

  // Browsers can't set an Authorization header on a WebSocket handshake, so -
  // same as the hand-rolled hook this replaces (useReactiveSearchOld, now
  // deleted) - the token/workspace/locale ride along as query params
  // instead. lazy: true keeps the pre-existing behavior of never opening a
  // connection until the user actually types something (this component
  // lives in the navbar, rendered on effectively every page).
  const { messages, restart } = useWebSocketX(
    () => {
      const wsRemote = (options.prefix || "")
        .replace("https", "wss")
        .replace("http", "ws");
      // NewUrl (unlike Create) correctly folds qs into the path, so build the
      // full url here and hand it to Create as overrideUrl rather than
      // letting Create build it - Create silently ignores qs whenever
      // overrideUrl is also given.
      const path = ReactiveSearchAction.NewUrl(qsRef.current);
      return ReactiveSearchAction.Create(`${wsRemote}${path}`);
    },
    { lazy: true }
  );

  const router = useRouter();
  const input = useRef<HTMLInputElement | null>();
  const [value, setValue] = useState("");
  const { locale } = useLocale();

  // Clear the search box, from somewhere else in the scope
  useEffect(() => {
    if (!phrase) {
      setValue("");
    }
  }, [phrase]);

  // Set the results into the reactive search context
  useEffect(() => {
    const parsed = messages
      .map((raw) => {
        try {
          return JSON.parse(raw as unknown as string) as ReactiveSearchResultDtoType;
        } catch {
          return null;
        }
      })
      .filter((item): item is ReactiveSearchResultDtoType => item !== null);
    setResult(parsed);
  }, [messages]);

  const oninput = (value: string) => {
    withDebounce(() => {
      setPhrase(value);
      qsRef.current = new URLSearchParams({
        searchPhrase: value,
        acceptLanguage: (options.headers as any)?.["accept-language"] ?? "",
        token: options.headers?.authorization ?? "",
        workspaceId: (options.headers as any)?.["workspace-id"] ?? "",
      });
      restart();
    }, 500);
  };

  useKeyPress("s", () => {
    input.current?.focus();
  });

  const { isMobileView } = detectDeviceType();

  if (isMobileView) {
    return null;
  }

  return (
    <form
      className="navbar-search-box"
      onSubmit={(e) => {
        e.preventDefault();
        if (result.length > 0) {
          if (result[0].actionFn === "navigate" && result[0].uiLocation) {
            router.push(`/${locale}${result[0].uiLocation}`);
            reset();
          }
        }
      }}
    >
      <input
        ref={(ref) => {
          input.current = ref;
        }}
        value={value}
        placeholder={t.reactiveSearch.placeholder}
        onInput={(e) => {
          setValue((e.target as any).value);
          oninput((e.target as any).value);
        }}
        className="form-control"
      />
    </form>
  );
}
