import { UserRoleWorkspaceEntity } from "src/sdk/fireback";

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
  selectedUrw?: UserRoleWorkspaceEntity;
  asPath?: string;
  userRoleWorkspaces?: UserRoleWorkspaceEntity[];
}

export interface MenuItem {
  label: string;
  href?: string;
  color?: string;
  icon?: string;
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
  name: string;
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
