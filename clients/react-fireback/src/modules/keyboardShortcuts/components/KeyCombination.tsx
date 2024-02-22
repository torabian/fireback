import { useT } from "@/hooks/useT";
import { Shortcut } from "../KeyboardShortcutDefinitions";

export function KeyCombination({
  shortcut,
  inCaptureState,
}: {
  shortcut?: Partial<Shortcut>;
  inCaptureState: boolean;
}) {
  const t = useT();
  if (!shortcut) {
    return <span>{t.keyboardShortcut.pressToDefine}</span>;
  }

  return (
    <span className="keybinding-combination">
      {shortcut.altKey && <span>alt</span>}
      {shortcut.ctrlKey && <span>ctrl</span>}
      {shortcut.shiftKey && <span>shift</span>}
      {shortcut.metaKey && <span>super</span>}
      {shortcut.key && !inCaptureState && <span>{shortcut.key}</span>}
      {inCaptureState && <span>??</span>}
    </span>
  );
}
