import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { useGetUserPassports } from "../../sdk/modules/workspaces/useGetUserPassports";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack } = useRouter();
  const { items, query } = useGetUserPassports({});

  return {
    items,
    goBack,
    query,
    s,
  };
};
