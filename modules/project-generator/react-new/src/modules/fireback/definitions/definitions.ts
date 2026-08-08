import { type IResponse } from "../sdk/core/http-tools";
import { type FormikProps } from "formik";

export enum KeyboardAction {
  NewEntity = "new_entity",
  SidebarToggle = "sidebarToggle",
  NewChildEntity = "new_child_entity",
  EditEntity = "edit_entity",
  ViewQuestions = "view_questions",
  ExportTable = "export_table",
  CommonBack = "common_back",
  StopStart = "StopStart",
  Delete = "delete",
  Select1Index = "select1_index",
  Select2Index = "select2_index",
  Select3Index = "select3_index",
  Select4Index = "select4_index",
  Select5Index = "select5_index",
  Select6Index = "select6_index",
  Select7Index = "select7_index",
  Select8Index = "select8_index",
  Select9Index = "select9_index",
  ToggleLock = "l",
}

export const NumericKeys = [
  KeyboardAction.Select1Index,
  KeyboardAction.Select2Index,
  KeyboardAction.Select3Index,
  KeyboardAction.Select4Index,
  KeyboardAction.Select5Index,
  KeyboardAction.Select6Index,
  KeyboardAction.Select7Index,
  KeyboardAction.Select8Index,
  KeyboardAction.Select9Index,
];

export interface KeyValue {
  label?: string;
  value?: string | number;
}

export interface StringKeyValue {
  label?: string;
  value?: string;
}

/**
 * Use this for selects, which provide more details than a key pair
 */
export interface OptionItem<T> extends KeyValue {
  data: Partial<T>;
  icon?: string;
  title?: string;
  description?: string;
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
