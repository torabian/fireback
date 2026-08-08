import { type ReactNode } from "react";
import {
  type BaseModalOpenParams,
  type OverlayInstanceComponentProps,
} from "./OverlayDefinitions";
import classNames from "classnames";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export const OverlayBaseModal = ({
  children,
  close,
  visible,
  params,
}: {
  children: ReactNode;
} & OverlayInstanceComponentProps<unknown, BaseModalOpenParams>) => {
  const s = useS(strings);
  return (
    <div
      className={classNames(
        "modal d-block with-fade-in modal-overlay",
        visible ? "visible" : "invisible"
      )}
    >
      <div className="modal-dialog">
        <div className="modal-content">
          <div className="modal-header">
            <h5 className="modal-title">{params?.title}</h5>
            <button
              type="button"
              id="cls"
              className="btn-close"
              onClick={close}
              aria-label={s.components.close}
            ></button>
          </div>
          {children}
        </div>
      </div>
    </div>
  );
};
