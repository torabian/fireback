import { osResources } from "../resources/resources";
import { BUILD_VARIABLES } from "./build-variables";

export function source(uri: string) {
  const prefix = BUILD_VARIABLES.PUBLIC_URL || "";

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
