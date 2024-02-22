import { KeyboardShortcutNavigationTools } from "src/sdk/fireback/modules/keyboardActions/keyboard-shortcut-navigation-tools";
import { Route } from "react-router-dom";
import { KeyboardShortcutArchiveScreen } from "./KeyboardShortcutArchiveScreen";

export function useKeyboardShortcutRoutes() {
  return (
    <>
      <Route
        element={<KeyboardShortcutArchiveScreen />}
        path={KeyboardShortcutNavigationTools.Rquery}
      ></Route>
    </>
  );
}
