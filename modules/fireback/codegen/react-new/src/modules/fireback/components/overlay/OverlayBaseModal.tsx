import { ReactNode } from "react";
import {
  BaseModalOpenParams,
  OverlayInstanceComponentProps,
} from "./OverlayDefinitions";

export const OverlayBaseModal = ({
  children,
  close,
  params,
}: {
  children: ReactNode;
} & OverlayInstanceComponentProps<unknown, BaseModalOpenParams>) => {
  return (
    <div className="modal d-block with-fade-in modal-overlay">
      <div className="modal-dialog">
        <div className="modal-content">
          <div className="modal-header">
            <h5 className="modal-title">{params?.title}</h5>
            <button
              type="button"
              id="cls"
              className="btn-close"
              onClick={close}
              aria-label="Close"
            ></button>
          </div>
          {children}
        </div>
      </div>
    </div>
  );
};
