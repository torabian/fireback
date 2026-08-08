export interface MenuItem {
  label?: string | null;
  href?: string | null;
  color?: string | null;
  icon?: string | null;
  onClick?: () => void;
  activeMatcher?: RegExp;
  displayFn?: (props: any) => boolean;
  forceActive?: boolean;
  children: MenuItem[];
  key?: string;
}

export interface MenuItemRendered extends MenuItem {
  isActive: boolean;
  isVisible: boolean;
}

export interface MenuRendered {
  name?: string | null;
  href?: string | null;
  children: MenuItemRendered[];
}
