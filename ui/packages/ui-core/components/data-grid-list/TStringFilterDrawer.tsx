import { useState } from "react";
import { FormText } from "../forms/form-text/FormText";
import { strings as coreStrings } from "../strings/translations";
import { useS } from "../../hooks/useS";
import { useSupportedLocales } from "../../hooks/useSupportedLocales";

export interface TStringFilterDrawerResult {
  /** locale -> search text, e.g. {"en": "home", "fa": "خانه"} */
  values: Record<string, string>;
}

// Opened by DataGridListHeaderCell for a "tstring" column (see DatatableColumn.tsx's
// filterType doc comment) instead of a plain text filter input - a complexes.TString
// value (modules/fireback/complexes/TString.go: a locale -> text map, e.g.
// {"en": "Home", "fa": "خانه"}) is stored as one JSON object per row, so there's no
// single flat string a plain "contains" filter could target; this lets the user type
// search text per language instead, one input each.
//
// Same render-prop shape (close/resolve) as every other openDrawer(...) caller in this
// app - see WorkspaceAddUserDrawer.tsx or UserPassportsList.tsx's SetPasswordDrawer.
// Locales offered come from useSupportedLocales() (the same list the rest of the app's
// own translations are built from), not a hardcoded set, so this stays in sync with
// whatever languages a given deployment actually ships. FormTString's own edit modal
// (forms/form-tstring/TStringEditModal.tsx) is the form-input analogue of this same
// idea - one field per locale - just resolving the TString record itself instead of a
// filter condition, and opened via openModal instead of openDrawer.
export const TStringFilterDrawer = ({
  close,
  resolve,
  fieldLabel,
  initialValues,
}: {
  close: () => void;
  resolve: (result?: TStringFilterDrawerResult) => void;
  fieldLabel: string;
  initialValues?: Record<string, string>;
}) => {
  const cs = useS(coreStrings);
  const locales = useSupportedLocales();

  const [values, setValues] = useState<Record<string, string>>(initialValues ?? {});

  return (
    <div className="confirm-drawer-container p-3">
      <h2>{fieldLabel}</h2>
      {locales.map((locale) => (
        <FormText
          key={locale}
          value={values[locale] ?? ""}
          onChange={(value) => setValues((prev) => ({ ...prev, [locale]: value }))}
          label={locale.toUpperCase()}
          placeholder={cs.table.filter.filterPlaceholder}
        />
      ))}
      <div>
        <button
          className="d-block w-100 btn btn-primary"
          onClick={() => resolve({ values })}
        >
          {cs.common.save}
        </button>
        <button className="d-block w-100 btn" onClick={() => close()}>
          {cs.common.cancel}
        </button>
      </div>
    </div>
  );
};

export default TStringFilterDrawer;
