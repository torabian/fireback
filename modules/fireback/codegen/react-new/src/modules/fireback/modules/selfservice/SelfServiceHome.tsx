import { Link } from "react-router-dom";
import { PageSection } from "../../components/page-section/PageSection";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const SelfServiceHome = () => {
  const s = useS(strings);
  return (
    <>
      <PageSection title={s.home.title} description={s.home.description} />

      <h2>
        <Link to="passports">{s.home.passportsTitle}</Link>
      </h2>
      <p>{s.home.passportsDescription}</p>
      <Link to="passports" className="btn btn-success btn-sm">
        {s.home.passportsTitle}
      </Link>
    </>
  );
};
