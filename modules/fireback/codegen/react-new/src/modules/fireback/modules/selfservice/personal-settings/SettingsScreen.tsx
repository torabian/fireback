import { useT } from "../../../hooks/useT";

import { usePageTitle } from "../../../components/page-title/PageTitle";
import { AccessiblitySettings } from "./AccessiblitySettings";
import { DebuggerSettings } from "./DebuggerSettings";
import { InterfaceSettings } from "./InterfaceSettings";
import { NotificationSettings } from "./NotificationSettings";
import { RichTextEditorSettings } from "./RichTextEditorSettings";
import { ThemeSettings } from "./ThemeSettings";
import { BUILD_VARIABLES } from "@/modules/fireback/hooks/build-variables";

export function SettingsScreen({}: {}) {
  const t = useT();
  usePageTitle(t.menu.settings);
 
  return (
    <div>
      <NotificationSettings />
      {BUILD_VARIABLES.FORCED_LOCALE ? null : <InterfaceSettings />}
      <RichTextEditorSettings />
      <AccessiblitySettings />
      {BUILD_VARIABLES.FORCE_APP_THEME !== "true" ? (
        <ThemeSettings />
      ) : null}
      <DebuggerSettings />
    </div>
  );
}
