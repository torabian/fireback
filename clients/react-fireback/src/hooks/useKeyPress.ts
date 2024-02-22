import { KeyboardAction } from "@/definitions/definitions";
import { useEffect } from "react";

// App predefined combination that user can change
export function useKeyCombination(
  action?: KeyboardAction,
  handler?: () => void
) {
  // Later use the backend information for this, this is where you need to manage also CTRL+x
  const mapper = {
    [KeyboardAction.NewEntity]: "n",
    [KeyboardAction.NewChildEntity]: "n",
    [KeyboardAction.EditEntity]: "e",
    [KeyboardAction.ViewQuestions]: "q",
    [KeyboardAction.Delete]: "Backspace",
    [KeyboardAction.StopStart]: " ",
    [KeyboardAction.ExportTable]: "x",
    [KeyboardAction.CommonBack]: "Escape",
  };

  return useKeyPress(action ? mapper[action] : undefined, handler);
}

export function useKeyPress(key?: string, handler?: () => void) {
  useEffect(() => {
    if (!key || !handler) {
      return;
    }
    function keyupHanlder(event: any) {
      var e = event || window.event,
        target = e.target || e.srcElement;

      if (
        ["INPUT", "TEXTAREA", "SELECT"].includes(target.tagName.toUpperCase())
      ) {
        if (event.key === "Escape") {
          e.target.blur();
        }
        return;
      }

      if (event.key === key) {
        handler && handler();
      }
    }

    if (!handler) {
      return;
    }

    window.addEventListener("keyup", keyupHanlder);

    return () => {
      window.removeEventListener("keyup", keyupHanlder);
    };
  }, [handler, key]);
}
