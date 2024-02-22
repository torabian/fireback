/**
 * Action menu stands for those menus, which can accept some buttons, and change
 * based on the page user is.
 */
import React, { useContext, useEffect, useMemo, useState } from "react";
import { uniqBy } from "lodash";
import { useT } from "@/hooks/useT";
import { KeyboardAction, PermissionLevel } from "@/definitions/definitions";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { userMeetsAccess } from "@/hooks/accessLevels";
import { osResources } from "../mulittarget/multitarget-resource";
import { useKeyCombination, useKeyPress } from "@/hooks/useKeyPress";
import classNames from "classnames";

export function ActionMenuManager({
  filter,
}: {
  filter?: (data: any) => boolean;
}) {
  const t = useContext(ActionMenuContext);

  return (
    <>
      {(t.refs || []).filter(filter ? filter : Boolean).map((item) => (
        <ActionMenu key={item.id} mref={item} />
      ))}
    </>
  );
}

export type onTriggerFn = (action: string) => void;

export function ActionMenu({ mref }: { mref: ActionMenuRef }) {
  return (
    <div className="action-menu">
      <ul className="navbar-nav">
        {mref.actions?.map((t) => (
          <ActionMenuItem item={t} key={t.uniqueActionKey} />
        ))}
      </ul>
    </div>
  );
}

export function ActionMenuItem({ item }: { item: IMenuActionItem }) {
  if (item.Component) {
    const Component = item.Component;
    return (
      <li className="action-menu-item">
        <Component />
      </li>
    );
  }
  return (
    <li
      className={classNames("action-menu-item", item.className)}
      onClick={item.onSelect}
    >
      {item.icon ? (
        <span>
          <img
            src={(process.env.REACT_APP_PUBLIC_URL || "") + item.icon}
            title={item.label}
            alt={item.label}
          />
        </span>
      ) : (
        <span>{item.label}</span>
      )}
    </li>
  );
}

export interface IMenuActionItem {
  label?: string;
  icon?: string;
  className?: string;
  uniqueActionKey: string;
  Component?: any;
  keyboardAction?: KeyboardAction;
  onSelect?: () => void;
}

export interface ActionMenuRef {
  id: string;
  actions?: IMenuActionItem[];
}

export type SetActionMenuFn = (
  menuName: string,
  items: IMenuActionItem[]
) => void;

export interface IActionMenuContext {
  refs: Array<ActionMenuRef>;
  setActionMenu: SetActionMenuFn;
  removeActionMenu: (menuName: string) => void;
  removeActionMenuItems: (menuName: string, items: string[]) => void;
}

export const ActionMenuContext = React.createContext<IActionMenuContext>({
  setActionMenu() {},
  removeActionMenu() {},
  removeActionMenuItems(menuName, items) {},
  refs: [],
});

export interface ActionMenuOptions {
  onTrigger: (actionKey: string) => void;
}

export function useMenuTools() {
  const t = useContext(ActionMenuContext);

  const addActions = (menu: string, items: IMenuActionItem[]) => {
    t.setActionMenu(menu, items);

    return () => t.removeActionMenu(menu);
  };

  const deleteActions = (menu: string, items: string[]) => {
    t.removeActionMenuItems(menu, items);
  };

  const removeActionMenu = (menu: string) => {
    t.removeActionMenu(menu);
  };

  return {
    addActions,
    removeActionMenu,
    deleteActions,
  };
}

export function useActions(
  menuName: string,
  items: Array<IMenuActionItem | undefined>,
  options?: ActionMenuOptions,
  deps?: any[]
) {
  const t = useContext(ActionMenuContext);

  useEffect(() => {
    t.setActionMenu(menuName, items.filter((t) => t !== undefined) as any);

    return () => {
      t.removeActionMenu(menuName);
    };
  }, deps || []);

  return {
    addActions(items: IMenuActionItem[], m = menuName) {
      t.setActionMenu(m, items);
    },
    deleteActions(items: string[], m = menuName) {
      t.removeActionMenuItems(m, items);
    },
  };
}

function combineArray(
  items?: IMenuActionItem[],
  items2?: IMenuActionItem[]
): IMenuActionItem[] {
  if (!items && !items2) {
    return [];
  }
  if (!items) {
    return items2 || [];
  }
  if (!items2) {
    return items || [];
  }

  let newItems: IMenuActionItem[] = [...items, ...items2];

  return uniqBy(newItems.reverse(), "uniqueActionKey");
}

export function ActionMenuProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [refs, setItemsRefs] = useState<Array<ActionMenuRef>>([]);

  const setActionMenu: SetActionMenuFn = (menuName, actions) => {
    setItemsRefs((refs) => {
      const existingMenu = refs.find((t) => t.id === menuName);

      if (!existingMenu) {
        refs.push({ id: menuName, actions });
      } else {
        refs = refs.map((t) => {
          if (t.id === menuName) {
            return { ...t, actions };
          }
          return t;
        });
      }

      return [...refs];
    });
  };

  const removeActionMenu = (menuName: string) => {
    setItemsRefs((refs) => [...refs.filter((t) => t.id !== menuName)]);
  };

  const removeActionMenuItems = (menuName: string, toDelete: string[]) => {
    for (let i = 0; i < refs.length; i++) {
      const x = refs[i];
      if (x.id === menuName && x.actions) {
        let newMenu: IMenuActionItem[] = [];
        for (const y of x.actions) {
          const isInDeleteList = !!toDelete.find(
            (t) => t === y.uniqueActionKey
          );

          if (!isInDeleteList) {
            newMenu.push(y);
          }
        }
        x.actions = newMenu;
      }
    }

    const newRefs = [...refs];

    setItemsRefs(newRefs);
  };

  return (
    <ActionMenuContext.Provider
      value={{
        refs,
        setActionMenu,
        removeActionMenuItems,
        removeActionMenu,
      }}
    >
      {children}
    </ActionMenuContext.Provider>
  );
}

export function useCommonCrudActions({
  onCancel,
  onSave,
  access,
}: {
  onSave: () => void;
  onCancel?: () => void;
  access?: PermissionLevel;
}) {
  const { selectedUrw } = useContext(RemoteQueryContext);

  const meets = useMemo(() => {
    if (!access) {
      return true;
    }

    if (selectedUrw?.workspaceId !== "root" && access?.onlyRoot) {
      return false;
    }

    if (!access?.permissions || access.permissions.length === 0) {
      return true;
    }

    return userMeetsAccess(selectedUrw as any, access.permissions[0]);
  }, [selectedUrw, access]);

  const t = useT();
  const editingCore = ({
    onSave,
    onCancel,
  }: {
    onSave: () => void;
    onCancel?: () => void;
  }) => {
    if (!meets) {
      return [];
    }

    return [
      {
        icon: "",
        label: t.common.save,
        uniqueActionKey: "save",
        onSelect: () => {
          onSave();
        },
      },
      onCancel && {
        icon: "",
        label: t.common.cancel,
        uniqueActionKey: "cancel",
        onSelect: () => {
          onCancel();
        },
      },
    ];
  };

  useActions(
    "editing-core",
    editingCore({
      onCancel,
      onSave,
    })
  );
}

export function useNewAction(
  onSelect: (() => void) | undefined,
  keyPressEventName?: KeyboardAction
) {
  const t = useT();

  useKeyCombination(keyPressEventName, onSelect);

  useActions("commonEntityActions", [
    onSelect && {
      icon: osResources.add,
      label: t.actions.new,
      uniqueActionKey: "new",
      onSelect,
    },
  ]);
}

export function useBackButton(
  onSelect: (() => void) | undefined,
  keyPressEventName?: KeyboardAction
) {
  const t = useT();

  useKeyCombination(keyPressEventName, onSelect);

  useActions("navigation", [
    onSelect && {
      icon: osResources.left,
      label: t.actions.back,
      uniqueActionKey: "back",
      className: "navigator-back-button",
      onSelect,
    },
  ]);
}

export function useCommonArchiveExportTools() {
  const { session, options } = useContext(RemoteQueryContext);

  useExportActions(() => {
    function toBinaryString(data: any) {
      var ret = [];
      var len = data.length;
      var byte;
      for (var i = 0; i < len; i++) {
        byte = (data.charCodeAt(i) & 0xff) >>> 0;
        ret.push(String.fromCharCode(byte));
      }

      return ret.join("");
    }

    var xhr = new XMLHttpRequest();

    xhr.open("GET", options.prefix + "roles/export");

    xhr.addEventListener(
      "load",
      function () {
        var data = toBinaryString(this.responseText);
        data = "data:application/text;base64," + btoa(data);
        document.location = data;
      },
      false
    );

    const h: any = options?.headers;
    xhr.setRequestHeader("Authorization", h.authorization || "");
    xhr.setRequestHeader("Workspace-Id", h["workspace-id"] || "");
    xhr.setRequestHeader("role-Id", h["role-id"] || "");
    xhr.overrideMimeType("application/octet-stream; charset=x-user-defined;");
    xhr.send(null);
  }, KeyboardAction.ExportTable);
}

export function useExportActions(
  onSelect: (() => void) | undefined,
  keyPressEventName?: KeyboardAction
) {
  const t = useT();

  useKeyCombination(keyPressEventName, onSelect);

  useActions("exportTools", [
    onSelect && {
      icon: osResources.export,
      label: t.actions.new,
      uniqueActionKey: "export",
      onSelect,
    },
  ]);
}

export function useEditAction(
  onSelect: (() => void) | undefined,
  keyPressEventName?: KeyboardAction
) {
  const t = useT();
  useKeyCombination(keyPressEventName, onSelect);

  useActions("commonEntityActions", [
    onSelect && {
      icon: osResources.edit,
      label: t.actions.edit,
      uniqueActionKey: "new",
      onSelect,
    },
  ]);
}
