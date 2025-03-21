import { useContext } from "react";
import { RemoteQueryContext } from "./sdk/core/react-tools";

export const LoggedIn = () => {
  const { signout } = useContext(RemoteQueryContext);

  return (
    <div>
      You are logged in!
      <br />
      <button onClick={signout}>Logout</button>
    </div>
  );
};
