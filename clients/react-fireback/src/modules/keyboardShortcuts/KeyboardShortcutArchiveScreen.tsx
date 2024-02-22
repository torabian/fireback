import { SmartHead } from "@/components/layouts/SmartHead";

import { getStaticPaths, makeStaticProps } from "@/hooks/getStatic";
import { useGetKeyboardShortcuts } from "src/sdk/fireback/modules/keyboardActions/useGetKeyboardShortcuts";
import { useQueryClient } from "react-query";
import { KeyBinding, Shortcut } from "./KeyboardShortcutDefinitions";
import { KeyBindings } from "./components/KeyBindings";
import { QueryErrorView } from "@/components/error-view/QueryError";
const getStaticProps = makeStaticProps(["common", "footer"]);
export { getStaticPaths, getStaticProps };

export function KeyboardShortcutArchiveScreen() {
  const queryClient = useQueryClient();

  const { query: queryKeyboardShortcuts } = useGetKeyboardShortcuts({
    queryClient,
    query: {
      itemsPerPage: 9999,
      startIndex: 0,
    },
  });

  const shortcuts: KeyBinding[] =
    queryKeyboardShortcuts.data?.data?.items || [];

  const onChange = (
    item: KeyBinding,
    shortcut?: Partial<Shortcut> | undefined
  ) => {
    // patch(
    //   { uniqueId: item.uniqueId, userCombination: shortcut } as any,
    //   undefined as any
    // );
  };

  return (
    <>
      <QueryErrorView query={queryKeyboardShortcuts} />
      <KeyBindings onChange={onChange} items={shortcuts} />
    </>
  );
}
