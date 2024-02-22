import { useEffect, useState } from "react";
import "react-phone-input-2/lib/style.css";
var jalaali = require("jalaali-js");

import { useT } from "@/hooks/useT";
import { localizeNumber } from "@/hooks/fonts";
import { useLocale } from "@/hooks/useLocale";

function getYears(range = 2) {
  const currentYear = 1402;
  const years = [];
  for (
    let i = Math.ceil(currentYear - range / 2);
    i < Math.ceil(currentYear + range / 2);
    i++
  ) {
    years.push(i);
  }
  return years;
}

function valueToDateNano(
  year: number,
  month: number,
  day: number,
  type: "european" | "jalali" = "european"
) {
  if (type === "european") {
    const d = new Date(year, month - 1, day);
    return d;
  }
  if (type === "jalali") {
    const m = jalaali.toGregorian(year, month, day);
    const d = new Date(m.gy, m.gm, m.gd);
    return d;
  }

  return new Date();
}

export function ReactRealDatePicker({
  type,
  onChange,
  value,
}: {
  type: "jalali" | "european";
  onChange?: (v: number) => void;
  value?: number | null;
}) {
  const { locale } = useLocale();
  const t = useT();

  const [pickerContent] = useState({
    years: getYears(5),
    days: [
      1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
      22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
    ],
    months: Object.values(t.jalaliMonths),
  });

  useEffect(() => {
    if (value) {
      if (type === "jalali") {
        const current = new Date(value / 1000000);
        const m = jalaali.toJalaali(
          current.getFullYear(),
          current.getMonth(),
          current.getDate()
        );
        if (m) {
          setInnerValue({
            day: m.jd,
            month: m.jm,
            year: m.jy,
          });
        }
      }
    }
  }, [value]);

  const [innerValue, setInnerValue] = useState({
    year: "1402",
    month: "5",
    day: "5",
  });

  const patchField = (field: string, value: any) => {
    const newV = {
      ...innerValue,
      [field]: value,
    };

    console.log(newV);

    const val = valueToDateNano(+newV.year, +newV.month, +newV.day, type);
    // Add 6 hours to cope with timezone issue.
    // Real fix is to add time zone
    const nano = (val.getTime() + 38400000) * 1000000;
    onChange && onChange(nano);
    setInnerValue((v) => newV);
  };

  return (
    <div className="form-control date-picker-inline">
      <select
        value={innerValue.day || ""}
        onChange={(e) => patchField("day", e.target.value)}
      >
        {pickerContent.days.map((i) => (
          <option value={i} key={i}>
            {localizeNumber("" + i, locale)}
          </option>
        ))}
      </select>

      <select
        value={innerValue.month || ""}
        onChange={(e) => patchField("month", e.target.value)}
      >
        {pickerContent.months.map((name, i) => (
          <option value={i} key={i}>
            {name}
          </option>
        ))}
      </select>

      <select
        value={innerValue.year || ""}
        onChange={(e) => patchField("year", e.target.value)}
      >
        {pickerContent.years.map((y, i) => (
          <option value={+y} key={y}>
            {localizeNumber("" + y, locale)}
          </option>
        ))}
      </select>
    </div>
  );
}
