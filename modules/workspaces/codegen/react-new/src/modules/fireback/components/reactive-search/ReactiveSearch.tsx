import { useDebouncedEffect } from "@/modules/fireback/hooks/useDebouncedEffect";
import { useKeyPress } from "@/modules/fireback/hooks/useKeyPress";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useT } from "@/modules/fireback/hooks/useT";
import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { useReactiveReactiveSearch } from "../../sdk/modules/workspaces/useReactiveReactiveSearch";
import { useContext, useEffect, useRef, useState } from "react";
import { ReactiveSearchContext } from "./ReactiveSearchContext";

export function ReactiveSearch() {
  const t = useT();
  const { withDebounce } = useDebouncedEffect();
  const { setResult, setPhrase, phrase, result, reset } = useContext(
    ReactiveSearchContext
  );
  const { operate, data } = useReactiveReactiveSearch({});
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
    setResult(data);
  }, [data]);

  const oninput = (value: string) => {
    withDebounce(() => {
      setPhrase(value);
      operate({ searchPhrase: encodeURIComponent(value) } as any);
    }, 500);
  };

  useKeyPress("s", () => {
    input.current?.focus();
  });

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
