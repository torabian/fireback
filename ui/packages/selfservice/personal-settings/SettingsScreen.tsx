import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";

import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";
import { AccessiblitySettings } from "./AccessiblitySettings";
import { DebuggerSettings } from "./DebuggerSettings";
import { InterfaceSettings } from "./InterfaceSettings";
import { NotificationSettings } from "./NotificationSettings";
import { RichTextEditorSettings } from "./RichTextEditorSettings";
import { ThemeSettings } from "./ThemeSettings";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";

export function SettingsScreen({}: {}) {
  const s = useS(strings);
  usePageTitle(s.pageTitle);

  return (
    <div>
      <NotificationSettings />
      {BUILD_VARIABLES.FORCED_LOCALE ? null : <InterfaceSettings />}
      <RichTextEditorSettings />
      <AccessiblitySettings />
      {BUILD_VARIABLES.FORCE_APP_THEME !== "true" ? <ThemeSettings /> : null}
      <DebuggerSettings />
    </div>
  );
}
