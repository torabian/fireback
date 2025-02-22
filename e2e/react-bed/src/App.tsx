import { useRef } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import "./App.css";
import { QueryCapabilitiesTest } from "./components/basic-test/QueryCapabilitiesTest";
import { SignupTest } from "./components/basic-test/SignupTest";
import logo from "./logo.svg";
import { RemoteQueryProvider } from "./sdk/core/react-tools";

function App() {
  const queryClient = useRef(new QueryClient());
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          This is a way to know if fireback typescript is compiled correctly for
          react.
        </p>
        <p>If you see this screen without errors, then we successed</p>

        <QueryClientProvider client={queryClient.current}>
          <RemoteQueryProvider
            identifier="fireback"
            queryClient={queryClient.current}
          >
            <SignupTest />
          </RemoteQueryProvider>
          <QueryCapabilitiesTest />
        </QueryClientProvider>
      </header>
    </div>
  );
}

export default App;
