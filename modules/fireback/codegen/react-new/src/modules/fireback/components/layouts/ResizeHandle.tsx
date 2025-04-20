import classNames from "classnames";
import { PanelResizeHandle } from "react-resizable-panels";

export function ResizeHandle({
  className = "",
  id,
  onDragComplete,
  minimal,
}: {
  className?: string;
  id?: string;
  onDragComplete?: () => void;
  minimal?: boolean;
}) {
  return (
    <PanelResizeHandle
      id={id}
      onDragging={(dragging) => {
        if (dragging === false) {
          onDragComplete?.();
        }
      }}
      className={classNames("panel-resize-handle", minimal ? "minimal" : "")}
    ></PanelResizeHandle>
  );
}
