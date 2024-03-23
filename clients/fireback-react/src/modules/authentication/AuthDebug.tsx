import { useState } from "react";
import { SessionDebug } from "../desktop-app-settings/DebuggerSettings";

export function AuthDebug() {
  const [visible, setVisibility] = useState(false);

  return (
    <div>
      <div>
        <button
          className="btn btn-sm btn-success"
          style={{ margin: "20px auto" }}
          onClick={() => setVisibility((x) => !x)}
        >
          Nerd
        </button>
      </div>
      {visible && <SessionDebug />}
    </div>
  );
}
