import classNames from "classnames";
import { KeyboardEventHandler, useState } from "react";
import { Shortcut } from "../KeyboardShortcutDefinitions";
import { KeyCombination } from "./KeyCombination";

export function KeyViewer({
  shortcut,
  onChange,
  readonly,
}: {
  shortcut?: Shortcut;
  onChange?: (shortcut?: Partial<Shortcut>) => void;
  readonly?: boolean;
}) {
  const [inCaptureState, setCaptureState] = useState(false);
  const [previewBinding, setPreviewBinding] = useState<Partial<Shortcut>>({});

  function keydownHandler(e: KeyboardEvent) {
    const newBinding: Partial<Shortcut> = {
      altKey: e.altKey,
      ctrlKey: e.ctrlKey,
      metaKey: e.metaKey,
      shiftKey: e.shiftKey,
      key: e.code,
    };

    setPreviewBinding(newBinding);

    if (!["Shift", "Meta", "Alt", "Control"].includes(e.key)) {
      window.removeEventListener("keydown", keydownHandler);
      setCaptureState(false);

      if (onChange) {
        onChange(newBinding);
      }
    }
  }

  const onClick = () => {
    if (readonly === true) {
      return;
    }

    setPreviewBinding({});
    setCaptureState(true);
    window.addEventListener("keydown", keydownHandler);
  };

  const onKeyUp = (e: any) => {
    if (e.code === "Enter") {
      onClick();
    }
  };

  return (
    <div
      className={classNames({ "in-capture-state": inCaptureState })}
      onClick={onClick}
      onKeyUp={onKeyUp}
      tabIndex={1}
    >
      <KeyCombination
        inCaptureState={inCaptureState}
        shortcut={inCaptureState ? previewBinding : shortcut}
      />
    </div>
  );
}
