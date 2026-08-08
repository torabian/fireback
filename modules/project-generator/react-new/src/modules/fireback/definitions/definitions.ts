export interface KeyValue {
  label?: string;
  value?: string | number;
}

export interface DatatableColumn {
  name?: string;
  title?: string;
  width?: number;
  filterable?: boolean;
  sortable?: boolean;
  filterType?: "string" | "date";
  getCellValue?: (dto: any) => any;
}
