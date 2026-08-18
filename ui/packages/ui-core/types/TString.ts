/**
 * Front-end mirror of complexes.TString (modules/fireback/complexes/TString.go):
 * a locale -> text map, e.g. {"en": "Home", "fa": "خانه"}, JSON-marshaled as a flat
 * object. Shared by every place that edits or filters one - FormTString
 * (forms/form-tstring/FormTString.tsx) and DataGridList's "tstring" column filter
 * (data-grid-list/TStringFilterDrawer.tsx) - so they agree on the same shape instead
 * of each declaring their own `Record<string, string>` alias.
 */
export type TString = Record<string, string>;
