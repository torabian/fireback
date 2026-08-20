export interface DatatableColumn {
  name?: string;
  title?: string;
  width?: number;
  filterable?: boolean;
  /**
   * The field name to filter by, when it differs from `name` - e.g. `name`
   * is the frontend/DTO field ("firstName") used for react-data-grid's
   * column key and for reading the cell value, while the backend's
   * ApplyQueryFilter/jsonlogic2sql expects the actual DB column name
   * ("first_name"). Defaults to `name` when omitted.
   */
  filterKey?: string;
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
