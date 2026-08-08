export interface QueryArchiveColumn {
  name?: string;
  width?: number;
  title?: string;
  getCellValue?: (m: any) => any;
}
