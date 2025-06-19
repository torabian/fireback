import { source } from "../../hooks/source";
import { osResources } from "../../resources/resources";
import ActiveLink from "../link/ActiveLink";

const tabs = [
  {
    to: "/dashboard",
    label: "Home",
    icon: source("/common/home.svg"),
  },
  { to: "/selfservice", label: "Profile", icon: source("/common/user.svg") },
  { to: "/settings", label: "Settings", icon: source(osResources.settings) },
];

export const TabbarMenu = () => {
  return (
    <nav className="bottom-nav-tabbar">
      {tabs.map((tab) => (
        <ActiveLink
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
