import { PanelResizeHandle } from "react-resizable-panels";

export function ResizeHandle({
  className = "",
  id,
  onDragComplete,
}: {
  className?: string;
  id?: string;
  onDragComplete?: () => void;
}) {
  return (
    <PanelResizeHandle
      id={id}
      onDragging={(dragging) => {
        if (dragging === false) {
          onDragComplete?.();
        }
      }}
      className="panel-resize-handle"
    ></PanelResizeHandle>
  );
}
