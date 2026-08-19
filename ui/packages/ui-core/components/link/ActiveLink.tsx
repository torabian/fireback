import { Children } from "react";
import Link from "./Link";
import { useLocale } from "../../hooks/useLocale";

const ActiveLink: any = (prp: any): any => {
  const { children, forceActive, ...props } = prp;

  const { asPath } = useLocale();
  const child = Children.only(children);

  // Bug fix: this used to compare asPath against a "/${locale}" + href -
  // routes no longer carry a locale segment (see Link.tsx's own history), so
  // href is compared as-is now.
  const active =
    asPath === props.href || asPath + "/" === props.href || forceActive;

  if (prp.disabled) {
    return <span className="disabled">{child}</span>;
  }

  return (
    <Link {...props} isActive={active}>
      {child}
    </Link>
  );
};

export default ActiveLink;
