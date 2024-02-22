import { useT } from "@/hooks/useT";
import { RemoteServiceSetting } from "./RemoteServiceSetting";

import { useRef, useState } from "react";
import { InterfaceSettings } from "./InterfaceSettings";
import { RichTextEditorSettings } from "./RichTextEditorSettings";
import { DebuggerSettings } from "./DebuggerSettings";
import { ThemeSettings } from "./ThemeSettings";
import { AccessiblitySettings } from "./AccessiblitySettings";

export function SettingsScreen({}: {}) {
  const t = useT();

  const editorRef: any = useRef(null);

  return (
    <div>
      {/* {process.env.REACT_APP_FORCE_REMOTE_SERVICE !== "true" ? (
        <RemoteServiceSetting />
      ) : null} */}
      {process.env.REACT_APP_FORCED_LOCALE ? null : <InterfaceSettings />}
      <RichTextEditorSettings />
      <AccessiblitySettings />
      {process.env.REACT_APP_FORCE_APP_THEME !== "true" ? (
        <ThemeSettings />
      ) : null}
      <DebuggerSettings />
    </div>
  );
}
