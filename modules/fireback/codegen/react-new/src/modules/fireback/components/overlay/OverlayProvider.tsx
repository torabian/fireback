import React, {
  createContext,
  FC,
  ReactNode,
  useContext,
  useEffect,
  useRef,
  useState,
} from "react";
import {
  BaseModalOpenParams,
  DialogResult,
  DrawerOpenParams,
  OpenOverlayConfig,
  OverlayContextType,
  OverlayController,
  OverlayInstance,
  OverlayInstanceComponentProps,
} from "./OverlayDefinitions";
import { OverlayBaseModal } from "./OverlayBaseModal";
import { OverlayDrawerImp } from "./OverlayDrawer";

const OverlayContext = createContext<OverlayContextType | null>(null);

let uniqueId = 0;

export const OverlayProvider = ({
  children,
  BaseModalWrapper = OverlayBaseModal,
  OverlayWrapper = OverlayDrawerImp,
}: {
  children: ReactNode;
  BaseModalWrapper?: FC<OverlayInstanceComponentProps<unknown>>;
  OverlayWrapper?: FC<OverlayInstanceComponentProps<unknown>>;
}) => {
  const [sheets, setSheets] = useState<OverlayInstance<unknown>[]>([]);
  const sheetsRef = useRef(sheets);
  sheetsRef.current = sheets;

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === "Escape" && sheets.length > 0) {
        const topSheet = sheets[sheets.length - 1];
        topSheet?.close?.();
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [sheets]);

  const openOverlay = <T = void,>(
    Component: OverlayInstance<unknown>["Component"],
    params?: OpenOverlayConfig
  ): OverlayController<T> => {
    const id = uniqueId++;
    const ref = React.createRef<any>();

    let resolveFn!: (result?: DialogResult<T>) => void;
    let rejectFn!: (reason?: DialogResult<any>) => void;

    const promise = new Promise<DialogResult<T>>((resolve, reject) => {
      resolveFn = resolve;
      rejectFn = reject;
    });

    const dismiss = () => {
      setSheets((prev) =>
        prev.map((s) => (s.id === id ? { ...s, visible: false } : s))
      );
      setTimeout(() => {
        setSheets((prev) => prev.filter((s) => s.id !== id));
      }, 300);
    };

    const overlayInstance: OverlayInstance<T> = {
      id,
      ref,
      Component,
      type: params?.type || "modal",
      params: params?.params,
      data: {},
      visible: false,
      onBeforeClose: undefined,
      resolve: (result?: T) => {
        setTimeout(() => resolveFn({ type: "resolved", data: result }), 50);
        dismiss();
      },
      close: async () => {
        const current = sheetsRef.current.find((item) => item.id === id);

        if (current?.onBeforeClose) {
          const allow = await current.onBeforeClose();
          if (!allow) return;
        }
        const shouldClose = await (overlayInstance.onBeforeClose?.() ?? true);
        if (!shouldClose) return;
        setTimeout(() => resolveFn({ data: null, type: "closed" }), 50);
        dismiss();
      },
      reject: (reason?: any) => {
        setTimeout(() => rejectFn({ data: reason, type: "rejected" }), 50);
        dismiss();
      },
    };

    setSheets((prev) => [...prev, overlayInstance]);
    setTimeout(() => {
      setSheets((prev) =>
        prev.map((s) => (s.id === id ? { ...s, visible: true } : s))
      );
    }, 50);

    const updateData = (newData: Partial<any>) => {
      setSheets((prev) =>
        prev.map((s) =>
          s.id === id ? { ...s, data: { ...s.data, ...newData } } : s
        )
      );
    };

    return {
      id,
      ref,
      promise,
      close: overlayInstance.close,
      resolve: overlayInstance.resolve,
      reject: overlayInstance.reject,
      updateData,
    };
  };

  const openModal = <T = void,>(
    Component: OverlayInstance<T>["Component"],
    params?: BaseModalOpenParams
  ): OverlayController<T> =>
    openOverlay<T>(Component, { type: "modal", params });

  const openDrawer = <T = void,>(
    Component: OverlayInstance<T>["Component"],
    params?: DrawerOpenParams
  ): OverlayController<T> =>
    openOverlay<T>(Component, { type: "drawer", params });

  const dismissAll = () => {
    sheetsRef.current.forEach((s) => s.reject?.("dismiss-all"));
    setSheets([]);
  };

  return (
    <OverlayContext.Provider
      value={{ openOverlay, openDrawer, openModal, dismissAll }}
    >
      {children}
      {sheets.map(
        ({
          id,
          type,
          Component,
          resolve,
          reject,
          close,
          params,
          visible,
          data,
        }) => {
          const C = type === "drawer" ? OverlayWrapper : BaseModalWrapper;
          return (
            <C
              key={id}
              visible={visible}
              close={close}
              reject={reject}
              resolve={resolve}
              params={params}
            >
              <Component
                resolve={resolve}
                reject={reject}
                close={close}
                data={data}
                setOnBeforeClose={(fn) => {
                  setSheets((prev) =>
                    prev.map((s) =>
                      s.id === id ? { ...s, onBeforeClose: fn } : s
                    )
                  );
                }}
              />
            </C>
          );
        }
      )}
    </OverlayContext.Provider>
  );
};

export const useOverlay = () => {
  const ctx = useContext(OverlayContext);
  if (!ctx) {
    throw new Error("useOverlay must be inside OverlayProvider");
  }
  return ctx;
};
