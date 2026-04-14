import { useContext } from "react";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { RemoteQueryContext } from "../../sdk/core/react-tools";

import { strings } from "./strings/translations";
import { useUserPassportsActionQuery } from "../../sdk/modules/abac/UserPassports";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack } = useRouter();
  const query = useUserPassportsActionQuery({});
  const { signout } = useContext(RemoteQueryContext);

  return {
    goBack,
    signout,
    query,
    s,
  };
};
