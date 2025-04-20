import { Fallback } from "@/modules/fireback/components/fallback/Fallback";

import { ErrorBoundary } from "react-error-boundary";
import { HashRouter, Route, Routes, useLocation } from "react-router-dom";
import { CSSTransition, TransitionGroup } from "react-transition-group";
import { ModuleScreen } from "./ModuleScreen";
import "./styles.scss";

const Router = HashRouter;

function App() {
  return (
    <ErrorBoundary
      FallbackComponent={Fallback}
      onReset={(details) => {
        // Reset the state of your app so the error doesn't happen again
      }}
    >
      <Router future={{ v7_startTransition: true }}>
        <AnimatedRoutes />
      </Router>
    </ErrorBoundary>
  );
}

function AnimatedRoutes() {
  const location = useLocation();

  return (
    <TransitionGroup>
      <CSSTransition
        key={location.key}
        classNames="fade"
        timeout={140}
        unmountOnExit
      >
        <div className="page-wrapper">
          <Routes location={location}>
            <Route path="/" element={<ModuleScreen />} />
          </Routes>
        </div>
      </CSSTransition>
    </TransitionGroup>
  );
}

export default App;
