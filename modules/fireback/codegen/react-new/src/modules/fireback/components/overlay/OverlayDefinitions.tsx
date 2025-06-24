import { ReactNode } from "react";

export type OverlayPresentationType = "drawer" | "modal";

export interface BaseModalOpenParams {
  title?: string;
}

export interface DrawerOpenParams {
  direction?: "right" | "left" | "top" | "bottom";
  size?: string | number;
}

export type OpenOverlayProps = BaseModalOpenParams | DrawerOpenParams;

export interface OpenOverlayConfig {
  type?: OverlayPresentationType;
  params?: OpenOverlayProps;
}

export type OverlayContextType = {
  openOverlay: <T = void>(
    Component: OverlayInstance<T>["Component"],
    params?: OpenOverlayConfig
  ) => OverlayController<T>;
  openModal: <T = void>(
    Component: OverlayInstance<T>["Component"],
    params?: BaseModalOpenParams
  ) => OverlayController<T>;
  openDrawer: <T = void>(
    Component: OverlayInstance<T>["Component"],
    params?: DrawerOpenParams
  ) => OverlayController<T>;

  dismissAll: () => void;
};

export interface DialogResult<T> {
  data: T;
  type: "closed" | "resolved" | "rejected";
}

interface OverlayControlActions<T> {
  resolve: (result?: T) => void;
  close: () => void;
  reject: (reason?: T) => void;
}

export type OverlayInstance<T, M = {}> = {
  id: number;
  ref: React.RefObject<any>;
  Component: React.ComponentType<OverlayInstanceComponentProps<T>>;
  params?: OpenOverlayProps;
  data?: M;
  visible: boolean;
  type: OverlayPresentationType;
  onBeforeClose?: () => boolean | Promise<boolean>; // ✅ support hook
} & OverlayControlActions<T>;

export type OverlayInstanceComponentProps<T, V = OpenOverlayProps, M = {}> = {
  visible?: boolean;
  params?: V | undefined;
  data?: M;
  children?: ReactNode;
  setOnBeforeClose?: (fn: () => boolean | Promise<boolean>) => void; // ✅ inject setter
} & OverlayControlActions<T>;

export type OverlayController<T, M = undefined> = {
  id: number;
  ref: React.RefObject<any>;
  promise: Promise<DialogResult<T>>;
  close: () => void;
  reject: (reason?: any) => void;
  resolve: (result?: T) => void;
  updateData: (data: M) => void;
};
