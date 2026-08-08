import classNames from "classnames";

export function LineLoader({ className }: { className?: string }) {
  return (
    <div className={classNames("loader", className)}>
      <div className="loader__element"></div>
    </div>
  );
}
