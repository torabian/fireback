import { useOverlay } from "./OverlayProvider";

/**
 * @description A set of very frequent used dialogs ui
 * for confirming, showing message, etc.
 * expand this if you think there are useful and repeatable.
 * for custom needs, build one using useOverlay hook and provider.
 */
export const commonDialogs = () => {
  const { openDrawer, openModal } = useOverlay();

  /**
   * Use for yes/no dialogs, where user needs to decide between 'yes' or 'no'
   * to an operation.
   * @param title
   * @param description
   * @param buttonLabel
   */
  const confirmDrawer = ({
    title,
    description,
    cancelLabel,
    confirmLabel,
  }: {
    title: string;
    description: string;
    cancelLabel?: string;
    confirmLabel?: string;
  }) => {
    return openDrawer(({ close, resolve }) => (
      <div className="confirm-drawer-container p-3">
        <h2>{title}</h2>
        <span>{description}</span>
        <div>
          <button
            className="d-block w-100 btn btn-primary"
            onClick={() => resolve()}
          >
            {confirmLabel}
          </button>
          <button className="d-block w-100 btn" onClick={() => close()}>
            {cancelLabel}
          </button>
        </div>
      </div>
    ));
  };
  /**
   * Use for yes/no dialogs, where user needs to decide between 'yes' or 'no'
   * to an operation.
   * @param title
   * @param description
   * @param buttonLabel
   */
  const confirmModal = ({
    title,
    description,
    cancelLabel,
    confirmLabel,
  }: {
    title: string;
    description: string;
    cancelLabel?: string;
    confirmLabel?: string;
  }) => {
    return openModal(
      ({ close, resolve }) => (
        <div className="confirm-drawer-container p-3">
          <span>{description}</span>
          <div className="row mt-4">
            <div className="col-md-6">
              <button
                className="d-block w-100 btn btn-primary"
                onClick={() => resolve()}
              >
                {confirmLabel}
              </button>
            </div>
            <div className="col-md-6">
              <button className="d-block w-100 btn" onClick={() => close()}>
                {cancelLabel}
              </button>
            </div>
          </div>
        </div>
      ),
      { title }
    );
  };

  return { confirmDrawer, confirmModal };
};
