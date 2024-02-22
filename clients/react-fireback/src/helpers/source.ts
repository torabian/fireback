import { osResources } from "@/components/mulittarget/multitarget-resource";

export function source(uri: string) {
  const prefix = process.env.REACT_APP_PUBLIC_URL || "";

  if (uri.startsWith("$")) {
    // console.log(88, uri, uri.substr(1), osResources[uri.substr(1)]);
    return prefix + (osResources as any)[uri.substr(1)];
  }

  if (!prefix) {
    return uri;
  }

  // Maybe somehow already applied. Think about this rule
  if (uri.startsWith(prefix)) {
    return uri;
  }

  return prefix + uri;
}
