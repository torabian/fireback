import { Toast } from "./toast";
import { strings } from "../components/strings/translations";

/*
 * Converts all errors, network, api into an object that can
 * be passed to setErrors of formik ref.
 *
 * `errors` can be either:
 * - a rejected exception (network failure, thrown Error, ...), or
 * - a *resolved* response envelope whose `.error` is populated - our fetch
 *   layer doesn't reject the mutation promise for backend-returned errors
 *   (4xx/5xx bodies still resolve), so callers that check
 *   `res.error`/`res instanceof GResponse` themselves pass the envelope
 *   straight in here instead of catching it.
 *
 * Field-level messages go under their `location` key (matched up with a
 * form field's name); anything that isn't tied to a specific field goes
 * under `form`, so a form-level error banner can render it.
 */
export function mutationErrorsToFormik(errors: any): any {
  const err: { [key: string]: string } = {};

  if (errors?.error && Array.isArray(errors.error?.errors)) {
    for (const field of errors.error.errors) {
      err[field.location] = field.messageTranslated || field.message;
    }
  }

  if (errors?.error?.messageTranslated || errors?.error?.message) {
    err.form = errors.error.messageTranslated || errors.error.message;
  } else if (errors?.status && errors?.ok === false) {
    // A raw, non-JSON HTTP failure (no parsed error envelope to read from).
    err.form = `${errors.status}`;
  } else if (errors?.message) {
    err.form = `${errors.message}`;
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
      ).toString(16),
  );
}

const errorToStirng = (error: any, s: typeof strings) => {
  if (!error) {
    return s.errors.UNKOWN_ERRROR;
  }
  let message = error?.messageTranslated || error?.message;
  for (let erritem of error?.errors || []) {
    message += " " + (erritem?.messageTranslated || erritem?.message);
  }
  return message;
};

export function httpErrorHanlder(res: any, s: typeof strings) {
  // Toast has errors
  // Toast("failed", { type: "error" });
}
