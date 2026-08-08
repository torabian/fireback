import type { GResponse } from "@/modules/sdk/sdk/envelopes";
import { type FormikProps } from "formik";

export interface EntityFormProps<T> {
  enabledFields?: {
    [key in keyof Partial<T>]: boolean;
  };
  form: FormikProps<Partial<T>>;
  initialData?: Partial<T>;
  isEditing?: boolean;
}

export interface EntityManagerProps<T, V> {
  data?: Partial<T> | null;
  enabledFields?: {
    [key in keyof Partial<T>]: boolean;
  };
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  onSuccess?: (response: GResponse<T>) => void;
  context?: V;
}
