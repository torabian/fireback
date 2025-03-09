import { QueryErrorView } from "../../components/error-view/QueryError";
import { usePresenter } from "./UserPassports.presenter";

export const UserPassportsScreen = ({}: {}) => {
  const { query, items, s } = usePresenter();

  return (
    <div className="signin-form-container">
      <h1>{s.userPassports.title}</h1>
      <p>{s.userPassports.description}</p>
      <QueryErrorView query={query} />

      <pre>{JSON.stringify(items, null, 2)}</pre>
    </div>
  );
};
