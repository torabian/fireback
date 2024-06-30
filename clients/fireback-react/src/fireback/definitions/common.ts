export interface QueryArchiveColumn {
  name?: string;
  width?: number;
  title?: string;
  getCellValue?: (m: any) => any;
}
export interface Timestamp {
  seconds: number;
  nanos: number;
}

export interface DisplayDetectionProps {
  location?: string;
  selectedUrw?: any;
  asPath?: string;
  userRoleWorkspaces?: any[];
}

export interface MenuItem {
  label?: string | null;
  href?: string | null;
  color?: string | null;
  icon?: string | null;
  onClick?: () => void;
  activeMatcher?: RegExp;
  displayFn?: (props: DisplayDetectionProps) => boolean;
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
  children: MenuItemRendered[];
}

export enum MacTagsColor {
  Green = "#00bd00",
  Red = "#ff0313",
  Orange = "#fa7a00",
  Yellow = "#f4b700",
  Blue = "#0072ff",
  Purple = "#ad41d1",
  Grey = "#717176",
}
