/*
 * Converts all errors, network, api into an object that can
 * be passed to setErrors of formik ref.
 */
export function mutationErrorsToFormik(errors: any): {
  [key: string]: string;
} {
  const err: {[key: string]: string} = {};

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
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    var r = (Math.random() * 16) | 0,
      v = c == 'x' ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}

const errorToStirng = (error: any, t: any) => {
  if (!error) {
    return t.errors.UNKOWN_ERRROR;
  }
  let message = error?.messageTranslated || error?.message;
  for (let erritem of error?.errors || []) {
    message += ' ' + (erritem?.messageTranslated || erritem?.message);
  }
  return message;
};
