import { RouterLink, useRouter } from "../../hooks/useRouter";
import { useCompiler } from "../../hooks/useEnvironment";
import { useLocale } from "../../hooks/useLocale";

const Link = ({
  children,
  isActive,
  skip,
  activeClassName,
  inActiveClassName,
  ...rest
}: any) => {
  const router = useRouter();
  const { locale } = useLocale();
  const locale$ = rest.locale || locale || "en";
  const { compiler } = useCompiler();
  const noPrefix = process.env.REACT_APP_NO_LOCALE_PREFIX === "true";

  let href: string = rest?.href || router?.asPath || "";

  if (typeof href === "string" && href?.indexOf && href.indexOf("http") === 0)
    skip = true;
  if (typeof href === "string" && locale$ && !skip && !href.startsWith(".")) {
    href = href
      ? (!noPrefix ? `/${locale}` : "") + href
      : router.pathname?.replace("[locale]", locale$);
  }

  if (isActive) {
    rest.className = `${rest.className || ""} ${activeClassName || "active"}`;
  }

  if (!isActive && inActiveClassName) {
    rest.className = `${rest.className || ""} ${inActiveClassName}`;
  }

  return (
    <RouterLink {...rest} href={href} compiler={compiler}>
      {children}
    </RouterLink>
  );
};

export default Link;
