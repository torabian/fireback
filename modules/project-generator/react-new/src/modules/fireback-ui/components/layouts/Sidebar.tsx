import { type MenuItem } from "../../types/MenuItem";
import { source } from "../../hooks/source";
import { useUiState } from "../../hooks/uiStateContext";

import classNames from "classnames";
import React, { useContext } from "react";
import { BUILD_VARIABLES } from "../../hooks/build-variables";
import { detectDeviceType } from "../../hooks/deviceInformation";
import { useRemoteMenuResolver } from "../../hooks/useRemoteMenuResolver";
import { useSortableOrder } from "../../hooks/useSortableOrder";
import { osResources } from "../../hooks/resources";
import type { AppMenuOptionalDto } from "../../../fireback/sdk/interfacetools/AppMenuOptionalDto";
import { AppMenuDto } from "../../../fireback/sdk/interfacetools/AppMenuDto";
import { ReactiveSearchContext } from "../reactive-search/ReactiveSearchContext";
import { CurrentUser } from "./CurrentUser";
import { MenuParticle } from "./MenuParticle";
import { useWorkspacesMenuPresenter } from "./useWorkspacesMenuPresenter";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import {
  DndContext,
  PointerSensor,
  closestCenter,
  useSensor,
  useSensors,
  type DragEndEvent,
} from "@dnd-kit/core";
import {
  SortableContext,
  useSortable,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import { restrictToVerticalAxis } from "@dnd-kit/modifiers";
import { GripVertical } from "lucide-react";

export function dataMenuToMenu(
  data: AppMenuOptionalDto,
  permissionCheck: (permissionKey?: string | null) => boolean = () => true,
  locale: string,
): MenuItem | null {
  if (!permissionCheck(data.capabilityId)) {
    return null;
  }

  const children = (data.children || [])
    .map((v: AppMenuOptionalDto) => dataMenuToMenu(v, permissionCheck, locale))
    .filter(Boolean) as MenuItem[];

  let label = data.label || "";

  if (typeof label !== "string") {
    label = label[locale];
  }

  return {
    label,

    children,
    displayFn: castMenuDefinitionToDisplayFn(data),
    icon: data.icon,
    href: data.href,
    activeMatcher: data.activeMatcher
      ? new RegExp(data.activeMatcher)
      : undefined,
  };
}

function castMenuDefinitionToDisplayFn(data: AppMenuDto) {
  return () => true;
}

export const defaultNavbar: MenuItem = {
  label: "Navbar",
  children: [],
};

function Sidebar({
  miniSize,
  onClose,
  sidebarItemSelectedExtra,
}: {
  miniSize: boolean;
  onClose?: () => void;
  sidebarItemSelectedExtra?: () => void;
}) {
  const {
    sidebarVisible,
    toggleSidebar: toggleSidebar$,
    sidebarItemSelected,
  } = useUiState();
  const menu = useRemoteMenuResolver("sidebar");
  const s = useS(strings);

  const { reset } = useContext(ReactiveSearchContext);

  const toggleSidebar = () => {
    reset();
    toggleSidebar$();
  };

  if (!menu) {
    return null;
  }

  let menus: MenuItem[] = [];
  if (Array.isArray(menu)) {
    menus = [...menu];
  } else if ((menu as any).children?.length) {
    menus.push(menu);
  }

  const { menus: workspaceMenus } = useWorkspacesMenuPresenter();
  menus.push(workspaceMenus[0]);

  const menuGroups = menus.map((m, index) => ({
    id: (typeof m.label === "string" && m.label) || `group-${index}`,
    menu: m,
  }));

  // Session-only reordering of the top-level menu groups — see
  // useSortableOrder. Nothing is persisted, so this resets on refresh.
  const { ordered: orderedGroups, reorder: reorderGroups } = useSortableOrder(
    menuGroups,
    (g) => g.id,
  );
  const groupSensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 4 } }),
  );

  function handleGroupDragEnd(event: DragEndEvent) {
    const { active, over } = event;
    if (over && active.id !== over.id) {
      reorderGroups(String(active.id), String(over.id));
    }
  }

  return (
    <div
      data-wails-drag
      className={classNames(
        miniSize ? "sidebar-extra-small" : "",
        "sidebar",
        sidebarVisible ? "open" : "",
        "scrollable-element",
        detectDeviceType().isMobileView ? "has-bottom-tab" : undefined,
      )}
    >
      <button
        className="sidebar-close"
        onClick={() => (onClose ? onClose() : toggleSidebar())}
      >
        <img src={source(osResources.cancel)} />
      </button>

      <DndContext
        sensors={groupSensors}
        collisionDetection={closestCenter}
        onDragEnd={handleGroupDragEnd}
        // See MenuParticle.tsx — the sidebar is a scroll container, and
        // dnd-kit's default auto-scroll fights with that while dragging.
        autoScroll={false}
        modifiers={[restrictToVerticalAxis]}
      >
        <SortableContext
          items={orderedGroups.map((g) => g.id)}
          strategy={verticalListSortingStrategy}
        >
          {orderedGroups.map((g) => (
            <SortableMenuGroup
              key={g.id}
              id={g.id}
              menu={g.menu}
              onClick={() => {
                sidebarItemSelected();
                sidebarItemSelectedExtra?.();
              }}
            />
          ))}
        </SortableContext>
      </DndContext>
      {BUILD_VARIABLES.GITHUB_DEMO === "true" && (
        <MenuParticle
          onClick={() => {
            sidebarItemSelected();
            sidebarItemSelectedExtra?.();
          }}
          menu={{
            label: s.components.demo,
            children: [
              {
                label: s.components.formSelect,
                icon: "/ios-theme/icons/settings.svg",
                children: [],
                href: "/demo/form-select",
              },
              {
                label: s.components.formDateTime,
                icon: "/ios-theme/icons/settings.svg",
                children: [],
                href: "/demo/form-date",
              },
              {
                label: s.components.overlaysAndModal,
                icon: "/ios-theme/icons/settings.svg",
                children: [],
                href: "/demo/modals",
              },
            ],
          }}
        />
      )}
      <CurrentUser
        onClick={() => {
          sidebarItemSelected();
          sidebarItemSelectedExtra?.();
        }}
      />
    </div>
  );
}

function SortableMenuGroup({
  id,
  menu,
  onClick,
}: {
  id: string;
  menu: MenuItem;
  onClick: () => void;
}) {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  const dragHandle = (
    <span
      className="drag-handle"
      aria-hidden="true"
      {...attributes}
      {...listeners}
    >
      <GripVertical size={14} />
    </span>
  );

  return (
    <div
      ref={setNodeRef}
      style={style}
      className={classNames("sortable-menu-group", isDragging && "is-dragging")}
    >
      <MenuParticle menu={menu} onClick={onClick} dragHandle={dragHandle} />
    </div>
  );
}

export default React.memo(Sidebar);
