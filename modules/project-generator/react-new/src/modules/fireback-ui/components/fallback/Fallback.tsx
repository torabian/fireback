import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export function Fallback({ error, resetErrorBoundary }: any) {
  // Call resetErrorBoundary() to reset the error boundary and retry the render.
  const s = useS(strings);

  return (
    <div role="alert">
      <p>{s.components.somethingWentWrong}</p>
      <div style={{ color: "red", padding: "30px" }}>{error.message}</div>
    </div>
  );
}
