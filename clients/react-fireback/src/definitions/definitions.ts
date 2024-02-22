import { IResponse } from "src/sdk/fireback";
import { FormikProps } from "formik";

export enum KeyboardAction {
  NewEntity = "new_entity",
  NewChildEntity = "new_child_entity",
  EditEntity = "edit_entity",
  ViewQuestions = "view_questions",
  ExportTable = "export_table",
  CommonBack = "common_back",
  StopStart = "StopStart",
  Delete = "delete",
}

export interface KeyValue {
  label?: string;
  value?: string | number;
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

export function toKeyValue<T>(
  items?: Array<T>,
  value?: string,
  label?: string
): KeyValue[] {
  if (!items) {
    return [];
  }

  return items.map((item: any) => {
    return {
      label: item[label || "name"],
      value: item[value || "uniqueId"],
    };
  });
}

export interface IPriceTag {
  amounts: Array<{
    value: number;
    currency: string;
  }>;
}

/**
 * Every fireback entity must have these fields implemented, in all entites
 */
export interface BaseRecord {
  uniqueId?: string;
  parentId?: string;
  roleId?: string;
  workspaceId?: string;
  createdAt?: string;
  deletedAt?: string;
  updatedAt?: string;
}

export class BaseRecord2 {
  static Fields = {
    uniqueId: "uniqueId",
    parentId: "parentId",
    roleId: "roleId",
    workspaceId: "workspaceId",
    createdAt: "createdAt",
    deletedAt: "deletedAt",
    updatedAt: "updatedAt",
  };
  uniqueId?: string;
  parentId?: string;
  roleId?: string;
  workspaceId?: string;
  createdAt?: string;
  deletedAt?: string;
  updatedAt?: string;
}

export interface Hierarchy {
  id: string;
  label?: string;
  chidlren?: Hierarchy[];
}

export type IndeterminateCheck = "checked" | "unchecked" | "indeterminate";

export interface DatatableColumn {
  name?: string;
  title?: string;
  width?: number;
  getCellValue?: (dto: any) => string | undefined;
}

export interface PermissionLevel {
  onlyRoot?: boolean;
  permissions: string[];
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
  isEditing?: boolean;
}

export type QuestionAnswerState = "correct" | "incorrect" | "blank";

export type FilterOperation =
  | `contains`
  | `notContains`
  | `startsWith`
  | `endsWith`
  | `equal`
  | `notEqual`
  | `greaterThan`
  | `greaterThanOrEqual`
  | `lessThan`
  | `lessThanOrEqual`;

export interface Filter {
  /** Specifies the name of a column whose value is used for filtering. */
  columnName: string;
  /** Specifies the operation name. The value is 'contains' if the operation name is not set. */
  operation?: FilterOperation;
  /** Specifies the filter value. */
  value?: any;
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
