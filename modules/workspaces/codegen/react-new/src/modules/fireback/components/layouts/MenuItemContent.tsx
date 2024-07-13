import { MenuItem } from "@/modules/fireback/definitions/common";
import { source } from "@/modules/fireback/hooks/source";
import { osResources } from "@/modules/fireback/resources/resources";

export function MenuItemContent({ item }: { item: MenuItem }) {
  return (
    <span>
      {item.icon && <img className="menu-icon" src={source(item.icon)} />}
      {item.color && !item.icon ? (
        <span
          className="tag-circle"
          style={{ backgroundColor: item.color }}
        ></span>
      ) : null}
      <span className="nav-link-text">{item.label}</span>
    </span>
  );
}
