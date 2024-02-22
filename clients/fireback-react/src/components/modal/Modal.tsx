import { useKeyPress } from "@/hooks/useKeyPress";
import { useT } from "@/hooks/useT";
import classNames from "classnames";
import React, { useContext, useEffect, useState } from "react";
declare var bootstrap: any;

export interface ModalProps {
  component?: any;
  onSubmit?: () => Promise<any>;
  title?: string;
  confirmButtonLabel?: string;
}

export interface ModalRef extends ModalProps {
  id: string;
}

export interface Dialog {}

export interface IModalContext {
  refs: Array<ModalRef>;
  openModal: (modal: ModalProps) => void;
  confirm: () => Promise<any>;
  closeModal: (key: string) => void;
}

export const ModalContext = React.createContext<IModalContext>({
  openModal() {},
  closeModal(key: string) {},
  confirm() {
    return new Promise((r) => {
      r(false);
    });
  },

  refs: [],
});

export function ModalView({
  mref,
  context,
}: {
  mref: ModalRef;
  context: IModalContext;
}) {
  const t = useT();
  const Component = mref.component;
  const onSubmit = async () => {
    if (mref.onSubmit) {
      if ((await mref.onSubmit()) === true) {
        context.closeModal(mref.id);
      }
    }
  };

  return (
    <div className="modal d-block">
      <div className="modal-dialog">
        <div className="modal-content">
          <div className="modal-header">
            <h5 className="modal-title">{mref.title}</h5>
            <button
              type="button"
              id="cls"
              className="btn-close"
              onClick={() => context.closeModal(mref.id)}
              aria-label="Close"
            ></button>
          </div>
          <div className="modal-body">
            <p>
              <Component />
            </p>
          </div>
          <div className="modal-footer">
            <button
              type="button"
              className="btn btn-secondary"
              onClick={() => context.closeModal(mref.id)}
            >
              {t.close}
            </button>
            <button
              onClick={onSubmit}
              type="button"
              className="btn btn-primary"
            >
              {mref.confirmButtonLabel || t.saveChanges}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export function ModalManager() {
  const t = useContext(ModalContext);

  return (
    <>
      {t.refs.map((item) => (
        <ModalView key={item.id} context={t} mref={item} />
      ))}
      {t.refs.length ? (
        <div
          className={classNames("modal-backdrop fade", t.refs.length && "show")}
        ></div>
      ) : null}
    </>
  );
}

export function ModalProvider({ children }: { children: React.ReactNode }) {
  const [modalRefs, setModalRefs] = useState<Array<ModalRef>>([]);
  const openModal = (modal: ModalProps) => {
    let r = (Math.random() + 1).toString(36).substring(2);

    const newRef: ModalRef = { ...modal, id: r };
    setModalRefs((r) => [...r, newRef]);
  };

  const closeModal = (id: string) => {
    setModalRefs((r) => r.filter((t) => t.id !== id));
  };

  const confirm2 = (): Promise<any> => {
    return new Promise((r) => {
      // Confirm
      r(true);
    });
    // setModalRefs((r) => r.filter((t) => t.id !== id));
  };

  const closeFocused = () => {
    setModalRefs((r) => r.filter((t, index) => index !== r.length - 1));
  };

  useKeyPress("Escape", closeFocused);

  return (
    <ModalContext.Provider
      value={{
        confirm: confirm2,
        refs: modalRefs,
        closeModal,
        openModal,
      }}
    >
      {children}
    </ModalContext.Provider>
  );
}
