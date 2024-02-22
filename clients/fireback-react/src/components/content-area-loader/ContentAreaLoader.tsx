import classNames from "classnames";
import { useEffect, useState } from "react";

import { LineLoader } from "@/components/line-loader/LineLoader";

/**
 * Use native html manipulation for this later, to improve it's performance.
 * @returns
 */
export function ContentAreaLoader({ show }: { show: boolean }) {
  const [visibleState, setVisibilityState] = useState<
    "HIDE" | "HIDING" | "SHOW"
  >("SHOW");
  useEffect(() => {
    if (show === false) {
      setTimeout(() => {
        setVisibilityState("HIDE");
      }, 501);
      setVisibilityState("HIDING");
    } else {
      setVisibilityState("SHOW");
    }
  }, [show]);

  if (visibleState !== "HIDE") {
    return (
      <>
        <div
          className={classNames("content-area-loader", {
            fadeout: visibleState === "HIDING",
          })}
        >
          <LineLoader />
        </div>
      </>
    );
  }

  return null;
}
