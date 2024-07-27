import { osResources } from "@/modules/fireback/resources/resources";

export function source(uri: string) {
  const prefix = process.env.REACT_APP_PUBLIC_URL || "";

  if (uri.startsWith("$")) {
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
