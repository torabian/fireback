/**
 * Handy tool to create a final callable url, from query string, query params,
 * and the actual url.
 * @param template
 * @param params
 * @param qs
 * @returns
 */
export function buildUrl(
  url: string,
  params?: Record<string, unknown>,
  qs?: URLSearchParams
) {
  // Replace :placeholders
  if (params) {
    Object.entries(params as Record<string, string>).forEach(([key, value]) => {
      url = url.replace(
        new RegExp(`:${key}`, "g"),
        encodeURIComponent(String(value))
      );
    });
  }

  if (qs && qs instanceof URLSearchParams) {
    url += `?${qs.toString()}`;
  } else if (qs && Object.keys(qs).length) {
    const query = new URLSearchParams(
      Object.entries(qs).map(([k, v]) => [k, String(v)])
    ).toString();
    url += `?${query}`;
  }

  return url;
}
