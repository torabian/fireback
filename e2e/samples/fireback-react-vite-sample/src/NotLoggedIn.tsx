import { useContext, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { RemoteQueryContext } from "./sdk/core/react-tools";

export const NotLoggedIn = () => {
  const { setSession } = useContext(RemoteQueryContext);
  const navigate = useNavigate();

  const login = () => {
    window.location.href =
      "http://localhost:4508/selfservice?redirect=" +
      encodeURIComponent("http://localhost:5173");
  };

  const [searchParams] = useSearchParams();

  useEffect(() => {
    try {
      const session = searchParams.get("session");
      if (session) {
        const b = JSON.parse(session);
        if (b && b.token) {
          setSession(b);
        }
        navigate(window.location.pathname, { replace: true });
      }
    } catch (er) {}
  }, [searchParams]);

  return (
    <div>
      Not logged in! Click to Login
      <br />
      <button onClick={login}>Login</button>
    </div>
  );
};
