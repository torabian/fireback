import {
  type MenuItem,
  type MenuItemRendered,
  type MenuRendered,
} from "../../types/MenuItem";
import { useLocale } from "../../hooks/useLocale";
import { useSortableOrder } from "../../hooks/useSortableOrder";
import classNames from "classnames";
import { useContext, type ReactNode } from "react";
import ActiveLink from "../link/ActiveLink";
import { MenuItemContent } from "./MenuItemContent";
import { useUserWorkspaceBrowseActionQuery } from "../../../sdk/abac/UserWorkspaceBrowseAction";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
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

function renderMenu(
  menu: MenuItem,
  data: {
    urw?: any;
    urws: any[];
    asPath: string;
  },
): MenuRendered | null {
  let hasChildren = false;
  const children = menu.children?.map((item): MenuItemRendered => {
    let forceActive = item.activeMatcher
      ? item.activeMatcher.test(data.asPath)
      : undefined;

    if (item.forceActive) {
      forceActive = true;
    }

    let isVisible = item.displayFn
      ? item.displayFn({
          location: "here",
          asPath: data.asPath,
          selectedUrw: data.urw,
          userRoleWorkspaces: data.urws,
        })
      : true;

    if (isVisible) {
      hasChildren = true;
    }

    return {
      ...item,
      isActive: forceActive || false,
      isVisible,
    };
  });

  if (hasChildren === false && !menu.href) {
    return null;
  }

  return {
    name: menu.label,
    href: menu.href,
    children,
  };
}
export function MenuParticle({
  menu,
  onClick,
  dragHandle,
}: {
  menu: MenuItem;
  onClick: () => void;
  // Rendered by the caller (it owns the useSortable() attributes/listeners)
  // and placed right next to the group's own header/row here, so it lines
  // up with the header instead of floating at the top of the whole block.
  dragHandle?: ReactNode;
}) {
  const { asPath } = useLocale();
  const { selectedUrw } = useContext(RemoteQueryContext) as any;
  const queryWorkspaces = useUserWorkspaceBrowseActionQuery({});

  const menuRendered = renderMenu(menu, {
    asPath,
    urw: selectedUrw as any,
    urws: queryWorkspaces.data?.data?.items || [],
  });

  if (!menuRendered) {
    return null;
  }

  return (
    <div className="sidebar-menu-particle" onClick={onClick}>
      {menu.children?.length > 0 ? (
        <>
          <span className="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none sidebar-menu-header">
            <span className="category">{menu.label}</span>
            {dragHandle}
          </span>
          <MenuUl items={menuRendered.children} />
        </>
      ) : (
        <li className="nav-item">
          <div className="nav-item-row">
            <LiItemLink item={menu as any} />
            {dragHandle}
          </div>
        </li>
      )}
    </div>
  );
}

function menuItemId(item: MenuItemRendered) {
  return `${item.label}_${item.href ?? ""}`;
}

function MenuUl({ items }: { items: MenuItemRendered[] }) {
  const { ordered, reorder } = useSortableOrder(items || [], menuItemId);
  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 4 } }),
  );

  // Sorting here is a local, session-only display order — see
  // useSortableOrder. It does not persist and does not write back to the
  // remote menu config.
  function handleDragEnd(event: DragEndEvent) {
    const { active, over } = event;
    if (over && active.id !== over.id) {
      reorder(String(active.id), String(over.id));
    }
  }

  if (ordered.length === 0) {
    return null;
  }

  return (
    <DndContext
      sensors={sensors}
      collisionDetection={closestCenter}
      onDragEnd={handleDragEnd}
      // The sidebar itself scrolls (overflow-y: auto); dnd-kit's default
      // auto-scroll walks up to that ancestor and scrolls it whenever the
      // pointer nears its edges, which reads as an unwanted "jump" while
      // dragging sideways in a narrow list. These lists are short and
      // never need auto-scrolling, so it's simplest to turn it off.
      autoScroll={false}
      // These are single vertical lists — lock the drag to the Y axis so
      // horizontal mouse movement can't drag items sideways.
      modifiers={[restrictToVerticalAxis]}
    >
      <SortableContext
        items={ordered.map(menuItemId)}
        strategy={verticalListSortingStrategy}
      >
        <ul className="nav nav-pills flex-column mb-auto">
          {ordered.map((item) => (
            <SortableLiItem
              key={menuItemId(item)}
              id={menuItemId(item)}
              item={item}
            />
          ))}
        </ul>
      </SortableContext>
    </DndContext>
  );
}

function LiItemLink({ item }: { item: MenuItemRendered }) {
  return item.href && !item.onClick ? (
    <ActiveLink
      replace
      href={item.href}
      className="nav-link"
      aria-current="page"
      forceActive={item.isActive}
      scroll={null}
      inActiveClassName="text-white"
      activeClassName="active"
    >
      <MenuItemContent item={item} />
    </ActiveLink>
  ) : (
    <a
      className={classNames("nav-link", item.isActive && "active")}
      onClick={item.onClick}
    >
      <MenuItemContent item={item} />
    </a>
  );
}

function SortableLiItem({ id, item }: { id: string; item: MenuItemRendered }) {
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

  return (
    <li
      ref={setNodeRef}
      style={style}
      className={classNames("nav-item", isDragging && "is-dragging")}
    >
      <div className="nav-item-row">
        <span
          className="drag-handle"
          aria-hidden="true"
          {...attributes}
          {...listeners}
        >
          <GripVertical size={12} />
        </span>
        <LiItemLink item={item} />
      </div>
      {item.children && item.children.length > 0 && (
        <MenuUl items={item.children as any} />
      )}
    </li>
  );
}
