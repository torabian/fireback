import { KeyboardAction } from "@/definitions/definitions";
import { useEffect } from "react";

// App predefined combination that user can change
export function useKeyCombination(
  action?: KeyboardAction | KeyboardAction[],
  handler?: (key: string) => void
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
    [KeyboardAction.Select1Index]: "1",
    [KeyboardAction.Select2Index]: "2",
    [KeyboardAction.Select3Index]: "3",
    [KeyboardAction.Select4Index]: "4",
    [KeyboardAction.Select5Index]: "5",
    [KeyboardAction.Select6Index]: "6",
    [KeyboardAction.Select7Index]: "7",
    [KeyboardAction.Select8Index]: "8",
    [KeyboardAction.Select9Index]: "9",
  };

  let mappedAction;
  if (typeof action === "object") {
    mappedAction = action.map((v) => mapper[v]);
  } else {
    mappedAction = mapper[action];
  }

  return useKeyPress(mappedAction, handler);
}

export function useKeyPress(
  key?: string | string[],
  handler?: (key: any) => void
) {
  useEffect(() => {
    if (!key || key.length === 0 || !handler) {
      return;
    }
    function keyupHanlder(event: any) {
      var e = event || window.event,
        target = e.target || e.srcElement;

      const tag = target.tagName.toUpperCase();
      const Type = target.type;
      if (["TEXTAREA", "SELECT"].includes(tag)) {
        if (event.key === "Escape") {
          e.target.blur();
        }
        return;
      }

      if (tag === "INPUT" && (Type === "text" || Type === "password")) {
        return;
      }

      const matched =
        typeof key === "string" ? event.key === key : key.includes(event.key);
      if (matched) {
        handler && handler(event.key);
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
