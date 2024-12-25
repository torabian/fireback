import { Route } from "react-router-dom";
import { KeyboardShortcutArchiveScreen } from "./KeyboardShortcutArchiveScreen";
import { KeyboardShortcutEntity } from "../../sdk/modules/accessibility/KeyboardShortcutEntity";

export function useKeyboardShortcutRoutes() {
  return (
    <>
      <Route
        element={<KeyboardShortcutArchiveScreen />}
        path={KeyboardShortcutEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
