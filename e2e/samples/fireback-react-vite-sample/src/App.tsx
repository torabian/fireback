import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { LoggedIn } from "./LoggedIn";
import { useContext, useRef } from "react";
import {
  RemoteQueryContext,
  RemoteQueryProvider,
} from "./sdk/core/react-tools";
import { NotLoggedIn } from "./NotLoggedIn";
import { QueryClient } from "react-query";

function Application() {
  const { session } = useContext(RemoteQueryContext);

  return (
    <Router>
      {session ? (
        <Routes>
          <Route path="/" Component={LoggedIn} />
        </Routes>
      ) : (
        <Routes>
          <Route path="/" Component={NotLoggedIn} />
        </Routes>
      )}
    </Router>
  );
}

function App() {
  const queryClient = useRef(new QueryClient());
  return (
    <RemoteQueryProvider
      identifier="sample_app"
      queryClient={queryClient.current}
    >
      <Application />
    </RemoteQueryProvider>
  );
}

export default App;
