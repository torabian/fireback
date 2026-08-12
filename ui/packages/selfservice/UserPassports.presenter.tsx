import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useAuthentication } from "@fireback/auth-client/AuthenticationContext";

import { strings } from "./strings/translations";
import { useUserPassportsActionQuery } from "@fireback/selfservice/sdk/abac/UserPassportsAction";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack } = useRouter();
  const query = useUserPassportsActionQuery({});
  const { signout } = useAuthentication();

  return {
    goBack,
    signout,
    query,
    s,
  };
};
