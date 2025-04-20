import classNames from "classnames";
import React from "react";

export function PageSection({
  title,
  children,
  className,
  description,
}: {
  className?: string;
  description?: string;
  title: string;
  children?: React.ReactNode;
}) {
  return (
    <div className={classNames("page-section", className)}>
      {title ? <h2 className="">{title}</h2> : null}
      {description ? <p className="">{description}</p> : null}
      <div className="mt-4">{children}</div>
    </div>
  );
}
