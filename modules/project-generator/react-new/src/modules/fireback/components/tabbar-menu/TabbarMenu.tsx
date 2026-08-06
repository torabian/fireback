import { source } from "../../hooks/source";
import { useS } from "../../hooks/useS";
import { osResources } from "../../resources/resources";
import ActiveLink from "../link/ActiveLink";
import { strings } from "../strings/translations";

const tabs = (s: typeof strings) => [
  {
    to: "/dashboard",
    label: s.components.home,
    icon: source("/common/home.svg"),
  },
  { to: "/selfservice", label: s.components.profile, icon: source("/common/user.svg") },
  { to: "/settings", label: s.components.settings, icon: source(osResources.settings) },
];

export const TabbarMenu = () => {
  const s = useS(strings);
  return (
    <nav className="bottom-nav-tabbar">
      {tabs(s).map((tab) => (
        <ActiveLink
          state={{ animated: true }}
          key={tab.to}
          href={tab.to}
          className={({ isActive }) =>
            isActive ? "nav-link active" : "nav-link"
          }
        >
          <span className="nav-link">
            <img className="nav-img" src={tab.icon} />

            {tab.label}
          </span>
        </ActiveLink>
      ))}
    </nav>
  );
};
