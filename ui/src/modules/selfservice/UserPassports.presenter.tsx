import { useRouter } from "../fireback-ui/hooks/useRouter";
import { useS } from "../fireback-ui/hooks/useS";
import { useAuthentication } from "../fireback-ui/auth/AuthenticationContext";

import { strings } from "./strings/translations";
import { useUserPassportsActionQuery } from "../sdk/abac/UserPassportsAction";

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
