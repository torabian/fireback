import { Children } from "react";
import Link from "./Link";
import { useLocale } from "../../hooks/useLocale";

const ActiveLink: any = (prp: any): any => {
  const { children, forceActive, ...props } = prp;

  const { locale, asPath } = useLocale();
  const child = Children.only(children);

  const noPrefix = process.env.REACT_APP_NO_LOCALE_PREFIX === "true";

  const active =
    asPath === (!noPrefix ? `/${locale}` : "") + props.href ||
    asPath + "/" === (!noPrefix ? `/${locale}` : "") + props.href ||
    forceActive;

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
