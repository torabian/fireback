// @ts-nocheck

import { loadFont, toFarsiNumber } from "@/hooks/fonts";
import { debounce } from "lodash";
import { useEffect, useRef } from "react";

const playSoundDebounced = debounce(playSound, 100, {
  trailing: true,
  leading: 100,
  maxWait: 100,
});

function playSound() {
  var SnapEffect = new Audio(
    process.env.REACT_APP_PUBLIC_URL + "soundeffects/snap.mp3"
  );
  SnapEffect.load();
  //   SnapEffect.pause();
  //   SnapEffect.currentTime = 0;
  SnapEffect.play();
}

interface NumericRange {
  position: number;
  index: number;
  value: number;
}

function createRange({
  absoluteDistance,
  fractionalCount,
}: {
  absoluteDistance;
  fractionalCount;
}) {
  const range: NumericRange[] = [];
  for (let i = 0; i <= 140; i++) {
    range.push({
      index: i,
      value: 80 + i,
      type: "major",
      position: 10 + i * absoluteDistance,
      label: toFarsiNumber(i + 80),
    });

    const distanceBetweenEach = absoluteDistance / (fractionalCount + 1);

    for (let f = 0; f < fractionalCount; f++) {
      const eachDistance = distanceBetweenEach * f;
      const value = i + 80 + (1 / (fractionalCount + 1)) * (f + 1);
      range.push({
        index: i,
        value,
        label: toFarsiNumber(("" + value).replace(".", "/")),
        type: "fraction",
        position:
          10 + i * absoluteDistance + eachDistance + distanceBetweenEach,
      });
    }
  }

  return range;
}

export function NumberPicker({
  canvasId,
  horizontal,
  onChange,
}: {
  canvasId: string;
  horizontal?: boolean;
  onChange?: (v: any) => void;
}) {
  const absoluteDistance = 40;
  let value = useRef(null);
  let positionY = useRef(-5000);
  const canvas_width = 250;
  const canvas_height = 1200;
  let savedPosition = useRef(0);
  let filterRange = useRef([]);
  let beginMouseY = useRef(0);
  const mouseDown = useRef(false);
  const fractionalCount = 1;
  const range = createRange({ absoluteDistance, fractionalCount });
  const distanceBetweenEach = absoluteDistance / (fractionalCount + 1);

  function filterToViewport(
    range: NumericRange[],
    scrollY: number,
    canvasHeight: number
  ) {
    const filterRange: NumericRange = [];

    for (let point of range) {
      //   if (point.position + scrollY < 0) {
      //     continue;
      //   }
      //   if (point.position + scrollY > canvasHeight / 2) {
      //     continue;
      //   }

      const currentRender = point.position + scrollY;
      const center = canvasHeight / 4;
      const slop = distanceBetweenEach / 2;
      if (currentRender - slop < center && currentRender + slop > center) {
        if (value.current !== point.value) {
          onChange && onChange(point.value);
          value.current = point.value;
          playSound();
        }
      }
      filterRange.push(point);
    }

    return filterRange;
  }

  function clock() {
    //   positionY.current += 1;
    const now = new Date();
    const canvas: HTMLCanvasElement = document.getElementById(canvasId);
    canvas.width = canvas_width;
    canvas.height = canvas_height;
    canvas.style.width = canvas.width / 2 + "px";
    canvas.style.height = canvas.height / 2 + "px";
    const ctx = canvas.getContext("2d") as CanvasRenderingContext2D;
    ctx.scale(2, 2);
    ctx.save();
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    ctx.beginPath();
    ctx.moveTo(95, canvas.height / 4);
    ctx.lineTo(100, canvas.height / 4 + 5);
    ctx.lineTo(100, canvas.height / 4 - 5);
    ctx.fill();

    const absoluteColor = "red";
    const fractionalColor = "black";
    const fractionalTextColor = "#d0cfd4";

    // console.log(filterRange);
    for (let point of filterRange.current || range) {
      if (point.type === "major") {
        ctx.beginPath();
        ctx.rect(10, point.position + positionY.current, 30, 1);
        ctx.fillStyle = absoluteColor;
        ctx.fill();

        ctx.font = "12px shabnam";
        ctx.fillStyle = "black";
        ctx.fillText(point.label, 60, point.position + positionY.current + 4);
      } else if (point.type === "fraction") {
        ctx.beginPath();
        ctx.rect(10, point.position + positionY.current, 15, 1);
        ctx.fillStyle = fractionalColor;
        ctx.fill();

        ctx.font = "10px shabnam";
        ctx.fillStyle = fractionalTextColor;
        ctx.fillText(point.label, 40, point.position + positionY.current + 4);
      }
    }

    // for (let i = 0; i <= 140; i++) {
    //   ctx.beginPath();
    //   const y = 10 + i * absoluteDistance + positionY.current;
    //   ctx.rect(10, y, 30, 1);
    //   ctx.fillStyle = absoluteColor;
    //   ctx.fill();

    //   ctx.font = "12px shabnam";
    //   ctx.fillText(toFarsiNumber(i + 80), 60, y + 4);

    //   for (let m = 0; m < fractionalCount; m++) {
    //     const distanceBetweenEach = absoluteDistance / (fractionalCount + 1);
    //     const eachDistance = distanceBetweenEach * m;
    //     ctx.beginPath();
    //     ctx.rect(
    //       10,
    //       10 +
    //         i * absoluteDistance +
    //         positionY.current +
    //         eachDistance +
    //         distanceBetweenEach,
    //       15,
    //       1
    //     );
    //     ctx.fillStyle = fractionalColor;
    //     ctx.fill();
    //   }
    // }

    if (positionY > 100) {
      positionY = 0;
    }

    ctx.save();

    window.requestAnimationFrame(clock);
  }

  useEffect(() => {
    loadFont(
      "shabnam",
      process.env.REACT_APP_PUBLIC_URL + "fonts/shabnam/Shabnam.ttf"
    ).then(() => {
      window.requestAnimationFrame(clock);
    });

    filterRange.current = filterToViewport(
      range,
      positionY.current,
      canvas_height
    );
  }, []);

  const moveUp = () => {
    const ref = setInterval(() => {
      positionY.current -= 5;

      if (positionY.current <= 0) {
        clearInterval(ref);
      }
    }, 1);
  };

  const moveDown = () => {
    const ref = setInterval(() => {
      positionY.current += 5;

      if (positionY.current >= -5000) {
        clearInterval(ref);
      }
    }, 1);
  };

  const onMouseMove = (e) => {
    const Y = getEventY(e);

    if (!mouseDown.current) {
      return;
    }

    filterRange.current = filterToViewport(
      range,
      positionY.current,
      canvas_height
    );

    positionY.current =
      savedPosition.current +
      ((horizontal ? e.clientX : Y) - beginMouseY.current);

    //   console.log(positionY.current, savedPosition.current);
  };

  const getEventY = (e) => {
    let Y = null;
    if (
      e.type == "touchstart" ||
      e.type == "touchmove" ||
      e.type == "touchend" ||
      e.type == "touchcancel"
    ) {
      const touch = e.touches[0] || e.changedTouches[0];
      Y = touch.clientY;
    } else if (
      e.type == "mousedown" ||
      e.type == "mouseup" ||
      e.type == "mousemove" ||
      e.type == "mouseover" ||
      e.type == "mouseout" ||
      e.type == "mouseenter" ||
      e.type == "mouseleave"
    ) {
      Y = e.clientY;
    }
    return Y;
  };

  const onMouseDown = (e) => {
    const Y = getEventY(e);
    mouseDown.current = true;
    beginMouseY.current = horizontal ? e.clientX : Y;
    savedPosition.current = positionY.current;
  };

  const onMouseUp = (e) => {
    const lastItem = filterRange.current[filterRange.current.length - 1];

    if (positionY.current > 0) {
      moveUp();
    } else if (positionY.current < -lastItem.position + canvas_height / 4) {
      moveDown();
    }

    mouseDown.current = false;
  };

  return (
    <>
      <canvas
        onMouseMove={onMouseMove}
        onMouseDown={onMouseDown}
        onMouseUp={onMouseUp}
        onTouchStart={onMouseDown}
        onTouchMove={onMouseMove}
        onTouchEnd={onMouseUp}
        // onPointerDown={onMouseDown}
        // onPointerUp={onMouseUp}
        id={canvasId}
        width="300"
        height="500"
      ></canvas>
    </>
  );
}
