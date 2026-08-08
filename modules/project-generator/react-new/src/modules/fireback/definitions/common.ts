export interface QueryArchiveColumn {
  name?: string;
  width?: number;
  title?: string;
  getCellValue?: (m: any) => any;
}
export interface Timestamp {
  seconds: number;
  nanos: number;
}

export enum MacTagsColor {
  Green = "#00bd00",
  Red = "#ff0313",
  Orange = "#fa7a00",
  Yellow = "#f4b700",
  Blue = "#0072ff",
  Purple = "#ad41d1",
  Grey = "#717176",
}
