// @ts-nocheck

/**
 * VITE_TARGET_APP is an environment variables that you can
 * use for spliting app into a different product.
 * when you import App.tsx from an specific project,
 * only routes and modules defined there will be added to your
 * pack. This is similar to C preprocessors idea.
 */
import React from "react";
import ReactDOM from "react-dom/client";

// #v-ifdef VITE_TARGET_APP == 'projectname'
import App from "./apps/projectname/App";
// #v-endif

// #v-ifdef VITE_TARGET_APP == 'designer'
import App from "./apps/designer/App";
// #v-endif

// #v-ifdef VITE_TARGET_APP == 'self-service'
import App from "./apps/self-service/App";
// #v-endif


// @fireback-append-app

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
