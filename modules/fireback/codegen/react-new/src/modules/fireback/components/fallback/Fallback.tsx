export function Fallback({ error, resetErrorBoundary }: any) {
  // Call resetErrorBoundary() to reset the error boundary and retry the render.

  return (
    <div role="alert">
      <p>Something went wrong:</p>
      <div style={{ color: "red", padding: "30px" }}>{error.message}</div>
    </div>
  );
}
