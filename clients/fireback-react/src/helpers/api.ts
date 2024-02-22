import { IError } from "src/sdk/fireback";
import { toast } from "react-toastify";
import { enTranslations } from "@/translations/en";

/*
 * Converts all errors, network, api into an object that can
 * be passed to setErrors of formik ref.
 */
export function mutationErrorsToFormik(errors: any): {
  [key: string]: string;
} {
  const err: { [key: string]: string } = {};

  if (errors.error && Array.isArray(errors.error?.errors)) {
    for (const field of errors.error?.errors) {
      err[field.location] = field.message;
    }
  }

  // This is when a network failure happens
  if (errors.status && errors.ok === false) {
    return {
      form: `${errors.status}`,
    };
  }

  if (errors?.error?.message) {
    err.form = errors?.error?.message;
  }

  if (errors.message) {
    return {
      form: `${errors.message}`,
    };
  }

  return err;
}

export function uuidv4() {
  return (([1e7] as any) + -1e3 + -4e3 + -8e3 + -1e11).replace(
    /[018]/g,
    (c: any) =>
      (
        c ^
        (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (c / 4)))
      ).toString(16)
  );
}

const errorToStirng = (error: IError, t: typeof enTranslations) => {
  if (!error) {
    return t.errors.UNKOWN_ERRROR;
  }
  let message = error?.messageTranslated || error?.message;
  for (let erritem of error?.errors || []) {
    message += " " + (erritem?.messageTranslated || erritem?.message);
  }
  return message;
};

export function httpErrorHanlder(res: any, t: typeof enTranslations) {
  toast(errorToStirng(res?.error, t), { type: "error" });
}
