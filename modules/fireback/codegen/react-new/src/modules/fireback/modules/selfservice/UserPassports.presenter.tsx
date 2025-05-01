import { useContext } from "react";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useGetUserPassports } from "../../sdk/modules/abac/useGetUserPassports";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack } = useRouter();
  const { items, query } = useGetUserPassports({});
  const { signout } = useContext(RemoteQueryContext);

  return {
    items,
    goBack,
    signout,
    query,
    s,
  };
};
