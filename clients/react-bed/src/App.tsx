import React, { useRef } from 'react';
import logo from './logo.svg';
import './App.css';
import { SignupTest } from './components/basic-test/SignupTest';
import { QueryClient, QueryClientProvider } from 'react-query';
import { QueryCapabilitiesTest } from './components/basic-test/QueryCapabilitiesTest';

function App() {
  const queryClient = useRef(new QueryClient())
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          This is a way to know if fireback typescript is compiled correctly for react.
        </p>
        <p>
          If you see this screen without errors, then we successed
        </p>


        <QueryClientProvider client={queryClient.current}>
          <SignupTest />
          <QueryCapabilitiesTest />
        </QueryClientProvider>
        
      </header>

    </div>
  );
}

export default App;
