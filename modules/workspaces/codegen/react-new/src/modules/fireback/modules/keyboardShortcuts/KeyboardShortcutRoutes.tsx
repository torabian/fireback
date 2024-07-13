import { KeyboardShortcutEntity } from "../../sdk/modules/keyboardActions/KeyboardShortcutEntity";
import { Route } from "react-router-dom";
import { KeyboardShortcutArchiveScreen } from "./KeyboardShortcutArchiveScreen";

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
