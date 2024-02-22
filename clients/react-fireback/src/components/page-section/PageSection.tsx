import classNames from "classnames";
import React from "react";

export function PageSection({
  title,
  children,
  className,
}: {
  className?: string;
  title: string;
  children?: React.ReactNode;
}) {
  return (
    <div className={classNames("page-section", className)}>
      <h2>{title}</h2>
      {children}
    </div>
  );
}
