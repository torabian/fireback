import { ReactNode } from "react";

/**
 * The type of overlay presentation.
 * - `"drawer"`: A side panel that slides in from an edge.
 * - `"modal"`: A centered popup dialog.
 */
export type OverlayPresentationType = "drawer" | "modal";

/**
 * Base configuration options for opening a modal.
 */
export interface BaseModalOpenParams {
  /**
   * Optional title to be shown on the modal.
   */
  title?: string;
}

/**
 * Configuration options for opening a drawer.
 */
export interface DrawerOpenParams {
  /**
   * The direction from which the drawer should appear.
   * Defaults typically to "right".
   */
  direction?: "right" | "left" | "top" | "bottom";

  /**
   * The size of the drawer. Can be a string (e.g., "400px") or a number.
   */
  size?: string | number;
}

/**
 * The union of all possible overlay configuration params.
 */
export type OpenOverlayProps = BaseModalOpenParams | DrawerOpenParams;

/**
 * Overlay configuration object passed to `openOverlay`.
 */
export interface OpenOverlayConfig {
  /**
   * The type of overlay to open — either a drawer or modal.
   */
  type?: OverlayPresentationType;

  /**
   * Parameters specific to the chosen overlay type.
   */
  params?: OpenOverlayProps;
}

/**
 * The shape of the context used to control overlays from anywhere in the app.
 */
export type OverlayContextType = {
  /**
   * Opens an overlay (either modal or drawer) with the given component.
   *
   * @param Component The React component to be rendered inside the overlay.
   * @param params Optional configuration for the overlay.
   * @returns A controller with methods and a promise to handle the overlay.
   */
  openOverlay: <T = void>(
    Component: OverlayInstance<T>["Component"],
    params?: OpenOverlayConfig
  ) => OverlayController<T>;

  /**
   * Opens a modal overlay.
   *
   * @param Component The modal component to render.
   * @param params Optional modal configuration (e.g., title).
   * @returns A controller to resolve, reject, or close the modal.
   */
  openModal: <T = void>(
    Component: OverlayInstance<T>["Component"],
    params?: BaseModalOpenParams
  ) => OverlayController<T>;

  /**
   * Opens a drawer overlay.
   *
   * @param Component The drawer component to render.
   * @param params Drawer configuration such as size and direction.
   * @returns A controller to manage the drawer's lifecycle.
   */
  openDrawer: <T = void>(
    Component: OverlayInstance<T>["Component"],
    params?: DrawerOpenParams
  ) => OverlayController<T>;

  /**
   * Dismisses all currently active overlays.
   */
  dismissAll: () => void;
};

/**
 * The result of an overlay interaction.
 */
export interface DialogResult<T> {
  /**
   * Optional result data returned by the overlay.
   */
  data: T;

  /**
   * The type of closure that occurred:
   * - `"resolved"`: User confirmed.
   * - `"rejected"`: User explicitly canceled.
   * - `"closed"`: User passively dismissed (e.g., backdrop click).
   */
  type: "closed" | "resolved" | "rejected";
}

/**
 * Internal actions available to overlays for controlling their own lifecycle.
 */
interface OverlayControlActions<T> {
  /**
   * Marks the overlay as resolved, returning optional result data.
   */
  resolve: (result?: T) => void;

  /**
   * Closes the overlay passively (e.g., user dismisses it without confirming).
   */
  close: () => void;

  /**
   * Marks the overlay as rejected, returning optional reason data.
   */
  reject: (reason?: T) => void;
}

/**
 * Internal representation of an overlay instance.
 */
export type OverlayInstance<T, M = {}> = {
  /**
   * Unique ID of the overlay.
   */
  id: number;

  /**
   * Ref to the overlay component instance.
   */
  ref: React.RefObject<any>;

  /**
   * The component rendered in the overlay.
   */
  Component: React.ComponentType<OverlayInstanceComponentProps<T>>;

  /**
   * Configuration parameters for the overlay.
   */
  params?: OpenOverlayProps;

  /**
   * Custom data passed down to the overlay.
   */
  data?: M;

  /**
   * Visibility flag — true if overlay is shown.
   */
  visible: boolean;

  /**
   * Whether it's a "modal" or "drawer".
   */
  type: OverlayPresentationType;

  /**
   * Optional hook to determine if overlay can close.
   * Returning false or a Promise resolving to false will block close.
   */
  onBeforeClose?: () => boolean | Promise<boolean>;
} & OverlayControlActions<T>;

/**
 * Props passed into an overlay component.
 */
export type OverlayInstanceComponentProps<T, V = OpenOverlayProps, M = {}> = {
  /**
   * Whether the overlay is currently visible.
   */
  visible?: boolean;

  /**
   * Modal/drawer configuration parameters.
   */
  params?: V | undefined;

  /**
   * Optional custom data passed in by caller.
   */
  data?: M;

  /**
   * Optional children for rendering inside overlay.
   */
  children?: ReactNode;

  /**
   * Used to set a function that runs before close and can block it.
   */
  setOnBeforeClose?: (fn: () => boolean | Promise<boolean>) => void;
} & OverlayControlActions<T>;

/**
 * Controller object returned by `openOverlay`, `openModal`, and `openDrawer`.
 */
export type OverlayController<T, M = undefined> = {
  /**
   * Unique ID of the overlay instance.
   */
  id: number;

  /**
   * Ref to the rendered overlay component.
   */
  ref: React.RefObject<any>;

  /**
   * Promise that resolves with a DialogResult when overlay is dismissed.
   */
  promise: Promise<DialogResult<T>>;

  /**
   * Closes the overlay (marked as `"closed"`).
   */
  close: () => void;

  /**
   * Rejects the overlay (marked as `"rejected"`).
   */
  reject: (reason?: any) => void;

  /**
   * Resolves the overlay (marked as `"resolved"`).
   */
  resolve: (result?: T) => void;

  /**
   * Updates custom data passed to the overlay.
   */
  updateData: (data: M) => void;
};
