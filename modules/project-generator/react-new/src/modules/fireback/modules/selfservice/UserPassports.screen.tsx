import { QueryErrorView } from "../../components/error-view/QueryError";
import ActiveLink from "../../components/link/ActiveLink";
import { useS } from "../../hooks/useS";
import type { UserPassportsActionResType } from "../../sdk/abac/UserPassportsAction";
import { strings } from "./strings/translations";
import { usePresenter } from "./UserPassports.presenter";

export const UserPassportsScreen = ({ }: {}) => {
  const { query, s, signout } = usePresenter();
  const items = query?.data?.data?.items || []

  return (
    <div className="signin-form-container">
      <h1>{s.userPassports.title}</h1>
      <p>{s.userPassports.description}</p>
      <QueryErrorView query={query} />

      <PassportList passports={items} />

      <button className="btn btn-danger mt-3 w-100" onClick={signout}>
        {s.signout}
      </button>
    </div>
  );
};

const PassportList = ({
  passports,
}: {
  passports: UserPassportsActionResType[];
}) => {
  const s = useS(strings);
  return (
    <div className="d-flex ">
      {passports.map((passport) => (
        <div key={passport.uniqueId} className="card p-3 w-100">
          <h3 className="card-title">{passport.type.toUpperCase()}</h3>
          <p className="card-text">{passport.value}</p>
          <p className="text-muted">
            {s.totpLabel} {passport.totpConfirmed ? s.yes : s.no}
          </p>
          <ActiveLink href={`../change-password/${passport.uniqueId}`}>
            <button className="btn btn-primary">
              {s.changePassword.submit}
            </button>
          </ActiveLink>
        </div>
      ))}
    </div>
  );
};
