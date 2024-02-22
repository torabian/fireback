import { Timestamp } from "src/sdk/fireback";
import moment from "moment";
import React, { useEffect, useRef, useState } from "react";

export const TimerUntil = ({
  until,
  onResend,
}: {
  until: string | Date | moment.Moment | Timestamp | number;
  onResend: () => void;
}) => {
  const [secondsLeft, setSecondsLeft] = useState(0);
  const ref = useRef<NodeJS.Timer>();

  const clear = () => {
    if (ref.current) {
      clearInterval(ref.current);
    }
  };

  useEffect(() => {
    const seconds = moment(until).diff(moment(), "second");

    if (seconds <= 0) {
      setSecondsLeft(0);
      clear();
    } else {
      setSecondsLeft(seconds);
      ref.current = setInterval(() => {
        setSecondsLeft((s) => s - 1);
        if (seconds <= 0) {
          clear();
        }
      }, 1000);
    }

    return () => {
      clear();
    };
  }, [until]);

  if (secondsLeft <= 0) {
    return (
      <div onClick={onResend}>
        <span>Tap to resend the activation code</span>
      </div>
    );
  }

  return <span>Please wait {secondsLeft} seconds more</span>;
};
