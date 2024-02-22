// @ts-nocheck

import React from "react";
import ReactDOM from "react-dom/client";

/// #if TARGET_APP == 'fireback'
import App from "./apps/fireback/App";

/// #endif

// @fireback-append-app

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
