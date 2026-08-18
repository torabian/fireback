import "./FormTString.css";

import classNames from "classnames";
import { Languages } from "lucide-react";
import { useEffect, useState } from "react";
import {
  BaseFormElement,
  type BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import { useOverlay } from "../../overlay/OverlayProvider";
import { useSupportedLocales } from "../../../hooks/useSupportedLocales";
import { strings as coreStrings } from "../../strings/translations";
import { useS } from "../../../hooks/useS";
import { type TString } from "../../../types/TString";
import { TStringEditModal } from "./TStringEditModal";

export interface FormTStringProps extends Omit<BaseFormElementProps, "value"> {
  value?: TString | null;
  onChange?: (value: TString) => void;
  placeholder?: string;
  disabled?: boolean;
}

const CYCLE_MS = 2000;

// Form field for editing a TString value (types/TString.ts - a locale -> text map,
// mirroring complexes.TString.go), reusing the same "one input per language" idea
// DataGridList's TString column filter already uses (TStringFilterDrawer.tsx) - just
// through a modal instead of a drawer, and resolving the record itself instead of a
// filter condition. See TStringEditModal.tsx for the actual edit form.
//
// The closed field is a button, not a text input - there's no single string to show.
// With 2+ non-empty languages it crossfades between them every CYCLE_MS (2s) so the
// user can tell there's more than one without opening the modal; with 0 or 1 it just
// shows the placeholder or that one value, no animation (see FormTString.css).
export const FormTString = (props: FormTStringProps) => {
  const { value, onChange, disabled, placeholder, label, ...rest } = props;
  const cs = useS(coreStrings);
  const { openModal } = useOverlay();
  const locales = useSupportedLocales();

  const entries = Object.entries(value ?? {}).filter(
    ([, text]) => (text ?? "").trim() !== "",
  );

  const [activeIndex, setActiveIndex] = useState(0);

  useEffect(() => {
    if (entries.length < 2) {
      setActiveIndex(0);
      return;
    }
    const id = setInterval(() => {
      setActiveIndex((i) => (i + 1) % entries.length);
    }, CYCLE_MS);
    return () => clearInterval(id);
    // Only the *count* of non-empty entries should restart the cycle - re-keying it
    // on the entries array itself (a new array every render) would reset the timer
    // on every keystroke made elsewhere in the form.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [entries.length]);

  const active = entries[activeIndex % Math.max(entries.length, 1)];
  const cycling = entries.length > 1;

  const open = () => {
    if (disabled) return;
    openModal<TString>(
      (modalProps) => (
        <TStringEditModal
          {...modalProps}
          locales={locales}
          initialValues={value ?? {}}
        />
      ),
      { title: label || cs.actions.edit },
    ).promise.then(({ type, data }) => {
      if (type !== "resolved" || !data) return;
      onChange?.(data);
    });
  };

  return (
    <BaseFormElement {...rest} label={label} onClick={open}>
      <button
        type="button"
        className={classNames("form-control", "form-tstring")}
        onClick={open}
        disabled={disabled}
      >
        <Languages className="form-tstring__icon" />
        {active ? (
          <span
            key={cycling ? `${active[0]}-${activeIndex}` : active[0]}
            className={classNames(
              "form-tstring__value",
              cycling && "form-tstring__value--cycling",
            )}
          >
            <span className="form-tstring__locale">{active[0].toUpperCase()}</span>
            {active[1]}
          </span>
        ) : (
          <span className="form-tstring__placeholder">
            {placeholder || cs.common.isNUll}
          </span>
        )}
      </button>
    </BaseFormElement>
  );
};

export default FormTString;
