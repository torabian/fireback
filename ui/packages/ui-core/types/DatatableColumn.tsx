export interface DatatableColumn {
  name?: string;
  title?: string;
  width?: number;
  filterable?: boolean;
  sortable?: boolean;
  /**
   * "tstring" is for a complexes.TString column (a locale -> text map, e.g.
   * {"en": "Home", "fa": "خانه"} - see modules/fireback/complexes/TString.go) -
   * DataGridListHeaderCell renders a button that opens TStringFilterDrawer
   * instead of a plain text input for these, since "what does the user even
   * type" is ambiguous for a multi-language value otherwise.
   */
  filterType?: "string" | "date" | "tstring";
  getCellValue?: (dto: any) => any;
}
