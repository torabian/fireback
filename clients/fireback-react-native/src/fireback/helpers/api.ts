const toLowerFirst = (input: string): string => {
  if (!input || typeof input !== 'string') {
    return '';
  }

  const str = input.split('');
  str[0] = str[0].toLowerCase();

  return str.join('');
};

/**
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

const HOST = 'http://localhost:4500';
// const HOST = 'https://auth.pixelplux.com'

export const execApi = (
  method: 'post' | 'get' | 'put' | 'delete' | 'patch',
  affix: string,
  body?: any,
) => {
  return fetch(HOST + '/' + affix, {
    method,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  }).then(response => {
    const contentType = response.headers.get('content-type');
    if (contentType && contentType.indexOf('application/json') !== -1) {
      return response.json().then(data => {
        if (response.ok) {
          return data;
        } else {
          throw data;
        }
      });
    } else {
      console.log(2, response);
      throw response;
    }
  });
};
