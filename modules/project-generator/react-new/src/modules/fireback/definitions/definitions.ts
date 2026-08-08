import { type IResponse } from "../sdk/core/http-tools";
import { type FormikProps } from "formik";

export interface KeyValue {
  label?: string;
  value?: string | number;
}

export type IndeterminateCheck = "checked" | "unchecked" | "indeterminate";

export interface DatatableColumn {
  name?: string;
  title?: string;
  width?: number;
  filterable?: boolean;
  sortable?: boolean;
  filterType?: "string" | "date";
  getCellValue?: (dto: any) => any;
}

export interface EntityManagerProps<T, V> {
  data?: Partial<T> | null;
  enabledFields?: {
    [key in keyof Partial<T>]: boolean;
  };
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  onSuccess?: (response: IResponse<T>) => void;
  context?: V;
}

export interface EntityFormProps<T> {
  enabledFields?: {
    [key in keyof Partial<T>]: boolean;
  };
  form: FormikProps<Partial<T>>;
  initialData?: Partial<T>;
  isEditing?: boolean;
}

export type JsonQuery = any;

export interface Context {
  url: string;
  token: string;
  workspaceId: string;
  body: any;
  acceptLanguage: string;
  method: string;
  itemsPerPage: number;
  startIndex?: number;
  paramValues: Array<string>;
}
