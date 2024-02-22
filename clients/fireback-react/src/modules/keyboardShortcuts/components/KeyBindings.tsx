import { useT } from "@/hooks/useT";
import { KeyBinding, Shortcut } from "../KeyboardShortcutDefinitions";
import { KeyViewer } from "./KeyViewer";

export function KeyBindings({
  items,
  onChange,
}: {
  items: KeyBinding[];
  onChange?: (item: KeyBinding, shortcut?: Partial<Shortcut>) => void;
}) {
  const t = useT();
  return (
    <div>
      <table className="table dto-view-table">
        <thead>
          <tr>
            <th style={{ width: "250px" }}>{t.keyboardShortcut.action}</th>
            <th style={{ width: "350px" }}>
              {t.keyboardShortcut.defaultBinding}
            </th>
            <th>{t.keyboardShortcut.userDefinedBinding}</th>
          </tr>
        </thead>
        <tbody>
          {items.map((item) => {
            return (
              <tr key={item.action}>
                <th scope="row">{item.action}</th>
                <td className="table-active">
                  <KeyViewer readonly shortcut={item.defaultCombination} />
                </td>
                <td className="table-active">
                  <KeyViewer
                    onChange={(shortcut) =>
                      onChange && onChange(item, shortcut)
                    }
                    shortcut={item.userCombination}
                  />
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
